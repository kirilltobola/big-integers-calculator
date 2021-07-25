package handlers

import (
	"html/template"
	"net/http"
)

type Data struct {
	Result string
}

func IndexGetHandler(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("html/index.html"))
	template.Execute(writer, nil)
}

func IndexPostHandler(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("html/index.html"))
	request.ParseForm()

	data := Data{
		Result: request.FormValue("expression"),
	}
	template.Execute(writer, data)
}
