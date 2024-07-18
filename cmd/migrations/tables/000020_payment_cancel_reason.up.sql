CREATE TABLE IF NOT EXISTS fin.payment_cancel_reason (
    id SERIAL PRIMARY KEY,
    name VARCHAR(250) NOT NULL,
    active BOOL DEFAULT true,
    created_by INT NOT NULL REFERENCES users(id),
    updated_by INT REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
