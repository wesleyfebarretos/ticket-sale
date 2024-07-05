-- name: Create :one
INSERT INTO events
    (product_id, start_at, end_at, city, state, location)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: Update :one
UPDATE events SET
    start_at = $2,
    end_at = $3,
    city = $4,
    state = $5,
    location = $6
WHERE id = $1
RETURNING product_id;

-- name: SoftDelete :exec
UPDATE products p 
SET 
    is_deleted = true
FROM 
    events e
WHERE 
    p.id = e.product_id
AND 
    e.id = $1;

-- name: GetAll :many
SELECT * FROM events_get_all 
WHERE 
    (product->>'isDeleted')::boolean IS FALSE;

-- name: GetOneById :one
SELECT * from events_with_relations
WHERE
    id = $1
AND
    (product->>'isDeleted')::boolean IS FALSE;
