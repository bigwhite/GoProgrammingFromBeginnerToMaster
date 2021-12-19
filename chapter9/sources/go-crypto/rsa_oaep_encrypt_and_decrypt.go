package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.PublicKey

	// 待加密的明文
	plaintext := []byte("I love go programming language!!")
	fmt.Printf("待加密明文: %s\n", plaintext)

	// 使用公钥加密
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader,
		&publicKey, plaintext, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("密文: %x\n", ciphertext)

	// 使用私钥解密
	textDecrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader,
		privateKey, ciphertext, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密出的明文: %s\n", textDecrypted)
}
