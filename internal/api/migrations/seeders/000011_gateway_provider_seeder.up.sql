INSERT INTO fin.gateway_provider
(id, name, created_by, updated_by)
VALUES
(1, 'stripe', 1, 1);

ALTER SEQUENCE fin.gateway_provider_id_seq RESTART WITH 2;
