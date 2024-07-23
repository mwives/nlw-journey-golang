-- name: InsertTrip :one
INSERT INTO trips (
    "destination",
    "owner_email",
    "owner_name",
    "starts_at",
    "ends_at"
  )
VALUES ($1, $2, $3, $4, $5)
RETURNING "id";
-- name: GetTrip :one
SELECT "id",
  "destination",
  "owner_email",
  "owner_name",
  "is_confirmed",
  "starts_at",
  "ends_at"
FROM trips
WHERE "id" = $1;
-- name: UpdateTrip :exec
UPDATE trips
SET "destination" = $2,
  "starts_at" = $3,
  "ends_at" = $4,
  "is_confirmed" = $5
WHERE "id" = $1;