package response

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusSuccess = "success"
	StatusError = "error"
)

func WriteJson(w http.ResponseWriter,status int,data interface{}) error {


	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GenralError(err error) Response {
	return Response{
		Status: StatusError,
		Error: err.Error(),
	}
}

func validationError(errs validator.ValidationErrors) Response {
	
}