INSERT INTO fin.payment_type
(id, name, active, allow_installment, allow_recurrence, created_by, uuid)
VALUES
(1, 'Credit Card', true, true, true, 1, 'b7259db0-a938-4b81-ae74-6b1f9460be29'),
(2, 'Payment Slip', true, false, false, 1, 'bc2c90b5-e5f1-4928-ad1d-18d1cd0bbab7'),
(3, 'Pix', true, false, false, 1, '4406694e-d323-49d7-8ccd-e894512b1317');

ALTER SEQUENCE fin.payment_type_id_seq RESTART WITH 4;
