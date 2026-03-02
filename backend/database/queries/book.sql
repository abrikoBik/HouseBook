-- name: GetBook :one
select * from books
where id = @id
limit 1;

-- name: ListBooks :many
select * from books
limit $1
offset $2;

-- name: CreateBook :one
insert into books (
    title,
    author,
    isbn,
    genre,
    publisher,
    year,
    language,
    cover_type,
    page_count,
    description
) values (
    @title,
    @author,
    @isbn,
    @genre,
    @publisher,
    @year,
    @language,
    @coverType,
    @pageCount,
    @description
)
returning *;

-- name: UpdateBook :one
update books
set
    title = coalesce(sqlc.narg('title'), title),
    author = coalesce(sqlc.narg('author'), author),
    isbn = coalesce(sqlc.narg('isbn'), isbn),
    genre = coalesce(sqlc.narg('genre'), genre),
    publisher = coalesce(sqlc.narg('publisher'), publisher),
    year = coalesce(sqlc.narg('year'), year),
    language = coalesce(sqlc.narg('language'), language),
    cover_type = coalesce(sqlc.narg('coverType'), cover_type),
    page_count = coalesce(sqlc.narg('pageCount'), page_count),
    description = coalesce(sqlc.narg('description'), description)
where id = @id
returning *;

-- name: DeleteBook :exec
delete from books
where id = @id;

-- name: SearchBooks :many
select * from books
where
    (title = sqlc.narg('title') or sqlc.narg('title') = '') and
    (author = sqlc.narg('author') or sqlc.narg('author') = '') and
    (genre = sqlc.narg('genre') or sqlc.narg('genre') = '') and
    (publisher = sqlc.narg('publisher') or sqlc.narg('publisher') = '') and
    (year = sqlc.narg('year') or sqlc.narg('year') = 0) and
    (language = sqlc.narg('language') or sqlc.narg('language') = '')
limit $1
offset $2;
