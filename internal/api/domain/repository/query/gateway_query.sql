-- name: GetActive :one
SELECT 
    *
FROM
    fin.gateway
WHERE
    active IS TRUE
AND
    is_deleted IS FALSE
LIMIT 1;
