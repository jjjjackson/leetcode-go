package helper

import (
	"fmt"
	"reflect"
)

type (
	// Value reflects a placeholder for any type
	Value interface{}
	// Slice reflects a slice or array of any type
	Slice interface{}
	// Pointer reflects a pointer that references any type
	Pointer interface{}
	// Map reflects a map of any type
	Map interface{}
)

// IsSlice returns true if the specified argument is a slice or array and false
// otherwise.
func IsSlice(value Value) bool {
	kind := reflect.ValueOf(value).Kind()
	return kind == reflect.Slice || kind == reflect.Array
}

func SliceEqual(slice1, slice2 Slice) (bool, error) {
	if !IsSlice(slice1) {
		return false, fmt.Errorf("argument type '%T' is not a slice", slice1)
	} else if !IsSlice(slice2) {
		return false, fmt.Errorf("argument type '%T' is not a slice", slice2)
	}

	slice1Val := reflect.ValueOf(slice1)
	slice2Val := reflect.ValueOf(slice2)

	if slice1Val.Type().Elem() != slice2Val.Type().Elem() {
		return false, fmt.Errorf("type of '%v' does not match type of '%v'", slice1Val.Type().Elem(), slice2Val.Type().Elem())
	}

	if slice1Val.Len() != slice2Val.Len() {
		return false, nil
	}

	result := true
	i, n := 0, slice1Val.Len()

	for i < n {
		j := 0
		e := false
		for j < n && !e {
			if slice1Val.Index(i).Interface() == slice2Val.Index(j).Interface() {
				e = true
			}
			j++
		}
		if !e {
			result = false
		}
		i++
	}

	return result, nil
}
