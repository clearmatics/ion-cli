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

/*
========================================================================================================================
	BOOLEAN TYPE TESTS
========================================================================================================================
*/

func Test_ConvertBool(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bool"].Inputs.NonIndexed()[0].Type

	input := "false"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bool", result)
	assert.NilError(t, err)

}

func Test_ConvertBoolArray(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bools"].Inputs.NonIndexed()[0].Type

	input := "false,false,true"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bools", result)
	assert.NilError(t, err)
}

func Test_ConvertBool2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bool2"].Inputs.NonIndexed()[0].Type

	input := "false,true"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bool2", result)
	assert.NilError(t, err)
}

func Test_ConvertBool4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bool4"].Inputs.NonIndexed()[0].Type

	input := "true,false,false,true"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bool4", result)
	assert.NilError(t, err)
}

func Test_ConvertBool222Array(t *testing.T) {
	// All functions have a single non indexed input arg
	solidityType := compiledTestContract.Abi.Methods["Bool222"].Inputs.NonIndexed()[0].Type

	input := "[[[true,true], [true,true]], [[true, false], [true, false]]]"

	result, err := utils.SolidityToStaticGoType(input, solidityType)
	assert.NilError(t, err)

	result = result

	_, err = compiledTestContract.Abi.Pack("Bool222", result)
	assert.NilError(t, err)
}

func Test_ConvertBool2222Array(t *testing.T) {
	// All functions have a single non indexed input arg
	solidityType := compiledTestContract.Abi.Methods["Bool2222"].Inputs.NonIndexed()[0].Type

	input := "[[[[true,true], [true,true]], [[true, false], [true, false]]], [[[true,true], [true,true]], [[true, false], [true, false]]]]"

	result, err := utils.SolidityToStaticGoType(input, solidityType)
	assert.NilError(t, err)

	result = result

	_, err = compiledTestContract.Abi.Pack("Bool2222", result)
	assert.NilError(t, err)
}

/*
========================================================================================================================
	INT TYPE TESTS
========================================================================================================================
*/

func Test_ConvertInt8(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int8"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int8", result)
	assert.NilError(t, err)

	// Now test incorrect int type passed
	expectedType = compiledTestContract.Abi.Methods["Int64"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int8", result)
	assert.ErrorContains(t, err, "abi: cannot use int64 as type int8 as argument")

	// Now test incorrect int type passed again
	expectedType = compiledTestContract.Abi.Methods["Int32"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int8", result)
	assert.ErrorContains(t, err, "abi: cannot use int32 as type int8 as argument")
}

func Test_ConvertInt8Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int8s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int8s", result)
	assert.NilError(t, err)
}

func Test_ConvertInt8_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int8_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int8_2", result)
	assert.NilError(t, err)
}

func Test_ConvertInt8_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int8_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int8_4", result)
	assert.NilError(t, err)
}

func Test_ConvertInt16(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int16"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int16", result)
	assert.NilError(t, err)

	// Now test incorrect int type passed
	expectedType = compiledTestContract.Abi.Methods["Int64"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int16", result)
	assert.ErrorContains(t, err, "abi: cannot use int64 as type int16 as argument")

	// Now test incorrect int type passed again
	expectedType = compiledTestContract.Abi.Methods["Int32"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int16", result)
	assert.ErrorContains(t, err, "abi: cannot use int32 as type int16 as argument")
}

func Test_ConvertInt16Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int16s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int16s", result)
	assert.NilError(t, err)
}

func Test_ConvertInt16_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int16_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int16_2", result)
	assert.NilError(t, err)
}

func Test_ConvertInt16_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int16_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int16_4", result)
	assert.NilError(t, err)
}

func Test_ConvertInt32(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int32"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int32", result)
	assert.NilError(t, err)

	// Now test incorrect int type passed
	expectedType = compiledTestContract.Abi.Methods["Int64"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int32", result)
	assert.ErrorContains(t, err, "abi: cannot use int64 as type int32 as argument")

	// Now test incorrect int type passed again
	expectedType = compiledTestContract.Abi.Methods["Int16"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int32", result)
	assert.ErrorContains(t, err, "abi: cannot use int16 as type int32 as argument")
}

func Test_ConvertInt32Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int32s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int32s", result)
	assert.NilError(t, err)
}

func Test_ConvertInt32_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int32_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int32_2", result)
	assert.NilError(t, err)
}

func Test_ConvertInt32_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int32_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int32_4", result)
	assert.NilError(t, err)
}

