-- +goose Up
ALTER TABLE public.users
ADD COLUMN role VARCHAR;

-- +goose Down
ALTER TABLE public.users
DROP COLUMN IF EXISTS role,
