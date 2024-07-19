CREATE VIEW products_details AS
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
    END as category,
    json_build_object(
        'creditcard', COALESCE(
            json_agg(
                json_build_object(
                    'installmentTimeId', it.id,
                    'installmentTimeName', it.name,
                    'fee', ppi.fee,
                    'tariff', ppi.tariff
                ) ORDER BY it.id ASC
            ) FILTER (WHERE ppi.payment_type_id = 1),
            '[]'::json
        ),
        'paymentSlip', COALESCE(
            json_agg(
                json_build_object(
                    'installmentTimeId', it.id,
                    'installmentTimeName', it.name,
                    'fee', ppi.fee,
                    'tariff', ppi.tariff
                ) ORDER BY it.id ASC
            ) FILTER (WHERE ppi.payment_type_id = 2),
            '[]'::json
        ),
        'pix', COALESCE(
            json_agg(
                json_build_object(
                    'installmentTimeId', it.id,
                    'installmentTimeName', it.name,
                    'fee', ppi.fee,
                    'tariff', ppi.tariff
                ) ORDER BY it.id ASC
            ) FILTER (WHERE ppi.payment_type_id = 3),
            '[]'::json
        )
    ) as installments
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
LEFT JOIN
    fin.product_payment_type_installment_time as ppi
ON
    ppi.product_id = p.id
LEFT JOIN
    fin.installment_time as it
ON
    it.id = ppi.installment_time_id
GROUP BY 
	p.id, 
	ps.id, 
	pc.id;
