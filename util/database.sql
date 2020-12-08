CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE status AS ENUM ('inserted', 'consumed');

CREATE TABLE workflows (
    uuid VARCHAR PRIMARY KEY DEFAULT uuid_generate_v4(),
    status STATUS NOT NULL DEFAULT 'inserted',
    data jsonb NOT NULL,
    steps jsonb NOT NULL
);