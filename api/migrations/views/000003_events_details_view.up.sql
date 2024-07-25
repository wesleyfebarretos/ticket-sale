CREATE VIEW events_details AS
SELECT 
    e.*,
    json_build_object(
        'id',pd.id,
        'name', pd.name,
        'description', pd.description,
        'uuid', pd.uuid,
        'price', pd.price,
        'active', pd.active,
        'image', pd.image,
        'discountPrice', pd.discount_price,
        'imageMobile', pd.image_mobile,
        'imageThumbnail', pd.image_thumbnail,
        'isDeleted', pd.is_deleted,
        'categoryId', pd.category_id,
        'category', pd.category,
        'stock', pd.stock,
        'installments', pd.installments
    ) as product
FROM events e
INNER JOIN 
    products_details pd ON pd.id = e.product_id;
