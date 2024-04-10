package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/golang-jwt/jwt/v5"
	"go-boot/config"
	"time"
)

var signingMethodMap = map[string]jwt.SigningMethod{
	"ES256": jwt.SigningMethodES256,
}

// GetSigningMethod 解析签名方法
func GetSigningMethod(signingMethod string) jwt.SigningMethod {
	if method, ok := signingMethodMap[signingMethod]; ok {
		return method
	}
	return signingMethodMap["ES256"]
}

// CreateToken 创建token
func CreateToken(_jwt config.Jwt) string {
	claims := jwt.MapClaims{
		"exp":    time.Now().Add(time.Second * time.Duration(_jwt.Expires)).Unix(),
		"issuer": _jwt.Issuer,
	}
	token := jwt.NewWithClaims(GetSigningMethod(_jwt.SigningMethod), claims)
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}
	return tokenString
}

// CreateTokenWithExpires 创建token
func CreateTokenWithExpires(_jwt config.Jwt, expires int) string {
	_jwt.Expires = expires
	return CreateToken(_jwt)
}
