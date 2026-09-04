-- Filename: 000008_create_images_table.up.sql

--holds information and metadata about every uploaded original image

BEGIN;

CREATE TABLE IF NOT EXISTS images(
    id                  uuid          PRIMARY KEY DEFAULT uuidv7(),
    original_filename   text          NOT NULL,
    stored_filename     text          NOT NULL UNIQUE,
    media_type          text          NOT NULL,
    size_bytes          bigint        NOT NULL,
    created_at          timestamptz   NOT NULL DEFAULT now()
);

COMMIT;