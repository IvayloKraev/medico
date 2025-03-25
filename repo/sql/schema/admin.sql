CREATE TABLE admin_auth
(
    id       uuid  NOT NULL DEFAULT gen_random_uuid(),
    email    text  NOT NULL,
    password bytea NOT NULL,
    PRIMARY KEY (id)
);
