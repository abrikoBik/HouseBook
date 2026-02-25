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
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	stream, header, err := request.FormFile(consts.Key)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer stream.Close()
	// who knows what might happen
	if header.Size > consts.MaxSize {
		log.Println("file is too big")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("receiving", header.Filename)
	id := uuid.New()
	path := filepath.Join(consts.StoragePath, id.String())
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	written, err := io.Copy(file, stream)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(written, "bytes written")
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(id.String()))
}
