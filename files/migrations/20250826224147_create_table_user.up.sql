-- +migrate Up
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS public.users (
    id                    BIGSERIAL PRIMARY KEY,
    full_name             VARCHAR(255),
    email                 CITEXT UNIQUE NOT NULL,              -- case-insensitive unique email
    password              VARCHAR(255),
    pin                   VARCHAR(255),
    token_activation      VARCHAR(255),
    token_activation_expired_at TIMESTAMP(6),
    token_forgot_password VARCHAR(255),
    token_forgot_pin      VARCHAR(255),
    status                SMALLINT NOT NULL DEFAULT 0 CHECK (status IN (0,1,2)), -- 0=pending,1=active,2=deleted
    role_id               BIGINT NOT NULL,                     -- Postgres has no UNSIGNED; use BIGINT
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_users_to_role
    FOREIGN KEY (role_id) REFERENCES public.roles(id)
    ON UPDATE CASCADE ON DELETE RESTRICT
);