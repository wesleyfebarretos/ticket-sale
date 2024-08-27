-- name: Create :one
INSERT INTO fin.creditcard 
    (name, "number", expiration, priority, notify_expiration, user_id, creditcard_type_id, creditcard_flag_id)
VALUES
    ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING
    *;

-- name: Update :exec
UPDATE 
    fin.creditcard
SET
    name = $1,
    "number" = $2,
    expiration = $3,
    priority = $4,
    notify_expiration = $5,
    user_id = $6,
    creditcard_type_id = $7,
    creditcard_flag_id = $8,
    updated_at = $9
WHERE
    uuid = $10;

-- name: SoftDelete :exec
UPDATE 
    fin.creditcard
SET
    is_deleted = true,
    updated_at = $2
WHERE
    uuid = $1;

-- name: GetAllUserCreditcards :many
SELECT 
    * 
FROM 
    user_creditcards
WHERE 
    user_id = $1;

-- name: GetByUuid :one
SELECT * FROM fin.creditcard WHERE uuid = $1;
