-- name: CreateDiary :execresult
INSERT INTO diary (
    content,
    user_email
) VALUES (
    $1, $2
);

-- name: GetDiary :one
SELECT * FROM diary
WHERE id = $1 LIMIT 1;

-- name GetDiarys :many
SELECT * FROM diary
WHERE user_email = $1
ORDER BY id;

-- name UpdateDiary :one
UPDATE diary
SET content = $1
WHERE id = $2;

-- name DeleteDiary :exec
DELETE FROM diary
WHERE id = $1;