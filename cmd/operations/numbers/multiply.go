package numbers

import (
	"big-integers-calculator/cmd/operations/polynomials"
	"big-integers-calculator/cmd/types"
)

func Multiply(number, otherNumber []complex128) types.Number {
	res := polynomials.Multiply(number, otherNumber)
	normalize(res)
	reverse(res)
	return types.Number(res)
}

func normalize(number []int) {
	carry := 0
	for i := 0; i < len(number); i++ {
		number[i] += carry
		carry = number[i] / 10
		number[i] %= 10
	}
}

func reverse(number []int) {
	i, j := 0, len(number)-1
	for i < j {
		(number)[i], (number)[j] = (number)[j], (number)[i]
		i++
		j--
	}
}
