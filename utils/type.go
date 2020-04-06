// Copyright (c) 2018 Clearmatics Technologies Ltd
package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/apaxa-go/eval"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// SolidityToStaticGoType takes some input string that needs to be converted and an abi.Type that it should be converted to
//
// This function is designed to be used to convert inputs to type-safe variables to be used by abi.Pack() to construct
// transaction payloads.
//
// It converts the input to an array format where necessary, then uses the converted input and the type to create a set
// of initialisers in Go expressions that are then evaluated to produce statically typed variables which are returned
func SolidityToStaticGoType(input string, ty abi.Type) (interface{}, error) {
	convertedInput, err := ConvertStringArray(input)
	if err != nil {
		return nil, err
	}

	initialiser, err := constructInitialiser(convertedInput, ty)
	if err != nil {
		return nil, err
	}

	expr, err := eval.ParseString(initialiser, "")
	if err != nil {
		return nil, err
	}

	a := eval.Args{
		"reflect.ValueOf":     eval.MakeDataRegularInterface(reflect.ValueOf),
		"reflect.Value":       eval.MakeTypeInterface(reflect.Value{}),
		"big.Int":             eval.MakeTypeInterface(*big.NewInt(0)),
		"big.NewInt":          eval.MakeDataRegularInterface(big.NewInt),
		"common.Address":      eval.MakeTypeInterface(common.HexToAddress("")),
		"common.HexToAddress": eval.MakeDataRegularInterface(common.HexToAddress),
	}

	r, err := expr.EvalToInterface(a)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// constructInitialiser takes an string/[]string input and an expected solidity type and returns an initialiser go
// expression as a string so that the expression can be evaluated elsewhere to produce type-safe input values to pack.
//
// It works by checking if the input is a slice of values or just a single value, and checks whether the intended target
// type is also of the same structure (array/slice/single value) and passes the string representation of the value to
// be used as the data to be initialised to the target type.
//
// If the input is a slice/array, we iteratively construct initialisation expressions for each item in the slice as per the
// element type expected in the solidity type passed.
//
// If the input is a single value, we return an initialiser expression for the supported types.
func constructInitialiser(input interface{}, ty abi.Type) (string, error) {
	switch reflect.TypeOf(input).Kind() {
	case reflect.Slice: // If input type is a slice, then we expect the expected solidity type to be a slice or array
		if ty.Kind != reflect.Array && ty.Kind != reflect.Slice {
			return "", errors.New("type mismatch: expected type is regular identifier, got array of items")
		}

		var initialisers string
		inputSlice := reflect.ValueOf(input)

		// Check that the lengths of the input slice is the same as the expected array type
		// ty.Size = 0 if the expected type is a slice
		if ty.Size > 0 && inputSlice.Len() != ty.Size {
			return "", errors.New(fmt.Sprintf("constructInitialiser: input length mismatch: input has %d values but expected type has %d items", inputSlice.Len(), ty.Size))
		}

		// Construct initialiser expression for each item in the slice
		for i := 0; i < inputSlice.Len(); i++ {
			var elementType abi.Type
			if ty.Elem == nil {
				if strings.Contains(ty.Type.String(), "[]uint8") {
					elementType = abi.Type{
						Elem:          nil,
						Kind:          reflect.Uint8,
						Type:          reflect.TypeOf(uint8(1)),
						Size:          8,
						T:             1,
						TupleRawName:  "",
						TupleElems:    nil,
						TupleRawNames: nil,
					}
				}
			} else {
				elementType = *ty.Elem
			}


			elementInitialiser, err := constructInitialiser(inputSlice.Index(i).Interface(), elementType)
			if err != nil {
				return "", err
			}

			if i > 0 {
				elementInitialiser = "," + elementInitialiser
			}
			initialisers = initialisers + elementInitialiser
		}

		return fmt.Sprintf("%s{%s}", ty.Type.String(), initialisers), nil
	case reflect.String: // Single element, directly construct initialiser
		if ty.Elem != nil {
			return "", errors.New("constructInitialiser: type mismatch error: end of input reached but type has nested elements")
		}
		switch ty.Type.String() {
		case "*big.Int":
			return fmt.Sprintf("big.NewInt(%s)", input), nil
		case "common.Address":
			return fmt.Sprintf("common.HexToAddress(\"%s\")", input), nil
		case "bool":
			return fmt.Sprintf("%s", input), nil
		case "string":
			return fmt.Sprintf("\"%s\"", input), nil
		default: // Else it could be various intx, uintx, byte array combinations or an unsupported type
			if strings.Contains(ty.String(), "byte") { // byte[n], bytesn, byte[n][m], bytesn[m], etc.
				stringInput, ok := input.(string)
				if !ok {
					return "", errors.New("constructInitialiser: expected single item string")
				}
				bytes, err := fromHex(stringInput)
				if err != nil {
					return "", errors.New(fmt.Sprintf("constructInitialiser: %s", err.Error()))
				}
				fmt.Println("bytes:", bytes)
				fmt.Println(ty.Size)

				if ty.Size > 0 && ty.Size != len(bytes) {
					return "", errors.New(fmt.Sprintf("constructInitialiser: incorrect byte array length error: expected length %d but input array has length %d", ty.Size, len(bytes)))
				}

				bytesString := strings.ReplaceAll(fmt.Sprintf("%d", bytes), "[", "")
				bytesString = strings.ReplaceAll(bytesString, "]", "")
				bytesString = strings.ReplaceAll(bytesString, " ", ",")

				return fmt.Sprintf("%s{%s}", ty.Type.String(), bytesString), nil
			} else if strings.Contains(ty.Type.String(), "int") { // (u)int8, (u)int16, (u)int32, (u)int64
				return fmt.Sprintf("%s(%s)", ty.Type.String(), input), nil
			}
			return "", errors.New(fmt.Sprintf("constructInitialiser: unexpected abi type %s", ty.Type.String()))
		}
	}

	return "", errors.New(fmt.Sprintf("constructInitialiser: unexpected input type, expected %s or %s, got %s", reflect.Slice, reflect.String, reflect.TypeOf(input).Kind()))
}

func fromHex(s string) ([]byte, error) {
	if has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return hex.DecodeString(s)
}

func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}
