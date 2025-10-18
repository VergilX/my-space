-- name: CreateUser :exec
INSERT INTO users(username, password) VALUES (
    ?, ?
);

-- name: GetUser :one
SELECT * FROM users
    WHERE id = ?;

-- name: UpdateUser :exec
UPDATE users
SET
    username = ?,
    password = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
    WHERE id = ?;