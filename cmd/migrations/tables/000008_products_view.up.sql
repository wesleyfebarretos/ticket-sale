CREATE VIEW products_with_relation AS
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
    pc.id = p.category_id;
