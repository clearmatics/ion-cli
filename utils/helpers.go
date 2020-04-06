package utils

import (
	"errors"
	"strings"
)

// Takes a string representation of an array (including nested array) and converts to array of strings
// Only supports directly n-dimensional arrays where the items of any array is either an array or a value
// Meaning we do not support arrays like [item, [arrayitem, arrayitem], item]
// Only [item, item, item] or [[arrayitem, arrayitem], [array2item, array2item]]
// Random spacings between list items in a list are removed: [[item,item] , [item,item]] -> [[item,item],[item,item]]
// but spacings between regular items in a list are not removed and are considered as part of the intended item:
// [[ item ,   item] , [item,item]] -> [[ item ,   item],[item,item]]
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
	} else { // If input is a debraced list of items, return the array form. Else return the single unarrayed item
		list := strings.Split(strings.ReplaceAll(input, " ", ""), ",")
		if len(list) > 1 {
			return list, nil
		}
		return list[0], nil
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
