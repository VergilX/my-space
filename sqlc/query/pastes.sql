-- name: CreatePaste :exec
INSERT INTO pastes(userid, text) VALUES (
    ?, ?
);

-- name: GetAllPastes :many
SELECT text FROM pastes
    WHERE userid = ?;

-- name: UpdatePaste :exec
UPDATE pastes
    SET text = ?
WHERE id = ?;

-- name: DeletePaste :exec
DELETE FROM pastes
    WHERE id = ?;

-- name: DeleteAllClipsFromUser :exec
DELETE FROM pastes
    WHERE userid = ?;