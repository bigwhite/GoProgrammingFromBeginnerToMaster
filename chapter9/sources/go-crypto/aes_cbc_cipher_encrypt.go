package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	// 密钥(key) 32字节 => AES-256
	key := []byte("12345678123456781234567812345678")

	// 创建AES分组密码算法实例
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// 待加密的明文字符串(长度恰为分组长度的整数倍)
	plaintext := []byte("I love go programming language!!")

	// 存储密文的切片，预留出在密文头部放置IV的空间
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// 这里IV采用固定字符串，这样多次运行结果相同，便于示例说明
	iv := []byte("abcdefghijklmnop")

	// 创建分组模式的实例，这里使用CBC模式
	cbcModeEncrypter := cipher.NewCBCEncrypter(aesCipher, iv)

	// 对明文进行加密
	cbcModeEncrypter.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// 这里将IV放在密文的头部(IV的长度 = block length)
	copy(ciphertext[:aes.BlockSize], []byte("abcdefghijklmnop"))

	fmt.Printf("明文：%s\n", plaintext)
	fmt.Printf("密文(包含IV)：%x\n", ciphertext)
}
