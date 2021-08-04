package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

// go mod init updev.co.th/hash-salt

// Define salt size
const saltSize = 16

// Generate 16 bytes randomly and securely using the
// Cryptographically secure pseudorandom number generator (CSPRNG)
// in the crypto.rand package
func generateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	/*
		salt[:]  : Slice defaults
		ref : https://tour.golang.org/moretypes/10

	*/

	fmt.Printf(" salt[:] %v \n", salt[:])
	// output => [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

func hashPassword(password string, salt []byte) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a base64 encoded string
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash
}

func doPasswordsMatch(hashedPassword, currPassword string,
	salt []byte) bool {
	var currPasswordHash = hashPassword(currPassword, salt)

	return hashedPassword == currPasswordHash
}

func main() {
	fmt.Println("TEST updev.co.th/hash-salt")
	var salt = generateRandomSalt(saltSize)
	fmt.Println("Salt:", salt)
	// output => [50 110 98 156 26 36 176 216 59 141 227 106 138 126 163 65]

	// Hash password using the salt
	var hashedPassword = hashPassword("hello", salt)

	fmt.Println("Password Hash:", hashedPassword)
	fmt.Println("Salt:", salt)

	// Check if passed password matches the original password by hashing it
	// with the original password's salt and check if the hashes match
	fmt.Println("Password Match:",
		doPasswordsMatch(hashedPassword, "hello", salt))
}
