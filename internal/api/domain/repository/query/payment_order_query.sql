-- name: Create :one
INSERT INTO fin.payment_order (
    creditcard_uuid,
    user_id,
    total_price,
    payment_type_id,
    installment_time_id,
    gateway_id,
    payment_status_id,
    total_price,
    added_value,
    base_value,
    created_by,
    updated_by
)
VALUES
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: Cancel :exec
UPDATE fin.payment_order
SET
    payment_status_id = 5,
    cancel_at = $2,
    updated_by = $3,
    updated_at = $4
WHERE
    uuid = $1;

-- name: GetOneByUuid :one
SELECT * FROM fin.payment_order WHERE uuid = $1;
