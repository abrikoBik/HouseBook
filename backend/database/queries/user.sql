-- name: GetUser :one
select * from users
where id = @id
limit 1;

-- name: CreateUser :one
insert into users (
    email,
    login,
    name,
    salt,
    hash
) values (
    @email,
    @login,
    @name,
    @salt,
    @hash
)
returning *;

-- name: UpdateUser :one
update users
set
    email = coalesce(sqlc.narg('email'), email),
    login = coalesce(sqlc.narg('login'), login),
    name = coalesce(sqlc.narg('name'), name)
where id = @id
returning *;

-- name: DeleteUser :exec
delete from users
where id = @id;
