package fft

import (
	"math"
)

var PI float64 = math.Acos(-1)

func RecursiveFft(poly []int) {
	if len(poly) == 1 {
		return
	}
	evenIndexes, oddIndexes := halvePolynomial(poly)
	RecursiveFft(evenIndexes)
	RecursiveFft(oddIndexes)

	recursiveFft(poly, evenIndexes, oddIndexes)
}

func halvePolynomial(poly []int) (even []int, odd []int) {
	polyLen := len(poly)
	halfPolyLen := polyLen / 2
	evenIndexes := make([]int, halfPolyLen)
	oddIndexes := make([]int, halfPolyLen)

	i := 0
	j := 0
	for i < polyLen {
		evenIndexes[j] = poly[i]
		oddIndexes[j] = poly[i+1]
		j++
		i += 2
	}
	return evenIndexes, oddIndexes
}

func recursiveFft(poly, evenIndexes, oddIndexes []int) {
	polyLen := len(poly)
	halfPolyLen := polyLen / 2

	angle := 2.0 * PI / float64(polyLen)
	powPrimitiveRoot := 1 + 0i
	primitiveRoot := complex(math.Cos(angle), math.Sin(angle))
	for i := 0; i < halfPolyLen; i++ {
		evenIndex := evenIndexes[i]
		oddIndexMulRoot := int(real(powPrimitiveRoot)) * oddIndexes[i]
		poly[i] = evenIndex + oddIndexMulRoot
		poly[i+halfPolyLen] = evenIndex - oddIndexMulRoot

		powPrimitiveRoot *= primitiveRoot
	}
}
