INSERT INTO fin.creditcard_type
(id, name, description, created_by)
VALUES
(1, 'credit', 'CREDIT', 1),
(2, 'debit', 'DEBIT', 1);

ALTER SEQUENCE fin.creditcard_type_id_seq RESTART WITH 3;
