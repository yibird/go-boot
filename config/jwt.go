package config

import "time"

type Jwt struct {
	// jwt签名
	SigningKey string `mapstructure:"signingKey" json:"signingKey" yaml:"signing-key"`
	// 过期时间
	ExpiresTime time.Time `mapstructure:"expiresTime" json:"expiresTime" yaml:"expires-time"`
	// 缓冲时间
	BufferTime time.Time `mapstructure:"bufferTime" json:"bufferTime" yaml:"buffer-time"`
	// 签发者
	Issuer string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
}
