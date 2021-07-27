package handlers

import (
	"big-integers-calculator/operations/numbers"
	"big-integers-calculator/operations/polynomials"
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

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
	left, right := parse(inputExpression)
	poly1, poly2 := createPolys(left, right)
	if request.FormValue("multiplyNumbers") == MULTIPLY_NUMBERS {
		fillPolys(poly1, poly2, left, right, true)
		res = numbers.Multiply(poly1, poly2)
	} else {
		fillPolys(poly1, poly2, left, right, false)
		res = polynomials.Multiply(poly1, poly2)
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

func parse(input string) (left, right string) {
	data := strings.Split(input, "*")
	left, right = data[0], data[1]
	return left, right
}

func createPolys(left, right string) (poly, otherPoly []complex128) {
	greaterLen := getGreaterLen(len(left), len(right))
	mulSize := getMulSize(greaterLen)

	poly = make([]complex128, mulSize)
	otherPoly = make([]complex128, mulSize)
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
