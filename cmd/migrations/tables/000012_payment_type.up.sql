CREATE TABLE IF NOT EXISTS fin.payment_type (
    id SERIAL PRIMARY KEY,
    uuid UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500),
    active BOOL NOT NULL DEFAULT true,
    allow_installment NOT NULL DEFAULT false,
    created_by INT NOT NULL REFERENCES users(id),
    updated_by INT REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE INDEX ON fin.payment_type(uuid);
