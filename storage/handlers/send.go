package handlers

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"storage/consts"

	"github.com/google/uuid"
)

func SendFile(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")
	if err := uuid.Validate(id); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	path := filepath.Join(consts.StoragePath, id)
	log.Println("sending", path)
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		if os.IsNotExist(err) {
			writer.WriteHeader(http.StatusNotFound)
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	defer file.Close()
	encoder := base64.NewEncoder(base64.StdEncoding, writer)
	defer encoder.Close()
	written, err := io.Copy(encoder, file)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(written, "bytes written")
}
