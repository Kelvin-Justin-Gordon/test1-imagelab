## Week 1

**Completed since last check-in**

- Copied and understood the supplied Gatekeeper starter code (Go backend,
  original migrations, project structure)

- Wrote migrations `000008`–`000010` by hand: `images` table, `variants`
  table, and extended the existing `jobs` table with a nullable `image_id`,
  two `CHECK` constraints.

- Implemented `POST /v1/images`: server-side validation via actually
  decoding the file, server-controlled random filename generation, original
  storage on the filesystem, and the durable `images` database record.

- Built the frontend: initial / image-selected / uploading states, local
  file preview with no premature upload, and a disabled-button +
  `isSubmitting` guard against overlapping submissions.

- Fixed several real bugs found during review: an orphaned-file cleanup
  gap when the database insert fails after the file is already written, a
  missing scan destination and wrong field types in the variants model,
  and a couple of JavaScript typos that silently broke the preview and
  submit flow.

  - Restyled the frontend toward a two-column card layout matching the
  visual direction shown before the assignment started.

- Ran the server, uploaded real images via the browser and via `curl`,
  and confirmed the stored file and database row matched in each case