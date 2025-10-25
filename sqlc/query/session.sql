-- name: CreateSessionToken :exec
INSERT INTO session(userid, token) VALUES (
    ?, ?
);

-- name: GetSessionToken :one
SELECT token FROM session
    WHERE userid = ?;

-- name: GetUserIDFromSessionToken :one
SELECT userid FROM session
    WHERE token = ?;

-- name: RenewSessionToken :exec
UPDATE session
SET
    token = ?,
    expiry = ?
WHERE userid = ?;

-- name: DeleteSessionToken :exec
DELETE FROM session
    WHERE userid = ?;