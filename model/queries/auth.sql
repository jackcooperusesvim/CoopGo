-- name: PubliclyExecuteTokens :exec
DELETE FROM session 
WHERE session.expiration_datetime <= ?;

-- name: CreateSessionToken :one
INSERT INTO session (token,expiration_datetime,account_id) VALUES (?,?,?) RETURNING *;

-- name: ValidateSessionToken :one
SELECT account.id, account.priviledge_type FROM session
LEFT JOIN account
ON account.id = session.account_id
WHERE account.priviledge_type = ? AND session.token = ?;

-- name: UnsafeCreateAccount :one
INSERT INTO account (email,password_hash, priviledge_type,last_updated) VALUES (?,?,?,CURRENT_TIMESTAMP) RETURNING *;
