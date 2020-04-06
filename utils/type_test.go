package utils_test

import (
	contract "github.com/clearmatics/ion-cli/contracts"
)

// Type Test
// Testing type conversions to expected solidity types by retrieving expected types and converting string
// representations of them to the expected type.
// The converted output is passed to Pack() function to see if the conversion was done correctly

const TypeTestContractFile = "TypeTest.sol"

var compiledTestContract *contract.ContractInstance
/*
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

/*
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

/*
========================================================================================================================
	INT TYPE TESTS
========================================================================================================================
*/

/*
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
	assert.ErrorContains(t, err, "abi: cannot use int64 as type int8 as argument")

	// Now test incorrect int type passed again
	expectedType = compiledTestContract.Abi.Methods["Int32"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int8", result)
	assert.ErrorContains(t, err, "abi: cannot use int32 as type int8 as argument")
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

func Test_ConvertInt16(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int16"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int16", result)
	assert.NilError(t, err)

	// Now test incorrect int type passed
	expectedType = compiledTestContract.Abi.Methods["Int64"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int16", result)
	assert.ErrorContains(t, err, "abi: cannot use int64 as type int16 as argument")

	// Now test incorrect int type passed again
	expectedType = compiledTestContract.Abi.Methods["Int32"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int16", result)
	assert.ErrorContains(t, err, "abi: cannot use int32 as type int16 as argument")
}

func Test_ConvertInt16Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int16s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int16s", result)
	assert.NilError(t, err)
}

func Test_ConvertInt16_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int16_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int16_2", result)
	assert.NilError(t, err)
}

func Test_ConvertInt16_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int16_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int16_4", result)
	assert.NilError(t, err)
}

func Test_ConvertInt32(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int32"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int32", result)
	assert.NilError(t, err)

	// Now test incorrect int type passed
	expectedType = compiledTestContract.Abi.Methods["Int64"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int32", result)
	assert.ErrorContains(t, err, "abi: cannot use int64 as type int32 as argument")

	// Now test incorrect int type passed again
	expectedType = compiledTestContract.Abi.Methods["Int16"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int32", result)
	assert.ErrorContains(t, err, "abi: cannot use int16 as type int32 as argument")
}

func Test_ConvertInt32Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int32s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int32s", result)
	assert.NilError(t, err)
}

func Test_ConvertInt32_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int32_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int32_2", result)
	assert.NilError(t, err)
}

func Test_ConvertInt32_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int32_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int32_4", result)
	assert.NilError(t, err)
}

func Test_ConvertInt64(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int64"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int64", result)
	assert.NilError(t, err)

	// Now test incorrect int type passed
	expectedType = compiledTestContract.Abi.Methods["Int16"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int64", result)
	assert.ErrorContains(t, err, "abi: cannot use int16 as type int64 as argument")

	// Now test incorrect int type passed again
	expectedType = compiledTestContract.Abi.Methods["Int32"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int64", result)
	assert.ErrorContains(t, err, "abi: cannot use int32 as type int64 as argument")
}

func Test_ConvertInt64Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int64s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int64s", result)
	assert.NilError(t, err)
}

func Test_ConvertInt64_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int64_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int64_2", result)
	assert.NilError(t, err)
}

func Test_ConvertInt64_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int64_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int64_4", result)
	assert.NilError(t, err)
}

func Test_ConvertInt128(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int128"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int128", result)
	assert.NilError(t, err)

	// Now test incorrect int type passed
	expectedType = compiledTestContract.Abi.Methods["Int16"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int128", result)
	assert.ErrorContains(t, err, "abi: cannot use int16 as type ptr as argument")

	// Now test incorrect int type passed again
	expectedType = compiledTestContract.Abi.Methods["Int32"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int128", result)
	assert.ErrorContains(t, err, "abi: cannot use int32 as type ptr as argument")
}

func Test_ConvertInt128Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int128s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int128s", result)
	assert.NilError(t, err)
}

func Test_ConvertInt128_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int128_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int128_2", result)
	assert.NilError(t, err)
}

func Test_ConvertInt128_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int128_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int128_4", result)
	assert.NilError(t, err)
}

func Test_ConvertInt256(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int256"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int256", result)
	assert.NilError(t, err)

	// Now test incorrect int type passed
	expectedType = compiledTestContract.Abi.Methods["Int16"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int256", result)
	assert.ErrorContains(t, err, "abi: cannot use int16 as type ptr as argument")

	// Now test incorrect int type passed again
	expectedType = compiledTestContract.Abi.Methods["Int32"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Int256", result)
	assert.ErrorContains(t, err, "abi: cannot use int32 as type ptr as argument")
}

func Test_ConvertInt256Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int256s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int256s", result)
	assert.NilError(t, err)
}

func Test_ConvertInt256_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int256_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int256_2", result)
	assert.NilError(t, err)
}

func Test_ConvertInt256_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Int256_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Int256_4", result)
	assert.NilError(t, err)
}

/*
========================================================================================================================
	UINT TYPE TESTS
========================================================================================================================
*/

