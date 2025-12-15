package models

import (
	"housebook/database"
	"net/http"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int64 `bun:",pk,autoincrement" json:"id"`

	Email string `bun:",unique,notnull" json:"email"`
	Login string `bun:",unique,notnull" json:"login"`
	Name  string `bun:",notnull" json:"name"`

	Salt []byte `bun:",notnull"`
	Hash []byte `bun:",notnull"`
}

type UserRequest struct {
	*User
	Password string `json:"password"`
}

type UserResponse struct {
	*User
}

func (user *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (user *User) Bind(r *http.Request) error {
	return nil
}

func (user *User) Insert(request UserRequest) error {
	user = request.User
	_, err := database.BunDatabase.
		NewInsert().
		Model(user).
		Exec(database.Context)
	return err
}

func (user *User) Select(id int64) error {
	// TODO: add selection by login
	return database.BunDatabase.
		NewSelect().
		Model(user).
		Where("id = ?", id).
		Scan(database.Context)
}

func (user *User) Update(request UserRequest, columns []string) error {
	q := database.BunDatabase.NewUpdate().Model(user)
	for _, column := range columns {
		q = q.Column(column)
	}
	_, err := q.WherePK().Exec(database.Context)
	return err
}

func (user *User) Delete() error {
	_, err := database.BunDatabase.
		NewDelete().
		Model(user).
		WherePK().
		Exec(database.Context)
	return err
}
