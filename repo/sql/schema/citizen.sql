CREATE TABLE citizen_auth (
    id uuid NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE citizen (
    id uuid NOT NULL,
    first_name text NOT NULL
)