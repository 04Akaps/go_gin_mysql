-- name: CreateDiary :execresult
INSERT INTO diary (
    content,
    user_email
) VALUES (
    ?, ?
);

-- name: GetDiary :one
SELECT * FROM diary
WHERE id = ? LIMIT 1;

-- name: GetDiarys :many
SELECT * FROM diary
WHERE user_email = ?
ORDER BY id;

-- name: UpdateDiary :exec
UPDATE diary
SET content = ?
WHERE id = ?;

-- name: DeleteDiary :exec
DELETE FROM diary
WHERE id = ?;