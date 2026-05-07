package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"housebook/consts"
	"housebook/database"
	"housebook/database/db"
	"housebook/util"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/unrolled/render"
)

type SessionRequest struct {
	db.CreateSessionParams
	Login    string
	Password string
}

// CreateSession godoc
//
//	@summary		Создать сессию
//	@description	Создаёт новую сессию (вход в аккаунт).
//	@accept			json
//	@produce		json
//
//	@tags			Сессии
//
//	@param			name		body		string	true	"Название сессии"		Example(Браузер на ПК)
//	@param			login		body		string	true	"Логин пользователя"	Example(ivanov)
//	@param			password	body		string	true	"Пароль пользователя"	Example(ivanov123qwerty)
//
//	@success		201			{object}	db.Session
//	@failure		400			{object}	ErrorResponse
//	@failure		401			{object}	ErrorResponse
//	@failure		404			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/users/{user-id}/sessions [post]
func CreateSession(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data SessionRequest
		r    = render.New()
		q    = db.New(database.Pool)
	)
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		RenderError(r, writer, ErrorBadJSON(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}

	user, err := q.GetUserByLogin(database.Context, data.Login)
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
	Logger.Info().Interface("userid", user.ID).Msg("user retrieved")

	hash, err := util.GenerateHash([]byte(data.Password), user.Salt)
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	if !bytes.Equal(hash, user.Hash) {
		RenderError(r, writer, ErrorWrongPassword(errors.New("hash mismatch")))
		return
	}

	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  data.Login,
		"iat":  jwt.NewNumericDate(now),
		"exp":  jwt.NewNumericDate(now.Add(time.Hour * 24 * 7 * 5)),
		"hash": hash,
		"name": data.Name,
	})
	data.Login = user.Login
	data.Userid = user.ID
	data.Jwt, err = token.SignedString([]byte(util.Getenv("JWT_KEY")))
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Debug().Interface("data", data).Msg("data set")
	session, err := q.CreateSession(database.Context, data.CreateSessionParams)
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Interface("id", session.ID).Msg("session created")

	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(session); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// GetSession godoc
//
//	@summary		Получить сессию
//	@description	Возвращает информацию о сессии.
//	@accept			plain
//	@produce		json
//
//	@tags			Сессии
//
//	@param			session-id	path		pgtype.UUID	true	"UUID сессии"
//
//	@success		200			{object}	db.GetSessionByIDRow
//	@failure		400			{object}	ErrorResponse
//	@failure		404			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/sessions/{session-id} [get]
func GetSession(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	uuid, err := util.PathUUID(consts.SessionIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.SessionIDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	session, err := q.GetSessionByID(database.Context, uuid)
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
	Logger.Info().Interface("id", session.ID).Msg("session retrieved")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(session); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// UpdateSession godoc
//
//	@summary		Переименовать сессию
//	@description	Изменяет название сессии.
//	@accept			json
//	@produce		json
//
//	@tags			Сессии
//
//	@param			session-id	path		pgtype.UUID	true	"UUID сессии"
//	@param			name		body		string		true	"Название сессии"	Example(Браузер на ПК)
//
//	@success		200			{object}	db.RenameSessionByIDRow
//	@failure		400			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/sessions/{session-id} [patch]
func UpdateSession(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.RenameSessionByIDParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	data.ID, err = util.PathUUID(consts.SessionIDPathValue, request)
	if err = errors.Join(err, json.NewDecoder(request.Body).Decode(&data)); err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.SessionIDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	session, err := q.RenameSessionByID(database.Context, data)
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Interface("id", session.ID).Msg("session updated")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(session); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// DeleteSession godoc
//
//	@summary		Удалить сессию
//	@description	Удаляет сессию из базы данных.
//	@accept			plain
//	@produce		json
//
//	@tags			Сессии
//
//	@param			session-id	path		pgtype.UUID	true	"UUID сессии"
//
//	@success		200			{string}	string		""
//	@failure		400			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/sessions/{session-id} [delete]
func DeleteSession(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	uuid, err := util.PathUUID(consts.SessionIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.SessionIDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	if err := q.DeleteSessionByID(database.Context, uuid); err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Interface("id", uuid).Msg("session deleted")

	writer.WriteHeader(http.StatusOK)
}

// ListUserSessions godoc
//
//	@summary		Список сессий пользователя
//	@description	Возвращает список всех сессий пользователя.
//	@accept			plain
//	@produce		json
//
//	@tags			Сессии,Пользователи
//
//	@param			user-id	path		pgtype.UUID	true	"UUID пользователя"
//	@param			limit	query		int			false	"Лимит количества возвращаемых элементов"	Example(10)
//	@param			offset	query		int			false	"Смещение возвращаемых элементов"			Example(5)
//
//	@success		200		{array}		db.ListUserSessionsRow
//	@failure		400		{object}	ErrorResponse
//	@failure		404		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/users/{user-id}/sessions [get]
func ListUserSessions(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.ListUserSessionsParams
		query  = request.URL.Query()
		r      = render.New()
		q      = db.New(database.Pool)
	)
	if params.Userid, err = util.PathUUID(consts.UserIDPathValue, request); err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.UserIDPathValue))
		return
	}
	if params.Limit, params.Offset, err = util.GetSearchPagination(query); err != nil {
		RenderError(r, writer, ErrorBadPaginationParams(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	sessions, err := q.ListUserSessions(database.Context, params)
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
	Logger.Info().Int("len(sessions)", len(sessions)).Msg("sessions listed")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(sessions); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}
