package handlers

import (
	"big-integers-calculator/operations/numbers"
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

// const (
// 	NUMBERS = iota
// 	POLYNOMIALS
// )

const MULTIPLY_NUMBERS string = "on"

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
	validateInput(inputExpression)

	if request.FormValue("multiplyNumbers") == MULTIPLY_NUMBERS {
		num1, num2 := parseNumbers(inputExpression)
		res = numbers.Multiply(num1, num2)
	} else {
		return
	}

	data := Data{
		Expression: request.FormValue("expression"),
		Result:     res,
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

func parseNumbers(input string) (poly, otherPoly []complex128) {
	data := strings.Split(input, "*")

	greaterLen := getGreaterLen(len(data[0]), len(data[1]))
	mulSize := getMulSize(greaterLen)
	poly = make([]complex128, mulSize)
	otherPoly = make([]complex128, mulSize)

	left, right := data[0], data[1]
	for i := 0; i < len(left); i++ {
		poly[i] = complex(float64(rune(left[len(left)-1-i])-'0'), 0)
	}
	for i := 0; i < len(right); i++ {
		otherPoly[i] = complex(float64(rune(right[len(right)-1-i])-'0'), 0)
	}
	return poly, otherPoly
}

func getGreaterLen(len1, len2 int) int {
	if len1 > len2 {
		return len1
	}
	return len2
}

func getMulSize(maxPolyLen int) int {
	mulSize := 1
	for mulSize < maxPolyLen+1 {
		mulSize <<= 1
	}
	mulSize <<= 1
	return mulSize
}
