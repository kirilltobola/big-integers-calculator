package main

import (
	"big-integers-calculator/fft"
)

func main() {
	poly := []int{1, 2, 3, 4, 1, 2, 3, 4}
	poly_copy := make([]int, len(poly))
	copy(poly_copy, poly)

	fft.RecursiveFft(poly_copy)
	fft.IterativeFft(poly)
	for i := 0; i < len(poly); i++ {
		if poly[i] != poly_copy[i] {
			panic("fft broken!")
		}
	}
}
