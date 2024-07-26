CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    product_id INT UNIQUE NOT NULL REFERENCES products(id),
    start_at TIMESTAMP,
    end_at TIMESTAMP,
    city VARCHAR(100),
    state CHAR(2),
    location VARCHAR(255),
    created_by INT NOT NULL REFERENCES users(id),
    updated_by INT REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
