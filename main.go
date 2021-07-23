package main

import (
	"big-integers-calculator/fft"
	"fmt"
)

func main() {
	poly := []int{1, 2, 3, 4}
	fft.Fftrec(poly)

	for _, item := range poly {
		fmt.Println(item)
	}
}
