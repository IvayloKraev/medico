CREATE TYPE moderator_type AS ENUM (
    'doctor',
    'pharmacy',
    'citizen',
    'medicament'
    );


CREATE TABLE moderator_auth
(
    id       uuid           NOT NULL DEFAULT gen_random_uuid(),
    email    text           NOT NULL,
    password bytea          NOT NULL,
    type     moderator_type NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE moderator
(
    id         uuid           NOT NULL,
    first_name text           NOT NULL,
    last_name  text           NOT NULL,
    email      text           NOT NULL,
    password   bytea          NOT NULL,
    type       moderator_type NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (id) REFERENCES moderator_auth (id) ON DELETE CASCADE
);