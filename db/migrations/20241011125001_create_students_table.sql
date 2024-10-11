-- +goose Up
CREATE TABLE IF NOT EXISTS public.students (
    id_students SERIAL PRIMARY KEY,
    name_students VARCHAR NOT NULL,
    email_students VARCHAR NOT NULL,
    password_students VARCHAR NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS public.students;
