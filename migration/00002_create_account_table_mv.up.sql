BEGIN;

-- Create a materialized view for account events

CREATE MATERIALIZED VIEW IF NOT EXISTS account AS
SELECT DISTINCT ON (event_data ->> 'account_name')
    event_data ->> 'account_name' AS account_name,
    SPLIT_PART(event_type,'_',2) AS status,
    SPLIT_PART(stream_name,'.',2) AS ID,
    created_at AS updated_at
FROM events
WHERE event_type ~~ 'account%'
ORDER BY event_data ->> 'account_name', created_at DESC;

-- Create a index for querying events by stream

-- unique index needed for refresh materialized view concurrently

CREATE UNIQUE INDEX account_mv_idx ON account (account_name);

COMMIT;