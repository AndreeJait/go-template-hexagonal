-- +migrate Up
INSERT INTO roles (id, name)
VALUES (1, 'FINANCE'),
       (2, 'TELLER'),
       (3, 'MERCHANT'),
       (4, 'ADMIN');