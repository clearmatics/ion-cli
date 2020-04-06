package utils_test

import (
	"github.com/clearmatics/ion-cli/utils"
	"gotest.tools/assert"
	"testing"
)

func Test_ConvertStringToArray(t *testing.T) {
	testString := "[[[sub,sub],[sub2,sub2]],[[sub3,sub3],[sub4,sub4]]]"
	expectedResult := []interface{}{[]interface{}{[]string{"sub", "sub"}, []string{"sub2", "sub2"}}, []interface{}{[]string{"sub3", "sub3"}, []string{"sub4", "sub4"}}}

	result, err := utils.ConvertStringArray(testString)
	assert.NilError(t, err)

	assert.DeepEqual(t, result, expectedResult)
}
