-- name: CreatePaste :exec
INSERT INTO pastes(userid, text, expires) VALUES (
    ?, ?, ?
);

-- name: GetAllPastes :many
SELECT * FROM pastes
    WHERE
        userid = ?
        AND
        expires > CURRENT_TIMESTAMP;

-- name: UpdatePaste :one
UPDATE pastes
    SET text = ?
WHERE id = ?
AND
expires > CURRENT_TIMESTAMP
RETURNING id;

-- name: DeletePaste :one
DELETE FROM pastes
    WHERE id = ?
    AND expires > CURRENT_TIMESTAMP
RETURNING id;

-- name: DeleteAllPastesFromUser :exec
DELETE FROM pastes
    WHERE userid = ? AND expires > CURRENT_TIMESTAMP;