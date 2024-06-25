DO $$
BEGIN
    -- Cria o ENUM roles apenas se não existir
    IF NOT EXISTS (
        SELECT 1
        FROM pg_type
        WHERE typname = 'roles' AND typtype = 'e'
    ) THEN
        EXECUTE 'CREATE TYPE roles AS ENUM (''admin'', ''user'', ''webservice'')';
    END IF;

    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.tables
        WHERE table_name = 'users'
    ) THEN
        CREATE TABLE users (
            id SERIAL PRIMARY KEY,
            first_name VARCHAR(50) NOT NULL,
            last_name VARCHAR(50) NOT NULL,
            email VARCHAR(255) NOT NULL UNIQUE,
            role roles NOT NULL,  -- Utiliza o ENUM roles criado acima
            password VARCHAR(500) NOT NULL,
            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        );
    END IF;
END $$;
