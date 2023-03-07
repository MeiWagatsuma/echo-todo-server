// Package lib is some funcion
package lib

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

// Sha3Hash if for password hashing
func Sha3Hash(input string) []byte {
	hash := sha3.New256()
	salt := "my salt"
	_, _ = hash.Write([]byte(input + salt))
	sha3 := hash.Sum(nil)

	return sha3
}

func main() {
	input := "hello"
	output := Sha3Hash(input)
	fmt.Println(output)
}
