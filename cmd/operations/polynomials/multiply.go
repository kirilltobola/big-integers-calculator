package polynomials

import (
	"big-integers-calculator/cmd/fft"
	"big-integers-calculator/cmd/types"
)

func Multiply(poly, otherPoly []complex128) types.Poly {
	fft.Fft(poly, false)
	fft.Fft(otherPoly, false)
	for i := 0; i < len(poly); i++ {
		(poly)[i] *= (otherPoly)[i]
	}

	fft.Fft(poly, true)
	return castToInt(poly)
}

func castToInt(poly []complex128) []int {
	res := make([]int, len(poly))
	for i := 0; i < len(res); i++ {
		res[i] = int(real(poly[i]) + 0.5)
	}

	return res
}
