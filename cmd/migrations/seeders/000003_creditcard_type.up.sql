INSERT INTO fin.creditcard_type
(id, name, created_by)
VALUES
(1, 'Credit', 1),
(2, 'Debit', 1);

ALTER SEQUENCE fin.creditcard_type_id_seq RESTART WITH 3;
