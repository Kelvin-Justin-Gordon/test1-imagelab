-- Filename: 000010_add_image_id_to_jobs_table.down.sql

BEGIN;

DROP INDEX IF EXISTS idx_jobs_image_id;

ALTER TABLE jobs
    DROP CONSTRAINT IF EXISTS jobs_owner_matches_type;

ALTER TABLE jobs
    DROP CONSTRAINT IF EXISTS jobs_single_owner_chk;

ALTER TABLE jobs
    DROP COLUMN IF EXISTS image_id;

ALTER TABLE jobs
    ALTER COLUMN consumer_id SET NOT NULL;

COMMIT; 