package utils_test

import (
	contract "github.com/clearmatics/ion-cli/contracts"
	"github.com/clearmatics/ion-cli/utils"
	"gotest.tools/assert"
	"os"
	"testing"
)

// Type Test
// Testing type conversions to expected solidity types by retrieving expected types and converting string
// representations of them to the expected type.
// The converted output is passed to Pack() function to see if the conversion was done correctly

const TypeTestContractFile = "TypeTest.sol"

var compiledTestContract *contract.ContractInstance

func TestMain(m *testing.M) {
	compiled, err := compileTestContract()
	if err != nil {
		panic(err.Error())
	}

	compiledTestContract = compiled
	os.Exit(m.Run())
}

func Test_ConvertBool(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bool"].Inputs.NonIndexed()[0].Type

	input := "false"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bool", result)
	assert.NilError(t, err)

}

func Test_ConvertBoolArray(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bools"].Inputs.NonIndexed()[0].Type

	input := "false,false,true"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bools", result)
	assert.NilError(t, err)
}

func Test_ConvertBool2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bool2"].Inputs.NonIndexed()[0].Type

	input := "false,true"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bool2", result)
	assert.NilError(t, err)
}

func Test_ConvertBool4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bool4"].Inputs.NonIndexed()[0].Type

	input := "true,false,false,true"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bool4", result)
	assert.NilError(t, err)
}

func Test_ConvertInt8(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int8"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int8", result)
	assert.NilError(t, err)

	// Now test incorrect int type passed
	expectedType = compiledTestContract.Abi.Methods["Int64"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int8", result)
	assert.Assert(t, err != nil)

	// Now test incorrect int type passed again
	expectedType = compiledTestContract.Abi.Methods["Int32"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int8", result)
	assert.Assert(t, err != nil)
}

func Test_ConvertInt8Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int8s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int8s", result)
	assert.NilError(t, err)
}

func Test_ConvertInt8_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int8_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int8_2", result)
	assert.NilError(t, err)
}

func Test_ConvertInt8_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int8_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int8_4", result)
	assert.NilError(t, err)
}

func Test_ConvertBytes8(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes8"].Inputs.NonIndexed()[0].Type
	expectedResult := [8]uint8{52, 113, 85, 90, 185, 169, 149, 40}

	input := "0x3471555ab9a99528"

	result, err := utils.ApplySolidityType(input, expectedType)
	assert.Equal(t, result, expectedResult)
	assert.NilError(t, err)
}

func Test_ConvertBytes32(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes32"].Inputs.NonIndexed()[0].Type
	expectedResult := [32]uint8{52, 113, 85, 90, 185, 169, 149, 40, 240, 47, 156, 221, 143, 0, 23, 254, 47, 86, 224, 17, 22, 172, 196, 254, 127, 120, 174, 233, 0, 68, 47, 53}

	input := "0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35"

	result, err := utils.ApplySolidityType(input, expectedType)
	assert.Equal(t, result, expectedResult)
	assert.NilError(t, err)
}

