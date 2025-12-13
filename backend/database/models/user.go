package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID    int64  `bun:",pk,autoincrement"`
	Login string `bun:",unique,notnull"`
	// Full name
	Name string `bun:",notnull"`
}

// TODO: decide if this struct is really needed
type UserRequest struct {
	Login string
	Name  string
}

func (user *User) Insert(request UserRequest) error {
	user.Login = request.Login
	user.Name = request.Name
	_, err := db.NewInsert().Model(user).Exec(ctx)
	return err
}

// TODO: add selection by login
func (user *User) Select(id int64) error {
	return db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
}

func (user *User) Update(request UserRequest, columns []string) error {
	q := db.NewUpdate().Model(user)
	for _, column := range columns {
		q = q.Column(column)
	}
	_, err := q.WherePK().Exec(ctx)
	return err
}

func (user *User) Delete() error {
	_, err := db.NewDelete().Model(user).WherePK().Exec(ctx)
	return err
}
