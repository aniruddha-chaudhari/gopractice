package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"studentapi/internal/types"
	"studentapi/internal/utils/response"

	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(fmt.Errorf("request body is empty")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(fmt.Errorf("failed to read request body")))
			return

		}

		//validate the student
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}
		
		slog.Info("new student request")

		response.WriteJson(w, http.StatusCreated, map[string]string{"message": "student created"})
	}
}
