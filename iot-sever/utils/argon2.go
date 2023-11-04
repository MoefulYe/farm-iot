package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

const memory = 64 * 1024
const iterations = 3
const parallelism = 2
const saltLength = 16
const keyLength = 32

func Argon2Generate(passwd string) (string, error) {
	salt, err := generateRandomBytes()
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey(
		[]byte(passwd), salt, iterations, memory, parallelism, keyLength,
	)
	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64Hash := base64.RawStdEncoding.EncodeToString(hash)
	return fmt.Sprintf("%s$%s", base64Salt, base64Hash), nil
}

func Argon2Verify(passwd string, passwdHash string) (bool, error) {
	hash, salt, err := destructHashedPasswd(passwdHash)
	if err != nil {
		return false, err
	}
	toCompare := argon2.IDKey(
		[]byte(passwd), salt, iterations, memory, parallelism, keyLength,
	)
	if subtle.ConstantTimeCompare(toCompare, hash) == 1 {
		return true, nil
	}
	return false, nil
}

func generateRandomBytes() ([]byte, error) {
	b := make([]byte, saltLength)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// [hash, salt, error]
func destructHashedPasswd(hashedPasswd string) ([]byte, []byte, error) {
	pair := strings.Split(hashedPasswd, "$")
	if len(pair) != 2 {
		return nil, nil, errors.New("hashedPasswd error")
	}
	hash, err := base64.RawStdEncoding.DecodeString(pair[1])
	if err != nil {
		return nil, nil, err
	}
	salt, err := base64.RawStdEncoding.DecodeString(pair[0])
	if err != nil {
		return nil, nil, err
	}
	return hash, salt, nil
}
