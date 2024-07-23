// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: link.sql

package pgstore

import (
	"context"

	"github.com/google/uuid"
)

const createTripLink = `-- name: CreateTripLink :one
INSERT INTO links ("trip_id", "title", "url")
VALUES ($1, $2, $3)
RETURNING "id"
`

type CreateTripLinkParams struct {
	TripID uuid.UUID
	Title  string
	Url    string
}

func (q *Queries) CreateTripLink(ctx context.Context, arg CreateTripLinkParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createTripLink, arg.TripID, arg.Title, arg.Url)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getTripLinks = `-- name: GetTripLinks :many
SELECT "id",
  "trip_id",
  "title",
  "url"
FROM links
WHERE "trip_id" = $1
`

func (q *Queries) GetTripLinks(ctx context.Context, tripID uuid.UUID) ([]Link, error) {
	rows, err := q.db.Query(ctx, getTripLinks, tripID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Link
	for rows.Next() {
		var i Link
		if err := rows.Scan(
			&i.ID,
			&i.TripID,
			&i.Title,
			&i.Url,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}