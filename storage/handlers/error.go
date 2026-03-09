package handlers

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Status int    `json:"-"`
	Err    error  `json:"-"`
	Error  string `json:"error"`
}

func RenderError(w http.ResponseWriter, err Error) {
	logger.Error().Err(err.Err).Send()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	err.Error = err.Err.Error()
	json.NewEncoder(w).Encode(err)
}
