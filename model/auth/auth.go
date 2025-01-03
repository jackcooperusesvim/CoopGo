package auth

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"log"
	"os"

	"github.com/jackcooperusesvim/coopGo/model/sqlgen"
	"github.com/jackcooperusesvim/coopGo/model/token"
	"github.com/jackcooperusesvim/coopGo/model/util"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
)

func Login(email, password string) (token, privledge_level string, account_id int64, err error) {

	q, ctx, err := util.DbInfo()
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
	token_hash, err := Hash([]byte(token))

	if err != nil {
		log.Println(err)
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

		AccountID: account_row.ID,
	})
	if err != nil {
		return "", "", 0, err
	}

	return token, account_row.PriviledgeType, account_row.ID, nil

}

func UnsafeCreateAccount(email, password, privledge_type string) error {
	password_hash, err := Hash([]byte(password))
	if err != nil {
		return err
	}
	q, ctx, err := util.DbInfo()
	if err != nil {
		return err
	}
	_, err = q.UnsafeCreateAccount(ctx,
		sqlgen.UnsafeCreateAccountParams{
			Email:          email,
			PasswordHash:   hex.EncodeToString(password_hash),
			PriviledgeType: privledge_type,
		})

	if err != nil {
		return err
	}
	return nil
}

type HashData struct {
	salt    []byte
	time    uint32
	memory  uint32
	threads uint8
}

func DefaultHashData() (hd HashData) {
	return HashData{}
}

func (hd *HashData) Hash(password string) string {
	return hex.EncodeToString(argon2.IDKey(
		[]byte(password),
		hd.salt,
		hd.time,
		hd.memory,
		hd.threads,
		uint32(len([]byte(password))),
	))
}

// StackOverflow Code : https://stackoverflow.com/questions/45267125/how-to-generate-unique-random-alphanumeric-tokens-in-golang
func GenerateSecureToken(length uint) string {
	bin, err := GenerateSecureTokenBin(length)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bin)
}

func GenerateSecureTokenBin(length uint) ([]byte, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return nil, errors.New("Token didn't generate???")
	}
	return b, nil
}

func Hash(pass []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}
