CREATE TABLE IF NOT EXISTS fin.product_payment_type_installment_time (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES products(id),
    payment_type_id INT NOT NULL REFERENCES fin.payment_type(id),
    installment_time_id INT NOT NULL REFERENCES fin.installment_time(id),
    fee DECIMAL(5,2),
    tariff DECIMAL(5,2),
    created_by INT NOT NULL REFERENCES users(id),
    updated_by INT REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
