package handlers

import (
	"encoding/json"
	"errors"
	"housebook/consts"
	"housebook/database"
	"housebook/database/db"
	"housebook/util"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/unrolled/render"
)

func CreateLibrary(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.CreateLibraryParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	userID, err := util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
		return
	}
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		RenderError(r, writer, ErrorBadJSON(err))
		return
	}
	data.Userid = userID

	if PingDB(r, writer) != nil {
		return
	}
	library, err := q.CreateLibrary(database.Context, data)
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Interface("id", library.ID).Msg("library created")

	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(library); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

func GetLibrary(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.GetLibraryParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	// TODO: clean up
	userID, err := util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
		return
	}
	data.Userid = userID
	libraryUUID, err := util.PathUUID("library-uuid", request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, "library-uuid"))
		return
	}
	data.ID = libraryUUID

	if PingDB(r, writer) != nil {
		return
	}
	library, err := q.GetLibrary(database.Context, data)
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
	Logger.Info().Interface("id", library.ID).Msg("library retrieved")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(library); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

func DeleteLibrary(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.DeleteLibraryParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	userID, err := util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
		return
	}
	data.Userid = userID
	libraryUUID, err := util.PathUUID("library-uuid", request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, "library-uuid"))
		return
	}
	data.ID = libraryUUID

	if PingDB(r, writer) != nil {
		return
	}
	if err := q.DeleteLibrary(database.Context, data); err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Interface("id", libraryUUID).Msg("library deleted")

	writer.WriteHeader(http.StatusOK)
}

func ListLibraries(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.ListLibrariesParams
		query  = request.URL.Query()
		r      = render.New()
		q      = db.New(database.Pool)
	)
	userID, err := util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
		return
	}
	params.Userid = userID
	params.Limit, params.Offset, err = util.GetSearchPagination(query)
	if err != nil {
		RenderError(r, writer, ErrorBadPaginationParams(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	libraries, err := q.ListLibraries(database.Context, params)
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
	Logger.Info().Int("len(libraries)", len(libraries)).Msg("libraries listed")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(libraries); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

func AddBookToLibrary(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.AddBookToLibraryParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		RenderError(r, writer, ErrorBadJSON(err))
		return
	}
	_, err = util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
		return
	}
	libraryUUID, err := util.PathUUID("library-uuid", request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, "library-uuid"))
		return
	}
	data.Libraryid = libraryUUID

	if PingDB(r, writer) != nil {
		return
	}
	if err = q.AddBookToLibrary(database.Context, data); err != nil {
		var (
			pgErr   *pgconn.PgError
			toWrite ErrorResponse
		)
		errors.As(err, &pgErr)
		if pgErr.Code == "23505" {
			toWrite = ErrorDuplicate(err, pgErr)
		} else {
			toWrite = ErrorGeneric(err)
		}
		RenderError(r, writer, toWrite)
		return
	}
	Logger.Info().
		Interface("libraryID", data.Libraryid).
		Interface("bookID", data.Bookid).
		Msg("book added to library")

	writer.WriteHeader(http.StatusCreated)
}

func ListLibraryBooks(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.ListLibraryBooksParams
		query  = request.URL.Query()
		r      = render.New()
		q      = db.New(database.Pool)
	)
	userID, err := util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
		return
	}
	params.Userid = userID
	libraryUUID, err := util.PathUUID("library-uuid", request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, "library-uuid"))
		return
	}
	params.ID = libraryUUID
	params.Limit, params.Offset, err = util.GetSearchPagination(query)
	if err != nil {
		RenderError(r, writer, ErrorBadPaginationParams(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	books, err := q.ListLibraryBooks(database.Context, params)
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
	Logger.Info().Int("len(books)", len(books)).Msg("books listed")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(books); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

func DeleteBookFromLibrary(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.DeleteBookFromLibraryParams
		r      = render.New()
		q      = db.New(database.Pool)
	)

	_, err = util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
		return
	}
	libraryUUID, err := util.PathUUID("library-uuid", request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, "library-uuid"))
		return
	}
	params.Libraryid = libraryUUID
	bookID, err := util.PathUUID("book-uuid", request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, "book-uuid"))
		return
	}
	params.Bookid = bookID

	if PingDB(r, writer) != nil {
		return
	}
	if err := q.DeleteBookFromLibrary(database.Context, params); err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().
		Interface("libraryID", libraryUUID).
		Interface("bookID", bookID).
		Msg("book deleted from library")

	writer.WriteHeader(http.StatusOK)
}
