-- name: GetSessionByID :one
select * from sessions
where id = @id
limit 1;

-- name: GetSessionByToken :one
select * from sessions
where jwt = @jwt
limit 1;

-- name: GetUserSessions :many
select * from sessions
where user_id = @userID
limit $1
offset $2;

-- name: CreateSession :one
insert into sessions (
    name,
    jwt,
    user_id
) values (
    @name,
    @jwt,
    @userID
)
returning *;

-- name: RenameSessionByID :one
update sessions
set name = @name
where id = @id
returning *;

-- name: RenameSessionByToken :one
update sessions
set name = @name
where jwt = @jwt
returning *;

-- name: DeleteSessionByID :exec
delete from sessions
where id = @id;

-- name: DeleteSessionByToken :exec
delete from sessions
where jwt = @jwt;
