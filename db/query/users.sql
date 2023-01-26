-- name: CreateUser :execresult
INSERT INTO users (
    email,
    gender,
    age,
    country
) VALUES (
    $1, $2, $3, $4
);

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE email = $1;