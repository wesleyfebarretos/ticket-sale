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

-- name: GetDifferentUserByEmail :one
SELECT 
    id, first_name, last_name,
    email, role, created_at, updated_at
FROM 
   users
WHERE
   email = $1 AND id != $2 LIMIT 1;

-- name: GetUserFullProfile :one
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
LIMIT 1;

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
    role = $5
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
