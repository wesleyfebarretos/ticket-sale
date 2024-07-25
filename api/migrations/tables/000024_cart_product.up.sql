CREATE TABLE IF NOT EXISTS fin.cart_product (
    id SERIAL PRIMARY KEY,
    cart_id UUID NOT NULL REFERENCES fin.cart(id),
    product_id INT NOT NULL REFERENCES products(id),
    qty INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);