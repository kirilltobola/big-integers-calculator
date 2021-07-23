package fft

import (
	"math"
)

var PI float64 = math.Acos(-1)

func Fftrec(poly []int) {
	polyLen := len(poly)
	if polyLen == 1 {
		return
	}
	evenIndexes, oddIndexes := halvePolynomial(poly)
	Fftrec(evenIndexes)
	Fftrec(oddIndexes)

	fft(poly, evenIndexes, oddIndexes)
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

func fft(poly, evenIndexes, oddIndexes []int) []int {
	polyLen := len(poly)
	halfPolyLen := polyLen / 2

	angle := float64(2) * PI / float64(polyLen)
	powPrimitiveRoot := complex(1.0, 0)
	primitiveRoot := complex(math.Cos(angle), math.Sin(angle))

	for i := 0; i < halfPolyLen; i++ {
		evenIndex := complex(float64(evenIndexes[i]), 0)
		oddIndexMulRoot := powPrimitiveRoot * complex(float64(oddIndexes[i]), 0)

		poly[i] = int(real(evenIndex + oddIndexMulRoot))
		poly[i+halfPolyLen] = int(real(evenIndex - oddIndexMulRoot))

		powPrimitiveRoot *= primitiveRoot
	}
	return poly
}
