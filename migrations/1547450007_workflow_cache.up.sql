CREATE TABLE CACHE_WORKFLOW (
    UUID uuid primary key,
    status status_enum,
    data jsonb,
    steps text[],
    csv text
);
