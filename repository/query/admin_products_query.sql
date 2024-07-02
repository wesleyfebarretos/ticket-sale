-- name: Create :one
INSERT INTO products 
(name, description, uuid, price, discount_price, active, image, image_mobile, image_thumbnail, category_id, created_by)
VALUES 
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
RETURNING *;

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

-- name: GetAllWithRelations :many
SELECT * FROM products_with_relation
WHERE 
    is_deleted IS FALSE 
AND
    active IS TRUE
ORDER BY 
    created_at DESC;

-- name: GetOneById :one
SELECT * FROM products_with_relation
WHERE 
    id = $1
LIMIT 1;

-- name: GetOneByUuid :one
SELECT * FROM products_with_relation
WHERE 
    uuid = $1
LIMIT 1;

-- name: CreateWithStock :one
BEGIN;
    INSERT INTO products 
        (name, description, uuid, price, discount_price, active, image, image_mobile, image_thumbnail, category_id, created_by)
    VALUES 
        ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
    RETURNING *;
    INSERT INTO product_stocks
        (product_id, qty, min_qty, created_by)
    VALUES
        ($12,$13,$14,$15)
    RETURNING *;
COMMIT;
