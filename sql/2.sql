CREATE TABLE stream_history (
	id SERIAL PRIMARY KEY,
	platform VARCHAR(255),
	user_identifier VARCHAR(255),
	stream_identifier VARCHAR(255),
	created_at TIMESTAMP
);
CREATE TABLE stream_notification (
    id SERIAL PRIMARY KEY,
    platform VARCHAR(255) NOT NULL,
    guild VARCHAR(255) NOT NULL,
    channel VARCHAR(255) NOT NULL,
    stream_platform VARCHAR(255) NOT NULL,
    user_identifier VARCHAR(255) NOT NULL,
    user_unique_id VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
ALTER TABLE stream_notification
ADD metadata JSONB; 



-- CREATE UNLOGGED TABLE stream_notif_delay (
--     key text UNIQUE NOT NULL,
--     expired_at timestamp);

-- CREATE INDEX idx_stream_notif_delay_key ON stream_notif_delay (key);

-- CREATE OR REPLACE PROCEDURE delete_notif_delay () AS
-- $$
-- BEGIN
--     DELETE FROM stream_notif_delay
--     WHERE expired_at < NOW();

--     COMMIT;
-- END;
-- $$ LANGUAGE plpgsql;

-- create EXTENSION pg_cron;

-- SELECT cron.schedule('* * * * *', $$CALL delete_notif_delay();$$);