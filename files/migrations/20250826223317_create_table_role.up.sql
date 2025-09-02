-- +migrate Up
CREATE TABLE public.roles
(
    id         bigserial PRIMARY KEY,
    name       varchar UNIQUE NOT NULL,
    created_at timestamp(6)   NOT NULL DEFAULT now(),
    updated_at timestamp(6)   NULL
);