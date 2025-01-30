package token

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"log"
	"os"
	"time"

	"github.com/jackcooperusesvim/coopGo/model/util"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
)

func ValidateToken(token string) (privledge_level string, account_id int64, err error) {

	start := time.Now()
	q, ctx, err := util.DbInfo()
	if err != nil {
		return "", 0, err
	}
	log.Println("open db:", time.Since(start))

	token_hash, err := Hash(token)

	if err != nil {
		return "", 0, err
	}

	hash_rows, err := q.GetSimilarSessionTokens(ctx, token_hash[:3]+"%")
	log.Println("hash_beginning", token_hash[:3]+"%")

	if err != nil {
		return "", 0, err
	}
	for _, hash_row := range hash_rows {

		if bcrypt.CompareHashAndPassword([]byte(hash_row.Token), []byte(token)) == nil {
			log.Println(token[:3] + "%")
			log.Println("quick-match SUCCESSFUL")
			return hash_row.PriviledgeType, hash_row.ID, nil
		}
	}

	log.Println("Can't quick-match")
	start = time.Now()
	hash_rows_all, err := q.GetSessionTokens(ctx)
	log.Println("query:", time.Since(start))

	if err != nil {
		return "", 0, err
	}

	for _, hash_row := range hash_rows_all {
		start = time.Now()
		if bcrypt.CompareHashAndPassword([]byte(hash_row.Token), []byte(token)) == nil {
			log.Println("query:", time.Since(start))
			return hash_row.PriviledgeType, hash_row.ID, nil
		}
	}

	return "", 0, errors.New("session token not found")

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
		return nil, errors.New("Token didn't generate?")
	}
	return b, nil
}

func ResetSalt(length uint) error {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return errors.New("Salt didn't generate???")
	}
	return nil
}
func AppendSalt(length uint) {
}

func GetSalt() []byte {
	return []byte(os.Getenv("salt"))
}
