-- +migrate Up
create table public.permissions (
    id bigserial primary key,
    name varchar(255) unique,
    created_at timestamp(6) not null default now(),
    updated_at timestamp(6) null
);