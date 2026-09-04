-- Filename: 000009_create_variants_table.down.sql

-- stores metadata for the three outputs (thumbnail, preview and display)

BEGIN;

CREATE TABLE IF NOT EXISTS variants(
    id               uuid               PRIMARY KEY DEFAULT uuidv7(),
    image_id         uuid               NOT NULL REFERENCES images(id) ON DELETE CASCADE,
    name             text               NOT NULL CHECK (name IN ('thumbnail', 'preview', 'display')),
    stored_filename  text               NOT NULL UNIQUE,
    width            integer            NOT NULL,
    height           integer            NOT NULL,
    size_bytes       bigint             NOT NULL,
    created_at       timestamptz        NOT NULL DEFAULT now(),
    UNIQUE (image_id, name)
);

COMMIT;