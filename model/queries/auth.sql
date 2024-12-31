-- name: PubliclyUnaliveTokens :exec
DELETE FROM session 
WHERE session.expiration_datetime <= datetime('now');

-- name: CreateSessionToken :one
INSERT INTO session (token,expiration_datetime,account_id) VALUES (?,
    datetime(
        'now',
        '+' || @days || ' days',
        '+' || @hours || ' hours',
        '+' || @minute || ' minutes'
    ),?)
RETURNING * ;

-- name: ValidateSessionToken :one
SELECT account.id, account.priviledge_type FROM session
LEFT JOIN account
ON account.id = session.account_id
WHERE session.token = ? AND session.expiration_datetime>datetime('now');

-- name: UnsafeCreateAccount :one
INSERT INTO account 
	(email,password_hash, priviledge_type,last_updated) 
VALUES 
	(?,?,?,datetime('now')) 
RETURNING *;

-- name: CheckPasswordAccount :one
SELECT id, priviledge_type, family_id 
FROM account 
WHERE password_hash = ? AND email = ?
LIMIT 1;
