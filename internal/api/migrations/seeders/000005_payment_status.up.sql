INSERT INTO fin.payment_status
(id, created_by, name, description)
VALUES
(1, 1, 'AP', 'AWAITING PAYMENT'),
(2, 1, 'TP', 'PAYMENT ATTEMPT'),
(3, 1, 'PA', 'PAID'),
(4, 1, 'PC', 'PENDING CANCELLATION'),
(5, 1, 'CA', 'CANCELLED'),
(6, 1, 'EX', 'EXPIRED - PAYMENT SLIP');

ALTER SEQUENCE fin.payment_status_id_seq RESTART WITH 7;
