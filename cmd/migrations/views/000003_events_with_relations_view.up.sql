CREATE VIEW events_with_relations AS
SELECT 
    e.*,
    json_build_object(
        'id', p.id,
        'name', p.name,
        'description', p.description,
        'uuid', p.uuid,
        'price', p.price,
        'discountPrice', p.discount_price,
        'active', p.active,
        'isDeleted', p.is_deleted,
        'image', p.image,
        'imageMobile', p.image_mobile,
        'imageThumbnail', p.image_thumbnail,
        'categoryId', p.category_id,
        'category',
            CASE
                WHEN pc.id IS NULL THEN NULL
                ELSE
                    json_build_object(
                        'id', pc.id,
                        'name', pc.name,
                        'description', pc.description
                    )
            END,
        'stock',
            CASE
                WHEN ps.id IS NULL THEN NULL
                ELSE
                    json_build_object(
                        'id', ps.id,
                        'qty', ps.qty,
                        'minQty', ps.min_qty,
                        'productId', ps.product_id
                    )
            END
    ) as product
FROM events e
INNER JOIN 
    products p ON p.id = e.product_id
LEFT JOIN
    product_categories pc ON pc.id = p.category_id
LEFT JOIN
    product_stocks ps ON ps.product_id = p.id;
