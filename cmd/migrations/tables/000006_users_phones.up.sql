BEGIN;
    CREATE TYPE phone_types AS ENUM ('phone', 'tellphone');
    CREATE TABLE IF NOT EXISTS users_phones (
        id SERIAL PRIMARY KEY,
        user_id INT NOT NULL REFERENCES users(id),
        ddd VARCHAR(5) NOT NULL,
        number VARCHAR(10) NOT NULL,
        type phone_types NOT NULL
    );
COMMIT;
