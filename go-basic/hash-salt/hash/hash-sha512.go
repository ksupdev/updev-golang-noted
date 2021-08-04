package hash

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"strings"
)

type HasherSHA512 struct {
	SaltString []byte
	SaltsSize  int
}

func NewHasherSHA512(saltStr []byte, saltSize int) *HasherSHA512 {
	return &HasherSHA512{
		SaltString: saltStr,
		SaltsSize:  saltSize}
}

func (h *HasherSHA512) EncodeWithSalt(value string) (string, error) {
	if h.SaltsSize == 0 || (h.SaltString == nil && len(h.SaltString) <= 0) || len(strings.TrimSpace(value)) <= 0 {

		return "", errors.New("please check requir argument")
	}

	// Convert password string to byte slice
	var valueByte = []byte(value)

	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Append salt to password
	valueByte = append(valueByte, h.SaltString...)

	// Write value bytes to the hasher
	sha512Hasher.Write(valueByte)

	// Get the SHA-512 hashed password
	var hashedValueBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a base64 encoded string
	var base64EncodedHash = base64.URLEncoding.EncodeToString(hashedValueBytes)

	return base64EncodedHash, nil
}

func (h *HasherSHA512) Encode(value string) (string, error) {
	if len(strings.TrimSpace(value)) <= 0 {

		return "", errors.New("please check requir argument")
	}
	// Convert password string to byte slice
	var valueByte = []byte(value)

	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Write value bytes to the hasher
	sha512Hasher.Write(valueByte)

	// Get the SHA-512 hashed password
	var hashedValueBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a base64 encoded string
	var base64EncodedHash = base64.URLEncoding.EncodeToString(hashedValueBytes)

	return base64EncodedHash, nil
}
