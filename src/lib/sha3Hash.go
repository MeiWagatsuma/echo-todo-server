// Package lib is some funcion
package main

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

func Sha3Hash(input string) string {
	hash := sha3.New256()
	salt := "my salt"
	_, _ = hash.Write([]byte(input + salt))
	sha3 := hash.Sum(nil)

	return fmt.Sprintf("%x", sha3)
}

func main() {
	input := "hello"
	output := Sha3Hash(input)
	fmt.Println(output)
}
