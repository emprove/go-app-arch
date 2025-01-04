package rest

import (
	"encoding/json"
	"net/http"

	"go-app-arch/internal/infrastructure/logging"
	"go-app-arch/internal/validation"
)

func JSON(w http.ResponseWriter, status int, data any) error {
	return JSONwithHeaders(w, status, data, nil)
}

func JSONwithHeaders(w http.ResponseWriter, status int, data any, headers http.Header) error {
	jsn, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	jsn = append(jsn, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsn)

	return nil
}

func ErrorMessage(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	err := JSONwithHeaders(w, status, map[string]string{"error": message}, headers)
	if err != nil {
		logging.LogRequestError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ServerError(w http.ResponseWriter, r *http.Request, err error) {
	logging.LogRequestError(r, err)
	ErrorMessage(w, r, http.StatusInternalServerError, "Internal server error", nil)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	ErrorMessage(w, r, http.StatusNotFound, "Not found", nil)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	ErrorMessage(w, r, http.StatusMethodNotAllowed, "Method not allowed", nil)
}

func BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	ErrorMessage(w, r, http.StatusBadRequest, err.Error(), nil)
}

func Forbidden(w http.ResponseWriter, r *http.Request) {
	ErrorMessage(w, r, http.StatusForbidden, "Forbidden", nil)
}

func FailedValidation(w http.ResponseWriter, r *http.Request, v validation.Validator) {
	err := JSON(w, http.StatusUnprocessableEntity, v)
	if err != nil {
		ServerError(w, r, err)
	}
}
