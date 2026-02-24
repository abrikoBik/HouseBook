package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"housebook/consts"
	"housebook/database"
	"housebook/database/db"
	"housebook/util"
	"net/http"

	"github.com/unrolled/render"
)

// TODO: move UserRequest somewhere appropriate

type UserRequest struct {
	db.CreateUserParams
	Password string
}

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	var (
		err         error
		requestData UserRequest
		data        db.CreateUserParams
		r           = render.New()
		q           = db.New(database.Pool)
	)
	if err := json.NewDecoder(request.Body).Decode(&requestData); err != nil {
		toWrite := ErrorBadRequest(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}

	if err := database.Pool.Ping(database.Context); err != nil {
		toWrite := ErrorDatabase(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
	requestData.Salt, err = util.RandomBytes(16)
	if err != nil {
		toWrite := ErrorGeneric(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
	requestData.Hash, err = util.GenerateHash([]byte(requestData.Password), requestData.Salt)
	if err != nil {
		toWrite := ErrorGeneric(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
	data = requestData.CreateUserParams

	book, err := q.CreateUser(database.Context, data)
	if err != nil {
		var toWrite *ErrorResponse
		switch err {
		default:
			toWrite = ErrorGeneric(err)
		}
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(book); err != nil {
		toWrite := ErrorRender(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	id, err := util.PathID(consts.IDPathValue, request)
	if err != nil {
		toWrite := ErrorBadRequest(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}

	if err := database.Pool.Ping(database.Context); err != nil {
		toWrite := ErrorDatabase(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
	user, err := q.GetUser(database.Context, id)
	if err != nil {
		var toWrite *ErrorResponse
		switch err {
		case sql.ErrNoRows:
			toWrite = ErrorNotFound(err)
		default:
			toWrite = ErrorGeneric(err)
		}
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(user); err != nil {
		toWrite := ErrorRender(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.UpdateUserParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	data.ID, err = util.PathID(consts.IDPathValue, request)
	if err = errors.Join(err, json.NewDecoder(request.Body).Decode(&data)); err != nil {
		toWrite := ErrorBadRequest(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}

	if err := database.Pool.Ping(database.Context); err != nil {
		toWrite := ErrorDatabase(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
	book, err := q.UpdateUser(database.Context, data)
	if err != nil {
		toWrite := ErrorGeneric(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(book); err != nil {
		toWrite := ErrorRender(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	id, err := util.PathID(consts.IDPathValue, request)
	if err != nil {
		toWrite := ErrorBadRequest(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}

	if err := database.Pool.Ping(database.Context); err != nil {
		toWrite := ErrorDatabase(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
	if err := q.DeleteUser(database.Context, id); err != nil {
		var toWrite *ErrorResponse
		switch err {
		default:
			toWrite = ErrorGeneric(err)
		}
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
