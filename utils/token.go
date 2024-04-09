package utils

import (
	"time"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)
//创建token
func GenerateToken(userID string) (string, error) {
	// 创建一个新的令牌声明
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间为 1 天后
	}

	// 创建一个签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥对令牌进行签名
	secret := "your_secret_key" // 自定义的密钥
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
// 验证 token
func ValidateToken(tokenString string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// 返回密钥用于验证签名
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	// 验证有效期
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}