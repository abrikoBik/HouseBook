package models

import (
	"housebook/database"
	"housebook/database/generated"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Book struct {
	UUID uuid.UUID `json:"id"`

	Title     string `json:"title"`
	Author    string `json:"author"`
	ISBN      string `json:"isbn"`
	Genre     string `json:"genre"`
	Publisher string `json:"publisher"`

	Year        int    `json:"year,omitempty"`
	Language    string `json:"language,omitempty"`
	CoverType   string `json:"coverType,omitempty"`
	PageCount   string `json:"pageCount,omitempty"`
	Description string `json:"description,omitempty"`
}

type BookRequest struct {
	Book
	UUID uuid.UUID `json:"-"`
}

type BookResponse struct {
	Book
}

func (book *Book) Create(request BookRequest) (generated.Book, error) {
	params := generated.CreateBookParams{
		Title:       request.Title,
		Author:      request.Author,
		Isbn:        request.ISBN,
		Genre:       request.Genre,
		Publisher:   request.Publisher,
		Year:        request.Year,
		Language:    request.Language,
		CoverType:   request.CoverType,
		PageCount:   request.PageCount,
		Description: request.Description,
	}
	return database.Queries.CreateBook(database.Context, params)
}

func (book *Book) Render(w http.ResponseWriter, r *http.Request) error {
	// TODO: add proper error handling
	return nil
}

func (book *BookRequest) Bind(r *http.Request) error {
	// TODO: add proper error handling
	return nil
}
