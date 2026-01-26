-- name: GetUser :one
select * from users
where id = $1
limit 1;

-- name: CreateUser :one
insert into users (
    email, login, name, salt, hash
) values (
    $1,    $2,    $3,   $4,   $5
)
returning *;

-- name: UpdateUser :one
update users
set
    email = $2,
    login = $3,
    name = $4,
    salt = $5,
    hash = $6
where id = $1
returning *;

-- name: DeleteUser :exec
delete from users
where id = $1;
