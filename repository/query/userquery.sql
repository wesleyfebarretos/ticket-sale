-- name: GetUsers :many
SELECT 
    id, first_name, last_name,
    email, role, created_at, updated_at
FROM 
   users ORDER BY id;

-- name: GetUser :one
SELECT 
    id, first_name, last_name,
    email, role, created_at, updated_at
FROM
   users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT 
    id, first_name, last_name,
    email, role, created_at, updated_at
FROM 
   users
WHERE
   email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users 
(first_name, last_name, email, password, role)
VALUES 
($1, $2, $3, $4, $5) 
RETURNING
    id, first_name, last_name,
    email, role, created_at, updated_at;

-- name: UpdateUser :exec
UPDATE users 
SET 
    first_name = $2,
    last_name = $3,
    email = $4,
    password = $5,
    role = $6
WHERE id = $1;

-- name: DestroyUser :exec
DELETE FROM users WHERE id = $1;
