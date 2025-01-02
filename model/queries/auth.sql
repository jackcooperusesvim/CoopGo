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

-- name: GetSimilarSessionTokens :many
SELECT account.id, account.priviledge_type, session.token FROM session
INNER JOIN account
ON account.id = session.account_id
WHERE session.token LIKE (@token_beginning_with_wildcard) AND session.expiration_datetime>datetime('now');

-- name: GetSessionTokens :many
SELECT account.id, account.priviledge_type, session.token FROM session
INNER JOIN account
ON account.id = session.account_id
WHERE session.expiration_datetime>datetime('now');

-- name: UnsafeCreateAccount :one
INSERT INTO account 
	(email,password_hash, priviledge_type,last_updated) 
VALUES 
	(?,?,?,datetime('now')) 
RETURNING *;

-- name: GetAccountInfo :one
SELECT id, priviledge_type, password_hash
FROM account
WHERE email = ?
LIMIT 1;
