package main

import (
	"big-integers-calculator/numbers"
	"big-integers-calculator/polynomial"
	"fmt"
)

func main() {
	testNumber()
}

func testPoly() {
	poly := []complex128{-5, 2, 8, -3, -3, 0, 1, 0, 1}
	otherPoly := []complex128{21, -9, -4, 0, 5, 0, 3}

	polynomial.Multiply(&poly, &otherPoly)
	for _, elem := range poly {
		fmt.Printf("%.2f ", real(elem))
	}
}

func testNumber() {
	number := []complex128{4, 3, 6, 3, 4, 5, 6, 3, 4, 5, 6, 4, 5, 6}
	otherNumber := []complex128{4, 6, 4, 5, 6, 4, 5, 6, 4, 5, 6}

	res := numbers.Multiply(number, otherNumber)
	for _, elem := range res {
		fmt.Print(elem)
	}
}
