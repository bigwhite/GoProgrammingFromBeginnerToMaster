package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	// 密钥(key) 32字节 => AES-256
	key := []byte("12345678123456781234567812345678")

	// 带有IV的密文数据(头部16字节为IV)
	ciphertextWithIV, err := hex.DecodeString("6162636465666768696a6b6c6d6e6f70bc93b5cb1a081b47357f73d40966e3ce53c29db21a13bec2f9be4f76d8f09f2b")
	if err != nil {
		panic(err)
	}

	// 从密文数据中提取IV数据
	iv := ciphertextWithIV[:aes.BlockSize]

	// 待解密的密文数据
	ciphertext := ciphertextWithIV[aes.BlockSize:]

	// 创建AES分组密码算法实例
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// 解密后存放明文的字节切片
	plaintext := make([]byte, len(ciphertext))

	// 创建分组模式的实例，这里使用CBC模式
	cbcModeDecrypter := cipher.NewCBCDecrypter(aesCipher, iv)

	// 对密文进行解密
	cbcModeDecrypter.CryptBlocks(plaintext, ciphertext)

	fmt.Printf("密文(包含IV)：%x\n", ciphertextWithIV)
	fmt.Printf("明文：%s\n", plaintext)
}
