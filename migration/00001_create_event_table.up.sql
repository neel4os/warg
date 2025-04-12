BEGIN;
CREATE TABLE IF NOT EXISTS events (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY, -- Unique identifier for the event
    stream_id UUID NOT NULL, -- Identifier for the event stream (e.g., aggregate ID)
    stream_name TEXT NOT NULL, -- Name of the event stream (e.g., aggregate type)
    event_type TEXT NOT NULL, -- Type of the event (e.g., "OrderCreated", "OrderShipped")
    event_data JSONB NOT NULL, -- Event payload (stored as JSON)
    metadata JSONB, -- Optional metadata (e.g., correlation ID, causation ID)
    version INT NOT NULL, -- Version of the stream (used for optimistic concurrency)
    initiator_type TEXT NOT NULL, -- Type of the initiator (e.g., "User", "System")
    initiator_name TEXT NOT NULL, -- Name of the initiator (e.g., "John Doe", "Order Service")
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() -- Timestamp of when the event was created
);

-- Ensure that each stream has a unique version for optimistic concurrency
CREATE UNIQUE INDEX idx_stream_version ON events (stream_id, version);

-- Create an index for querying events by stream
CREATE INDEX idx_stream_id ON events (stream_id);

-- Create an index for querying events by stream name
CREATE INDEX idx_stream_name ON events (stream_name);

-- Create an index for querying events by event type
CREATE INDEX idx_event_type ON events (event_type);

-- Create an index for querying events by creation time
CREATE INDEX idx_created_at ON events (created_at);

-- Create a trigger function to auto-increment the version
CREATE OR REPLACE FUNCTION increment_version()
RETURNS TRIGGER AS $$
BEGIN
    NEW.version := COALESCE(
        (SELECT MAX(version) + 1 FROM events WHERE stream_id = NEW.stream_id),
        1
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Attach the trigger to the events table
CREATE TRIGGER set_version
BEFORE INSERT ON events
FOR EACH ROW
EXECUTE FUNCTION increment_version();
COMMIT;