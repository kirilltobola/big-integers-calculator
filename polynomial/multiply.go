package polynomial

import (
	"big-integers-calculator/fft"
	"math"
)

func Multiply(poly, otherPoly *[]complex128) []int {
	greaterLen := getGreaterLen(poly, otherPoly)
	mulSize := getMulSize(greaterLen)

	resize(poly, mulSize)
	resize(otherPoly, mulSize)

	fft.IterativeFft(*poly, false)
	fft.IterativeFft(*otherPoly, false)
	for i := 0; i < mulSize; i++ {
		(*poly)[i] *= (*otherPoly)[i]
	}

	fft.IterativeFft(*poly, true)
	res := make([]int, len(*poly))
	for i := 0; i < len(res); i++ {
		res[i] = int(real((*poly)[i]) + 0.5)
	}
	return res
}

func getGreaterLen(poly, otherPoly *[]complex128) int {
	floatPolyLen := float64(len(*poly))
	floatOtherPolyLen := float64(len(*otherPoly))
	return int(math.Max(floatPolyLen, floatOtherPolyLen))
}

func getMulSize(maxPolyLen int) int {
	mulSize := 1
	for mulSize < maxPolyLen+1 {
		mulSize <<= 1
	}
	mulSize <<= 1
	return mulSize
}

func resize(poly *[]complex128, size int) {
	for len(*poly) < size {
		*poly = append(*poly, 0+0i)
	}
}
