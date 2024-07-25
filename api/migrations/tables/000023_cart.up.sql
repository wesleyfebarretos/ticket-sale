CREATE TYPE cart_status AS ENUM ('active', 'checked', 'abandoned');

CREATE TABLE fin.cart (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id INT NOT NULL REFERENCES users(id),
    status cart_status NOT NULL DEFAULT 'active',
    total_amount DECIMAL(10, 2) NOT NULL,
    session_id VARCHAR(255),
    discount DECIMAL(10, 2) DEFAULT 0.00,
    promo_code VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
