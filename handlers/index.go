package handlers

import (
	"html/template"
	"net/http"
	"regexp"
)

const (
	NUMBERS = iota
	POLYNOMIALS
)

type Data struct {
	Expression string
	Result     string
}

func IndexGetHandler(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("html/index.html"))
	template.Execute(writer, nil)
}

func IndexPostHandler(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("html/index.html"))
	request.ParseForm()
	inputExpression := request.FormValue("expression")
	switch validateInput(inputExpression) {
	case NUMBERS:
		return
	case POLYNOMIALS:
		return
	}

	data := Data{
		Expression: request.FormValue("expression"),
		Result:     "",
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
