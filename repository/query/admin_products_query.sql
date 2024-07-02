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
SELECT 
    p.*,
    CASE
        WHEN ps.id IS NULL THEN NULL
        ELSE
            json_build_object(
                'id', ps.id,
                'productId', ps.product_id,
                'qty', ps.qty,
                'minQty', ps.min_qty
            )
    END as stock,
    CASE
        WHEN pc.id IS NULL THEN NULL
        ELSE
            json_build_object(
                'id', pc.id,
                'name', pc.name,
                'description', pc.description
            )
    END as category
FROM 
    products as p
LEFT JOIN 
    product_stocks as ps 
ON 
    ps.product_id = p.id
LEFT JOIN
    product_categories as pc
ON
    pc.id = p.category_id
WHERE 
    p.is_deleted IS FALSE 
AND
    p.active IS TRUE
ORDER BY 
    p.created_at DESC;

-- name: GetOneById :one
SELECT 
    p.*,
    CASE
        WHEN ps.id IS NULL THEN NULL
        ELSE
            json_build_object(
                'id', ps.id,
                'productId', ps.product_id,
                'qty', ps.qty,
                'minQty', ps.min_qty
            )
    END as stock,
    CASE
        WHEN pc.id IS NULL THEN NULL
        ELSE
            json_build_object(
                'id', pc.id,
                'name', pc.name,
                'description', pc.description
            )
    END as category
FROM 
    products as p
LEFT JOIN 
    product_stocks as ps 
ON 
    ps.product_id = p.id
LEFT JOIN
    product_categories as pc
ON
    pc.id = p.category_id
WHERE 
    p.id = $1
LIMIT 1;

-- name: GetOneByUuid :one
SELECT 
    p.*,
    CASE
        WHEN ps.id IS NULL THEN NULL
        ELSE
            json_build_object(
                'id', ps.id,
                'productId', ps.product_id,
                'qty', ps.qty,
                'minQty', ps.min_qty
            )
    END as stock,
    CASE
        WHEN pc.id IS NULL THEN NULL
        ELSE
            json_build_object(
                'id', pc.id,
                'name', pc.name,
                'description', pc.description
            )
    END as category
FROM 
    products as p
LEFT JOIN 
    product_stocks as ps 
ON 
    ps.product_id = p.id
LEFT JOIN
    product_categories as pc
ON
    pc.id = p.category_id
WHERE 
    p.uuid = $1
LIMIT 1;
