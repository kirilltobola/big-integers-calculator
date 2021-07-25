package handlers

import (
	"big-integers-calculator/numbers"
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

const (
	NUMBERS = iota
	POLYNOMIALS
)

type Data struct {
	Expression string
	Result     []int
}

func IndexGetHandler(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("html/index.html"))
	template.Execute(writer, nil)
}

func IndexPostHandler(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("html/index.html"))
	request.ParseForm()
	inputExpression := request.FormValue("expression")
	var res []int
	switch validateInput(inputExpression) {
	case NUMBERS:
		num1, num2 := parseNumbers(inputExpression)
		res = numbers.Multiply(num1, num2)
	case POLYNOMIALS:
		return
	}

	data := Data{
		Expression: request.FormValue("expression"),
		Result:     res,
	}
	template.Execute(writer, data)
}

func validateInput(input string) int {
	numbersPattern := `^\d+\*\d+$`
	polysPattern := `^\(\d+(\s\d+)*\)\*\(\d+(\s\d+)*\)$`
	polynomials, _ := regexp.Match(polysPattern, []byte(input))
	numbers, _ := regexp.Match(numbersPattern, []byte(input))

	if polynomials {
		return POLYNOMIALS
	} else if numbers {
		return NUMBERS
	}
	panic("Incorrect input!")
}

func parseNumbers(input string) (poly []complex128, otherPoly []complex128) {
	poly = make([]complex128, 0)
	otherPoly = make([]complex128, 0)

	data := strings.Split(input, "*")
	left, right := data[0], data[1]
	for i := 0; i < len(left); i++ {
		poly = append(poly, complex(float64(rune(left[i])-'0'), 0))
	}
	for i := 0; i < len(right); i++ {
		otherPoly = append(otherPoly, complex(float64(rune(right[i])-'0'), 0))
	}
	return poly, otherPoly
}
