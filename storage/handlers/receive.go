package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"storage/consts"

	"github.com/google/uuid"
)

func ReceiveFile(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseMultipartForm(consts.MaxSize); err != nil {
		logger.Error().Err(err).Msg("exiting at request.ParseMultipartForm")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	stream, header, err := request.FormFile(consts.Key)
	if err != nil {
		logger.Error().Err(err).Msg("exiting at request.FormFile")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer stream.Close()
	if header.Size > consts.MaxSize {
		logger.Error().
			Int64("header.Size", header.Size).
			Int("consts.MaxSize", consts.MaxSize).
			Msg("file is too big")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("receiving", header.Filename)
	id := uuid.New()
	path := filepath.Join(consts.StoragePath, id.String())
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		logger.Error().Err(err).Msg("exiting at os.OpenFile")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	written, err := io.Copy(file, stream)
	if err != nil {
		logger.Error().Err(err).Msg("exiting at io.Copy")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info().Int64("written", written).Str("id", id.String()).Send()
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(id.String()))
}
