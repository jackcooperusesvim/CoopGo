-- name: PubliclyExecuteTokens :exec
DELETE FROM session 
WHERE session.expiration_datetime <= ?;

-- name: CreateSessionToken :one
INSERT INTO session (token,expiration_datetime,account_id) VALUES (?,?,?,?) RETURNING *;

-- name: ValidateSessionToken :one
SELECT account.id FROM session
LEFT JOIN account
ON account.id = session.account_id
WHERE account.priviledge_type = $1 AND session.token = $2

