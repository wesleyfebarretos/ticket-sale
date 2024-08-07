// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: creditcard_query.sql

package creditcard_connection

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const create = `-- name: Create :one
INSERT INTO fin.creditcard 
    (name, "number", expiration, priority, notify_expiration, user_id, creditcard_type_id, creditcard_flag_id)
VALUES
    ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING
    id, uuid, name, number, expiration, priority, notify_expiration, user_id, creditcard_type_id, creditcard_flag_id, is_deleted, created_at, updated_at
`

type CreateParams struct {
	Name             string    `json:"name"`
	Number           string    `json:"number"`
	Expiration       time.Time `json:"expiration"`
	Priority         int32     `json:"priority"`
	NotifyExpiration bool      `json:"notifyExpiration"`
	UserID           int32     `json:"userId"`
	CreditcardTypeID int32     `json:"creditcardTypeId"`
	CreditcardFlagID int32     `json:"creditcardFlagId"`
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (FinCreditcard, error) {
	row := q.db.QueryRow(ctx, create,
		arg.Name,
		arg.Number,
		arg.Expiration,
		arg.Priority,
		arg.NotifyExpiration,
		arg.UserID,
		arg.CreditcardTypeID,
		arg.CreditcardFlagID,
	)
	var i FinCreditcard
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Name,
		&i.Number,
		&i.Expiration,
		&i.Priority,
		&i.NotifyExpiration,
		&i.UserID,
		&i.CreditcardTypeID,
		&i.CreditcardFlagID,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllUserCreditcards = `-- name: GetAllUserCreditcards :many
SELECT 
    uuid, name, number, expiration, user_id, created_at, "creditcardFlag", "creditcardType" 
FROM 
    user_creditcards
WHERE 
    user_id = $1
`

func (q *Queries) GetAllUserCreditcards(ctx context.Context, userID int32) ([]UserCreditcard, error) {
	rows, err := q.db.Query(ctx, getAllUserCreditcards, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserCreditcard{}
	for rows.Next() {
		var i UserCreditcard
		if err := rows.Scan(
			&i.Uuid,
			&i.Name,
			&i.Number,
			&i.Expiration,
			&i.UserID,
			&i.CreatedAt,
			&i.CreditcardFlag,
			&i.CreditcardType,
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

const softDelete = `-- name: SoftDelete :exec
UPDATE 
    fin.creditcard
SET
    is_deleted = true,
    updated_at = $2
WHERE
    uuid = $1
`

type SoftDeleteParams struct {
	Uuid      uuid.UUID `json:"uuid"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (q *Queries) SoftDelete(ctx context.Context, arg SoftDeleteParams) error {
	_, err := q.db.Exec(ctx, softDelete, arg.Uuid, arg.UpdatedAt)
	return err
}

const update = `-- name: Update :exec
UPDATE 
    fin.creditcard
SET
    name = $1,
    "number" = $2,
    expiration = $3,
    priority = $4,
    notify_expiration = $5,
    user_id = $6,
    creditcard_type_id = $7,
    creditcard_flag_id = $8,
    updated_at = $9
WHERE
    uuid = $10
`

type UpdateParams struct {
	Name             string    `json:"name"`
	Number           string    `json:"number"`
	Expiration       time.Time `json:"expiration"`
	Priority         int32     `json:"priority"`
	NotifyExpiration bool      `json:"notifyExpiration"`
	UserID           int32     `json:"userId"`
	CreditcardTypeID int32     `json:"creditcardTypeId"`
	CreditcardFlagID int32     `json:"creditcardFlagId"`
	UpdatedAt        time.Time `json:"updatedAt"`
	Uuid             uuid.UUID `json:"uuid"`
}

func (q *Queries) Update(ctx context.Context, arg UpdateParams) error {
	_, err := q.db.Exec(ctx, update,
		arg.Name,
		arg.Number,
		arg.Expiration,
		arg.Priority,
		arg.NotifyExpiration,
		arg.UserID,
		arg.CreditcardTypeID,
		arg.CreditcardFlagID,
		arg.UpdatedAt,
		arg.Uuid,
	)
	return err
}
