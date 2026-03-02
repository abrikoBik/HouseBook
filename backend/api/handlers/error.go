package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/unrolled/render"
)

// ErrorResponse is a helper struct for rendering errors.
type ErrorResponse struct {
	Status      int    `json:"-"`
	Error       string `json:"error"`
	Explanation string `json:"explanation,omitempty"`

	err error `json:"-"`
}

// RenderError logs the error, sets the Content-Type header to application/json and writes err as JSON to writer.
func RenderError(renderer *render.Render, writer http.ResponseWriter, err ErrorResponse) {
	Logger.Err(err.err).Msg(err.Error)
	writer.Header().Set("Content-Type", "application/json")
	renderer.JSON(writer, err.Status, err)
}

// TODO: add more errors

func ErrorBadJSON(err error) ErrorResponse {
	return ErrorResponse{
		Status: http.StatusBadRequest,
		Error:  "bad JSON",
		err:    err,
	}
}

func ErrorBadID(err error, key string) ErrorResponse {
	return ErrorResponse{
		Status:      http.StatusBadRequest,
		Error:       "bad ID",
		Explanation: key,
		err:         err,
	}
}

func ErrorBadSearchParams(err error, object string) ErrorResponse {
	return ErrorResponse{
		Status:      http.StatusBadRequest,
		Error:       "bad parameters",
		Explanation: object,
		err:         err,
	}
}

func ErrorBadPaginationParams(err error) ErrorResponse {
	return ErrorResponse{
		Status:      http.StatusBadRequest,
		Error:       "bad parameters",
		Explanation: "limit or offset",
		err:         err,
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

func ErrorDuplicate(err error, pgErr *pgconn.PgError) ErrorResponse {
	return ErrorResponse{
		Status:      http.StatusConflict,
		Error:       "duplicate key provided",
		Explanation: pgErr.ConstraintName,
		err:         err,
	}
}

func ErrorGeneric(err error) ErrorResponse {
	return ErrorResponse{
		Status: http.StatusInternalServerError,
		Error:  "unknown error",
		err:    err,
	}
}
