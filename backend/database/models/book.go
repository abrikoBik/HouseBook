package models

import (
	"database/sql"
	"fmt"
	"housebook/database"
	"housebook/util"
	"net/http"

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

	Year        sql.NullString `json:"year"`
	Language    sql.NullString `json:"language"`
	CoverType   sql.NullBool   `json:"coverType"`
	Color       sql.NullBool   `json:"color"`
	PageCount   sql.NullInt32  `json:"pageCount"`
	Description sql.NullString `json:"description"`
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

func SearchBooks(params util.SearchTable) (books []*Book, err error) {
	q := database.BunDatabase.NewSelect().Model(&books)
	for n, v := range params {
		if v != "" {
			q = q.Where(fmt.Sprintf("%s = ?", n), v)
		}
	}
	err = q.Scan(database.Context)
	return
}

func ListBooks(limit, offset int) (books []*Book, err error) {
	err = database.BunDatabase.
		NewSelect().
		Model(&books).
		Limit(limit).
		Offset(offset).
		Scan(database.Context)
	return
}
