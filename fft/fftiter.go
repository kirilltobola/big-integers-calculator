package fft

import (
	"math"
)

func IterativeFft(poly []int) {
	putInPlaceElements(poly)
	iterativeFft(poly)
}

func putInPlaceElements(poly []int) {
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

func iterativeFft(poly []int) {
	polyLen := len(poly)
	for len := 2; len <= polyLen; len <<= 1 {
		angle := 2.0 * PI / float64(len)
		primitiveRoot := complex(math.Cos(angle), math.Sin(angle))
		for i := 0; i < polyLen; i += len {
			powPrimitiveRoot := 1 + 0i
			for j := 0; j < len/2; j++ {
				evenIndex := poly[i+j]
				oddIndexMulRoot := int(real(powPrimitiveRoot)) * poly[i+j+len/2]
				poly[i+j] = evenIndex + oddIndexMulRoot
				poly[i+j+len/2] = evenIndex - oddIndexMulRoot
				powPrimitiveRoot *= primitiveRoot
			}
		}
	}
}
