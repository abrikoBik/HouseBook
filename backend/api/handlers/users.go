package handlers

import (
	"encoding/json"
	"errors"
	"housebook/consts"
	"housebook/database"
	"housebook/database/db"
	"housebook/util"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/unrolled/render"
)

type UserRequest struct {
	db.CreateUserParams
	Password string
}

// CreateUser godoc
//
//	@summary		Создать пользователя
//	@description	Создаёт нового пользователя.
//	@accept			json
//	@produce		json
//
//	@tags			Пользователи
//
//	@param			email		body		string	true	"Электронная почта пользователя"	Example(ivanov@example.com)
//	@param			login		body		string	true	"Логин пользователя"				Example(ivanov)
//	@param			name		body		string	true	"ФИО пользователя"					Example(Иванов Иван Иванович)
//	@param			password	body		string	true	"Пароль пользователя"				Example(ivanov123qwerty)
//
//	@success		201			{object}	db.User
//	@failure		400			{object}	ErrorResponse
//	@failure		409			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/users [post]
func CreateUser(writer http.ResponseWriter, request *http.Request) {
	var (
		err         error
		requestData UserRequest
		data        db.CreateUserParams
		r           = render.New()
		q           = db.New(database.Pool)
	)
	if err := json.NewDecoder(request.Body).Decode(&requestData); err != nil {
		RenderError(r, writer, ErrorBadJSON(err))
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
		var (
			pgErr   *pgconn.PgError
			toWrite ErrorResponse
		)
		errors.As(err, &pgErr)
		if pgErr.Code == "23505" {
			toWrite = ErrorDuplicate(err, pgErr)
		} else {
			toWrite = ErrorGeneric(err)
		}
		RenderError(r, writer, toWrite)
		return
	}
	Logger.Info().Int64("id", user.ID).Msg("user created")

	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(user); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// GetUser godoc
//
//	@summary		Получить пользователя
//	@description	Получает информацию о пользователе.
//	@accept			plain
//	@produce		json
//
//	@tags			Пользователи
//
//	@success		200	{object}	db.User
//	@failure		400	{object}	ErrorResponse
//	@failure		404	{object}	ErrorResponse
//	@failure		500	{object}	ErrorResponse
//	@router			/api/users/{id} [get]
func GetUser(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	id, err := util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	user, err := q.GetUser(database.Context, id)
	if err != nil {
		var toWrite ErrorResponse
		if errors.Is(err, pgx.ErrNoRows) {
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

// UpdateUser godoc
//
//	@summary		Изменить пользователя
//	@description	Изменяет информацию о существующем пользователе.
//	@accept			json
//	@produce		json
//
//	@tags			Пользователи
//
//	@param			email		body		string	false	"Электронная почта пользователя"	Example(ivanov@example.com)
//	@param			login		body		string	false	"Логин пользователя"				Example(ivanov)
//	@param			name		body		string	false	"ФИО пользователя"					Example(Иванов Иван Иванович)
//	@param			password	body		string	false	"Пароль пользователя"				Example(ivanov123qwerty)
//
//	@success		200			{object}	db.User
//	@failure		400			{object}	ErrorResponse
//	@failure		409			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/users/{id} [patch]
func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.UpdateUserParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	data.ID, err = util.PathID(consts.IDPathValue, request)
	if err = errors.Join(err, json.NewDecoder(request.Body).Decode(&data)); err != nil {
		RenderError(r, writer, ErrorBadJSON(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	user, err := q.UpdateUser(database.Context, data)
	if err != nil {
		var (
			pgErr   *pgconn.PgError
			toWrite ErrorResponse
		)
		errors.As(err, &pgErr)
		if pgErr.Code == "23505" {
			toWrite = ErrorDuplicate(err, pgErr)
		} else {
			toWrite = ErrorGeneric(err)
		}
		RenderError(r, writer, toWrite)
		return
	}
	Logger.Info().Int64("id", user.ID).Msg("user updated")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(user); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// SetProfileImage godoc
//
//	@summary		Задать аватар
//	@description	Задаёт аватар пользователя.
//	@accept			mpfd
//	@produce		json
//
//	@tags			Пользователи
//
//	@param			image	formData	file	true	"Файл аватара (PNG или JPEG)"
//
//	@success		200		{string}	string	""
//	@failure		400		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/users/{id}/image [put]
func SetProfileImage(writer http.ResponseWriter, request *http.Request) {
	// TODO: rename all functions so their names match those generated by sqlc
	var (
		err  error
		data db.SetProfilePictureParams
		r    = render.New()
		q    = db.New(database.Pool)
	)

	if data.ID, err = util.PathID(consts.IDPathValue, request); err != nil {
		RenderError(r, writer, ErrorBadJSON(err))
		return
	}

	var client http.Client
	req, err := http.NewRequest("POST", consts.StorageURL, nil)
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	req.Body = request.Body
	req.Header.Set("Content-Type", request.Header.Get("Content-Type"))
	response, err := client.Do(req)
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	defer response.Body.Close()

	var errored bool
	switch response.StatusCode {
	case http.StatusCreated:
		idJSON := StorageIDResponse{}
		if err := json.NewDecoder(response.Body).Decode(&idJSON); err != nil {
			RenderError(r, writer, ErrorGeneric(err))
			errored = true
			return
		}
		if err := data.Imageuuid.Scan(idJSON.ID); err != nil {
			RenderError(r, writer, ErrorGeneric(err))
			errored = true
		}
	default:
		errorJSON := StorageErrorResponse{}
		if err := json.NewDecoder(response.Body).Decode(&errorJSON); err != nil {
			RenderError(r, writer, ErrorGeneric(err))
			errored = true
			return
		}
		RenderError(r, writer, ErrorImage(errorJSON.Error))
		errored = true
	}
	if errored {
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	if err := q.SetProfilePicture(database.Context, data); err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().
		Int64("id", data.ID).
		Str("imageUUID", data.Imageuuid.String()).
		Msg("profile picture set")

	writer.WriteHeader(http.StatusOK)
}

// DeleteUser godoc
//
//	@summary		Удалить пользователя
//	@description	Удаляет пользователя из базы данных.
//	@accept			plain
//	@produce		json
//
//	@tags			Пользователи
//
//	@success		200	{string}	string	""
//	@failure		400	{object}	ErrorResponse
//	@failure		500	{object}	ErrorResponse
//	@router			/api/users/{id} [delete]
func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	id, err := util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
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
