package handlers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"housebook_storage/consts"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"golang.org/x/image/draw"
)

func SendFile(writer http.ResponseWriter, request *http.Request) {
	var (
		base64Query = request.URL.Query().Get("base64")
		widthQuery  = request.URL.Query().Get("width")
		heightQuery = request.URL.Query().Get("height")
		err         error
	)
	id := request.PathValue("id")
	if err := uuid.Validate(id); err != nil {
		RenderError(writer, Error{http.StatusBadRequest, err, ""})
		return
	}
	path := filepath.Join(consts.StoragePath, id)
	logger.Debug().Str("path", path).Send()
	file, err := os.Open(path)
	if err != nil {
		var status int
		if os.IsNotExist(err) {
			status = http.StatusNotFound
		} else {
			status = http.StatusInternalServerError
		}
		RenderError(writer, Error{status, err, ""})
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
		RenderError(writer, Error{http.StatusInternalServerError, err, ""})
		return
	}
	if widthQuery != "" && heightQuery != "" {
		width, werr := strconv.Atoi(widthQuery)
		height, herr := strconv.Atoi(heightQuery)
		if err := errors.Join(werr, herr); err != nil {
			RenderError(writer, Error{http.StatusBadRequest, err, ""})
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
		writer.Header().Set("Cache-Control", "public, max-age=86400, immutable")
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
		RenderError(writer, Error{http.StatusInternalServerError, err, ""})
		return
	}
	logger.Info().Str("id", id).Send()
}
