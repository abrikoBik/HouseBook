package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrorResponse is a helper struct for rendering errors.
type ErrorResponse struct {
	Status int    `json:"-"`
	Error  string `json:"error"`
	// TODO: remove in production
	Debug string `json:"debug"`
}

func (err *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// TODO: add logging
	render.Status(r, err.Status)
	return nil
}

func ErrorBadRequest(err error) *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusBadRequest,
		Error:  "bad request",
		Debug:  err.Error(),
	}
}

func ErrorRender(err error) *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusInternalServerError,
		Error:  "render error",
		Debug:  err.Error(),
	}
}

func ErrorNotFound(err error) *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusNotFound,
		Error:  "not found",
		Debug:  err.Error(),
	}
}

func ErrorDatabase(err error) *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusInternalServerError,
		Error:  "database error",
		Debug:  err.Error(),
	}
}

func ErrorGeneric(err error) *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusInternalServerError,
		Error:  "unknown error",
		Debug:  err.Error(),
	}
}
