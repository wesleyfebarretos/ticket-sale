INSERT INTO fin.gateway_payment_type
(id, name, created_by)
VALUES
(1, 'Installment', 1),
(2, 'Recurrent', 1);

ALTER SEQUENCE fin.gateway_payment_type_id_seq RESTART WITH 3;
