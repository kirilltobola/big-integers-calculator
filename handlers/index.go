package handlers

import (
	"big-integers-calculator/operations/numbers"
	"big-integers-calculator/operations/polynomials"
	"big-integers-calculator/types"
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

const MULTIPLY_NUMBERS string = "on"

const (
	INDEX_PATH      string = "html/index.html"
	HTML_INPUT_NAME string = "expression"
)

func IndexGetHandler(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles(INDEX_PATH))
	template.Execute(writer, nil)
}

func IndexPostHandler(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles(INDEX_PATH))
	request.ParseForm()
	input := request.FormValue(HTML_INPUT_NAME)

	validateInput(input)
	var data types.Data = types.Data{
		Input: request.FormValue(HTML_INPUT_NAME),
	}
	left, right := parse(input)
	poly1, poly2 := createPolys(left, right)

	if request.FormValue("multiplyNumbers") == MULTIPLY_NUMBERS {
		fillNumber(poly1, left)
		fillNumber(poly2, right)
		var res types.Number = numbers.Multiply(poly1, poly2)
		data.Result = res.Trim().String()
	} else {
		fillPoly(poly1, left)
		fillPoly(poly2, right)
		var res types.Poly = polynomials.Multiply(poly1, poly2)
		data.Result = res.Trim().String()
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
	delimeter := "*"
	data := strings.Split(input, delimeter)
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

func fillPoly(poly []complex128, data string) {
	dataSize := len(data)
	for i := 0; i < dataSize; i++ {
		poly[i] = complex(float64(rune(data[i])-'0'), 0)
	}
}

func fillNumber(number []complex128, data string) {
	dataSize := len(data)
	for i := 0; i < dataSize; i++ {
		number[i] = complex(float64(rune(data[dataSize-1-i])-'0'), 0)
	}
}
