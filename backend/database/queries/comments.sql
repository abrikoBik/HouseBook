-- name: GetComment :one
select * from comments
where id = @id
limit 1;

-- name: ListBookComments :many
select * from comments
where book_id = @bookID
limit $1
offset $2;

-- name: ListUserComments :many
select * from comments
where user_id = @userID
limit $1
offset $2;

-- name: CreateComment :one
insert into comments (
    user_id,
    book_id,
    body
) values (
    @userID,
    @bookID,
    @body
)
returning *;

-- name: DeleteComment :exec
delete from comments
where id = @id;
