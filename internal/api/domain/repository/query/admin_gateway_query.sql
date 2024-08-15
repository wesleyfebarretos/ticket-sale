-- name: Create :one
INSERT INTO fin.gateway
    ("name", description, client_id, client_secret, "order", active, test_environment, notif_user, notif_password, soft_descriptor, gateway_process_id, webhook_url, url, auth_type, use_3ds, adq_code_3ds, default_adq_code, use_antifraud, created_by, updated_by, gateway_provider_id)
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
RETURNING *;

-- name: CreatePaymentTypes :batchone
INSERT INTO fin.gateway_payment_type_association
    (gateway_id, gateway_payment_type_id, created_by, updated_by)
VALUES
    ($1,$2,$3,$4)
RETURNING  *;

-- name: Update :exec
UPDATE fin.gateway SET
    "name" = $2,
    description = $3,
    client_id = $4,
    client_secret = $5,
    "order" = $6,
    active = $7,
    test_environment = $8,
    notif_user = $9,
    notif_password = $10,
    soft_descriptor = $11,
    gateway_process_id = $12,
    webhook_url = $13,
    url = $14,
    auth_type = $15,
    use_3ds = $16,
    adq_code_3ds = $17,
    default_adq_code = $18,
    use_antifraud = $19,
    updated_by = $20,
    gateway_provider_id = $21
WHERE id = $1;

-- name: GetAll :many
SELECT * FROM gateway_details;

-- name: GetOneById :one
SELECT 
    *
FROM 
    gateway_details
WHERE id = $1;

-- name: SoftDelete :exec
UPDATE fin.gateway SET 
    active = false,
    is_deleted = true,
    updated_by = $2
WHERE id = $1;


