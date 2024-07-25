BEGIN;
    CREATE TABLE fin.payment_order_product (
        id SERIAL PRIMARY KEY,
        payment_order_id INT NOT NULL REFERENCES fin.payment_order(id),
        product_id INT NOT NULL REFERENCES products(id),
        qty INT NOT NULL,
        unit_price DOUBLE PRECISION NOT NULL,
        subtotal_price DOUBLE PRECISION NOT NULL,
        discount DOUBLE PRECISION NOT NULL,
        total_price DOUBLE PRECISION NOT NULL,
        created_by INT NOT NULL REFERENCES users(id),
        updated_by INT REFERENCES users(id),
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

    CREATE INDEX ON fin.payment_order_product(payment_order_id);
    CREATE INDEX ON fin.payment_order_product(product_id);
COMMIT;
