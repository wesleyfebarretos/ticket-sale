CREATE TABLE IF NOT EXISTS fin.creditcard (
    id SERIAL PRIMARY KEY,
    uuid UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    "number" VARCHAR(20) NOT NULL,
    expiration DATE NOT NULL,
    priority INT NOT  NULL DEFAULT 0,
    notify_expiration BOOL NOT NULL DEFAULT FALSE,
    user_id INT NOT NULL REFERENCES users(id),
    creditcard_type_id INT NOT NULL REFERENCES fin.creditcard_type(id),
    creditcard_flag_id INT NOT NULL REFERENCES fin.creditcard_flag(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE INDEX ON fin.creditcard(uuid);
