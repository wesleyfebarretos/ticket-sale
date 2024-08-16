-- name: Create :one
INSERT INTO fin.gateway_customer
(user_id, gateway_id, gateway_customer_id)
VALUES
($1, $2, $3)
RETURNING *;

-- name: FindOneByUserId :one
SELECT
    *
FROM
    fin.gateway_customer
WHERE
    user_id = $1;

-- name: FindOneByGatewayAndUserId :one
SELECT
    *
FROM
    fin.gateway_customer
WHERE
    user_id = $1
AND
    gateway_id = $2;

