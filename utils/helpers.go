package utils

import (
	"errors"
	"fmt"
	"github.com/apaxa-go/eval"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"reflect"
	"regexp"
	"strings"
)

// Checks if input type is a regular type, or nested arrays or slices according to r
// Input will be in the form of slices of types or just the type itself
// This function will check, given the input, whether the slice lengths and nested structures match the static
// structure of r
// 		[ item, item ] has the same structure as [ 0, 1 ]
// 		[ [ item, item ], item ] does not have the same fingerprint as [ [ item, item ], [ item, item ] ]
// 		"something" has the same fingerprint as 1234
// 		[ "string", "string" ] has the same fingerprint as [ 1234, 5678 ]
//		[ [], [] ] does not have the same fingerprint as [ [ item, item ], [ item, item ] ]
func HasExpectedArrayStructure(sliceStructure interface{}, staticType reflect.Type) bool {
	fmt.Println("CHECKING ARRAY STRUCTURE OF", sliceStructure)
	fmt.Println("AGAINST", staticType)

	fmt.Println("Input:", sliceStructure)
	fmt.Println("Input Type:", reflect.TypeOf(sliceStructure))
	fmt.Println("Input Type.Kind:", reflect.TypeOf(sliceStructure).Kind())
	if isSlice(sliceStructure) || isArray(sliceStructure) {
		fmt.Println("Input Type.Elem:", reflect.TypeOf(sliceStructure).Elem())
	}

	fmt.Println("Static type:", staticType)
	fmt.Println("Static type Kind:", staticType.Kind())
	fmt.Println("Static type String:", staticType.String())

	switch staticType.Kind() {
	case reflect.Slice:
		fmt.Println("Expected type is slice")
		if !isSlice(sliceStructure) {
			return false
		}

		contents := reflect.ValueOf(sliceStructure)
		fmt.Println("Contents length:", contents.Len())

		// If the input slice has values, check that each item also matches the static structure
		for i := 0; i < contents.Len(); i++ {
			item := contents.Index(i).Interface()
			staticItemType := staticType.Elem()

			fmt.Println("Static type", staticItemType)
			fmt.Println("Item", item)
			fmt.Println("Item type", reflect.TypeOf(item))

			if !HasExpectedArrayStructure(item, staticItemType) {
				fmt.Println("Fingerprint check fail")

				fmt.Println("Input elem", item, "does not have same fingerprint as expected elem", staticItemType)
				return false
			}
		}

		return true

		break
	case reflect.Array:
		if staticType.String() == "common.Address" && reflect.TypeOf(sliceStructure).Kind() == reflect.String {
			fmt.Println("Input is a string representation of address")
			return true
		}
		fmt.Println("Expected type is array")
		if !isSlice(sliceStructure) && !isArray(sliceStructure) {
			return false
		}

		expectedLength := staticType.Len()
		contents := reflect.ValueOf(sliceStructure)

		fmt.Println("Expected length:", expectedLength)
		fmt.Println("input length:", contents.Len())

		// Ensure the length of the sliceStructure array is the same as the expected type
		if contents.Len() == expectedLength {
			// Check for all items in the sliceStructure array, that they have the same fingerprint as the respective expected type
			for i := 0; i < contents.Len(); i++ {
				inputElem := contents.Index(i).Interface()
				staticItemType := staticType.Elem()

				fmt.Println("Checking fingerprint of array", contents, i, "with", staticItemType)
				pass := HasExpectedArrayStructure(inputElem, staticItemType)
				fmt.Println("Does", contents, i, inputElem, "have same fingerprint as", staticItemType, "?:", pass)
				if !HasExpectedArrayStructure(inputElem, staticItemType) {
					fmt.Println("Fingerprint check fail")

					fmt.Println("Input elem", sliceStructure, "does not have same fingerprint as expected elem", staticItemType)
					return false
				}
			}

			return true
		}
		break
	default:
		fmt.Println("Expected type is not array/slice")
		if !isSlice(sliceStructure) && !isArray(sliceStructure) {
			fmt.Println("Input type is not slice")
			return true
		}
		fmt.Println("Input type is slice!")
	}

	return false
}

func isSlice(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.Slice
}

func isArray(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.Array
}

