CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    uuid UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    price DOUBLE PRECISION NOT NULL,
    discount_price DOUBLE PRECISION,
    active bool NOT NULL DEFAULT false,
    is_deleted bool NOT NULL DEFAULT false,
    image TEXT,
    image_mobile TEXT,
    image_thumbnail TEXT,
    category_id INT NOT NULL REFERENCES product_categories(id),
    created_by INT NOT NULL REFERENCES users(id),
    updated_by INT REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE INDEX ON products(uuid);
