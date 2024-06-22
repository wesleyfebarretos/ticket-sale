BEGIN;
CREATE TABLE users_addresses (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    street_address VARCHAR(255) NOT NULL,
    city VARCHAR(100) NOT NULL,
    complement VARCHAR(500),
    state CHAR(2) NOT NULL,
    postal_code VARCHAR(20),
    country VARCHAR(100) NOT NULL,
    address_type VARCHAR(50),
    favorite BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
COMMIT;
