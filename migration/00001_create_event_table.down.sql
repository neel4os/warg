BEGIN;
DROP TRIGGER IF EXISTS set_version ON events;

DROP FUNCTION IF EXISTS increment_version;

DROP INDEX IF EXISTS idx_created_at;
DROP INDEX IF EXISTS idx_event_type;
DROP INDEX IF EXISTS idx_stream_name;
DROP INDEX IF EXISTS idx_stream_id;
DROP INDEX IF EXISTS idx_stream_version;

DROP TABLE IF EXISTS events;
COMMIT;