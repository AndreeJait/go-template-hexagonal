-- +migrate Up
CREATE TABLE IF NOT EXISTS public.role_permissions (
    role_id               BIGINT NOT NULL,                     -- Postgres has no UNSIGNED; use BIGINT
    permission_id         BIGINT NOT NULL,
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT pk_role_id_and_permission_id PRIMARY KEY (role_id, permission_id),
    CONSTRAINT fk_role_permission_to_role
    FOREIGN KEY (role_id) REFERENCES public.roles(id)
    ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT fk_role_permission_to_permission
    FOREIGN KEY (permission_id) REFERENCES public.permissions(id)
    ON UPDATE CASCADE ON DELETE RESTRICT
);