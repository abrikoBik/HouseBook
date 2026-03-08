package consts

import "fmt"

const (
	DocsEndpoint      = "/docs"
	BooksEndpoint     = "/books"
	UsersEndpoint     = "/users"
	TradesEndpoint    = "/trades"
	LibrariesEndpoint = "/libraries"
	CommentsEndpoint  = "/comments"

	GenericIDPathValue = "id"
	BookIDPathValue    = "book-id"
	CommentIDPathValue = "comment-id"
	LibraryIDPathValue = "library-id"
	TradeIDPathValue   = "trade-id"
	UserIDPathValue    = "user-id"
)

var (
	GenericIDPath = fmt.Sprintf("/{%s}", GenericIDPathValue)
	BookIDPath    = fmt.Sprintf("/{%s}", BookIDPathValue)
	CommentIDPath = fmt.Sprintf("/{%s}", CommentIDPathValue)
	LibraryIDPath = fmt.Sprintf("/{%s}", LibraryIDPathValue)
	TradeIDPath   = fmt.Sprintf("/{%s}", TradeIDPathValue)
	UserIDPath    = fmt.Sprintf("/{%s}", UserIDPathValue)
)
