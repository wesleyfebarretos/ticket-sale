BEGIN;
    CREATE TYPE gateway_auth_type AS ENUM ('bearer', 'basic');

    CREATE TABLE fin.gateway (
        id SERIAL PRIMARY KEY NOT NULL,
        uuid UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
        name VARCHAR(50) NOT NULL,
        description TEXT NULL,
        client_id TEXT NULL,
        client_secret TEXT NULL,
        "order" INT NOT NULL DEFAULT 0,
        active BOOL NOT NULL DEFAULT true,
        is_deleted BOOL NOT NULL DEFAULT FALSE,
        test_environment BOOL NOT NULL DEFAULT false,
        notif_user TEXT NULL,
        notif_password TEXT NULL,
        soft_descriptor TEXT NULL,
        gateway_process_id INT NOT NULL REFERENCES fin.gateway_process(id),
        webhook_url TEXT NULL,
        url VARCHAR(255) NULL,
        auth_type_id gateway_auth_type NOT NULL,
        use_3ds BOOL NOT NULL DEFAULT false,
        adq_code_3ds TEXT NULL,
        default_adq_code TEXT NULL,
        use_antifraud BOOL NOT NULL DEFAULT false,
        created_by INT NOT NULL REFERENCES users(id),
        updated_by INT REFERENCES users(id),
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
    );

    CREATE INDEX ON fin.gateway(uuid);
COMMIT;

