// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user-query.sql

package sqlc

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users 
(first_name, last_name, email, password, role)
VALUES 
($1, $2, $3, $4, $5) 
RETURNING
    id, first_name, last_name,
    email, role, created_at, updated_at
`

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      Roles  `json:"role"`
}

type CreateUserRow struct {
	ID        int32      `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Role      Roles      `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.Role,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const destroyUser = `-- name: DestroyUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DestroyUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, destroyUser, id)
	return err
}

const getDifferentUserByEmail = `-- name: GetDifferentUserByEmail :one
SELECT 
    id, first_name, last_name,
    email, role, created_at, updated_at
FROM 
   users
WHERE
   email = $1 AND id != $2 LIMIT 1
`

type GetDifferentUserByEmailParams struct {
	Email string `json:"email"`
	ID    int32  `json:"id"`
}

type GetDifferentUserByEmailRow struct {
	ID        int32      `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Role      Roles      `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func (q *Queries) GetDifferentUserByEmail(ctx context.Context, arg GetDifferentUserByEmailParams) (GetDifferentUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getDifferentUserByEmail, arg.Email, arg.ID)
	var i GetDifferentUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT 
    id, first_name, last_name,
    email, role, created_at, updated_at
FROM
   users WHERE id = $1 LIMIT 1
`

type GetUserRow struct {
	ID        int32      `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Role      Roles      `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func (q *Queries) GetUser(ctx context.Context, id int32) (GetUserRow, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT 
    id, first_name, last_name,
    email, role, created_at, updated_at
FROM 
   users
WHERE
   email = $1 LIMIT 1
`

type GetUserByEmailRow struct {
	ID        int32      `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Role      Roles      `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserFullProfile = `-- name: GetUserFullProfile :one
SELECT 
    u.id, 
    u.first_name,
    u.last_name,
    u.email,
    u.role,
    u.created_at,
    u.updated_at,
    COALESCE(
        json_agg(
        json_build_object(
            'id', ua.id,
            'userId', ua.user_id,
            'streetAddress', ua.street_address,
            'city', ua.city,
            'complement', ua.complement,
            'state', ua.state,
            'postalCode', ua.postal_code,
            'country', ua.country,
            'addressType', ua.address_type,
            'favorite', ua.favorite
        ) ORDER BY ua.favorite DESC
    ) FILTER (WHERE ua.id IS NOT NULL), '[]'::json
    ) AS addresses
FROM 
    users AS u
LEFT JOIN 
    users_addresses AS ua
ON 
    u.id = ua.user_id
WHERE 
    u.id = $1 
GROUP BY 
	u.id 
LIMIT 1
`

type GetUserFullProfileRow struct {
	ID        int32       `json:"id"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Email     string      `json:"email"`
	Role      Roles       `json:"role"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt *time.Time  `json:"updatedAt"`
	Addresses interface{} `json:"addresses"`
}

func (q *Queries) GetUserFullProfile(ctx context.Context, id int32) (GetUserFullProfileRow, error) {
	row := q.db.QueryRow(ctx, getUserFullProfile, id)
	var i GetUserFullProfileRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Addresses,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT 
    id, first_name, last_name,
    email, role, created_at, updated_at
FROM 
   users ORDER BY id
`

type GetUsersRow struct {
	ID        int32      `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Role      Roles      `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func (q *Queries) GetUsers(ctx context.Context) ([]GetUsersRow, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUsersRow{}
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Role,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateUser = `-- name: UpdateUser :exec
UPDATE users 
SET 
    first_name = $2,
    last_name = $3,
    email = $4,
    role = $5
WHERE id = $1
`

type UpdateUserParams struct {
	ID        int32  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Role      Roles  `json:"role"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Role,
	)
	return err
}
