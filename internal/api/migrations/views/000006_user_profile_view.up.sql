CREATE VIEW user_profile AS
SELECT 
    u.id, 
    u.first_name,
    u.last_name,
    u.email,
    u.role,
    u.created_at,
    u.updated_at,
    COALESCE(
        json_agg(
            json_build_object(
                'id', ua.id,
                'userId', ua.user_id,
                'streetAddress', ua.street_address,
                'city', ua.city,
                'complement', ua.complement,
                'state', ua.state,
                'postalCode', ua.postal_code,
                'country', ua.country,
                'addressType', ua.address_type,
                'favorite', ua.favorite
            ) ORDER BY ua.favorite DESC
        ) FILTER (WHERE ua.id IS NOT NULL), '[]'::json
    ) AS addresses,
    COALESCE(
        json_agg(
            json_build_object(
                'id', up.id,
                'userId', up.user_id,
                'ddd', up.ddd,
                'number', up.number,
                'type', up.type
            ) ORDER BY up.id ASC
        ) FILTER (WHERE up.id IS NOT NULL), '[]'::json
    ) AS phones
FROM 
    users AS u
LEFT JOIN 
    users_addresses AS ua
ON 
    u.id = ua.user_id
LEFT JOIN
    users_phones as up
ON
    u.id = up.user_id
GROUP BY u.id;
