CREATE TABLE IF NOT EXISTS fin.gateway_customer_card(
    id SERIAL PRIMARY KEY,
    gateway_id INT NOT NULL REFERENCES fin.gateway(id),
    user_id INT NOT NULL REFERENCES users(id),
    card_id INT NOT NULL REFERENCES fin.creditcard(id),
    gateway_card_id VARCHAR(150) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON fin.gateway_customer_card(gateway_id);
CREATE INDEX ON fin.gateway_customer_card(user_id);
