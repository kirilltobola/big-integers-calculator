package types

import (
	"bytes"
	"strconv"
)

type Poly []int

func (p Poly) Trim() Poly {
	size := len(p)
	for i := 0; i < size; i++ {
		if p[size-1-i] != 0 {
			return p[:size-i]
		}
	}
	return []int{0}
}

func (p Poly) String() string {
	var b bytes.Buffer
	for _, elem := range p {
		b.WriteString(strconv.Itoa(elem))
		b.WriteString(" ")
	}
	return b.String()
}

func (p *Poly) Equal(other *Poly) bool {
	if len(*p) != len(*other) {
		return false
	}

	for i, elem := range *p {
		if elem != (*other)[i] {
			return false
		}
	}

	return true
}
