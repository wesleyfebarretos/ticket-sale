CREATE VIEW gateway_details AS
SELECT 
    g.id,
    g.uuid,
    g.name,
    g.description,
    g.client_id,
    g.client_secret,
    g."order",
    g.active,
    g.test_environment,
    g.notif_user,
    g.notif_password,
    g.soft_descriptor,
    g.gateway_process_id,
    g.webhook_url,
    g.url,
    g.auth_type,
    g.use_3ds,
    g.adq_code_3ds,
    g.default_adq_code,
    g.use_antifraud,
    g.created_by,
    g.updated_by,
    g.created_at,
    g.updated_at,
    json_build_object(
        'id', gp.id,
        'name', gp.name
    ) AS "gatewayProcess",
    json_agg(
        json_build_object(
            'id', gpt.id,
            'name', gpt.name
        ) ORDER BY gpt.id ASC
    ) as "gatewayPaymentTypes"
FROM 
    fin.gateway g
JOIN
    fin.gateway_process gp ON gp.id = g.gateway_process_id
JOIN
    fin.gateway_payment_type_association gpta ON gpta.gateway_id = g.id
JOIN
    fin.gateway_payment_type gpt ON gpt.id = gpta.gateway_payment_type_id
GROUP BY g.id, gp.id;
