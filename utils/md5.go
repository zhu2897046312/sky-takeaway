package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// HashPassword 函数用于将密码进行 md5 加密
func Md5(password string) (string, error) {
	// 将密码转换为字节数组
	passwordBytes := []byte(password)

	// 使用 md5 进行加密
	hash := md5.Sum(passwordBytes)

	// 将加密结果转换为十六进制字符串
	hashedPassword := hex.EncodeToString(hash[:])

	return hashedPassword, nil
}