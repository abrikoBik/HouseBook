package handlers

import (
	"encoding/base64"
	"fmt"
	"housebook_storage/consts"
	"image"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"golang.org/x/image/draw"

	"image/jpeg"
	"image/png"
)

func SendFile(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var (
		base64Query = request.URL.Query().Get("base64")
		widthQuery  = request.URL.Query().Get("width")
		heightQuery = request.URL.Query().Get("height")
		err         error
	)
	response := zerolog.New(writer)
	id := request.PathValue("id")
	if err := uuid.Validate(id); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		response.Error().Err(err).Send()
		return
	}
	path := filepath.Join(consts.StoragePath, id)
	logger.Debug().Str("path", path).Send()
	file, err := os.Open(path)
	if err != nil {
		logger.Error().Err(err).Msg("exiting at os.Open")
		if os.IsNotExist(err) {
			writer.WriteHeader(http.StatusNotFound)
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
		}
		response.Error().Err(err).Send()
		return
	}
	defer file.Close()

	var writeTo io.Writer = writer
	if base64Query != "" {
		base64Encoder := base64.NewEncoder(base64.StdEncoding, writer)
		defer base64Encoder.Close()
		writeTo = base64Encoder
	}

	var responseImage *image.RGBA
	img, format, err := image.Decode(file)
	if err != nil {
		logger.Error().Err(err).Msg("exiting at image.Decode")
		writer.WriteHeader(http.StatusInternalServerError)
		response.Error().Err(err).Send()
		return
	}

	if widthQuery != "" && heightQuery != "" {
		width, err := strconv.Atoi(widthQuery)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response.Error().Err(err).Send()
			return
		}
		height, err := strconv.Atoi(heightQuery)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response.Error().Err(err).Send()
			return
		}
		logger.Debug().Int("width", width).Int("height", height).Msg("resizing")
		responseImage = image.NewRGBA(image.Rect(0, 0, width, height))
		draw.NearestNeighbor.Scale(responseImage, responseImage.Rect, img, img.Bounds(), draw.Over, nil)
	} else {
		responseImage = image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))
		draw.Copy(responseImage, img.Bounds().Min, img, img.Bounds(), draw.Over, nil)
	}

	if base64Query == "" {
		writer.Header().Set("Content-Type", fmt.Sprintf("image/%s", format))
	} else {
		logger.Debug().Msg("encoding as base64")
		writer.Header().Set("Content-Type", "text/plain")
	}
	switch format {
	case "png":
		err = png.Encode(writeTo, responseImage)
	case "jpeg":
		err = jpeg.Encode(writeTo, responseImage, nil)
	}
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		logger.Error().Err(err).Str("format", format).Msg("exiting at image.Encode")
		writer.WriteHeader(http.StatusInternalServerError)
		response.Error().Err(err).Send()
		return
	}
	logger.Info().Str("id", id).Send()
}
