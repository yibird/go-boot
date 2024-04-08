package core

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

// Encryptor 是一个加密算法的接口
type Encryptor interface {
	Encrypt(string) (string, error)
}

// BCryptEncryptor 是 BCrypt 加密算法的实现
type BCryptEncryptor struct{}

// Encrypt 实现了 Encryptor 接口
func (e *BCryptEncryptor) Encrypt(input string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// NoopEncryptor 明文加密
type NoopEncryptor struct{}

func (e *NoopEncryptor) Encrypt(input string) (string, error) {
	return "{noop}" + input, nil
}

type ScryptEncryptor struct{}

func (e *ScryptEncryptor) Encrypt(input string) (string, error) {
	salt := []byte{0x12, 0x34, 0x56, 0x78, 0x90}
	key, err := scrypt.Key([]byte(input), salt, 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}
	// 这里你可能需要将生成的哈希值转换为十六进制字符串，以便后续存储或比较
	// 这里我们使用 %x 格式化谓词将字节数组转换为十六进制字符串
	return fmt.Sprintf("%x", key), nil
}

// EncryptorFactory 加密工厂
func EncryptorFactory(algorithm string) Encryptor {
	switch algorithm {
	case "bcrypt":
		return &BCryptEncryptor{}
	default:
		panic(fmt.Sprintf("Unknown encryption algorithm: %s", algorithm))
	}
}
