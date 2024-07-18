CREATE TABLE IF NOT EXISTS fin.gateway_payment_type_association (
    id SERIAL PRIMARY KEY,
    gateway_id INT NOT NULL REFERENCES fin.gateway(id),
    gateway_payment_type_id INT NOT NULL REFERENCES fin.gateway_payment_type(id),
    updated_by INT REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
