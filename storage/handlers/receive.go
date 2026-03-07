package handlers

import (
	"errors"
	"housebook_storage/consts"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func ReceiveFile(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	response := zerolog.New(writer)
	if err := request.ParseMultipartForm(consts.MaxSize); err != nil {
		logger.Error().Err(err).Msg("exiting at request.ParseMultipartForm")
		writer.WriteHeader(http.StatusInternalServerError)
		// TODO: delete in production?
		response.Error().Err(err).Send()
		return
	}
	stream, header, err := request.FormFile(consts.Key)
	if err != nil {
		logger.Error().Err(err).Msg("exiting at request.FormFile")
		writer.WriteHeader(http.StatusInternalServerError)
		response.Error().Err(err).Send()
		return
	}
	defer stream.Close()
	if header.Size > consts.MaxSize {
		logger.Error().
			Int64("header.Size", header.Size).
			Int("consts.MaxSize", consts.MaxSize).
			Msg("file is too big")
		writer.WriteHeader(http.StatusBadRequest)
		response.Error().Err(errors.New("file is too big")).Send()
		return
	}
	logger.Debug().Str("header.Filename", header.Filename).Msg("receiving")
	id := uuid.New()
	path := filepath.Join(consts.StoragePath, id.String())
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		logger.Error().Err(err).Msg("exiting at os.OpenFile")
		writer.WriteHeader(http.StatusInternalServerError)
		response.Error().Err(err).Send()
		return
	}
	defer file.Close()
	written, err := io.Copy(file, stream)
	if err != nil {
		logger.Error().Err(err).Msg("exiting at io.Copy")
		writer.WriteHeader(http.StatusInternalServerError)
		response.Error().Err(err).Send()
		return
	}
	logger.Info().Int64("written", written).Str("id", id.String()).Send()
	writer.WriteHeader(http.StatusCreated)
	response.Log().Str("id", id.String()).Send()
}
