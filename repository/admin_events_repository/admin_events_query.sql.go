// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: admin_events_query.sql

package admin_events_repository

import (
	"context"
	"time"
)

const create = `-- name: Create :one
INSERT INTO events
    (product_id, start_at, end_at, city, state, location, created_by)
VALUES
    ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, product_id, start_at, end_at, city, state, location, created_by, updated_by, created_at, updated_at
`

type CreateParams struct {
	ProductID int32      `json:"productId"`
	StartAt   *time.Time `json:"startAt"`
	EndAt     *time.Time `json:"endAt"`
	City      *string    `json:"city"`
	State     *string    `json:"state"`
	Location  *string    `json:"location"`
	CreatedBy int32      `json:"createdBy"`
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (Event, error) {
	row := q.db.QueryRow(ctx, create,
		arg.ProductID,
		arg.StartAt,
		arg.EndAt,
		arg.City,
		arg.State,
		arg.Location,
		arg.CreatedBy,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.StartAt,
		&i.EndAt,
		&i.City,
		&i.State,
		&i.Location,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAll = `-- name: GetAll :many
SELECT id, product_id, start_at, end_at, city, state, location, created_by, updated_by, created_at, updated_at, product FROM events_get_all 
WHERE 
    (product->>'isDeleted')::boolean IS FALSE
`

func (q *Queries) GetAll(ctx context.Context) ([]EventsGetAll, error) {
	rows, err := q.db.Query(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []EventsGetAll{}
	for rows.Next() {
		var i EventsGetAll
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.StartAt,
			&i.EndAt,
			&i.City,
			&i.State,
			&i.Location,
			&i.CreatedBy,
			&i.UpdatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Product,
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

const getOneById = `-- name: GetOneById :one
SELECT id, product_id, start_at, end_at, city, state, location, created_by, updated_by, created_at, updated_at, product from events_details
WHERE
    id = $1
AND
    (product->>'isDeleted')::boolean IS FALSE
`

func (q *Queries) GetOneById(ctx context.Context, id int32) (EventsDetail, error) {
	row := q.db.QueryRow(ctx, getOneById, id)
	var i EventsDetail
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.StartAt,
		&i.EndAt,
		&i.City,
		&i.State,
		&i.Location,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Product,
	)
	return i, err
}

const softDelete = `-- name: SoftDelete :exec
UPDATE products p 
SET 
    is_deleted = true
FROM 
    events e
WHERE 
    p.id = e.product_id
AND 
    e.id = $1
`

func (q *Queries) SoftDelete(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, softDelete, id)
	return err
}

const update = `-- name: Update :one
UPDATE events SET
    start_at = $2,
    end_at = $3,
    city = $4,
    state = $5,
    location = $6,
    updated_at = $7
WHERE id = $1
RETURNING product_id
`

type UpdateParams struct {
	ID        int32      `json:"id"`
	StartAt   *time.Time `json:"startAt"`
	EndAt     *time.Time `json:"endAt"`
	City      *string    `json:"city"`
	State     *string    `json:"state"`
	Location  *string    `json:"location"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

func (q *Queries) Update(ctx context.Context, arg UpdateParams) (int32, error) {
	row := q.db.QueryRow(ctx, update,
		arg.ID,
		arg.StartAt,
		arg.EndAt,
		arg.City,
		arg.State,
		arg.Location,
		arg.UpdatedAt,
	)
	var product_id int32
	err := row.Scan(&product_id)
	return product_id, err
}
