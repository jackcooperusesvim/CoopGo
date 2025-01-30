// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: auth.sql

package sqlgen

import (
	"context"
	"database/sql"
)

const createSessionToken = `-- name: CreateSessionToken :one
INSERT INTO session (token,expiration_datetime,account_id) VALUES (?,
    datetime(
        'now',
        '+' || ? || ' days',
        '+' || ? || ' hours',
        '+' || ? || ' minutes'
    ),?)
RETURNING id, token, expiration_datetime, account_id
`

type CreateSessionTokenParams struct {
	Token     string
	Days      sql.NullString
	Hours     sql.NullString
	Minute    sql.NullString
	AccountID int64
}

func (q *Queries) CreateSessionToken(ctx context.Context, arg CreateSessionTokenParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSessionToken,
		arg.Token,
		arg.Days,
		arg.Hours,
		arg.Minute,
		arg.AccountID,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.ExpirationDatetime,
		&i.AccountID,
	)
	return i, err
}

const getAccountInfo = `-- name: GetAccountInfo :one
SELECT id, priviledge_type, password_hash
FROM account
WHERE email = ?
LIMIT 1
`

type GetAccountInfoRow struct {
	ID             int64
	PriviledgeType string
	PasswordHash   string
}

func (q *Queries) GetAccountInfo(ctx context.Context, email string) (GetAccountInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getAccountInfo, email)
	var i GetAccountInfoRow
	err := row.Scan(&i.ID, &i.PriviledgeType, &i.PasswordHash)
	return i, err
}

const getSessionToken = `-- name: GetSessionToken :one
SELECT account.id, account.priviledge_type, session.token FROM session
INNER JOIN account
ON account.id = session.account_id
WHERE session.token = ?
LIMIT 1
`

type GetSessionTokenRow struct {
	ID             int64
	PriviledgeType string
	Token          string
}

func (q *Queries) GetSessionToken(ctx context.Context, token string) (GetSessionTokenRow, error) {
	row := q.db.QueryRowContext(ctx, getSessionToken, token)
	var i GetSessionTokenRow
	err := row.Scan(&i.ID, &i.PriviledgeType, &i.Token)
	return i, err
}

const publiclyUnaliveTokens = `-- name: PubliclyUnaliveTokens :exec
DELETE FROM session 
WHERE session.expiration_datetime <= datetime('now')
`

func (q *Queries) PubliclyUnaliveTokens(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, publiclyUnaliveTokens)
	return err
}

const unsafeCreateAccount = `-- name: UnsafeCreateAccount :one
INSERT INTO account 
	(email,password_hash, priviledge_type,last_updated) 
VALUES 
	(?,?,?,datetime('now')) 
RETURNING id, email, password_hash, priviledge_type, last_updated, family_id
`

type UnsafeCreateAccountParams struct {
	Email          string
	PasswordHash   string
	PriviledgeType string
}

func (q *Queries) UnsafeCreateAccount(ctx context.Context, arg UnsafeCreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, unsafeCreateAccount, arg.Email, arg.PasswordHash, arg.PriviledgeType)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.PriviledgeType,
		&i.LastUpdated,
		&i.FamilyID,
	)
	return i, err
}

const validateToken = `-- name: ValidateToken :one
;

SELECT account.id, account.priviledge_type, session.token FROM session
INNER JOIN account
ON account.id = session.account_id
WHERE session.token = ?
`

type ValidateTokenRow struct {
	ID             int64
	PriviledgeType string
	Token          string
}

func (q *Queries) ValidateToken(ctx context.Context, token string) (ValidateTokenRow, error) {
	row := q.db.QueryRowContext(ctx, validateToken, token)
	var i ValidateTokenRow
	err := row.Scan(&i.ID, &i.PriviledgeType, &i.Token)
	return i, err
}
