package handlers

import (
	"big-integers-calculator/operations/numbers"
	"big-integers-calculator/operations/polynomials"
	"bytes"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const MULTIPLY_NUMBERS string = "on"

type poly []int

func (p poly) Trim() poly {
	size := len(p)
	for i := 0; i < size; i++ {
		if p[size-1-i] != 0 {
			return p[:size-i]
		}
	}
	return []int{0}
}

func (p poly) String() string {
	var b bytes.Buffer
	for _, elem := range p {
		b.WriteString(strconv.Itoa(elem))
		b.WriteString(" ")
	}
	return b.String()
}

type number []int

func (n number) Trim() number {
	for i, elem := range n {
		if elem != 0 {
			return n[i:]
		}
	}
	return []int{0}
}

func (n number) String() string {
	var b bytes.Buffer
	for _, elem := range n {
		b.WriteString(strconv.Itoa(elem))
	}
	return b.String()
}

type Data struct {
	Input  string
	Result string
}

func IndexGetHandler(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("html/index.html"))
	template.Execute(writer, nil)
}

func IndexPostHandler(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("html/index.html"))
	request.ParseForm()
	input := request.FormValue("expression")
	validateInput(input)

	left, right := parse(input)
	poly1, poly2 := createPolys(parse(input))
	var data Data
	if request.FormValue("multiplyNumbers") == MULTIPLY_NUMBERS {
		fillPolys(poly1, poly2, left, right, true)
		var res number = numbers.Multiply(poly1, poly2)
		data = Data{
			Input:  request.FormValue("expression"),
			Result: res.Trim().String(),
		}
	} else {
		fillPolys(poly1, poly2, left, right, false)
		var res poly = polynomials.Multiply(poly1, poly2)
		data = Data{
			Input:  request.FormValue("expression"),
			Result: res.Trim().String(),
		}
	}

	template.Execute(writer, data)
}

func validateInput(input string) {
	pattern := `^\d+\*\d+$`
	correctInput, _ := regexp.Match(pattern, []byte(input))
	if !correctInput {
		panic("Incorrect input!")
	}
}

func parse(input string) (left, right string) {
	data := strings.Split(input, "*")
	left, right = data[0], data[1]
	return left, right
}

func createPolys(left, right string) (poly, otherPoly []complex128) {
	size := getSize(len(left), len(right))
	poly = make([]complex128, size)
	otherPoly = make([]complex128, size)
	return poly, otherPoly
}

func getSize(len1, len2 int) int {
	greaterLen := getGreaterLen(len1, len2)
	size := 1
	for size < greaterLen+1 {
		size <<= 1
	}
	size <<= 1
	return size
}

func getGreaterLen(len1, len2 int) int {
	if len1 > len2 {
		return len1
	}
	return len2
}

func fillPolys(poly, otherPoly []complex128, left, right string, mulNumbers bool) {
	var index int
	for i := 0; i < len(left); i++ {
		if mulNumbers {
			index = len(left) - 1 - i
		} else {
			index = i
		}
		poly[i] = complex(float64(rune(left[index])-'0'), 0)
	}

	for i := 0; i < len(right); i++ {
		if mulNumbers {
			index = len(right) - 1 - i
		} else {
			index = i
		}
		otherPoly[i] = complex(float64(rune(right[index])-'0'), 0)
	}
}
