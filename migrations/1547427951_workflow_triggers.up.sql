CREATE OR REPLACE FUNCTION notify_workflow()
    RETURNS trigger AS $$
DECLARE
BEGIN
    PERFORM pg_notify('workflow_insert', row_to_json(NEW)::text);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER notify_workflow
    AFTER INSERT ON workflow
    FOR EACH ROW
    EXECUTE PROCEDURE notify_workflow();