//func Test_ConvertToTypeString(t *testing.T) {
//	expectedType := abi.Argument{
//		Name:    "string",
//		Type:    abi.Type{
//			Elem:          nil,
//			Kind:          17,
//			Type:          nil,
//			Size:          10,
//			T:             8,
//			TupleRawName:  "",
//			TupleElems:    nil,
//			TupleRawNames: nil,
//		},
//		Indexed: false,
//	}
//	expectedResult := [10]uint8{52,113,85,90,185,169,149,40,240,47}
//
//	input := "0x3471555ab9a99528f02f"
//
//	result, err := utils.ApplySolidityType(input, expectedType.Type)
//	assert.Equal(t, result, expectedResult)
//	assert.NilError(t, err)
//}
//
//func Test_ConvertToTypeBool(t *testing.T) {
//	expectedType := abi.Argument{
//		Name:    "bytes10",
//		Type:    abi.Type{
//			Elem:          nil,
//			Kind:          17,
//			Type:          nil,
//			Size:          10,
//			T:             8,
//			TupleRawName:  "",
//			TupleElems:    nil,
//			TupleRawNames: nil,
//		},
//		Indexed: false,
//	}
//	expectedResult := [10]uint8{52,113,85,90,185,169,149,40,240,47}
//
//	input := "0x3471555ab9a99528f02f"
//
//	result, err := utils.ApplySolidityType(input, expectedType.Type)
//	assert.Equal(t, result, expectedResult)
//	assert.NilError(t, err)
//}
//
//func Test_ConvertToTypeInt8(t *testing.T) {
//	expectedType := abi.Argument{
//		Name:    "bytes10",
//		Type:    abi.Type{
//			Elem:          nil,
//			Kind:          17,
//			Type:          nil,
//			Size:          10,
//			T:             8,
//			TupleRawName:  "",
//			TupleElems:    nil,
//			TupleRawNames: nil,
//		},
//		Indexed: false,
//	}
//	expectedResult := [10]uint8{52,113,85,90,185,169,149,40,240,47}
//
//	input := "0x3471555ab9a99528f02f"
//
//	result, err := utils.ApplySolidityType(input, expectedType.Type)
//	assert.Equal(t, result, expectedResult)
//	assert.NilError(t, err)
//}
//
//func Test_ConvertToTypeInt16(t *testing.T) {
//	expectedType := abi.Argument{
//		Name:    "bytes10",
//		Type:    abi.Type{
//			Elem:          nil,
//			Kind:          17,
//			Type:          nil,
//			Size:          10,
//			T:             8,
//			TupleRawName:  "",
//			TupleElems:    nil,
//			TupleRawNames: nil,
//		},
//		Indexed: false,
//	}
//	expectedResult := [10]uint8{52,113,85,90,185,169,149,40,240,47}
//
//	input := "0x3471555ab9a99528f02f"
//
//	result, err := utils.ApplySolidityType(input, expectedType.Type)
//	assert.Equal(t, result, expectedResult)
//	assert.NilError(t, err)
//}
//
//func Test_ConvertToTypeInt32(t *testing.T) {
//	expectedType := abi.Argument{
//		Name:    "bytes10",
//		Type:    abi.Type{
//			Elem:          nil,
//			Kind:          17,
//			Type:          nil,
//			Size:          10,
//			T:             8,
//			TupleRawName:  "",
//			TupleElems:    nil,
//			TupleRawNames: nil,
//		},
//		Indexed: false,
//	}
//	expectedResult := [10]uint8{52,113,85,90,185,169,149,40,240,47}
//
//	input := "0x3471555ab9a99528f02f"
//
//	result, err := utils.ApplySolidityType(input, expectedType.Type)
//	assert.Equal(t, result, expectedResult)
//	assert.NilError(t, err)
//}
//
//func Test_ConvertToTypeInt64(t *testing.T) {
//	expectedType := abi.Argument{
//		Name:    "bytes10",
//		Type:    abi.Type{
//			Elem:          nil,
//			Kind:          17,
//			Type:          nil,
//			Size:          10,
//			T:             8,
//			TupleRawName:  "",
//			TupleElems:    nil,
//			TupleRawNames: nil,
//		},
//		Indexed: false,
//	}
//	expectedResult := [10]uint8{52,113,85,90,185,169,149,40,240,47}
//
//	input := "0x3471555ab9a99528f02f"
//
//	result, err := utils.ApplySolidityType(input, expectedType.Type)
//	assert.Equal(t, result, expectedResult)
//	assert.NilError(t, err)
//}
//
//func Test_ConvertToTypeUint8(t *testing.T) {
//	expectedType := abi.Argument{
//		Name:    "bytes10",
//		Type:    abi.Type{
//			Elem:          nil,
//			Kind:          17,
//			Type:          nil,
//			Size:          10,
//			T:             8,
//			TupleRawName:  "",
//			TupleElems:    nil,
//			TupleRawNames: nil,
//		},
//		Indexed: false,
//	}
//	expectedResult := [10]uint8{52,113,85,90,185,169,149,40,240,47}
//
//	input := "0x3471555ab9a99528f02f"
//
//	result, err := utils.ApplySolidityType(input, expectedType.Type)
//	assert.Equal(t, result, expectedResult)
//	assert.NilError(t, err)
//}
//
//func Test_ConvertToTypeUint16(t *testing.T) {
//	expectedType := abi.Argument{
//		Name:    "bytes10",
//		Type:    abi.Type{
//			Elem:          nil,
//			Kind:          17,
//			Type:          nil,
//			Size:          10,
//			T:             8,
//			TupleRawName:  "",
//			TupleElems:    nil,
//			TupleRawNames: nil,
//		},
//		Indexed: false,
//	}
//	expectedResult := [10]uint8{52,113,85,90,185,169,149,40,240,47}
//
//	input := "0x3471555ab9a99528f02f"
//
//	result, err := utils.ApplySolidityType(input, expectedType.Type)
//	assert.Equal(t, result, expectedResult)
//	assert.NilError(t, err)
//}
//
//func Test_ConvertToTypeUint32(t *testing.T) {
//	expectedType := abi.Argument{
//		Name:    "bytes10",
//		Type:    abi.Type{
//			Elem:          nil,
//			Kind:          17,
//			Type:          nil,
//			Size:          10,
//			T:             8,
//			TupleRawName:  "",
//			TupleElems:    nil,
//			TupleRawNames: nil,
//		},
//		Indexed: false,
//	}
//	expectedResult := [10]uint8{52,113,85,90,185,169,149,40,240,47}
//
//	input := "0x3471555ab9a99528f02f"
//
//	result, err := utils.ApplySolidityType(input, expectedType.Type)
//	assert.Equal(t, result, expectedResult)
//	assert.NilError(t, err)
//}
//
//func Test_ConvertToTypeUint64(t *testing.T) {
//	expectedType := abi.Argument{
//		Name:    "bytes10",
//		Type:    abi.Type{
//			Elem:          nil,
//			Kind:          17,
//			Type:          nil,
//			Size:          10,
//			T:             8,
//			TupleRawName:  "",
//			TupleElems:    nil,
//			TupleRawNames: nil,
//		},
//		Indexed: false,
//	}
//	expectedResult := [10]uint8{52,113,85,90,185,169,149,40,240,47}
//
//	input := "0x3471555ab9a99528f02f"
//
//	result, err := utils.ApplySolidityType(input, expectedType.Type)
//	assert.Equal(t, result, expectedResult)
//	assert.NilError(t, err)
//}
//
//func Test_ConvertToTypeInt(t *testing.T) {
//	expectedType := abi.Argument{
//		Name:    "bytes10",
//		Type:    abi.Type{
//			Elem:          nil,
//			Kind:          17,
//			Type:          nil,
//			Size:          10,
//			T:             8,
//			TupleRawName:  "",
//			TupleElems:    nil,
//			TupleRawNames: nil,
//		},
//		Indexed: false,
//	}
//	expectedResult := [10]uint8{52,113,85,90,185,169,149,40,240,47}
//
//	input := "0x3471555ab9a99528f02f"
//
//	result, err := utils.ApplySolidityType(input, expectedType.Type)
//	assert.Equal(t, result, expectedResult)
//	assert.NilError(t, err)
//}
//
//func Test_ConvertToTypeAddress(t *testing.T) {
//	expectedType := abi.Argument{
//		Name:    "bytes10",
//		Type:    abi.Type{
//			Elem:          nil,
//			Kind:          17,
//			Type:          nil,
//			Size:          10,
//			T:             8,
//			TupleRawName:  "",
//			TupleElems:    nil,
//			TupleRawNames: nil,
//		},
//		Indexed: false,
//	}
//	expectedResult := [10]uint8{52,113,85,90,185,169,149,40,240,47}
//
//	input := "0x3471555ab9a99528f02f"
//
//	result, err := utils.ApplySolidityType(input, expectedType.Type)
//	assert.Equal(t, result, expectedResult)
//	assert.NilError(t, err)
//}

func compileTestContract() (*contract.ContractInstance, error) {
	solcVersion, err := utils.GetSolidityContractVersion(TypeTestContractFile)
	if err != nil {
		return nil, err
	}
	solc, err := utils.GetSolidityCompilerVersion(solcVersion)
	if err != nil {
		return nil, err
	}

	compiledContract, err := utils.CreateContractInstance(TypeTestContractFile, solc.Name())
	if err != nil {
		return nil, err
	}

	return compiledContract, nil
}
