package model

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"log"

	"github.com/jackcooperusesvim/coopGo/model/sqlgen"
	"golang.org/x/crypto/bcrypt"
)

func ValidateToken(token string) (privledge_level string, account_id int64, err error) {

	q, ctx, err := DbInfo()
	if err != nil {
		return "", 0, err
	}

	token_hash, err := Hash(token)
	if err != nil {
		return "", 0, err
	}

	token_row, err := q.ValidateSessionToken(ctx, token_hash)

	if err != nil {
		return "", 0, err
	}
	return token_row.PriviledgeType, token_row.ID, nil

}
func Login(email, password string) (token, privledge_level string, account_id int64, err error) {

	q, ctx, err := DbInfo()
	if err != nil {
		return "", "", 0, err
	}
	password_hash, nil := Hash(password)

	res, err := q.CheckPasswordAccount(ctx, sqlgen.CheckPasswordAccountParams{
		Email:        email,
		PasswordHash: password_hash,
	})

	token = GenerateSecureToken(10)
	token_hash, err := Hash(token)

	if err != nil {
		return "", "", 0, err
	}

	_, err = q.CreateSessionToken(ctx, sqlgen.CreateSessionTokenParams{
		Token: token_hash,

		Days: sql.NullString{
			String: "1",
			Valid:  true,
		},

		Hours: sql.NullString{
			String: "0",
			Valid:  true,
		},

		Minute: sql.NullString{
			String: "0",
			Valid:  true,
		},

		AccountID: res.ID,
	})
	if err != nil {
		return "", "", 0, err
	}
	token_row, err := q.ValidateSessionToken(ctx, token_hash)

	if err != nil {
		return "", "", 0, err
	}

	return token, token_row.PriviledgeType, token_row.ID, nil

}

func UnsafeCreateAccount(email, password, privledge_type string) error {
	password_hash, err := Hash(password)
	if err != nil {
		return err
	}
	q, ctx, err := DbInfo()
	if err != nil {
		return err
	}
	_, err = q.UnsafeCreateAccount(ctx,
		sqlgen.UnsafeCreateAccountParams{
			Email:          email,
			PasswordHash:   password_hash,
			PriviledgeType: privledge_type,
		})

	if err != nil {
		return err
	}
	return nil
}

func Hash(i string) (string, error) {
	log.Println(i)
	log.Println([]byte(i))
	h, e := bcrypt.GenerateFromPassword([]byte(i), bcrypt.DefaultCost)
	return string(h), e
}

// StackOverflow Code : https://stackoverflow.com/questions/45267125/how-to-generate-unique-random-alphanumeric-tokens-in-golang
func GenerateSecureToken(length uint) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