/*
func Test_ConvertUint8(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint8"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint8", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint64"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint8", result)
	assert.ErrorContains(t, err, "abi: cannot use uint64 as type uint8 as argument")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["Uint32"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint8", result)
	assert.ErrorContains(t, err, "abi: cannot use uint32 as type uint8 as argument")
}

func Test_ConvertUint8Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint8s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint8s", result)
	assert.NilError(t, err)
}

func Test_ConvertUint8_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint8_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint8_2", result)
	assert.NilError(t, err)
}

func Test_ConvertUint8_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint8_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint8_4", result)
	assert.NilError(t, err)
}

func Test_ConvertUint16(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint16"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint16", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint64"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint16", result)
	assert.ErrorContains(t, err, "abi: cannot use uint64 as type uint16 as argument")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["Uint32"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint16", result)
	assert.ErrorContains(t, err, "abi: cannot use uint32 as type uint16 as argument")
}

func Test_ConvertUint16Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint16s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint16s", result)
	assert.NilError(t, err)
}

func Test_ConvertUint16_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint16_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint16_2", result)
	assert.NilError(t, err)
}

func Test_ConvertUint16_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint16_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint16_4", result)
	assert.NilError(t, err)
}

func Test_ConvertUint32(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint32"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint32", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint64"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint32", result)
	assert.ErrorContains(t, err, "abi: cannot use uint64 as type uint32 as argument")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["Uint16"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint32", result)
	assert.ErrorContains(t, err, "abi: cannot use uint16 as type uint32 as argument")
}

func Test_ConvertUint32Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint32s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint32s", result)
	assert.NilError(t, err)
}

func Test_ConvertUint32_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint32_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint32_2", result)
	assert.NilError(t, err)
}

func Test_ConvertUint32_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint32_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint32_4", result)
	assert.NilError(t, err)
}

func Test_ConvertUint64(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint64"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint64", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint16"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint64", result)
	assert.ErrorContains(t, err, "abi: cannot use uint16 as type uint64 as argument")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["Uint32"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint64", result)
	assert.ErrorContains(t, err, "abi: cannot use uint32 as type uint64 as argument")
}

func Test_ConvertUint64Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint64s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint64s", result)
	assert.NilError(t, err)
}

func Test_ConvertUint64_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint64_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint64_2", result)
	assert.NilError(t, err)
}

func Test_ConvertUint64_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint64_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint64_4", result)
	assert.NilError(t, err)
}

func Test_ConvertUint128(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint128"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint128", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint16"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint128", result)
	assert.ErrorContains(t, err, "abi: cannot use uint16 as type ptr as argument")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["Uint32"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint128", result)
	assert.ErrorContains(t, err, "abi: cannot use uint32 as type ptr as argument")
}

func Test_ConvertUint128Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint128s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint128s", result)
	assert.NilError(t, err)
}

func Test_ConvertUint128_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint128_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint128_2", result)
	assert.NilError(t, err)
}

func Test_ConvertUint128_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint128_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint128_4", result)
	assert.NilError(t, err)
}

func Test_ConvertUint256(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint256"].Inputs.NonIndexed()[0].Type

	input := "5"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint256", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint16"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint256", result)
	assert.ErrorContains(t, err, "abi: cannot use uint16 as type ptr as argument")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["Uint32"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Uint256", result)
	assert.ErrorContains(t, err, "abi: cannot use uint32 as type ptr as argument")
}

func Test_ConvertUint256Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint256s"].Inputs.NonIndexed()[0].Type

	input := "5,6,7"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint256s", result)
	assert.NilError(t, err)
}

func Test_ConvertUint256_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint256_2"].Inputs.NonIndexed()[0].Type

	input := "5,6"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint256_2", result)
	assert.NilError(t, err)
}

func Test_ConvertUint256_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Uint256_4"].Inputs.NonIndexed()[0].Type

	input := "5,6,7,8"

	result, err := utils.ApplySolidityType(input, expectedType)
	_, err = compiledTestContract.Abi.Pack("Uint256_4", result)
	assert.NilError(t, err)
}

/*
========================================================================================================================
	ADDRESS TYPE TESTS
========================================================================================================================
*/

