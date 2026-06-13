-- name: CreateUser :one
INSERT INTO users (name, dob)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetAllUsers :many
SELECT *
FROM users
ORDER BY id;

-- name: UpdateUser :exec
UPDATE users
SET name = $1,
    dob = $2
WHERE id = $3;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;