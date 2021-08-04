package api

import (
	"big-integers-calculator/cmd/types"
	"big-integers-calculator/web/handlers"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func MuliplyData(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var data types.Data = types.Data{
		Input: request.Form.Get("input"),
	}
	if handlers.ValidateInput(data.Input) {
		handlers.Multiply(&data, request)
	} else {
		data.Error = errors.New(handlers.INCORRECT_INPUT_MSG)
	}

	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(writer, "%+v", string(out))
}
