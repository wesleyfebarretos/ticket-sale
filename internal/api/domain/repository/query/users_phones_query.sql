-- name: Create :one
INSERT INTO users_phones
    (user_id, ddd, number, type)
VALUES 
    ($1, $2, $3, $4) 
RETURNING *;
