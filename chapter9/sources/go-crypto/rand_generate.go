package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	c := 32
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", b)
}
