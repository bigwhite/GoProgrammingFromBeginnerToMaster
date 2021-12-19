package main

import (
	"crypto/sha256"
	"fmt"
)

func sum256(data []byte) string {
	sum := sha256.Sum256(data)
	return fmt.Sprintf("%x", sum)
}

func main() {
	s := "I love go programming language!!"
	fmt.Println(sum256([]byte(s)))
}
