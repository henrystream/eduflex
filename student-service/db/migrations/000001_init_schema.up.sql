CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE students (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name   TEXT NOT NULL,
    last_name    TEXT NOT NULL,
    date_of_birth DATE NOT NULL,
    email        TEXT UNIQUE NOT NULL,
    phone        TEXT,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