func Test_ConvertInt64(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int64"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int64", result)
	assert.NilError(t, err)

	// Now test incorrect int type passed
	expectedType = compiledTestContract.Abi.Methods["Int16"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int64", result)
	assert.ErrorContains(t, err, "abi: cannot use int16 as type int64 as argument")

	// Now test incorrect int type passed again
	expectedType = compiledTestContract.Abi.Methods["Int32"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int64", result)
	assert.ErrorContains(t, err, "abi: cannot use int32 as type int64 as argument")
}

func Test_ConvertInt64Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int64s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int64s", result)
	assert.NilError(t, err)
}

func Test_ConvertInt64_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int64_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int64_2", result)
	assert.NilError(t, err)
}

func Test_ConvertInt64_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int64_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int64_4", result)
	assert.NilError(t, err)
}

func Test_ConvertInt64_22Array(t *testing.T) {
	// All functions have a single non indexed input arg
	solidityType := compiledTestContract.Abi.Methods["Int64_22"].Inputs.NonIndexed()[0].Type

	input := "[[5,6],[7,8]]"
	result, err := utils.SolidityToStaticGoType(input, solidityType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int64_22", result)
	assert.NilError(t, err)
}

func Test_ConvertInt128(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int128"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int128", result)
	assert.NilError(t, err)

	// Now test incorrect int type passed
	expectedType = compiledTestContract.Abi.Methods["Int16"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int128", result)
	assert.ErrorContains(t, err, "abi: cannot use int16 as type ptr as argument")

	// Now test incorrect int type passed again
	expectedType = compiledTestContract.Abi.Methods["Int32"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int128", result)
	assert.ErrorContains(t, err, "abi: cannot use int32 as type ptr as argument")
}

func Test_ConvertInt128Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int128s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int128s", result)
	assert.NilError(t, err)
}

func Test_ConvertInt128_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int128_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int128_2", result)
	assert.NilError(t, err)
}

func Test_ConvertInt128_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int128_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int128_4", result)
	assert.NilError(t, err)
}

func Test_ConvertInt256(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int256"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int256", result)
	assert.NilError(t, err)

	// Now test incorrect int type passed
	expectedType = compiledTestContract.Abi.Methods["Int16"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int256", result)
	assert.ErrorContains(t, err, "abi: cannot use int16 as type ptr as argument")

	// Now test incorrect int type passed again
	expectedType = compiledTestContract.Abi.Methods["Int32"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int256", result)
	assert.ErrorContains(t, err, "abi: cannot use int32 as type ptr as argument")
}

func Test_ConvertInt256Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int256s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int256s", result)
	assert.NilError(t, err)
}

func Test_ConvertInt256_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int256_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int256_2", result)
	assert.NilError(t, err)
}

func Test_ConvertInt256_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int256_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int256_4", result)
	assert.NilError(t, err)
}

/*
========================================================================================================================
	UINT TYPE TESTS
========================================================================================================================
*/

func Test_ConvertUint8(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint8"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint8", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint64"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint8", result)
	assert.ErrorContains(t, err, "abi: cannot use uint64 as type uint8 as argument")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["Uint32"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint8", result)
	assert.ErrorContains(t, err, "abi: cannot use uint32 as type uint8 as argument")
}

func Test_ConvertUint8Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint8s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint8s", result)
	assert.NilError(t, err)
}

func Test_ConvertUint8_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint8_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint8_2", result)
	assert.NilError(t, err)
}

func Test_ConvertUint8_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint8_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint8_4", result)
	assert.NilError(t, err)
}

func Test_ConvertUint16(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint16"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint16", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint64"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint16", result)
	assert.ErrorContains(t, err, "abi: cannot use uint64 as type uint16 as argument")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["Uint32"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint16", result)
	assert.ErrorContains(t, err, "abi: cannot use uint32 as type uint16 as argument")
}

func Test_ConvertUint16Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint16s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint16s", result)
	assert.NilError(t, err)
}

func Test_ConvertUint16_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint16_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint16_2", result)
	assert.NilError(t, err)
}

func Test_ConvertUint16_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint16_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint16_4", result)
	assert.NilError(t, err)
}

func Test_ConvertUint32(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint32"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint32", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint64"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint32", result)
	assert.ErrorContains(t, err, "abi: cannot use uint64 as type uint32 as argument")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["Uint16"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint32", result)
	assert.ErrorContains(t, err, "abi: cannot use uint16 as type uint32 as argument")
}

func Test_ConvertUint32Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint32s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint32s", result)
	assert.NilError(t, err)
}

