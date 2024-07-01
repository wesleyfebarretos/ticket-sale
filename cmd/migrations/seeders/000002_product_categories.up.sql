INSERT INTO product_categories
(id, name, description, created_by)
VALUES
(1, 'digital', 'DIGITAL', 1),
(2, 'physical', 'PHYSICAL', 1),
(3, 'event', 'EVENT', 1);

ALTER SEQUENCE product_categories_id_seq RESTART WITH 4;
