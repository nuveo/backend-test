CREATE TABLE TEST_WORKFLOW (
    UUID uuid primary key,
    status status_enum,
    data jsonb,
    steps text[]
);
