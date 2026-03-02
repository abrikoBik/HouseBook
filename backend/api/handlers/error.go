package handlers

import (
	"net/http"

	"github.com/unrolled/render"
)

// ErrorResponse is a helper struct for rendering errors.
type ErrorResponse struct {
	Status int    `json:"-"`
	Error  string `json:"error"`

	err error `json:"-"`
}

func RenderError(renderer *render.Render, writer http.ResponseWriter, err ErrorResponse) {
	Logger.Err(err.err).Msg(err.Error)
	writer.Header().Set("Content-Type", "application/json")
	renderer.JSON(writer, err.Status, err)
}

// TODO: add more errors

func ErrorBadRequest(err error) ErrorResponse {
	return ErrorResponse{
		Status: http.StatusBadRequest,
		Error:  "bad request",
		err:    err,
	}
}

func ErrorRender(err error) ErrorResponse {
	return ErrorResponse{
		Status: http.StatusNotAcceptable,
		Error:  "render error",
		err:    err,
	}
}

func ErrorNotFound(err error) ErrorResponse {
	return ErrorResponse{
		Status: http.StatusNotFound,
		Error:  "not found",
		err:    err,
	}
}

func ErrorDatabase(err error) ErrorResponse {
	return ErrorResponse{
		Status: http.StatusInternalServerError,
		Error:  "database error",
		err:    err,
	}
}

func ErrorDuplicate(err error) ErrorResponse {
	return ErrorResponse{
		Status: http.StatusConflict,
		Error:  "duplicate key provided",
		err:    err,
	}
}

func ErrorGeneric(err error) ErrorResponse {
	return ErrorResponse{
		Status: http.StatusInternalServerError,
		Error:  "unknown error",
		err:    err,
	}
}
