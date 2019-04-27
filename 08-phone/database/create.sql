CREATE SCHEMA "db";

CREATE TABLE db.phones
(
    id    SERIAL PRIMARY KEY NOT NULL,
    phone VARCHAR(30)        NOT NULL
);

INSERT INTO db.phones (phone)
VALUES ('1234567890');
INSERT INTO db.phones (phone)
VALUES ('123 456 7891');
INSERT INTO db.phones (phone)
VALUES ('(123) 456 7892');
INSERT INTO db.phones (phone)
VALUES ('(123) 456-7893');
INSERT INTO db.phones (phone)
VALUES ('123-456-7894');
INSERT INTO db.phones (phone)
VALUES ('123-456-7890');
INSERT INTO db.phones (phone)
VALUES ('1234567892');
INSERT INTO db.phones (phone)
VALUES ('(123)456-7892');
