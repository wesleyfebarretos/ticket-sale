BEGIN;
    CREATE TABLE fin.payment_order (
        id SERIAL PRIMARY KEY,
        uuid UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
        creditcard_uuid UUID REFERENCES fin.creditcard(uuid),
        user_id INT NOT NULL REFERENCES users(id),
        total_price DOUBLE PRECISION NOT NULL,
        payment_type_id INT NOT NULL REFERENCES fin.payment_type(id),
        payment_period_id INT NOT NULL REFERENCES fin.payment_period(id),
        gateway_id INT NOT NULL REFERENCES fin.gateway(id),
        payment_status_id INT NOT NULL REFERENCES fin.payment_status(id),
        payment_cancel_reason_id INT REFERENCES fin.payment_cancel_reason(id),
        extra_info TEXT NULL,
        payment_at TIMESTAMP NULL,
        cancel_at TIMESTAMP NULL,
        due_at TIMESTAMP NULL,
        expiration_at TIMESTAMP NULL,
        base_value DOUBLE PRECISION NOT NULL DEFAULT 0,
        reversed_value DOUBLE PRECISION NOT NULL DEFAULT 0,
        canceled_value DOUBLE PRECISION NOT NULL DEFAULT 0,
        added_value DOUBLE PRECISION NOT NULL DEFAULT 0,
        total_value DOUBLE PRECISION NOT NULL DEFAULT 0,
        created_by INT NOT NULL REFERENCES users(id),
        updated_by INT REFERENCES users(id),
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
    );

    CREATE INDEX ON fin.payment_order(uuid);
    CREATE INDEX ON fin.payment_order(user_id);
COMMIT;
