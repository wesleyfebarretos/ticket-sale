INSERT INTO product_categories
(id, name, created_by)
VALUES
(1, 'Digital', 1),
(2, 'Physical', 1),
(3, 'Event', 1);

ALTER SEQUENCE product_categories_id_seq RESTART WITH 4;