// Takes a string representation of an array (including nested array) and converts to array of strings
// Only supports directly n-dimensional arrays where the items of any array is either an array or a value
// Meaning we do not support arrays like [item, [arrayitem, arrayitem], item]
// Only [item, item, item] or [[arrayitem, arrayitem], [array2item, array2item]]
func ConvertStringArray(input string) (interface{}, error) {
	if string(input[0]) == "[" { // Item is array
		finalEnd := len([]rune(input)) - 1

		var array []interface{}

		start := 0
		for start < finalEnd {
			if string(input[start]) == "," || string(input[start]) == " " {
				start++
				continue
			}

			// Given the beginning of an array, find the end of this array
			end, err := findEndOfArray(input, start)
			if err != nil {
				return nil, err
			}

			// Remove the wrapping braces and take only the contents of this array and parse
			contents := input[start+1 : end]
			item, err := ConvertStringArray(contents)
			if err != nil {
				return nil, err
			}

			// If the item is the only item, no need to wrap in another array
			if start == 0 && end == finalEnd {
				return item, nil
			}

			array = append(array, item)
			start = end + 1
		}

		return array, nil
	} else {
		return strings.Split(strings.ReplaceAll(input, " ", ""), ","), nil
	}
}

func findEndOfArray(input string, start int) (int, error) {
	if string(input[start]) != "[" {
		return 0, errors.New("func findEndOfArray: invalid start index")
	}

	sub := 0
	for index, char := range input[start:] {
		if string(char) == "[" {
			sub++
		} else if string(char) == "]" {
			if sub > 1 {
				sub--
			} else {
				return start + index, nil
			}
		}
	}

	return 0, errors.New("func findEndOfArray: could not find end of array")
}

