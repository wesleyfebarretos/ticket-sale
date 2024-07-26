BEGIN;
    CREATE TABLE IF NOT EXISTS users_addresses (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL REFERENCES users(id),
        street_address VARCHAR(255) NOT NULL,
        city VARCHAR(100) NOT NULL,
        complement VARCHAR(500),
        state CHAR(2) NOT NULL,
        postal_code VARCHAR(20),
        country VARCHAR(100) NOT NULL,
        address_type VARCHAR(50),
        favorite BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );
    CREATE INDEX ON users_addresses(user_id);
COMMIT;
