-- Filename: 000010_add_image_id_to_jobs_table.up.sql

-- extends the existing jobs table from 000004
-- alters the table so it can hold either a consumer-report job or an image-processing job in the same table

BEGIN;

ALTER TABLE jobs
    ALTER COLUMN consumer_id DROP NOT NULL;

ALTER TABLE jobs
    ADD COLUMN IF NOT EXISTS image_id uuid REFERENCES images(id) ON DELETE CASCADE;

ALTER TABLE jobs
    ADD CONSTRAINT jobs_single_owner_chk CHECK (
        (consumer_id IS NOT NULL AND image_id IS NULL) OR
        (consumer_id IS NULL AND image_id IS NOT NULL)
    );

ALTER TABLE jobs
    ADD CONSTRAINT jobs_owner_matches_type CHECK (
        (job_type != 'consumer_activity_report' OR consumer_id IS NOT NULL) AND
        (job_type != 'image_processing' OR image_id IS NOT NULL)
    );

CREATE INDEX IF NOT EXISTS idx_jobs_image_id ON jobs (image_id);

COMMIT;