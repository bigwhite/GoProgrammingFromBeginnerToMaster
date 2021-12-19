package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.PublicKey

	fmt.Printf("Private Key's size = %d bits\n", privateKey.Size()*8)
	fmt.Printf("Public Key's size = %d bits\n", publicKey.Size()*8)
}
