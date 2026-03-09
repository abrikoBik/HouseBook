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

-- name: SetProfilePicture :one
update users
set image_uuid = @imageUUID
where id = @id
returning *;

-- name: SetUserLibraries :one
update users
set
    now_reading_id = coalesce(sqlc.narg('nowReading'), now_reading_id),
    finished_reading_id = coalesce(sqlc.narg('finishedReading'), finished_reading_id),
    trade_ready_id = coalesce(sqlc.narg('tradeReady'), trade_ready_id)
where id = @id
returning *;
