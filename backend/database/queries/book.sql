-- name: GetBook :one
select * from books
where id = $1
limit 1;

-- name: ListBooks :many
select * from books
limit $1
offset $2;

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
    title = coalesce($2, title),
    author = coalesce($3, author),
    isbn = coalesce($4, isbn),
    genre = coalesce($5, genre),
    publisher = coalesce($6, publisher),
    year = coalesce($7, year),
    language = coalesce($8, language),
    cover_type = coalesce($9, cover_type),
    page_count = coalesce($10, page_count),
    description = coalesce($11, description)
where id = $1
returning *;

-- name: DeleteBook :exec
delete from books
where id = $1;
