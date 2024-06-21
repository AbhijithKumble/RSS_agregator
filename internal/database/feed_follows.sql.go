// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollows = `-- name: CreateFeedFollows :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id) 
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, user_id, feed_id
`

type CreateFeedFollowsParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

func (q *Queries) CreateFeedFollows(ctx context.Context, arg CreateFeedFollowsParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollows,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const deleteFeedFollows = `-- name: DeleteFeedFollows :exec
DELETE FROM feed_follows 
WHERE id =$1 
AND user_id =$2
`

type DeleteFeedFollowsParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteFeedFollows(ctx context.Context, arg DeleteFeedFollowsParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollows, arg.ID, arg.UserID)
	return err
}

const getFeedFollows = `-- name: GetFeedFollows :many
SELECT id, created_at, updated_at, user_id, feed_id FROM feed_follows WHERE user_id=$1
`

func (q *Queries) GetFeedFollows(ctx context.Context, userID uuid.UUID) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
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