/*
func Test_ConvertAddress(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Address"].Inputs.NonIndexed()[0].Type

	input := "0xb8844cf76df596e746f360957aa3af954ef51605"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Address", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint16"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.ErrorContains(t, err, "strconv.ParseUint: parsing \"0xb8844cf76df596e746f360957aa3af954ef51605\": invalid syntax")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["String"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Address", result)
	assert.ErrorContains(t, err, "abi: cannot use string as type array as argument")
}

func Test_ConvertAddressArray(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Addresses"].Inputs.NonIndexed()[0].Type

	input := "0xb8844cf76df596e746f360957aa3af954ef51605,0xb8844cf76df596e746f360957aa3af954ef51605,0xb8844cf76df596e746f360957aa3af954ef51605"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Addresses", result)
	assert.NilError(t, err)
}

func Test_ConvertAddress_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Address2"].Inputs.NonIndexed()[0].Type

	input := "0xb8844cf76df596e746f360957aa3af954ef51605,0xb8844cf76df596e746f360957aa3af954ef51605"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Address2", result)
	assert.NilError(t, err)
}

func Test_ConvertAddress_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Address4"].Inputs.NonIndexed()[0].Type

	input := "0xb8844cf76df596e746f360957aa3af954ef51605,0xb8844cf76df596e746f360957aa3af954ef51605,0xb8844cf76df596e746f360957aa3af954ef51605,0xb8844cf76df596e746f360957aa3af954ef51605"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Address4", result)
	assert.NilError(t, err)
}

/*
========================================================================================================================
	STRING TYPE TESTS
========================================================================================================================
*/

/*
func Test_ConvertString(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["String"].Inputs.NonIndexed()[0].Type

	input := "somerandomstring123"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("String", result)
	assert.NilError(t, err)

	// Now test incorrect uint type passed
	expectedType = compiledTestContract.Abi.Methods["Uint16"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.ErrorContains(t, err, "strconv.ParseUint: parsing \"somerandomstring123\": invalid syntax")

	// Now test incorrect uint type passed again
	expectedType = compiledTestContract.Abi.Methods["Address"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("String", result)
	assert.ErrorContains(t, err, "abi: cannot use array as type string as argument")
}

func Test_ConvertStringArray(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Strings"].Inputs.NonIndexed()[0].Type

	input := "somerandomstring123,somerandomstring456,somerandomstring789"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Strings", result)
	assert.NilError(t, err)
}

func Test_ConvertString_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["String2"].Inputs.NonIndexed()[0].Type

	input := "somerandomstring123,somerandomstring456"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("String2", result)
	assert.NilError(t, err)
}

func Test_ConvertString_4Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["String4"].Inputs.NonIndexed()[0].Type

	input := "somerandomstring123,somerandomstring456,somerandomstring789,somerandomstring101112"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("String4", result)
	assert.NilError(t, err)
}

/*
========================================================================================================================
	BYTESN TYPE TESTS
========================================================================================================================
*/

/*
func Test_ConvertBytes(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes"].Inputs.NonIndexed()[0].Type

	input := "0x3471555ab9a9952172378abababa"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes", result)
	assert.NilError(t, err)

	input = "ajshduihuieh"
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.ErrorContains(t, err, "encoding/hex: invalid byte: U+0073 's'")

	input = "123456789"
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.ErrorContains(t, err, "encoding/hex: odd length hex string")

	input = "1234567891"
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	expectedType = compiledTestContract.Abi.Methods["Int256"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes", result)
	assert.ErrorContains(t, err, "abi: cannot use ptr as type slice as argument")
}

func Test_ConvertBytes8(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes8"].Inputs.NonIndexed()[0].Type

	input := "0x3471555ab9a99528"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes8", result)
	assert.NilError(t, err)
}

func Test_ConvertBytes32(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes32"].Inputs.NonIndexed()[0].Type

	input := "0x3471555ab9a99528f02f9cdd8f0017fe2f56e01116acc4fe7f78aee900442f35"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes32", result)
	assert.NilError(t, err)
}

/*
========================================================================================================================
	BYTE[] TYPE TESTS
========================================================================================================================
*/

