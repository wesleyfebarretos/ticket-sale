-- name: Create :one
INSERT INTO fin.gateway_customer_card
(gateway_id, user_id, card_id, gateway_card_id)
VALUES
($1,$2,$3,$4)
RETURNING *;

-- name: GetByUserAndGatewayId :many
SELECT
    id,
    gateway_id,
    user_id,
    card_id,
    gateway_card_id
FROM 
    fin.gateway_customer_card
WHERE 
    user_id = $1
AND
    gateway_id = $2;