func Test_ConvertUint32_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint32_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint32_2", result)
	assert.NilError(t, err)
}

func Test_ConvertUint32_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint32_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint32_4", result)
	assert.NilError(t, err)
}

func Test_ConvertUint64(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint64"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint64", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint16"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint64", result)
	assert.ErrorContains(t, err, "abi: cannot use uint16 as type uint64 as argument")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["Uint32"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint64", result)
	assert.ErrorContains(t, err, "abi: cannot use uint32 as type uint64 as argument")
}

func Test_ConvertUint64Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint64s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint64s", result)
	assert.NilError(t, err)
}

func Test_ConvertUint64_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint64_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint64_2", result)
	assert.NilError(t, err)
}

func Test_ConvertUint64_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint64_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint64_4", result)
	assert.NilError(t, err)
}

func Test_ConvertUint128(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint128"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint128", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint16"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint128", result)
	assert.ErrorContains(t, err, "abi: cannot use uint16 as type ptr as argument")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["Uint32"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint128", result)
	assert.ErrorContains(t, err, "abi: cannot use uint32 as type ptr as argument")
}

func Test_ConvertUint128Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint128s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint128s", result)
	assert.NilError(t, err)
}

func Test_ConvertUint128_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint128_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint128_2", result)
	assert.NilError(t, err)
}

func Test_ConvertUint128_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint128_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint128_4", result)
	assert.NilError(t, err)
}

func Test_ConvertUint256(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint256"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint256", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint16"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint256", result)
	assert.ErrorContains(t, err, "abi: cannot use uint16 as type ptr as argument")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["Uint32"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint256", result)
	assert.ErrorContains(t, err, "abi: cannot use uint32 as type ptr as argument")
}

func Test_ConvertUint256Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint256s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint256s", result)
	assert.NilError(t, err)
}

func Test_ConvertUint256_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint256_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint256_2", result)
	assert.NilError(t, err)
}

func Test_ConvertUint256_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint256_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.SolidityToStaticGoType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint256_4", result)
	assert.NilError(t, err)
}

func Test_ConvertUint256_22Array(t *testing.T) {
	// All functions have a single non indexed input arg
	solidityType := compiledTestContract.Abi.Methods["Uint256_22"].Inputs.NonIndexed()[0].Type

	input := "[[5,6],[7,8]]"
	result, err := utils.SolidityToStaticGoType(input, solidityType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint256_22", result)
	assert.NilError(t, err)
}

/*
========================================================================================================================
	ADDRESS TYPE TESTS
========================================================================================================================
*/

func Test_ConvertAddress(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Address"].Inputs.NonIndexed()[0].Type

	input := "0xb8844cf76df596e746f360957aa3af954ef51605"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Address", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint16"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "expression:1:1-1:50: cannot convert 1053404704927428982920127934424083581711615858181 (type untyped constant) to type uint16")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["String"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Address", result)
	assert.ErrorContains(t, err, "abi: cannot use string as type array as argument")
}

func Test_ConvertAddressArray(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Addresses"].Inputs.NonIndexed()[0].Type

	input := "0xb8844cf76df596e746f360957aa3af954ef51605,0xb8844cf76df596e746f360957aa3af954ef51605,0xb8844cf76df596e746f360957aa3af954ef51605"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Addresses", result)
	assert.NilError(t, err)
}

func Test_ConvertAddress_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Address2"].Inputs.NonIndexed()[0].Type

	input := "0xb8844cf76df596e746f360957aa3af954ef51605,0xb8844cf76df596e746f360957aa3af954ef51605"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Address2", result)
	assert.NilError(t, err)
}

func Test_ConvertAddress_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Address4"].Inputs.NonIndexed()[0].Type

	input := "0xb8844cf76df596e746f360957aa3af954ef51605,0xb8844cf76df596e746f360957aa3af954ef51605,0xb8844cf76df596e746f360957aa3af954ef51605,0xb8844cf76df596e746f360957aa3af954ef51605"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Address4", result)
	assert.NilError(t, err)
}

