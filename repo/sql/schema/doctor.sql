CREATE TABLE doctor_auth
(
    id       uuid NOT NULL DEFAULT gen_random_uuid(),
    email    text NOT NULL,
    password text NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE doctor
(
    id           uuid        NOT NULL,
    first_name   text        NOT NULL,
    second_name  text        NOT NULL,
    last_name    text        NOT NULL,
    uin          varchar(10) NOT NULL CHECK ( length(uin) == 10 ),
    phone_number varchar(13) NOT NULL CHECK ( length(uin) >= 10 ),
    email        text        NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (id) REFERENCES doctor_auth (id)
)