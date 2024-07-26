-- name: Create :one
INSERT INTO product_stocks
    (product_id, qty, min_qty, created_by)
VALUES
    ($1,$2,$3,$4)
RETURNING *;

-- name: Update :exec
UPDATE product_stocks SET
    qty = $1,
    min_qty = $2,
    updated_by = $3,
    updated_at = $4
WHERE
    id = $5;

