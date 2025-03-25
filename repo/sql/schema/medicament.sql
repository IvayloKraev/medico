CREATE TYPE unit_type AS ENUM (
    'mg',
    'g'
    );

CREATE TYPE application_type AS ENUM (
    'hard_tablets',
    'soft_tables'
    );

CREATE TYPE interaction_level_type AS ENUM (
    'minor',
    'moderate',
    'fatal'
    );

CREATE TABLE active_ingredient
(
    id             uuid        NOT NULL DEFAULT gen_random_uuid(),
    official_name  text        NOT NULL,
    bulgarian_name text        NOT NULL,
    description    text        NOT NULL,
    atc            varchar(10) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE active_ingredient_interaction
(
    active_ingredient_id_1 uuid                   NOT NULL,
    active_ingredient_id_2 uuid                   NOT NULL,
    description            text                   NOT NULL,
    level                  interaction_level_type NOT NULL,
    FOREIGN KEY (active_ingredient_id_1) REFERENCES active_ingredient (id),
    FOREIGN KEY (active_ingredient_id_2) REFERENCES active_ingredient (id)
);

CREATE TABLE active_ingredient_food_interaction
(
    active_ingredient_id uuid                   NOT NULL,
    food                 text                   NOT NULL,
    description          text                   NOT NULL,
    level                interaction_level_type NOT NULL,
    FOREIGN KEY (active_ingredient_id) REFERENCES active_ingredient (id)
);

CREATE TABLE medicament
(
    id             uuid NOT NULL DEFAULT gen_random_uuid(),
    official_name  text NOT NULL,
    bulgarian_name text NOT NULL,
    description    text NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE medicament_active_ingredient
(
    medicament_id        uuid      NOT NULL,
    active_ingredient_id uuid      NOT NULL,
    quantity             float8    NOT NULL,
    unit                 unit_type NOT NULL
)