// Package lib is some funcion
package lib

import (
	"golang.org/x/crypto/sha3"
)

// Sha3Hash if for password hashing
func Sha3Hash(input string) []byte {
	hash := sha3.New256()
	salt := "my salt"
	_, _ = hash.Write([]byte(input + salt))
	hashedBinary := hash.Sum(nil)

	return hashedBinary
}
