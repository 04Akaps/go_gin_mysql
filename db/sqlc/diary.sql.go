// Code generated by sqlc. DO NOT EDIT.
// source: diary.sql

package db

import (
	"context"
	"database/sql"
)

const createDiary = `-- name: CreateDiary :execresult
INSERT INTO diary (
    content,
    user_email
) VALUES (
    ?, ?
)
`

type CreateDiaryParams struct {
	Content   string `json:"content"`
	UserEmail string `json:"user_email"`
}

func (q *Queries) CreateDiary(ctx context.Context, arg CreateDiaryParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createDiary, arg.Content, arg.UserEmail)
}

const deleteDiary = `-- name: DeleteDiary :exec
DELETE FROM diary
WHERE id = ?
`

func (q *Queries) DeleteDiary(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteDiary, id)
	return err
}

const getDiary = `-- name: GetDiary :one
SELECT id, content, user_email, created_at FROM diary
WHERE id = ? LIMIT 1
`

func (q *Queries) GetDiary(ctx context.Context, id int64) (Diary, error) {
	row := q.db.QueryRowContext(ctx, getDiary, id)
	var i Diary
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.UserEmail,
		&i.CreatedAt,
	)
	return i, err
}

const getDiarys = `-- name: GetDiarys :many
SELECT id, content, user_email, created_at FROM diary
WHERE user_email = ?
ORDER BY id
`

func (q *Queries) GetDiarys(ctx context.Context, userEmail string) ([]Diary, error) {
	rows, err := q.db.QueryContext(ctx, getDiarys, userEmail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Diary{}
	for rows.Next() {
		var i Diary
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.UserEmail,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDiary = `-- name: UpdateDiary :exec
UPDATE diary
SET content = ?
WHERE id = ?
`

type UpdateDiaryParams struct {
	Content string `json:"content"`
	ID      int64  `json:"id"`
}

func (q *Queries) UpdateDiary(ctx context.Context, arg UpdateDiaryParams) error {
	_, err := q.db.ExecContext(ctx, updateDiary, arg.Content, arg.ID)
	return err
}
