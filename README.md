# ImageLab — Assessment 1 (Version 1)

Async `202 Accepted` + durable Job + one worker + 1-second short polling,
applied to image upload and variant generation. Built by extending the
Gatekeeper async-report starter provided for CMPS4191 — same
`202 → job → worker → poll` shape, applied to a new domain (image
uploads instead of consumer-activity reports).

## Prerequisites

- Go (see `go.mod` for version)
- PostgreSQL 18+ (for native `uuidv7()`/`uuidv4()` support)
- `psql`
- The `migrate` CLI

## Setup

```bash
# 1. Create the database and a dedicated role (not the postgres superuser)
psql -U postgres -h localhost -c "CREATE DATABASE imagelab;"
psql -U postgres -h localhost -c "CREATE ROLE imagelab_app WITH LOGIN PASSWORD 'yourpassword';"
psql -U postgres -h localhost -c "ALTER DATABASE imagelab OWNER TO imagelab_app;"

# 2. Set up your local environment file
cp .envrc.example .envrc
# edit .envrc with the real password you set above
source .envrc

# 3. Enable required extensions
psql "$IMAGELAB_DB_DSN" -c "CREATE EXTENSION IF NOT EXISTS citext;"
psql "$IMAGELAB_DB_DSN" -c "SELECT uuidv7(), uuidv4();"  # confirms native UUID support

# 4. Fetch Go dependencies
go mod tidy

# 5. Apply all migrations
migrate -path migrations -database "$IMAGELAB_DB_DSN" up
```

## Run it

```bash
go run ./cmd/api -db-dsn="$IMAGELAB_DB_DSN"
```

Then open `http://localhost:4000` in a browser. The Go server serves the
frontend as static files from the same origin as the API (`-frontend-dir`,
default `./frontend`), so there's no separate dev server and no CORS
configuration needed. Uploaded originals land in `./uploads/originals` by
default (`-upload-dir` to change it).

## Try the upload path

```bash
curl -i -X POST http://localhost:4000/v1/images \
  -F "image=@/path/to/a/photo.jpg"
```

Expect `201 Created` with the stored image's `id`, `media_type`,
`size_bytes`, and `created_at`. An oversized file gets
`413 Request Entity Too Large`; a file that isn't a decodable JPEG or PNG
gets `415 Unsupported Media Type`. Neither creates an `images` row, and any
partially-written file is cleaned up automatically if the database insert
fails.

