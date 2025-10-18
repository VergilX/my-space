-- name: CreateClip :exec
INSERT INTO clips(userid, text) VALUES (
    ?, ?
);

-- name: GetClipContent :one
SELECT text FROM clips
    WHERE id = ?;

-- name: UpdateClip :exec
UPDATE clips
    SET text = ?
WHERE userid = ?;

-- name: DeleteClip :exec
DELETE FROM clips
    WHERE userid = ?;