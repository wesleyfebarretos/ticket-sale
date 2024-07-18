// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: admin_users_query.sql

package admin_users_repository

import (
	"context"
	"time"
)

const checkIfEmailExists = `-- name: CheckIfEmailExists :one
SELECT 
    id, first_name, last_name,
    email, role, created_at, updated_at
FROM 
   users
WHERE
   email = $1 AND id != $2 LIMIT 1
`

type CheckIfEmailExistsParams struct {
	Email string `json:"email"`
	ID    int32  `json:"id"`
}

type CheckIfEmailExistsRow struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      Roles     `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (q *Queries) CheckIfEmailExists(ctx context.Context, arg CheckIfEmailExistsParams) (CheckIfEmailExistsRow, error) {
	row := q.db.QueryRow(ctx, checkIfEmailExists, arg.Email, arg.ID)
	var i CheckIfEmailExistsRow
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

const create = `-- name: Create :one
INSERT INTO users 
    (first_name, last_name, email, password, role)
VALUES 
    ($1, $2, $3, $4, $5) 
RETURNING
    id,
    first_name,
    last_name,
    email,
    role,
    created_at,
    updated_at
`

type CreateParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      Roles  `json:"role"`
}

type CreateRow struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      Roles     `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (CreateRow, error) {
	row := q.db.QueryRow(ctx, create,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.Role,
	)
	var i CreateRow
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

const delete = `-- name: Delete :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) Delete(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, delete, id)
	return err
}

const getAll = `-- name: GetAll :many
SELECT 
    id,
    first_name,
    last_name,
    email,
    role,
    created_at,
    updated_at
FROM 
   users 
WHERE
    role = $1
ORDER BY 
    id
`

type GetAllRow struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      Roles     `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (q *Queries) GetAll(ctx context.Context, role Roles) ([]GetAllRow, error) {
	rows, err := q.db.Query(ctx, getAll, role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllRow{}
	for rows.Next() {
		var i GetAllRow
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

const getOneByEmail = `-- name: GetOneByEmail :one
SELECT 
    id,
    first_name,
    last_name,
    email,
    role,
    created_at,
    updated_at
FROM 
   users
WHERE
   email = $1 
AND
    role = $2
LIMIT 1
`

type GetOneByEmailParams struct {
	Email string `json:"email"`
	Role  Roles  `json:"role"`
}

type GetOneByEmailRow struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      Roles     `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (q *Queries) GetOneByEmail(ctx context.Context, arg GetOneByEmailParams) (GetOneByEmailRow, error) {
	row := q.db.QueryRow(ctx, getOneByEmail, arg.Email, arg.Role)
	var i GetOneByEmailRow
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

const getOneByEmailAndRoles = `-- name: GetOneByEmailAndRoles :one
SELECT 
    id,
    first_name,
    last_name,
    email,
    role,
    created_at,
    updated_at
FROM 
   users
WHERE
   email = $1 
AND
    role IN($2,$3)
LIMIT 1
`

type GetOneByEmailAndRolesParams struct {
	Email  string `json:"email"`
	Role   Roles  `json:"role"`
	Role_2 Roles  `json:"role2"`
}

type GetOneByEmailAndRolesRow struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      Roles     `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (q *Queries) GetOneByEmailAndRoles(ctx context.Context, arg GetOneByEmailAndRolesParams) (GetOneByEmailAndRolesRow, error) {
	row := q.db.QueryRow(ctx, getOneByEmailAndRoles, arg.Email, arg.Role, arg.Role_2)
	var i GetOneByEmailAndRolesRow
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

const getOneById = `-- name: GetOneById :one
SELECT 
    id,
    first_name,
    last_name,
    email,
    role,
    created_at,
    updated_at
FROM
   users 
WHERE 
    id = $1 
AND 
    role = $2
LIMIT 1
`

type GetOneByIdParams struct {
	ID   int32 `json:"id"`
	Role Roles `json:"role"`
}

type GetOneByIdRow struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      Roles     `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (q *Queries) GetOneById(ctx context.Context, arg GetOneByIdParams) (GetOneByIdRow, error) {
	row := q.db.QueryRow(ctx, getOneById, arg.ID, arg.Role)
	var i GetOneByIdRow
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

const update = `-- name: Update :exec
UPDATE users 
SET 
    first_name = $2,
    last_name = $3,
    email = $4,
    role = $5
WHERE 
    id = $1
`

type UpdateParams struct {
	ID        int32  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Role      Roles  `json:"role"`
}

func (q *Queries) Update(ctx context.Context, arg UpdateParams) error {
	_, err := q.db.Exec(ctx, update,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Role,
	)
	return err
}
