-- name: CreateUser :one
insert into users (id, username, hashed_password, created_at, updated_at)
values ($1, $2, $3, $4, $5)
returning *;

-- name: GetUserByUsername :one
select * from users where username = $1;