-- +goose Up
ALTER TABLE public.users
ADD COLUMN email VARCHAR NOT NULL,
ADD COLUMN password VARCHAR NOT NULL;

-- +goose Down
ALTER TABLE public.users
DROP COLUMN IF EXISTS email,
DROP COLUMN IF EXISTS password;
