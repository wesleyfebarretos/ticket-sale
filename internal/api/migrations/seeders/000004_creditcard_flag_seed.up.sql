INSERT INTO fin.creditcard_flag
(id, created_by, name, description, regex)
VALUES
(1, 1, 'visa', 'VISA', '''^4d{5}'''),
(2, 1, 'mastercard', 'MASTER', '^5d{5}'),
(3, 1, 'american-express', 'AMEX', '^3[47]d{4}'),
(4, 1, 'diners', 'AMEX', '^(30[0-5]|36|38)d{3}'),
(5, 1, 'jcb', 'JCB', '^(2131|2132|1800|35d{1,2})d{3}'),
(6, 1, 'visa-electron', 'VISA ELECTRON', '^(4026|417500|4405)'),
(7, 1, 'maestro', 'MAESTRO', '^(5018|5020|5038|56|58)'),
(8, 1, 'discover', 'DISCOVER', '^(6011|622(12[6-9]|1[3-9][0-9]|[2-8][0-9]{2}|9[0-2][0-5])|64[4-9])'),
(9, 1, 'elo', 'ELO', '^(6363|6364|6365|504175)d{2}$'),
(10, 1, 'hipercard', 'HIPERCARD', '^606282');

ALTER SEQUENCE fin.creditcard_flag_id_seq RESTART WITH 11;
