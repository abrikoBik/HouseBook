package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
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
//	@param			login		body		string	true	"Логин пользователя"				Example(ivanov)
//	@param			password	body		string	true	"Пароль пользователя"				Example(ivanov123qwerty)
//
//	@success		201			{object}	db.Session
//	@failure		400			{object}	ErrorResponse
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
	})
	data.Login = user.Login
	data.Userid = user.ID
	data.Jwt, err = token.SignedString([]byte(os.Getenv("JWT_KEY")))
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
