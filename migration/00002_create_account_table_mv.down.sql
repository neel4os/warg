BEGIN;
-- Drop the materialized view if it exists
DROP MATERIALIZED VIEW IF EXISTS account;
-- Drop the index if it exists
DROP INDEX IF EXISTS account_mv_idx;

COMMIT;