package models

import (
	"housebook/database"
	"net/http"

	"github.com/emvi/null"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Book struct {
	bun.BaseModel `bun:"table:books"`
	UUID          uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()" json:"uuid"`

	Title     string `bun:",notnull" json:"title"`
	Author    string `bun:",notnull" json:"author"`
	ISBN      string `bun:",unique,notnull" json:"isbn"`
	Genre     string `bun:",notnull" json:"genre"`
	Publisher string `bun:",notnull" json:"publisher"`

	Year        null.Int32  `json:"year"`
	Language    null.String `json:"language"`
	CoverType   null.Bool   `json:"coverType"`
	PageCount   null.Int32  `json:"pageCount"`
	Description null.String `json:"description"`
}

type BookRequest struct {
	*Book
}

type BookResponse struct {
	*Book
}

func (book *Book) Render(w http.ResponseWriter, r *http.Request) error {
	// TODO: add proper error handling
	return nil
}

func (book *BookRequest) Bind(r *http.Request) error {
	// TODO: add proper error handling
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
	// TODO: reimplement
	// book = request.Book
	// _, err := database.BunDatabase.
	// 	NewUpdate().
	// 	Model(book).
	// 	Column(
	// 		"title",
	// 		"author",
	// 		"isbn",
	// 		"genre",
	// 		"publisher",
	// 		"year",
	// 		"language",
	// 		"cover_type",
	// 		"color",
	// 		"page_count",
	// 		"description",
	// 	).
	// 	WherePK().
	// 	Exec(database.Context)
	// return err
	return nil
}

func (book *Book) Delete() error {
	_, err := database.BunDatabase.
		NewDelete().
		Model(book).
		WherePK().
		Exec(database.Context)
	return err
}
