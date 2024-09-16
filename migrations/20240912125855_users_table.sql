-- +goose Up
create table if not exists users(
    id serial primary key,
    "name" varchar not null,
    age integer not null
);

-- +goose Down
drop table if exists users
