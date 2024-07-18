INSERT INTO fin.payment_period
(id, created_by, name, times)
VALUES
(1, 1, '1x', 1),
(2, 1, '2x', 2),
(3, 1, '3x', 3),
(4, 1, '4x', 4),
(5, 1, '5x', 5),
(6, 1, '6x', 6),
(7, 1, '7x', 7),
(8, 1, '8x', 8),
(9, 1, '9x', 9),
(10, 1, '10x', 10),
(11, 1, '11x', 11),
(12, 1, '12x', 12);

ALTER SEQUENCE fin.payment_period_id_seq RESTART WITH 13;
