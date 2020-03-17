// Copyright (c) 2018 Clearmatics Technologies Ltd
package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var NotArrayFormatError = errors.New("expected input as array format: <item>,<item>,<item>,...")

func ApplySolidityType(input string, argType abi.Type) (interface{}, error) {

	fmt.Println("Applying Solidity Type... ")
	fmt.Println("ArgType: ", argType)
	fmt.Println("ArgType.Elem: ", argType.Elem)
	fmt.Println("ArgType.Type: ", argType.Type)
	fmt.Println("ArgType.Kind: ", argType.Kind)
	fmt.Println("ArgType.Size: ", argType.Size)
	if argType.Kind == reflect.Array || argType.Kind == reflect.Slice { // Some solidity array type
		// bytes = []byte{} argument type = slice, no element, type equates to []uint8
		// byte[] = [][1]byte{} argument type = slice, element type = array, type equates to [][1]uint8
		// byte = bytes1
		// bytesn = [n]byte{} 0 < n < 33, argument type = array, no element, type equates to [n]uint8
		// bytesn[] = [][n]byte{} argument type = slice, element type = array, type equates to [][n]uint8
		// bytesn[m] = [m][n]byte{} argument type = array, element type = array, type equates to [m][n]uint8
		// Many annoying cases of byte arrays
		if argType.Elem == nil { // One dimensional byte array. Accepts all byte arrays as hex string with pre-pended '0x' only
			return ConvertOneDimensionalByteArray(input, argType)
		} else { // Elem has type, could be array of primitives or 2D byte array
			return ConvertGeneralArray(input, argType)
		}
	} else { // Is some simple solidity primitive type (including address/string which are byte-array aliases)
		return ConvertToType(input, &argType)
	}
}

func ConvertOneDimensionalByteArray(input string, argType abi.Type) (interface{}, error) {
	if argType.Type == reflect.TypeOf(common.Address{}) { // address solidity type
		item, err := ConvertToType(input, &argType)
		if err != nil {
			return nil, err
		}
		return item, nil
	} else if argType.Type == reflect.TypeOf([]byte{}) { // bytes solidity type
		bytes, err := hex.DecodeString(input[2:])
		if err != nil {
			return nil, err
		}
		return bytes, nil
	} else {
		// Fixed byte array of size n; bytesn solidity type
		// Any submitted bytes longer than the expected size will be truncated
		return ConvertToByteArray(input, argType)
	}
}

