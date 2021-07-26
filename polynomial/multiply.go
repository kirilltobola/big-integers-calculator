package polynomial

import (
	"big-integers-calculator/fft"
)

func Multiply(poly, otherPoly []complex128) []int {
	fft.IterativeFft(poly, false)
	fft.IterativeFft(otherPoly, false)
	for i := 0; i < len(poly); i++ {
		(poly)[i] *= (otherPoly)[i]
	}

	fft.IterativeFft(poly, true)
	return castToInt(poly)
}

func castToInt(poly []complex128) []int {
	res := make([]int, len(poly))
	for i := 0; i < len(res); i++ {
		res[i] = int(real(poly[i]) + 0.5)
	}

	return res
}
