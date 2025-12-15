package models

import (
	"housebook/database"
	"net/http"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Book struct {
	bun.BaseModel `bun:"table:books"`
	UUID          uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()" json:"uuid"`

	Title     string `bun:",notnull" json:"title"`
	Author    string `bun:",notnull" json:"author"` // TODO: create a separate Author model
	ISBN      string `bun:",unique,notnull" json:"isbn"`
	Genre     string `bun:",notnull" json:"genre"`
	Language  string `bun:",notnull" json:"language"`
	Publisher string `bun:",notnull" json:"publisher"`

	CoverType   string `json:"coverType"`
	Color       bool   `json:"color"`
	PageCount   int    `json:"pageCount"`
	Description string `json:"description"`
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
	book = request.Book
	_, err := database.BunDatabase.
		NewInsert().
		Model(book).
		Exec(database.Context)
	return err
}

func (book *Book) Select() error {
	return database.BunDatabase.
		NewSelect().
		Model(book).
		WherePK().
		Scan(database.Context)
}

func (book *Book) Update(request BookRequest) error {
	book = request.Book
	_, err := database.BunDatabase.
		NewUpdate().
		Model(book).
		Column(
			"title",
			"author",
			"isbn",
			"genre",
			"language",
			"publisher",
			"cover_type",
			"color",
			"page_count",
			"description",
		).
		WherePK().
		Exec(database.Context)
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
