package handlers

import (
	"encoding/json"
	"housebook/database"
	"housebook/database/db"
	"log"

	"net/http"

	"github.com/unrolled/render"
)

// func bookResponse(book *models.Book) *models.Book {
// 	return &models.BookResponse{Book: book}
// }

// func bookListResponse(books []*models.Book) (result []render.Renderer) {
// 	for _, book := range books {
// 		result = append(result, bookResponse(book))
// 	}
// 	return
// }

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	// TODO: create a small package for all rendering-related stuff
	var (
		data = db.CreateBookParams{}
		r    = render.New()
		q    = db.New(database.Pool)
	)
	if err := database.Pool.Ping(database.Context); err != nil {
		log.Println(err)
		toWrite := ErrorGeneric(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		log.Println(err)
		toWrite := ErrorBadRequest(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
	log.Println(data)
	book, err := q.CreateBook(database.Context, data)
	if err != nil {
		log.Println(err)
		toWrite := ErrorGeneric(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
	log.Println(book)
	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(book); err != nil {
		log.Println(err)
		toWrite := ErrorRender(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	// var (
	// 	book    = &models.Book{}
	// 	id, err = util.PathUUID(consts.UUIDPathValue, r)
	// )
	// if err != nil {
	// 	render.Render(w, r, ErrorBadRequest(err))
	// 	return
	// }
	// book.UUID = *id
	// if err := book.Select(); err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		render.Render(w, r, ErrorNotFound(err))
	// 		return
	// 	}
	// 	render.Render(w, r, ErrorGeneric(err))
	// 	return
	// }
	// render.Status(r, http.StatusOK)
	// if err := render.Render(w, r, bookResponse(book)); err != nil {
	// 	render.Render(w, r, ErrorRender(err))
	// 	return
	// }
	w.WriteHeader(http.StatusNotImplemented)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// var (
	// 	book    = &models.Book{}
	// 	data    = &models.BookRequest{}
	// 	id, err = util.PathUUID(consts.UUIDPathValue, r)
	// )
	// if err != nil {
	// 	render.Render(w, r, ErrorBadRequest(err))
	// 	return
	// }
	// book.UUID = *id
	// if err := book.Select(); err != nil {
	// 	render.Render(w, r, ErrorGeneric(err))
	// 	return
	// }
	// if err := render.Bind(r, data); err != nil {
	// 	render.Render(w, r, ErrorBadRequest(err))
	// 	return
	// }
	// if err := book.Update(*data); err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		render.Render(w, r, ErrorNotFound(err))
	// 		return
	// 	}
	// 	render.Render(w, r, ErrorGeneric(err))
	// 	return
	// }
	// render.Status(r, http.StatusOK)
	// if err := render.Render(w, r, bookResponse(book)); err != nil {
	// 	render.Render(w, r, ErrorRender(err))
	// 	return
	// }
	w.WriteHeader(http.StatusNotImplemented)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// var (
	// 	book    = &models.Book{}
	// 	id, err = util.PathUUID(consts.UUIDPathValue, r)
	// )
	// if err != nil {
	// 	render.Render(w, r, ErrorBadRequest(err))
	// 	return
	// }
	// book.UUID = *id
	// if err := book.Delete(); err != nil {
	// 	render.Render(w, r, ErrorGeneric(err))
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusNotImplemented)
}

func SearchBook(w http.ResponseWriter, r *http.Request) {
	// var (
	// 	params = util.SearchTable{}
	// 	query  = r.URL.Query()
	// )
	// for _, param := range consts.BookSearchParams {
	// 	params[param] = query.Get(param)
	// }
	// books, err := models.Search[models.Book](params)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		render.Render(w, r, ErrorNotFound(err))
	// 		return
	// 	}
	// 	render.Render(w, r, ErrorGeneric(err))
	// 	return
	// }
	// render.Status(r, http.StatusOK)
	// if err := render.RenderList(w, r, bookListResponse(books)); err != nil {
	// 	render.Render(w, r, ErrorRender(err))
	// 	return
	// }
	w.WriteHeader(http.StatusNotImplemented)
}

func ListBooks(w http.ResponseWriter, r *http.Request) {
	// var (
	// 	query              = r.URL.Query()
	// 	limit, offset, err = util.SearchParams(query)
	// )
	// if err != nil {
	// 	render.Render(w, r, ErrorBadRequest(err))
	// 	return
	// }
	// books, err := models.List[models.Book](limit, offset)
	// if err != nil {
	// 	// TODO: decide if 404 is applicable here
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		render.Render(w, r, ErrorNotFound(err))
	// 		return
	// 	}
	// 	render.Render(w, r, ErrorGeneric(err))
	// 	return
	// }
	// render.Status(r, http.StatusOK)
	// if err := render.RenderList(w, r, bookListResponse(books)); err != nil {
	// 	render.Render(w, r, ErrorRender(err))
	// 	return
	// }
	w.WriteHeader(http.StatusNotImplemented)
}
