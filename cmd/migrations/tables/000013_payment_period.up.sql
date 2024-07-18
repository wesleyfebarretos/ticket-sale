CREATE TABLE IF NOT EXISTS fin.payment_period (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    times INT NOT NULL,
    created_by INT NOT NULL REFERENCES users(id),
    updated_by INT REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