/*
func Test_ConvertByte(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Byte"].Inputs.NonIndexed()[0].Type

	input := "0x34"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byte", result)
	assert.NilError(t, err)

	input = "ajshduihuieh"
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.ErrorContains(t, err, "encoding/hex: invalid byte: U+0073 's'")

	input = "123456789"
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.ErrorContains(t, err, "encoding/hex: odd length hex string")

	input = "1234567891"
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	expectedType = compiledTestContract.Abi.Methods["Int256"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byte", result)
	assert.ErrorContains(t, err, "abi: cannot use ptr as type array as argument")

	input = "somestring"
	expectedType = compiledTestContract.Abi.Methods["String"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byte", result)
	assert.ErrorContains(t, err, "abi: cannot use string as type array as argument")
}

func Test_ConvertByteArray(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Byten"].Inputs.NonIndexed()[0].Type

	input := "0x3471555ab9a9952172378abababa"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byten", result)
	assert.NilError(t, err)

	input = "ajshduihuizxncjoijadseh"
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.ErrorContains(t, err, "encoding/hex: invalid byte: U+0073 's'")

	input = "123456789"
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.ErrorContains(t, err, "encoding/hex: odd length hex string")

	input = "1234567891"
	expectedType = compiledTestContract.Abi.Methods["Int256"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byten", result)
	assert.ErrorContains(t, err, "abi: cannot use *big.Int as type [0]slice as argument")

	input = "somestring"
	expectedType = compiledTestContract.Abi.Methods["String"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byten", result)
	assert.ErrorContains(t, err, "abi: cannot use string as type [0]slice as argument")

	input = "0x3471555ab9a9952172378a"
	expectedType = compiledTestContract.Abi.Methods["Bytes"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byten", result)
	assert.ErrorContains(t, err, "abi: cannot use []uint8 as type [0]array as argument")

	expectedType = compiledTestContract.Abi.Methods["Byten"].Inputs.NonIndexed()[0].Type
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byten", result)
	assert.NilError(t, err)
}

func Test_ConvertByte2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Byte2"].Inputs.NonIndexed()[0].Type

	input := "0x3471"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byte2", result)
	assert.NilError(t, err)

	input = "0x347111"
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byte2", result)
	assert.ErrorContains(t, err, "abi: cannot use [3]array as type [2]array as argument")
}

func Test_ConvertByte8Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Byte8"].Inputs.NonIndexed()[0].Type

	input := "0x3471112712816461"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byte8", result)
	assert.NilError(t, err)

	input = "0x3471112712"
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byte8", result)
	assert.ErrorContains(t, err, "abi: cannot use [5]array as type [8]array as argument")
}

func Test_ConvertByte64Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Byte64"].Inputs.NonIndexed()[0].Type

	input := "0x34711127128723648732687462387468723652345328665847236873468723653471112712872364873268746238746872365234532866584723687346872365"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byte64", result)
	assert.NilError(t, err)

	input = "0x3471112712872364873268746238746872365234532866584723687346872365"
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Byte64", result)
	assert.ErrorContains(t, err, "abi: cannot use [32]array as type [64]array as argument")
}

/*
========================================================================================================================
	BYTESN[] TYPE TESTS
========================================================================================================================
*/

/*
func Test_ConvertBytes1Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes1n"].Inputs.NonIndexed()[0].Type

	input := "0x347111271287236523453286658472368734687236"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes1n", result)
	assert.NilError(t, err)
}

func Test_ConvertBytes2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes2n"].Inputs.NonIndexed()[0].Type

	input := "0x34711127128723652345328665847236873468723669"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes2n", result)
	assert.NilError(t, err)

	input = "0x347111271287236523453286658472368734687236"
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.ErrorContains(t, err, "array length mismatch: cannot convert 21 bytes to [][2]byte array, please supply a multiple of 2")
}

/*
========================================================================================================================
	BYTESN[] TYPE TESTS
========================================================================================================================
*/

/*
func Test_ConvertBytes1_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes1_2"].Inputs.NonIndexed()[0].Type

	input := "0x34711127128723652345328665847236873468723669"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes1_2", result)
	assert.NilError(t, err)

	input = "0x347111271287236523453286658472368734687236"
	result, err = utils.ApplySolidityType(input, expectedType)
	assert.ErrorContains(t, err, "array length mismatch: cannot convert 21 bytes to [][2]byte array, please supply a multiple of 2")
}

func Test_ConvertBytes2_2Array(t *testing.T) {
	// All functions have a single non indexed input arg
	expectedType := compiledTestContract.Abi.Methods["Bytes2_2"].Inputs.NonIndexed()[0].Type

	input := "0x34711127"
	result, err := utils.ApplySolidityType(input, expectedType)
	assert.NilError(t, err)

	_, err = compiledTestContract.Abi.Pack("Bytes2_2", result)
	assert.NilError(t, err)

	//input = "0x347111271287236523453286658472368734687236"
	//result, err = utils.ApplySolidityType(input, expectedType)
	//assert.ErrorContains(t, err, "array length mismatch: cannot convert 21 bytes to [][2]byte array, please supply a multiple of 2")
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
*/