-- name: GetBook :one
select * from books
where id = $1
limit 1;

-- name: ListBooks :many
select * from books;

-- name: CreateBook :one
insert into books (
    title, author, isbn, genre, publisher, year, language, cover_type, page_count, description
) values (
    $1,    $2,     $3,   $4,    $5,        $6,   $7,       $8,         $9,         $10
)
returning *;

-- name: UpdateBook :one
update books
set
    title = $2,
    author = $3,
    isbn = $4,
    genre = $5,
    publisher = $6,
    year = $7,
    language = $8,
    cover_type = $9,
    page_count = $10,
    description = $11
where id = $1
returning *;

-- name: DeleteBook :exec
delete from books
where id = $1;