func SolidityToStaticGoType(input string, ty abi.Type) (interface{}, error) {
	convertedInput, err := ConvertStringArray(input)
	if err != nil {
		return nil, err
	}

	//if !HasExpectedArrayStructure(convertedInput, ty) {
	//	return nil, errors.New(fmt.Sprintf("input does not match the expected type structure %s", ty.Type.String()))
	//}

	fmt.Println("Converted input", convertedInput)

	initialiser, err := ConstructInitialiser(convertedInput, ty)
	fmt.Println("Initialiser expression", initialiser)

	//goType := BindTypeGo(ty)
	//
	//fmt.Println(fmt.Sprintf("Translated Solidity type to Go type: %s", goType))

	//var newType string

	//if strings.Contains(goType, "big.Int") {
	//	if strings.Contains(goType, "[") {
	//		newType = goType + "{}"
	//	} else {
	//		newType = "big.NewInt(0)"
	//	}
	//} else if strings.Contains(goType, "byte") {
	//	newType = goType + "{}"
	//} else if strings.Contains(goType, "int") {
	//	if strings.Contains(goType, "[") {
	//		newType = goType + "{}"
	//	} else {
	//		newType = goType + "(0)"
	//	}
	//} else if strings.Contains(goType, "bool") {
	//	if strings.Contains(goType, "[") {
	//		newType = goType + "{}"
	//	} else {
	//		newType = goType + "(true)"
	//	}
	//} else if strings.Contains(goType, "common.Address") {
	//	newType = goType + "{}"
	//} else if strings.Contains(goType, "string") {
	//	if strings.Contains(goType, "[") {
	//		newType = goType + "{}"
	//	} else {
	//		return "", nil
	//	}
	//}
	//
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

//func ConvertToGoType(input interface{}, output interface{}) (interface{}, error) {
//	outputType := reflect.TypeOf(output)
//
//	if !HasExpectedArrayStructure(input, outputType) {
//		return nil, errors.New("could not convert to go type: input has different type structure")
//	}
//
//	fmt.Println("output type", outputType)
//
//	switch outputType.Kind() {
//	case reflect.Slice, reflect.Array:
//		if outputType.String() == "common.Address" {
//			// If output kind is not array or slice, input type should be string or arraystructure check fails
//			inputString, ok := input.(string)
//			if !ok {
//				return nil, errors.New("convertToGoType: error should not occur. input could not be asserted to type string")
//			}
//			result, err := EvaluateGoTypeWithValue(inputString, outputType)
//			if err != nil {
//				return nil, err
//			}
//
//			return result, nil
//		}
//
//		inputItems := reflect.ValueOf(input)
//		outputItems := reflect.ValueOf(output)
//		outputItemsInterface := reflect.ValueOf(output).Interface()
//
//		for i := 0; i < inputItems.Len(); i++ {
//			item, err := convertToGoType(inputItems.Index(i).Interface(), outputItems.Index(i).Interface())
//			if err != nil {
//				return nil, err
//			}
//
//			fmt.Println("item", item)
//			fmt.Println("item type", reflect.TypeOf(item))
//
//			fmt.Println("output items", outputItems)
//			fmt.Println("output items interface", outputItemsInterface)
//			fmt.Println("output items interface type", reflect.TypeOf(outputItemsInterface))
//
//			if outputType.Kind() == reflect.Slice {
//				asserted, ok := outputItemsInterface.([]interface{})
//				if !ok {
//					return nil, errors.New("convertToGoType: type assertion error: failed to assert slice type to output variable")
//				}
//				outputItemsInterface = append(asserted, item)
//
//
//			} else if outputType.Kind() == reflect.Array {
//
//			}
//		}
//
//		break
//	default:
//		// If output kind is not array or slice, input type should be string or arraystructure check fails
//		inputString, ok := input.(string)
//		if !ok {
//			return nil, errors.New("convertToGoType: error should not occur. input could not be asserted to type string")
//		}
//		result, err := EvaluateGoTypeWithValue(inputString, outputType)
//		if err != nil {
//			return nil, err
//		}
//
//		return result, nil
//	}
//
//	return nil, errors.New("convertToGoType: unknown error: unable to convert")
//}

//func EvaluateGoTypeWithValue(input interface{}, ty reflect.Type) (interface{}, error) {
//	fmt.Println("TYPE", ty)
//
//	var newType string
//
//	if strings.Contains(ty.String(), "big.Int") {
//		newType = fmt.Sprintf("big.NewInt(%s)", input)
//	} else if strings.Contains(ty.String(), "byte") {
//		//bytes, err := hex.DecodeString(input)
//		//if err != nil {
//		//	return nil, err
//		//}
//		//singleByte := bytes[0]
//		//newType = goType + "{}"
//	} else if strings.Contains(ty.String(), "int") {
//		newType = ty.String() + fmt.Sprintf("(%s)", input)
//	} else if strings.Contains(ty.String(), "bool") {
//		newType = ty.String() + fmt.Sprintf("(%s)", input)
//	} else if strings.Contains(ty.String(), "common.Address") {
//		newType = fmt.Sprintf("common.HexToAddress(\"%s\")", input)
//	} else if strings.Contains(ty.String(), "string") {
//		return input, nil
//	} else {
//		return nil, errors.New(fmt.Sprintf("EvaluateGoTypeWithValue type error: type %s unsupported", ty.String()))
//	}
//
//	fmt.Println("NEW TYPE:", newType)
//
//	expr, err := eval.ParseString(fmt.Sprintf("reflect.ValueOf(%s).Interface()", newType), "")
//	if err != nil {
//		return nil, err
//	}
//
//	a := eval.Args{
//		"reflect.ValueOf": eval.MakeDataRegularInterface(reflect.ValueOf),
//		"reflect.Value": eval.MakeTypeInterface(reflect.Value{}),
//		"big.Int": eval.MakeTypeInterface(*big.NewInt(0)),
//		"big.NewInt": eval.MakeDataRegularInterface(big.NewInt),
//		"common.Address": eval.MakeTypeInterface(common.HexToAddress("")),
//		"common.HexToAddress": eval.MakeDataRegularInterface(common.HexToAddress),
//	}
//
//	r, err := expr.EvalToInterface(a)
//	if err != nil {
//		return nil, err
//	}
//
//	return r, nil
//}

func ConstructInitialiser(input interface{}, ty abi.Type) (string, error) {
	fmt.Println("Input Type", reflect.TypeOf(input))
	fmt.Println("Input Type.Kind", reflect.TypeOf(input).Kind())
	switch reflect.TypeOf(input).Kind() {
	case reflect.Slice:
		if ty.Kind != reflect.Array && ty.Kind != reflect.Slice {
			return "", errors.New("type mismatch: expected type is regular identifier, got array of items")
		}
		fmt.Println("ty Type", ty.Type.String())

		var initialisers string

		inputSlice := reflect.ValueOf(input)
		for i := 0; i < inputSlice.Len(); i++ {
			elementInitialiser, err := ConstructInitialiser(inputSlice.Index(i).Interface(), *ty.Elem)
			if err != nil {
				return "", err
			}
			fmt.Println("Element Initialiser", elementInitialiser)

			if i > 0 {
				elementInitialiser = "," + elementInitialiser
			}

			initialisers = initialisers + elementInitialiser
		}

		return fmt.Sprintf("%s{%s}", ty.Type.String(), initialisers), nil
	case reflect.String: // Single element
		fmt.Println("Type", ty.Type.String())
		switch ty.Type.String() {
		case "*big.Int":
			return fmt.Sprintf("big.NewInt(%s)", input), nil
		case "common.Address":
			return fmt.Sprintf("common.HexToAddress(\"%s\")", input), nil
		case "bool":
			return fmt.Sprintf("%s", input), nil
		case "string":
			return fmt.Sprintf("%s", input), nil
		default:
			if strings.Contains(ty.Type.String(), "int") {
				return fmt.Sprintf("%s", input), nil
			}
			return "", errors.New(fmt.Sprintf("ConstructInitialiser: unexpected abi type %s", ty.Type.String()))
		}
	}

	return "", errors.New(fmt.Sprintf("ConstructInitialiser: unexpected input type, expected %s or %s, got %s", reflect.Slice, reflect.String, reflect.TypeOf(input).Kind()))
}

func BindTypeGo(kind abi.Type) string {
	fmt.Println(fmt.Sprintf("Binding Type: %s", kind.String()))
	fmt.Println(fmt.Sprintf("Binding Type.Type: %s", kind.Type))
	fmt.Println(fmt.Sprintf("Binding Type.Elem: %s", kind.Elem))
	stringKind := kind.String()
	innerLen, innerMapping := bindUnnestedTypeGo(stringKind)
	fmt.Println(fmt.Sprintf("InnerLen: %d", innerLen))
	fmt.Println(fmt.Sprintf("InnerMapping: %s", innerMapping))

	innerMapping, parts := wrapArray(stringKind, innerLen, innerMapping)

	fmt.Println(fmt.Sprintf("Parts: %s", parts))

	arrayBinding := arrayBindingGo(innerMapping, parts)
	fmt.Println(fmt.Sprintf("Array Binding: %s", arrayBinding))

	fmt.Println("")
	return arrayBinding
}

func bindUnnestedTypeGo(stringKind string) (int, string) {
	switch {
	case strings.HasPrefix(stringKind, "address"):
		return len("address"), "common.Address"

	case strings.HasPrefix(stringKind, "bytes"):
		parts := regexp.MustCompile(`bytes([0-9]*)`).FindStringSubmatch(stringKind)
		return len(parts[0]), fmt.Sprintf("[%s]byte", parts[1])

	case strings.HasPrefix(stringKind, "int") || strings.HasPrefix(stringKind, "uint"):
		parts := regexp.MustCompile(`(u)?int([0-9]*)`).FindStringSubmatch(stringKind)
		switch parts[2] {
		case "8", "16", "32", "64":
			return len(parts[0]), fmt.Sprintf("%sint%s", parts[1], parts[2])
		}
		return len(parts[0]), "*big.Int"

	case strings.HasPrefix(stringKind, "bool"):
		return len("bool"), "bool"

	case strings.HasPrefix(stringKind, "string"):
		return len("string"), "string"

	default:
		return len(stringKind), stringKind
	}
}

func wrapArray(stringKind string, innerLen int, innerMapping string) (string, []string) {
	remainder := stringKind[innerLen:]
	//find all the sizes
	matches := regexp.MustCompile(`\[(\d*)\]`).FindAllStringSubmatch(remainder, -1)
	parts := make([]string, 0, len(matches))
	for _, match := range matches {
		//get group 1 from the regex match
		parts = append(parts, match[1])
	}
	return innerMapping, parts
}

func arrayBindingGo(inner string, arraySizes []string) string {
	out := ""
	//prepend all array sizes, from outer (end arraySizes) to inner (start arraySizes)
	for i := len(arraySizes) - 1; i >= 0; i-- {
		out += "[" + arraySizes[i] + "]"
	}
	out += inner
	return out
}
