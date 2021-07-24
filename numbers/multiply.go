package numbers

import "big-integers-calculator/polynomial"

func Multiply(number, otherNumber []complex128) []int {
	reverse(number)
	reverse(otherNumber)
	res := polynomial.Multiply(&number, &otherNumber)
	normalize(res)
	reverse(res)
	return res
}

func reverse(number interface{}) {
	switch number.(type) {
	case []int:
		reverseInt(number.([]int))
	case []complex128:
		reverseComplex(number.([]complex128))
	default:
		panic("Wrong type!")
	}
}

func reverseComplex(number []complex128) {
	i, j := 0, len(number)-1
	for i < j {
		(number)[i], (number)[j] = (number)[j], (number)[i]
		i++
		j--
	}
}

func reverseInt(number []int) {
	i, j := 0, len(number)-1
	for i < j {
		(number)[i], (number)[j] = (number)[j], (number)[i]
		i++
		j--
	}
}

func normalize(number []int) {
	carry := 0
	for i := 0; i < len(number); i++ {
		number[i] += carry
		carry = number[i] / 10
		number[i] %= 10
	}
}
