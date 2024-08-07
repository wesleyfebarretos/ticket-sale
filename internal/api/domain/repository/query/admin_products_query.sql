-- name: Create :one
INSERT INTO products 
(name, description, price, discount_price, active, image, image_mobile, image_thumbnail, category_id, created_by)
VALUES 
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
RETURNING *;

-- name: CreateInstallments :batchone
INSERT INTO fin.product_payment_type_installment_time
    (product_id, payment_type_id, installment_time_id, fee, tariff, created_by)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING
    *;

-- name: Update :exec
UPDATE products SET
    name = $1,
    description = $2,
    price = $3,
    discount_price = $4,
    active = $5,
    image = $6,
    image_mobile = $7,
    image_thumbnail = $8,
    category_id = $9,
    updated_by = $10
WHERE 
    id = $11;

-- name: SoftDelete :exec
UPDATE products SET
    is_deleted = true,
    updated_by = $2
WHERE 
    id = $1
AND 
    is_deleted IS FALSE;

-- name: GetAll :many
SELECT * FROM products 
WHERE 
    is_deleted IS FALSE 
AND
    active IS TRUE
ORDER BY 
    created_at DESC;

-- name: GetAllProductsDetails :many
SELECT * FROM products_details
WHERE 
    is_deleted IS FALSE 
AND
    active IS TRUE
ORDER BY 
    created_at DESC;

-- name: GetOneById :one
SELECT * FROM products_details
WHERE 
    id = $1
LIMIT 1;

-- name: GetOneByUuid :one
SELECT * FROM products_details
WHERE 
    uuid = $1
LIMIT 1;

-- name: GetAllProductInstallmentTimes :many
SELECT 
    id,
    fee,
    tariff,
    payment_type_id,
    installment_time_id
FROM 
    fin.product_payment_type_installment_time
WHERE
    product_id = $1;

-- name: DeleteAllProductInstallmentTimes :exec
DELETE FROM
    fin.product_payment_type_installment_time
WHERE
    product_id = $1;