func ConvertGeneralArray(input string, argType abi.Type) (interface{}, error) {
	elementType := argType.Elem
	fmt.Println("ElementType: ", elementType)
	fmt.Println("ElementType.Type: ", elementType.Type)
	fmt.Println("ElementType.Kind: ", elementType.Kind)
	fmt.Println("ElementType.Size: ", elementType.Size)
	//fmt.Println("Number of inputs: ", size)
	//fmt.Println("Expected number of inputs: ", argSize)

	// Elements cannot be kind slice                                        only mean slice
	if elementType != nil &&
		elementType.Kind == reflect.Array &&
		elementType.Type != reflect.TypeOf(common.Address{}) {
		// Is 2D byte array
		/* Nightmare to implement, have to account for:
		   * Slice of fixed byte arrays; bytes32[] in solidity for example, generally bytesn[]
		   * Fixed array of fixed byte arrays; bytes32[10] in solidity for example, generally bytesn[m]
		   * Slice or fixed array of string; identical to above two cases as string in solidity is array of bytes

		   Since the upper bound of elements in an array in solidity is 2^256-1, and each fixed byte array
		   has a limit of bytes32 (bytes1, bytes2, ..., bytes31, bytes32), and Golang array creation takes
		   constant length values, we would have to paste the switch-case containing 1-32 fixed byte arrays
		   2^256-1 times to handle every possibility. Since arrays of arrays in seldom used, we have not
		   implemented it.
		*/

		if elementType.Size == 1 {
			/*		Solidity byte[]
					 	Normal byte array but in the form of []bytes1 = [][1]uint8 which makes it look like 2D array

						Solidity byte[n]
						Looks like [n]bytes1 = [n][1]uint8. Abi Pack accepts creation of dynamic length array filled to expected length
			*/

			bytes, err := hex.DecodeString(input[2:])
			if err != nil {
				return nil, err
			}

			var array [][1]uint8
			for _, item := range bytes {
				array = append(array, [1]uint8{item})
			}

			return array, nil
		} /*else if elementType.Size == 2 {
			bytes, err := hex.DecodeString(input[2:])
			if err != nil {
				return nil, err
			}

			if len(bytes) % elementType.Size > 0 {
				return nil, errors.New(fmt.Sprintf("array length mismatch: cannot convert %d bytes to [][%d]byte array, please supply a multiple of %d", len(bytes), elementType.Size, elementType.Size))
			}

			var array [][]uint8
			for i := 0; i < len(bytes); i += elementType.Size  {
				array = append(array, []uint8{bytes[i], bytes[i+1]})
			}

			return array, nil
		} else {
			bytes, err := hex.DecodeString(input[2:])
			if err != nil {
				return nil, err
			}

			if len(bytes) % elementType.Size > 0 {
				return nil, errors.New(fmt.Sprintf("array length mismatch: cannot convert %d bytes to [][%d]byte array, please supply a multiple of %d", len(bytes), elementType.Size, elementType.Size))
			}

			var array []interface{}

			for i := 0; i < len(bytes); i += elementType.Size {
				subarray := makeDynamicSubArray(bytes, elementType.Size, i)
				array = append(array, subarray)
			}

			return array, nil
		} */

		return nil, errors.New("2d arrays unsupported, use \"bytes\" instead")

		/*
		   slice := make([]interface{}, 0, size)
		   err = addFixedByteArrays(array, elementType.Size, slice)
		   if err != nil {
		       return nil, err
		   }
		   args = append(args, slice)
		   continue
		*/
	} else {

		array := strings.Split(input, ",")
		argSize := argType.Size
		size := len(array)
		if argSize != 0 && size != argSize {
			return nil, NotArrayFormatError
		}

		switch elementType.Type {
		case reflect.TypeOf(false):
			convertedArray := make([]bool, 0, size)
			for _, item := range array {
				b, err := ConvertToBool(item)
				if err != nil {
					return nil, err
				}
				convertedArray = append(convertedArray, b)
			}
			return convertedArray, nil
		case reflect.TypeOf(int8(0)):
			convertedArray := make([]int8, 0, size)
			for _, item := range array {
				i, err := strconv.ParseInt(item, 10, 8)
				if err != nil {
					return nil, err
				}
				convertedArray = append(convertedArray, int8(i))
			}
			return convertedArray, nil
		case reflect.TypeOf(int16(0)):
			convertedArray := make([]int16, 0, size)
			for _, item := range array {
				i, err := strconv.ParseInt(item, 10, 16)
				if err != nil {
					return nil, err
				}
				convertedArray = append(convertedArray, int16(i))
			}
			return convertedArray, nil
		case reflect.TypeOf(int32(0)):
			convertedArray := make([]int32, 0, size)
			for _, item := range array {
				i, err := strconv.ParseInt(item, 10, 32)
				if err != nil {
					return nil, err
				}
				convertedArray = append(convertedArray, int32(i))
			}
			return convertedArray, nil
		case reflect.TypeOf(int64(0)):
			convertedArray := make([]int64, 0, size)
			for _, item := range array {
				i, err := strconv.ParseInt(item, 10, 64)
				if err != nil {
					return nil, err
				}
				convertedArray = append(convertedArray, int64(i))
			}
			return convertedArray, nil
		case reflect.TypeOf(uint8(0)):
			convertedArray := make([]uint8, 0, size)
			for _, item := range array {
				u, err := strconv.ParseUint(item, 10, 8)
				if err != nil {
					return nil, err
				}
				convertedArray = append(convertedArray, uint8(u))
			}
			return convertedArray, nil
		case reflect.TypeOf(uint16(0)):
			convertedArray := make([]uint16, 0, size)
			for _, item := range array {
				u, err := strconv.ParseUint(item, 10, 16)
				if err != nil {
					return nil, err
				}
				convertedArray = append(convertedArray, uint16(u))
			}
			return convertedArray, nil
		case reflect.TypeOf(uint32(0)):
			convertedArray := make([]uint32, 0, size)
			for _, item := range array {
				u, err := strconv.ParseUint(item, 10, 32)
				if err != nil {
					return nil, err
				}
				convertedArray = append(convertedArray, uint32(u))
			}
			return convertedArray, nil
		case reflect.TypeOf(uint64(0)):
			convertedArray := make([]uint64, 0, size)
			for _, item := range array {
				u, err := strconv.ParseUint(item, 10, 64)
				if err != nil {
					return nil, err
				}
				convertedArray = append(convertedArray, uint64(u))
			}
			return convertedArray, nil
		case reflect.TypeOf(&big.Int{}):
			convertedArray := make([]*big.Int, 0, size)
			for _, item := range array {
				newInt := new(big.Int)
				newInt, ok := newInt.SetString(item, 10)
				if !ok {
					return nil, errors.New("Could not convert string to big.int")
				}
				convertedArray = append(convertedArray, newInt)
			}
			return convertedArray, nil
		case reflect.TypeOf(common.Address{}):
			convertedArray := make([]common.Address, 0, size)
			for _, item := range array {
				a := common.HexToAddress(item)
				convertedArray = append(convertedArray, a)
			}
			return convertedArray, nil
		case reflect.TypeOf("String"):
			convertedArray := make([]string, 0, size)
			for _, item := range array {
				convertedArray = append(convertedArray, item)
			}
			return convertedArray, nil
		default:
			errStr := fmt.Sprintf("Type %s not found", elementType.Type)
			return nil, errors.New(errStr)
		}
	}
}

