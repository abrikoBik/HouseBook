package handlers

import (
	"net/http"
)

// ErrorResponse is a helper struct for rendering errors.
type ErrorResponse struct {
	Status int    `json:"-"`
	Error  string `json:"error"`
	// TODO: remove in production
	Debug string `json:"debug"`
}

func ErrorBadRequest(err error) *ErrorResponse {
	Logger.Err(err).Msg("bad request")
	return &ErrorResponse{
		Status: http.StatusBadRequest,
		Error:  "bad request",
		Debug:  err.Error(),
	}
}

func ErrorRender(err error) *ErrorResponse {
	Logger.Err(err).Msg("render error")
	return &ErrorResponse{
		Status: http.StatusInternalServerError,
		Error:  "render error",
		Debug:  err.Error(),
	}
}

func ErrorNotFound(err error) *ErrorResponse {
	Logger.Err(err).Msg("not found")
	return &ErrorResponse{
		Status: http.StatusNotFound,
		Error:  "not found",
		Debug:  err.Error(),
	}
}

func ErrorDatabase(err error) *ErrorResponse {
	Logger.Err(err).Msg("database error")
	return &ErrorResponse{
		Status: http.StatusInternalServerError,
		Error:  "database error",
		Debug:  err.Error(),
	}
}

func ErrorGeneric(err error) *ErrorResponse {
	Logger.Err(err).Msg("unknown error")
	return &ErrorResponse{
		Status: http.StatusInternalServerError,
		Error:  "unknown error",
		Debug:  err.Error(),
	}
}
