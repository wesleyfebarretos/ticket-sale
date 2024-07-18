INSERT INTO fin.gateway_process
(id, name, created_by)
VALUES
(1, 'Authorize and Capture', 1),
(2, 'Authorize Only', 1);

ALTER SEQUENCE fin.gateway_process_id_seq RESTART WITH 3;
