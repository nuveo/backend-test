CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE STATUS AS ENUM ('inserted', 'consumed');

CREATE TABLE IF NOT EXISTS workflow (
	uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
	status STATUS DEFAULT 'inserted' NOT NULL,
	data JSONB NOT NULL,
	steps text[] NOT NULL
);