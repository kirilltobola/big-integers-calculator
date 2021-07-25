package fft

import (
	"math"
)

var PI float64 = math.Acos(-1)

func IterativeFft(poly []complex128, interpolate bool) {
	putInPlaceElements(poly)
	iterativeFft(poly, interpolate)
}

func putInPlaceElements(poly []complex128) {
	polyLen := len(poly)
	bitSize := math.Log2(float64(polyLen))
	for i := 0; i < polyLen; i++ {
		reversedBitsIndex := reverseBits(i, int(bitSize))
		if i < reversedBitsIndex {
			poly[i], poly[reversedBitsIndex] = poly[reversedBitsIndex], poly[i]
		}
	}
}

func reverseBits(number, bitSize int) int {
	reversedBitsNumber := 0
	for i := 0; i < bitSize; i++ {
		if number&(1<<i) != 0 {
			reversedBitsNumber |= 1 << (bitSize - 1 - i)
		}
	}
	return reversedBitsNumber
}

func iterativeFft(poly []complex128, interpolate bool) {
	polyLen := len(poly)
	for len := 2; len <= polyLen; len <<= 1 {
		angle := 2.0 * PI / float64(len)
		if interpolate {
			angle *= -1.0
		}
		primitiveRoot := complex(math.Cos(angle), math.Sin(angle))
		for i := 0; i < polyLen; i += len {
			powPrimitiveRoot := 1 + 0i
			for j := 0; j < len/2; j++ {
				evenIndex := poly[i+j]
				oddIndexMulRoot := powPrimitiveRoot * poly[i+j+len/2]
				poly[i+j] = evenIndex + oddIndexMulRoot
				poly[i+j+len/2] = evenIndex - oddIndexMulRoot
				powPrimitiveRoot *= primitiveRoot
			}
		}
	}
	if interpolate {
		div := complex(float64(polyLen), 0)
		for i := 0; i < polyLen; i++ {
			poly[i] /= div
		}
	}
}