func ConvertToByteArray(input string, argType abi.Type) (interface{}, error) {
	// Fixed byte array of size n; bytesn solidity type
	// Any submitted bytes longer than the expected size will be truncated
	bytes, err := hex.DecodeString(input[2:])
	if err != nil {
		return nil, err
	}

	// Fixed sized arrays can't be created with variables as size
	switch argType.Size {
	case 1:
		var byteArray [1]byte
		copy(byteArray[:], bytes[:1])
		return byteArray, nil
	case 2:
		var byteArray [2]byte
		copy(byteArray[:], bytes[:2])
		return byteArray, nil
	case 3:
		var byteArray [3]byte
		copy(byteArray[:], bytes[:3])
		return byteArray, nil
	case 4:
		var byteArray [4]byte
		copy(byteArray[:], bytes[:4])
		return byteArray, nil
	case 5:
		var byteArray [5]byte
		copy(byteArray[:], bytes[:5])
		return byteArray, nil
	case 6:
		var byteArray [6]byte
		copy(byteArray[:], bytes[:6])
		return byteArray, nil
	case 7:
		var byteArray [7]byte
		copy(byteArray[:], bytes[:7])
		return byteArray, nil
	case 8:
		var byteArray [8]byte
		copy(byteArray[:], bytes[:8])
		return byteArray, nil
	case 9:
		var byteArray [9]byte
		copy(byteArray[:], bytes[:9])
		return byteArray, nil
	case 10:
		var byteArray [10]byte
		copy(byteArray[:], bytes[:10])
		return byteArray, nil
	case 11:
		var byteArray [11]byte
		copy(byteArray[:], bytes[:11])
		return byteArray, nil
	case 12:
		var byteArray [12]byte
		copy(byteArray[:], bytes[:12])
		return byteArray, nil
	case 13:
		var byteArray [13]byte
		copy(byteArray[:], bytes[:13])
		return byteArray, nil
	case 14:
		var byteArray [14]byte
		copy(byteArray[:], bytes[:14])
		return byteArray, nil
	case 15:
		var byteArray [15]byte
		copy(byteArray[:], bytes[:15])
		return byteArray, nil
	case 16:
		var byteArray [16]byte
		copy(byteArray[:], bytes[:16])
		return byteArray, nil
	case 17:
		var byteArray [17]byte
		copy(byteArray[:], bytes[:17])
		return byteArray, nil
	case 18:
		var byteArray [18]byte
		copy(byteArray[:], bytes[:18])
		return byteArray, nil
	case 19:
		var byteArray [19]byte
		copy(byteArray[:], bytes[:19])
		return byteArray, nil
	case 20:
		var byteArray [20]byte
		copy(byteArray[:], bytes[:20])
		return byteArray, nil
	case 21:
		var byteArray [21]byte
		copy(byteArray[:], bytes[:21])
		return byteArray, nil
	case 22:
		var byteArray [22]byte
		copy(byteArray[:], bytes[:22])
		return byteArray, nil
	case 23:
		var byteArray [23]byte
		copy(byteArray[:], bytes[:23])
		return byteArray, nil
	case 24:
		var byteArray [24]byte
		copy(byteArray[:], bytes[:24])
		return byteArray, nil
	case 25:
		var byteArray [25]byte
		copy(byteArray[:], bytes[:25])
		return byteArray, nil
	case 26:
		var byteArray [26]byte
		copy(byteArray[:], bytes[:26])
		return byteArray, nil
	case 27:
		var byteArray [27]byte
		copy(byteArray[:], bytes[:27])
		return byteArray, nil
	case 28:
		var byteArray [28]byte
		copy(byteArray[:], bytes[:28])
		return byteArray, nil
	case 29:
		var byteArray [29]byte
		copy(byteArray[:], bytes[:29])
		return byteArray, nil
	case 30:
		var byteArray [30]byte
		copy(byteArray[:], bytes[:30])
		return byteArray, nil
	case 31:
		var byteArray [31]byte
		copy(byteArray[:], bytes[:31])
		return byteArray, nil
	case 32:
		var byteArray [32]byte
		copy(byteArray[:], bytes[:32])
		return byteArray, nil
	default:
		errStr := fmt.Sprintf("Error parsing fixed size byte array. Array of size %d incompatible", argType.Type.Size())
		return nil, errors.New(errStr)
	}
}

