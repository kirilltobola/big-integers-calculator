package main

import (
	"big-integers-calculator/polynomial"
	"fmt"
)

func main() {
	poly := []complex128{-5, 2, 8, -3, -3, 0, 1, 0, 1}
	otherPoly := []complex128{21, -9, -4, 0, 5, 0, 3}

	polynomial.Multiply(&poly, &otherPoly)
	for _, elem := range poly {
		fmt.Printf("%.2f ", real(elem))
	}
}
