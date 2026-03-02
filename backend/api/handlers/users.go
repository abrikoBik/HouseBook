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
		RenderError(r, writer, ErrorBadRequest(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	requestData.Salt, err = util.RandomBytes(16)
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	requestData.Hash, err = util.GenerateHash([]byte(requestData.Password), requestData.Salt)
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	data = requestData.CreateUserParams
	user, err := q.CreateUser(database.Context, data)
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Int64("id", user.ID).Msg("user created")

	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(user); err != nil {
		RenderError(r, writer, ErrorRender(err))
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
		RenderError(r, writer, ErrorBadRequest(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	user, err := q.GetUser(database.Context, id)
	if err != nil {
		var toWrite ErrorResponse
		if errors.Is(err, sql.ErrNoRows) {
			toWrite = ErrorNotFound(err)
		} else {
			toWrite = ErrorGeneric(err)
		}
		RenderError(r, writer, toWrite)
		return
	}
	Logger.Info().Int64("id", user.ID).Msg("user retrieved")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(user); err != nil {
		RenderError(r, writer, ErrorRender(err))
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
		RenderError(r, writer, ErrorBadRequest(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	user, err := q.UpdateUser(database.Context, data)
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Int64("id", user.ID).Msg("user updated")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(user); err != nil {
		RenderError(r, writer, ErrorRender(err))
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
		RenderError(r, writer, ErrorBadRequest(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	if err := q.DeleteUser(database.Context, id); err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Int64("id", id).Msg("user deleted")

	writer.WriteHeader(http.StatusOK)
}
