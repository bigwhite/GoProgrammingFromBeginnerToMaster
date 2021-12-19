package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// 生成公钥密码的密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.PublicKey

	// 待签名消息
	msg := []byte("I love go programming language!!")

	// 计算摘要
	digest := sha256.Sum256(msg)

	// 用私钥签名
	sign, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, digest[:], nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("签名：%s\n", fmt.Sprintf("%x", sign))

	// 用公钥验证签名
	err = rsa.VerifyPSS(&publicKey, crypto.SHA256, digest[:], sign, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("签名验证成功!\n")
}
