package numbers_test

import (
	"big-integers-calculator/operations/numbers"
	"big-integers-calculator/types"
	"testing"
)

func TestMultiply(t *testing.T) {
	number1 := []complex128{0, 0, 1, 1}
	number2 := []complex128{0, 0, 1, 1}

	got := numbers.Multiply(number1, number2)
	expected := types.Number{0, 1, 2, 1}
	if !got.Equal(&expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}
