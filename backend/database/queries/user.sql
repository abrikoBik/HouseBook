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

-- name: SetNowReadingLibrary :one
update users
set now_reading_id = @libraryUUID
where id = @id
returning *;

-- name: SetFinishedReadingLibrary :one
update users
set finished_reading_id = @libraryUUID
where id = @id
returning *;

-- name: SetTradeReadyLibrary :one
update users
set trade_ready_id = @libraryUUID
where id = @id
returning *;
