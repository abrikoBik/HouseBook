package handlers

import (
	"encoding/json"
	"errors"
	"housebook_storage/consts"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	_ "golang.org/x/image/webp"
)

type IDResponse struct {
	ID string `json:"id"`
}

func ReceiveFile(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseMultipartForm(consts.MaxSize); err != nil {
		RenderError(writer, Error{http.StatusInternalServerError, err, ""})
		return
	}
	stream, header, err := request.FormFile(consts.Key)
	if err != nil {
		RenderError(writer, Error{http.StatusInternalServerError, err, ""})
		return
	}
	defer stream.Close()
	if header.Size > consts.MaxSize {
		RenderError(writer, Error{http.StatusBadRequest, errors.New("file is too big"), ""})
		return
	}
	logger.Debug().Str("header.Filename", header.Filename).Msg("receiving")
	id := uuid.New()
	path := filepath.Join(consts.StoragePath, id.String())
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		RenderError(writer, Error{http.StatusInternalServerError, err, ""})
		return
	}
	defer file.Close()
	img, format, err := image.Decode(stream)
	if err != nil {
		RenderError(writer, Error{http.StatusBadRequest, err, ""})
		return
	}
	switch format {
	case "png":
		err = png.Encode(file, img)
	case "jpeg":
		err = jpeg.Encode(file, img, nil)
	case "webp":
		err = png.Encode(file, img)
	}
	if err != nil {
		RenderError(writer, Error{http.StatusInternalServerError, err, ""})
		return
	}
	logger.Info().Str("format", format).Str("id", id.String()).Send()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(IDResponse{id.String()})
}
