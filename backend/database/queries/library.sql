-- name: GetLibrary :one
select
    libraries.*,
    count(library_books.book_id)::int as book_count
from libraries
join users on libraries.user_id = users.id
left join library_books on libraries.id = library_books.library_id
where libraries.id = @id
group by libraries.id
limit 1;

-- name: ListLibraryBooks :many
select books.id, library_books.exchangeable from libraries
join users on libraries.user_id = users.id
left join library_books on libraries.id = library_books.library_id
left join books on library_books.book_id = books.id
where libraries.id = @id
limit $1
offset $2;

-- name: ListLibraries :many
select * from libraries
limit $1
offset $2;

-- name: ListUserLibraries :many
select * from libraries
where libraries.user_id = @userID
limit $1
offset $2;

-- name: CreateLibrary :one
insert into libraries (
    user_id,
    name
) values (
    @userID,
    @name
)
returning *;

-- name: DeleteLibrary :exec
delete from libraries
where id = @id;

-- name: AddBookToLibrary :exec
insert into library_books (
    library_id,
    book_id
) values (
    @libraryID,
    @bookID
);

-- name: DeleteBookFromLibrary :exec
delete from library_books
where
library_id = @libraryID and
book_id = @bookID;

-- name: ChangeBookExchangeable :one
update library_books
set
    exchangeable = not exchangeable
where
    library_id = @libraryID and
    book_id = @bookID
returning *;
