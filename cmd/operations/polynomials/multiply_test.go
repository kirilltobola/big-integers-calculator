package polynomials_test

import (
	"big-integers-calculator/cmd/operations/polynomials"
	"big-integers-calculator/cmd/types"
	"testing"
)

func TestMultiply(t *testing.T) {
	poly1 := []complex128{3, 2, 1, 0, 0, 0, 0, 0}
	poly2 := []complex128{3, 2, 1, 0, 0, 0, 0, 0}
	got := polynomials.Multiply(poly1, poly2)
	expected := types.Poly{9, 12, 10, 4, 1, 0, 0, 0}

	if !got.Equal(&expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}
