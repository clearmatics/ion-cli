package utils_test

import (
	"github.com/clearmatics/ion-cli/utils"
	"github.com/ethereum/go-ethereum/common"
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func Test_ConvertStringToArray(t *testing.T) {
	testString := "[[[sub,sub],[sub2,sub2]],[[sub3,sub3],[sub4,sub4]]]"
	expectedResult := []interface{}{[]interface{}{[]string{"sub", "sub"}, []string{"sub2", "sub2"}}, []interface{}{[]string{"sub3", "sub3"}, []string{"sub4", "sub4"}}}

	result, err := utils.ConvertStringArray(testString)
	assert.NilError(t, err)

	assert.DeepEqual(t, result, expectedResult)
}

func Test_HasExpectedArrayStructure(t *testing.T) {
	assertedType := reflect.TypeOf([][2][]int{})

	// Passes
	testInput := [][2][]string{{{"some", "long", "string", "of", "dynamic", "strings"}, {"that", "should", "fill", "the", "test"}}}
	assert.Check(t, utils.HasExpectedArrayStructure(testInput, assertedType))

	// Fails
	testInput2 := [][3][]string{{{"some", "long", "string", "of", "dynamic", "strings"}, {"that", "should", "fill", "the", "test"}}}
	assert.Assert(t, !utils.HasExpectedArrayStructure(testInput2, assertedType))

	// Passes
	assert.Check(t, utils.HasExpectedArrayStructure("something", reflect.TypeOf(0)))

	// Passes
	testInput3 := [3][]int{{1, 2, 3}, {4, 5, 6, 7, 8, 9}, {7, 8, 9, 10}}
	assertedType = reflect.TypeOf(testInput3)
	assert.Assert(t, utils.HasExpectedArrayStructure(testInput3, assertedType))

	// Passes
	testInput4 := [3]int{1, 2, 3}
	assertedType = reflect.TypeOf([3]string{})
	assert.Assert(t, utils.HasExpectedArrayStructure(testInput4, assertedType))

	// Passes
	testInput4_1 := []int{1, 2, 3}
	assertedType = reflect.TypeOf([3]string{})
	assert.Assert(t, utils.HasExpectedArrayStructure(testInput4_1, assertedType))

	// Fails
	// Note here that if the expected asserted type is a slice and the input is an array, we fail this case
	// However if the expected asserted type is an array and input is a slice, we allow it
	testInput4_2 := [3]int{1, 2, 3}
	assertedType = reflect.TypeOf([]string{})
	assert.Assert(t, !utils.HasExpectedArrayStructure(testInput4_2, assertedType))

	// Passes
	testInput5 := []int{1, 2, 3}
	assertedType = reflect.TypeOf([]string{})
	assert.Assert(t, utils.HasExpectedArrayStructure(testInput5, assertedType))

	// Fails
	testInput6 := [][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}}
	assertedType = reflect.TypeOf([]string{})
	assert.Assert(t, !utils.HasExpectedArrayStructure(testInput6, assertedType))

	// Passes
	testInput7 := [][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}}
	assertedType = reflect.TypeOf([][3]string{})
	assert.Assert(t, utils.HasExpectedArrayStructure(testInput7, assertedType))

	testInput8 := []common.Address{common.HexToAddress("0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8"), common.HexToAddress("0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8")}
	assertedType = reflect.TypeOf([]common.Address{})
	assert.Check(t, utils.HasExpectedArrayStructure(testInput8, assertedType))

	testInput9 := []string{"0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8", "0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8"}
	assertedType = reflect.TypeOf([]common.Address{})
	assert.Check(t, utils.HasExpectedArrayStructure(testInput9, assertedType))
}
