package errors

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
	"sort"
)

type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func BadRequest(msg string) ErrorResponse {
	if msg == "" {
		msg = "Your request is in a bad format."
	}
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
}

func (e ErrorResponse) StatusCode() int {
	return e.Status
}

func InternalServerError(msg string) ErrorResponse {
	if msg == "" {
		msg = "We encountered an error while processing your request."
	}
	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: msg,
	}
}

func NotFound(msg string) ErrorResponse {
	if msg == "" {
		msg = "The requested resource was not found."
	}
	return ErrorResponse{
		Status:  http.StatusNotFound,
		Message: msg,
	}
}

func Unathorized(msg string) ErrorResponse {
	if msg == "" {
		msg = "You are not authenticated to perform the requested action."
	}
	return ErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: msg,
	}
}

type invalidField struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func InvalidInput(errs validation.Errors) ErrorResponse {
	var details []invalidField
	var fields []string
	for field := range errs {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	for _, field := range fields {
		details = append(details, invalidField{
			Field: field,
			Error: errs[field].Error(),
		})
	}

	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: "There is some problem with the data you submitted.",
		Details: details,
	}
}
