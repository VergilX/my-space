-- name: CreateCSRFToken :exec
INSERT INTO csrf(userid, token) VALUES (
    ?, ?
);

-- name: GetCSRFToken :one
SELECT * FROM csrf
    WHERE userid = ?;

-- name: RenewCSRFToken :exec
UPDATE csrf
SET
    token = ?,
    expiry = ?
WHERE userid = ?;

-- name: DeleteCSRFToken :exec
DELETE FROM csrf
    WHERE userid = ?;

-- name: VerifyCSRFToken :one
SELECT EXISTS (
    SELECT id
    FROM csrf
WHERE token = ?
);