-- name: CreateUser :one
INSERT INTO users(username, password) VALUES (
    ?, ?
)
RETURNING id;

-- name: GetUser :one
SELECT * FROM users
    WHERE username = ?;

-- name: UpdateUser :exec
UPDATE users
SET
    username = ?,
    password = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
    WHERE id = ?;

-- name: DoesUserExist :one
SELECT EXISTS (
    SELECT id
    FROM users
WHERE username = ?
);