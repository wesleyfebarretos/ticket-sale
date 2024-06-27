-- name: Create :one
INSERT INTO users_addresses
(user_id, street_address, city, complement, state, postal_code, country, address_type, favorite)
VALUES 
($1, $2, $3, $4, $5, $6, $7, $8, $9) 
RETURNING *;

-- name: Update :exec
UPDATE users_addresses
SET 
    street_address = $2,
    city = $3,
    complement = $4,
    state = $5,
    postal_code = $6,
    country = $7,
    address_type = $8,
    favorite = $9
WHERE id = $1;


