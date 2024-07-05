-- name: Create :one
INSERT INTO events
    (product_id, start_at, end_at, city, state, location)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: Update :exec
UPDATE events SET
    start_at = $2,
    end_at = $3,
    city = $4,
    state = $5,
    location = $6
WHERE id = $1;

-- name: Delete :exec
DELETE  FROM events WHERE id = $1;

-- name: GetAll :many
SELECT * FROM get_all_events
WHERE
    p.is_deleted IS FALSE;
