CREATE TABLE entities (
    id serial primary key,
    name varchar(255),
    support integer default 0,
    against integer default 0
)