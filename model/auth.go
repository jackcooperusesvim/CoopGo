package model

import (
	"github.com/jackcooperusesvim/coopGo/model/sqlgen"
	"golang.org/x/crypto/bcrypt"
)

func ValidateToken(token string) (privledge_level string, account_id int64, err error) {

	q, ctx, err := DbInfo()
	if err != nil {
		return "", 0, err
	}

	token_row, err := q.ValidateSessionToken(ctx, token)

	if err != nil {
		return "", 0, err
	}
	return token_row.PriviledgeType.String, token_row.ID.Int64, nil

}

func UnsafeCreateAccount(email, password, privledge_type string) error {
	password_hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
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
			PasswordHash:   string(password_hash),
			PriviledgeType: privledge_type,
		})

	if err != nil {
		return err
	}
	return nil
}
