-- name: CreateUser :execresult
INSERT INTO users (
    email,
    gender,
    age,
    country
) VALUES (
   ?, ?, ?, ?
);

-- name: GetUser :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE email = ?;

-- name: GetAllUsers :many
SELECT * FROM users;