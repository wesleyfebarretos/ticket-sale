INSERT INTO fin.payment_cancel_reason
(id, created_by, name)
VALUES
(1, 1, 'Withdrawal'),
(2, 1, 'Health'),
(3, 1, 'Travel'),
(4, 1, 'Work'),
(5, 1, 'Change of Payment Method'),
(6, 1, 'Change of Plan'),
(7, 1, 'Change of Due Date'),
(8, 1, 'Deceased'),
(9, 1, 'Change of Location'),
(10, 1, 'Does Not Wish Renewal'),
(11, 1, 'Chargeback'),
(12, 1, 'Others');

ALTER SEQUENCE fin.payment_cancel_reason_id_seq RESTART WITH 13;
