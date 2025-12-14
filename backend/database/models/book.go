package models

import (
	"housebook/database"
	"net/http"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Book struct {
	bun.BaseModel `bun:"table:books"`

	UUID  uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()" json:"uuid"`
	Title string    `bun:",notnull" json:"title"`
	// TODO: create a separate Author model
	Author string `bun:",notnull" json:"author"`
	ISBN   string `bun:",unique,notnull" json:"isbn"`
}

type BookRequest struct {
	*Book
}

type BookResponse struct {
	*Book
}

func (book *Book) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (book *Book) Bind(r *http.Request) error {
	return nil
}

func (book *Book) Insert(request BookRequest) error {
	book.Title = request.Title
	book.Author = request.Author
	book.ISBN = request.ISBN
	_, err := database.BunDatabase.
		NewInsert().
		Model(book).
		Exec(database.Context)
	return err
}

func (book *Book) Select(id uuid.UUID) error {
	// TODO: add selection by ISBN
	return database.BunDatabase.
		NewSelect().
		Model(book).
		Where("uuid = ?", id).
		Scan(database.Context)
}

func (book *Book) Update(request BookRequest) error {
	q := database.BunDatabase.NewUpdate().Model(book)
	// TODO: clean up
	if request.Title != "" {
		book.Title = request.Title
		q = q.Column("title")
	}
	if request.Author != "" {
		book.Author = request.Author
		q = q.Column("author")
	}
	if request.ISBN != "" {
		book.ISBN = request.ISBN
		q = q.Column("isbn")
	}
	_, err := q.WherePK().Exec(database.Context)
	return err
}

func (book *Book) Delete() error {
	_, err := database.BunDatabase.
		NewDelete().
		Model(book).
		WherePK().
		Exec(database.Context)
	return err
}
