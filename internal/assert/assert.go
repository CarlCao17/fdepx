package assert

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

// Copied from https://github.com/stretchr/testify/blob/v1.7.0/assert/assertions.go#L334
//
// isEqual determines if two objects are considered equal.
//
// This function does no assertion of any kind.
func isEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	switch exp := expected.(type) {
	// handle byte slices more efficiently
	case []byte:
		act, ok := actual.([]byte)
		if !ok {
			return false
		}
		return bytes.Equal(exp, act)
	// float point should not use ==
	case float32:
		act, ok := actual.(float32)
		if !ok {
			return false
		}
		return floatAlmostEqual(float64(exp), float64(act))
	case float64:
		act, ok := actual.(float64)
		if !ok {
			return false
		}
		return floatAlmostEqual(exp, act)
	case []float32:
		act, ok := actual.([]float32)
		if !ok {
			return false
		}
		if len(exp) != len(act) {
			return false
		}
		for i := range exp {
			if !floatAlmostEqual(float64(exp[i]), float64(act[i])) {
				return false
			}
		}
		return true
	case []float64:
		act, ok := actual.([]float64)
		if !ok {
			return false
		}
		if len(exp) != len(act) {
			return false
		}
		for i := range exp {
			if !floatAlmostEqual(exp[i], act[i]) {
				return false
			}
		}
		return true
	default:
		return reflect.DeepEqual(expected, actual)
	}
}

func floatAlmostEqual(f1, f2 float64) bool {
	const delta = 1e-6
	return math.Abs(f1-f2) < delta
}

func Equal(t *testing.T, expected, actual interface{}) bool {
	ok := isEqual(expected, actual)
	if !ok {
		t.Errorf("\nExpect : %#v\nbut got: %#v\n", expected, actual)
	}
	return ok
}
