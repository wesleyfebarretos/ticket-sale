CREATE TABLE IF NOT EXISTS fin.gateway_customer (
    user_id INT NOT NULL REFERENCES users(id),
    gateway_id INT NOT NULL REFERENCES fin.gateway(id),
    gateway_customer_id VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON fin.gateway_customer(user_id);
CREATE INDEX ON fin.gateway_customer(gateway_id);