func Test_ConvertAddress_22Array(t *testing.T) {
	// All functions have a single non indexed input arg
	solidityType := compiledTestContract.Abi.Methods["Address22"].Inputs.NonIndexed()[0].Type

	input := "[[0xb8844cf76df596e746f360957aa3af954ef51605,0xb8844cf76df596e746f360957aa3af954ef51605],[0xb8844cf76df596e746f360957aa3af954ef51605,0xb8844cf76df596e746f360957aa3af954ef51605]]"

	result, err := utils.SolidityToStaticGoType(input, solidityType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Address22", result)
	assert.NilError(t, err)
}

/*
========================================================================================================================
	STRING TYPE TESTS
========================================================================================================================
*/

func Test_ConvertString(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["String"].Inputs.NonIndexed()[0].Type

	input := "somerandomstring123"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("String", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint16"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "expression:1:8-1:26: undefined: somerandomstring123")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["Address"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("String", result)
	assert.ErrorContains(t, err, "abi: cannot use array as type string as argument")
}

func Test_ConvertStringArray(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Strings"].Inputs.NonIndexed()[0].Type

	input := "somerandomstring123,somerandomstring456,somerandomstring789"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Strings", result)
	assert.NilError(t, err)
}

func Test_ConvertString_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["String2"].Inputs.NonIndexed()[0].Type

	input := "somerandomstring123,somerandomstring456"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("String2", result)
	assert.NilError(t, err)
}

func Test_ConvertString_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["String4"].Inputs.NonIndexed()[0].Type

	input := "somerandomstring123,somerandomstring456,somerandomstring789,somerandomstring101112"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("String4", result)
	assert.NilError(t, err)
}

func Test_ConvertString_22Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["String22"].Inputs.NonIndexed()[0].Type

	input := "[[somerandomstring123,somerandomstring456],[somerandomstring789,somerandomstring101112]]"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("String22", result)
	assert.NilError(t, err)
}

/*
========================================================================================================================
	BYTESN TYPE TESTS
========================================================================================================================
*/

func Test_ConvertBytesArray(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes"].Inputs.NonIndexed()[0].Type

	input := "0x3471555ab9a9952172378abababa"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes", result)
	assert.NilError(t, err)

	input = "ajshduihuieh"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: encoding/hex: invalid byte: U+006A 'j'")

	input = "[34,44]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	input = "1234567891"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	expectedType = compiledTestContract.Abi.Methods["Int256"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes", result)
	assert.ErrorContains(t, err, "abi: cannot use ptr as type slice as argument")
}

func Test_ConvertBytes8(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes8"].Inputs.NonIndexed()[0].Type

	input := "0x3471555ab9a99528"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes8", result)
	assert.NilError(t, err)
}

func Test_ConvertBytes32(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes32"].Inputs.NonIndexed()[0].Type

	input := "0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes32", result)
	assert.NilError(t, err)
}

func Test_ConvertBytes32_2(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes32_2"].Inputs.NonIndexed()[0].Type

	input := "[0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35]"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes32_2", result)
	assert.NilError(t, err)
}

/*
========================================================================================================================
	BYTE[] TYPE TESTS
========================================================================================================================
*/

func Test_ConvertByte1(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Byte"].Inputs.NonIndexed()[0].Type

	input := "0x34"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byte", result)
	assert.NilError(t, err)

	input = "ajshduihuieh"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: encoding/hex: invalid byte: U+006A 'j'")

	input = "344"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: incorrect byte array length error: expected length 1 but input array has length 2")

	input = "[34,44]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: input length mismatch: input has 2 values but expected type has 1 items")

	input = "3444"
	expectedType = compiledTestContract.Abi.Methods["Int256"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byte", result)
	assert.ErrorContains(t, err, "abi: cannot use ptr as type array as argument")

	input = "somestring"
	expectedType = compiledTestContract.Abi.Methods["String"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byte", result)
	assert.ErrorContains(t, err, "abi: cannot use string as type array as argument")
}

func Test_ConvertByteArray(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Byten"].Inputs.NonIndexed()[0].Type

	input := "[34,71,55,5a,b9,a9,95,21,72,37,8a,ba,ba,ba]"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byten", result)
	assert.NilError(t, err)

	input = "0x3471555ab9a9952172378abababa"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: type mismatch error: end of input reached but type has nested elements")

	input = "[ss,ss,ss]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: encoding/hex: invalid byte: U+0073 's'")

	input = "[aaa,aaa,aaa]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: incorrect byte array length")

	input = "1234567891"
	expectedType = compiledTestContract.Abi.Methods["Int256"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byten", result)
	assert.ErrorContains(t, err, "abi: cannot use *big.Int as type [0]slice as argument")

	input = "somestring"
	expectedType = compiledTestContract.Abi.Methods["String"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byten", result)
	assert.ErrorContains(t, err, "abi: cannot use string as type [0]slice as argument")

	input = "0x3471555ab9a9952172378a"
	expectedType = compiledTestContract.Abi.Methods["Bytes"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byten", result)
	assert.ErrorContains(t, err, "abi: cannot use []uint8 as type [0]array as argument")

	expectedType = compiledTestContract.Abi.Methods["Byten"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: type mismatch error: end of input reached but type has nested elements")

	input = "[34,71,55,5a,b9,a9,95,21,72,37,8a]"
	expectedType = compiledTestContract.Abi.Methods["Byten"].Inputs.NonIndexed()[0].Type
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byten", result)
	assert.NilError(t, err)
}

func Test_ConvertByte2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Byte2"].Inputs.NonIndexed()[0].Type

	input := "[34,71]"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byte2", result)
	assert.NilError(t, err)

	input = "[34,71,11]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: input length mismatch: input has 3 values but expected type has 2 items")

	_, err = compiledTestContract.Abi.Pack("Byte2", [3]byte{34, 71, 11})
	assert.ErrorContains(t, err, "abi: cannot use [3]uint8 as type [2]array as argument")
}

func Test_ConvertByte8Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Byte8"].Inputs.NonIndexed()[0].Type

	input := "[34,71,11,27,12,81,64,61]"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byte8", result)
	assert.NilError(t, err)

	input = "[34,71,11,27,12]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: input length mismatch: input has 5 values but expected type has 8 items")
}

