CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

------------------------------------------------------------
-- SCHOOLS
------------------------------------------------------------
CREATE TABLE schools (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    address TEXT,
    contact_email TEXT,
    contact_phone TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

------------------------------------------------------------