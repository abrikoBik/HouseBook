package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrorResponse is a helper struct for rendering errors.
type ErrorResponse struct {
	Status int    `json:"-"`
	Error  string `json:"error"`
}

func (err *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// TODO: add logging
	render.Status(r, err.Status)
	return nil
}

func ErrorBadRequest(err error) *ErrorResponse {
	return &ErrorResponse{Status: http.StatusBadRequest, Error: err.Error()}
}

func ErrorRender(err error) *ErrorResponse {
	return &ErrorResponse{Status: http.StatusInternalServerError, Error: err.Error()}
}

func ErrorNotFound(err error) *ErrorResponse {
	return &ErrorResponse{Status: http.StatusNotFound, Error: err.Error()}
}

func ErrorGeneric(err error) *ErrorResponse {
	// TODO: add logging
	return &ErrorResponse{Status: http.StatusInternalServerError, Error: err.Error()}
}
