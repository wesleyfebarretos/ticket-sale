INSERT INTO users
(id, first_name, last_name, email, role, password, created_at, updated_at)
VALUES
(1, 'Ticket Sale', 'SA', 'ticketsale@gmail.com', 'super admin', '$2a$10$wefNx6VUum4nbovgKxkECeSL7.PX0SPcsMMEy9uJf6geXhMfTOSTS', NOW(), NOW());

ALTER SEQUENCE users_id_seq RESTART WITH 2;