func ConvertToType(str string, typ *abi.Type) (interface{}, error) {
	fmt.Println(typ)
	fmt.Println("Type: ", typ.Type)
	fmt.Println("Type.Kind: ", typ.Type.Kind())
	fmt.Println("Type.Size: ", typ.Type.Size())
	fmt.Println("Kind: ", typ.Kind)
	fmt.Println("Size: ", typ.Size)

	switch typ.Kind {
	case reflect.String:
		return str, nil
	case reflect.Bool:
		b, err := ConvertToBool(str)
		return b, err
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return ConvertToInt(true, typ.Size, str)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return ConvertToInt(false, typ.Size, str)
	case reflect.Ptr:
		i, err := ConvertToInt(false, typ.Size, str)
		return i, err
	case reflect.Array:
		if typ.Type == reflect.TypeOf(common.Address{}) {
			return common.HexToAddress(str), nil
		} else {
			return nil, errors.New("Conversion failed. Item is array type, cannot parse")
		}
	default:
		errStr := fmt.Sprintf("Error, type not found: %s", typ.Kind)
		return nil, errors.New(errStr)
	}
}

func ConvertToInt(signed bool, size int, value string) (interface{}, error) {
	if size%8 > 0 {
		return nil, errors.New("integer is not a multiple of 8")
	} else if !isGoIntSize(size) {
		newInt := new(big.Int)
		newInt, ok := newInt.SetString(value, 10)
		if !ok {
			return nil, errors.New("could not convert string to big.int")
		}

		return newInt, nil
	} else {
		if signed {
			i, err := strconv.ParseInt(value, 10, size)
			if err != nil {
				return nil, err
			}

			switch size {
			case 8:
				return int8(i), nil
			case 16:
				return int16(i), nil
			case 32:
				return int32(i), nil
			case 64:
				return int64(i), nil
			}
		} else {
			u, err := strconv.ParseUint(value, 10, size)
			if err != nil {
				return nil, err
			}

			switch size {
			case 8:
				return uint8(u), nil
			case 16:
				return uint16(u), nil
			case 32:
				return uint32(u), nil
			case 64:
				return uint64(u), nil
			}
		}
	}

	return 0, errors.New("integer conversion fatal error")
}

func makeDynamicArray(bytes [][]byte, superArraySize int, subArraySize int) interface{} {
	return nil
}

func makeDynamicSubArray(bytes []byte, size int, index int) interface{} {
	switch size {
	case 1:
		subarray := [1]uint8{}
		for j := 0; j < size; j++ {
			subarray[j] = bytes[index+j]
		}
		return subarray
	case 2:
		subarray := [2]uint8{}
		for j := 0; j < size; j++ {
			subarray[j] = bytes[index+j]
		}
		return subarray
	case 3:
		subarray := [3]uint8{}
		for j := 0; j < size; j++ {
			subarray[j] = bytes[index+j]
		}
		return subarray
	case 4:
		subarray := [4]uint8{}
		for j := 0; j < size; j++ {
			subarray[j] = bytes[index+j]
		}
		return subarray
	case 5:
		subarray := [5]uint8{}
		for j := 0; j < size; j++ {
			subarray[j] = bytes[index+j]
		}
		return subarray
	case 6:
		subarray := [6]uint8{}
		for j := 0; j < size; j++ {
			subarray[j] = bytes[index+j]
		}
		return subarray
	case 7:
		subarray := [7]uint8{}
		for j := 0; j < size; j++ {
			subarray[j] = bytes[index+j]
		}
		return subarray
	case 8:
		subarray := [8]uint8{}
		for j := 0; j < size; j++ {
			subarray[j] = bytes[index+j]
		}
		return subarray
	case 9:
		subarray := [9]uint8{}
		for j := 0; j < size; j++ {
			subarray[j] = bytes[index+j]
		}
		return subarray
	case 10:
		subarray := [10]uint8{}
		for j := 0; j < size; j++ {
			subarray[j] = bytes[index+j]
		}
		return subarray
	case 11:
		subarray := [11]uint8{}
		for j := 0; j < size; j++ {
			subarray[j] = bytes[index+j]
		}
		return subarray
	}

	return []byte{}
}

// MUST CHECK RETURNED ERROR ELSE WILL RETURN FALSE FOR ANY ERRONEOUS INPUT
func ConvertToBool(value string) (bool, error) {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return b, nil
}

func isGoIntSize(size int) (isGoPrimitive bool) {
	switch size {
	case 8, 16, 32, 64:
		return true
	default:
		return false
	}
}
