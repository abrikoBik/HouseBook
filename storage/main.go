package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const (
	hostname = "0.0.0.0"
	port     = 3001
	// About 128MB.
	maxSize     = 1 << 27
	key         = "file"
	storagePath = "/var/lib/storage"
)

func ReceiveFile(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseMultipartForm(maxSize); err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	stream, header, err := request.FormFile(key)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer stream.Close()
	// who knows what might happen
	if header.Size > maxSize {
		log.Println("file is too big")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("receiving", header.Filename)
	id := uuid.New()
	path := filepath.Join(storagePath, id.String())
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

func SendFile(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")
	if err := uuid.Validate(id); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	path := filepath.Join(storagePath, id)
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

func main() {
	ctx := context.Background()
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		log.Panicln(err.Error())
	}
	defer listener.Close()
	router := http.NewServeMux()
	router.HandleFunc("POST /", ReceiveFile)
	router.HandleFunc("GET /{id}", SendFile)
	server := http.Server{
		Addr:    listener.Addr().String(),
		Handler: router,
	}
	log.Printf("listening at %s:%d\n", hostname, port)
	if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
		log.Panicln(err.Error())
	}
	defer server.Shutdown(ctx)
}
