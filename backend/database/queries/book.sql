-- name: GetBook :one
select * from books
where id = @id
limit 1;

-- name: ListBooks :many
select * from books b
order by b.added_at desc
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
    (lower(title) = lower(sqlc.narg('title')) or sqlc.narg('title') = '') and
    (lower(author) = lower(sqlc.narg('author')) or sqlc.narg('author') = '') and
    (lower(genre) = lower(sqlc.narg('genre')) or sqlc.narg('genre') = '') and
    (lower(publisher) = lower(sqlc.narg('publisher')) or sqlc.narg('publisher') = '') and
    (year = sqlc.narg('year') or sqlc.narg('year') = 0) and
    (lower(language) = lower(sqlc.narg('language')) or sqlc.narg('language') = '')
limit $1
offset $2;

-- name: SetCoverImage :one
update books
set image_uuid = @imageUUID
where id = @id
returning *;

-- name: SearchBooksCombined :many
select * from books
where
    (lower(title) like lower('%' || @query || '%')) or
    (lower(author) like lower('%' || @query || '%')) or
    (lower(publisher) like lower('%' || @query || '%'))
limit $1
offset $2;

-- name: SearchExchangeableBooks :many
select
    u.id as user_id,
    l.id as library_id
from library_books lb
    join libraries l on lb.library_id = l.id
    join users u on l.user_id = u.id
where
    lb.exchangeable = true and
    lb.book_id = @bookID
limit $1
offset $2;

-- name: GetRecommendedBooks :many
select * from books
order by random()
limit $1
offset $2;
