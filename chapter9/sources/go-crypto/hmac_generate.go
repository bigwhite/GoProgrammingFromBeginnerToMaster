package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	// 密钥(key) 32字节
	key := []byte("12345678123456781234567812345678")

	// 要传递的消息
	message := []byte("I love go programming language!!")

	// 创建hmac实例(使用SHA-256单向散列函数)
	mac := hmac.New(sha256.New, key)
	mac.Write(message)

	// 计算mac值
	m := mac.Sum(nil)
	ms := fmt.Sprintf("%x", m) // mac to string
	fmt.Printf("mac值 = %s\n", ms)
}
