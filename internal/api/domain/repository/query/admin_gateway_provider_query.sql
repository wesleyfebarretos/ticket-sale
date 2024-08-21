-- name: Create :one
INSERT INTO fin.gateway_provider
    (name, created_by, updated_by)
VALUES
    ($1,$2,$3)
RETURNING *;
