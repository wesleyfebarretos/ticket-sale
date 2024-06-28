-- name: GetAll :many
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
    id;

-- name: GetOneById :one
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
LIMIT 1;

-- name: GetOneByEmail :one
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
LIMIT 1;

-- name: Create :one
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
    updated_at;

-- name: Update :exec
UPDATE users 
SET 
    first_name = $2,
    last_name = $3,
    email = $4,
    role = $5
WHERE 
    id = $1
AND
    role = $2;

-- name: Delete :exec
DELETE FROM users WHERE id = $1;

-- name: CheckIfEmailExists :one
SELECT 
    id, first_name, last_name,
    email, role, created_at, updated_at
FROM 
   users
WHERE
   email = $1 AND id != $2 LIMIT 1;
