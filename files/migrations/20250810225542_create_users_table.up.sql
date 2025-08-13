-- +migrate Up
create table public.users (
       id bigserial primary key,
       email varchar(255) not null unique,
       password varchar(255) not null,
       created_at timestamp default now(),
       updated_at timestamp default now()
);