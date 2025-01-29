package model

import (
	"crypto/rand"
	"crypto/sha256"
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

	// hash_rows, err := q.GetSimilarSessionTokens(ctx, token[:3]+"%")
	// log.Println(token[:3] + "%")

	// if err != nil {
	// 	return "", 0, err
	// }
	// for _, hash_row := range hash_rows {
	//
	// 	if bcrypt.CompareHashAndPassword([]byte(hash_row.Token), []byte(token)) == nil {
	// 		log.Println(token[:3] + "%")
	// 		return hash_row.PriviledgeType, hash_row.ID, nil
	// 	}
	// }

	log.Println("ValidateToken")
	val_row, err := q.ValidateToken(ctx, HashNoSalt(token))

	if err != nil {
		return "", 0, err
	}
	return val_row.PriviledgeType, val_row.ID, nil
}

func Login(email, password string) (token, privledge_level string, account_id int64, err error) {

	q, ctx, err := DbInfo()
	if err != nil {
		return "", "", 0, err
	}

	account_row, err := q.GetAccountInfo(ctx, email)

	if err != nil {
		return "", "", 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(account_row.PasswordHash), []byte(password))

	if err != nil {
		log.Println(err)
		return "", "", 0, err
	}

	token = GenerateSecureToken(10)
	token_hash := HashNoSalt(token)

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

		AccountID: account_row.ID,
	})
	if err != nil {
		return "", "", 0, err
	}

	return token, account_row.PriviledgeType, account_row.ID, nil

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
	h, e := bcrypt.GenerateFromPassword([]byte(i), bcrypt.DefaultCost)
	return string(h), e
}
func HashNoSalt(i string) string {
	hasher := sha256.New()
	hasher.Write([]byte(i))
	return hex.EncodeToString(hasher.Sum(nil))
}

// StackOverflow Code : https://stackoverflow.com/questions/45267125/how-to-generate-unique-random-alphanumeric-tokens-in-golang
func GenerateSecureToken(length uint) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
