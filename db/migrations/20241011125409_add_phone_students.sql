-- +goose Up
ALTER TABLE public.students
ADD COLUMN phone_students VARCHAR;

-- +goose Down
ALTER TABLE public.students
DROP COLUMN IF EXISTS phone_students;