func Test_ConvertByte64Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Byte64"].Inputs.NonIndexed()[0].Type

	input := "[34,71,11,27,12,87,23,64,87,32,68,74,62,38,74,68,72,36,52,34,53,28,66,58,47,23,68,73,46,87,23,65,34,71,11,27,12,87,23,64,87,32,68,74,62,38,74,68,72,36,52,34,53,28,66,58,47,23,68,73,46,87,23,65]"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byte64", result)
	assert.NilError(t, err)

	input = "[34,71,11,27,12,87,23,64,87,32,68,74,62,38,74,68,72,36,52,34,53,28,66,58,47,23,68,73,46,87,23,65]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: input length mismatch: input has 32 values but expected type has 64 items")
}

/*
========================================================================================================================
	BYTESN[] TYPE TESTS
========================================================================================================================
*/

func Test_ConvertBytes1Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes1n"].Inputs.NonIndexed()[0].Type

	input := "[[34],[71],[11]]"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes1n", result)
	assert.NilError(t, err)

	input = "[[34,71],[71],[11]]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: input length mismatch: input has 2 values but expected type has 1 items")
}

func Test_ConvertBytes2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes2n"].Inputs.NonIndexed()[0].Type

	input := "[3471,1127,1287,2365,2345,3286,6584,7236,8734,6872,3669]"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes2n", result)
	assert.NilError(t, err)

	input = "[347111,2712,8723,6523,4532,8665,8472,3687,3468,7236]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: incorrect byte array length error: expected length 2 but input array has length 3")
}

/*
========================================================================================================================
	BYTESN[] TYPE TESTS
========================================================================================================================
*/

func Test_ConvertBytes1_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes1_2"].Inputs.NonIndexed()[0].Type

	input := "[[34],[71]]"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes1_2", result)
	assert.NilError(t, err)

	input = "[34,71]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes1_2", result)
	assert.NilError(t, err)

	input = "[[34],[71],[11]]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: input length mismatch: input has 3 values but expected type has 2 items")

	input = "[[3471],[11]]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: incorrect byte array length error: expected length 1 but input array has length 2")
}

func Test_ConvertBytes2_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes2_2"].Inputs.NonIndexed()[0].Type

	input := "[3471,1127]"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes2_2", result)
	assert.NilError(t, err)

	input = "[3471,1127,1287]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: input length mismatch: input has 3 values but expected type has 2 items")

	input = "[347111,2712]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: incorrect byte array length error: expected length 2 but input array has length 3")
}

func Test_ConvertBytes32_4nArray(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes32_4n"].Inputs.NonIndexed()[0].Type

	input := "[[0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35],[0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35],[0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35]]"
	result, err := utils.SolidityToStaticGoType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes32_4n", result)
	assert.NilError(t, err)

	input = "[[0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35],[0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35]]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: input length mismatch: input has 5 values but expected type has 4 items")

	input = "[[0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f3536,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35],[0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35,0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35]]"
	result, err = utils.SolidityToStaticGoType(input, expectedType)
	assert.ErrorContains(t, err, "constructInitialiser: incorrect byte array length error: expected length 32 but input array has length 33")
}

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
