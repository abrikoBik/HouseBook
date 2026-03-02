-- name: GetLibrary :one
select libraries.* from libraries
join users on libraries.user_id = users.id
where libraries.user_id = @id
limit 1;

-- name: ListLibraryBooks :many
select books.id from libraries
join users on libraries.user_id = users.id
left join library_books on libraries.id = library_books.library_id
left join books on library_books.book_id = books.id
where libraries.id = @id
limit 1;

-- name: ListLibraries :many
select * from libraries
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
    book_id,
    quantity
) values (
    @libraryID,
    @bookID,
    @quantity
);

-- name: DeleteBookFromLibrary :exec
delete from library_books
where
library_id = @libraryID and
book_id = @bookID;
