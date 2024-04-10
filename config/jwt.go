package config

import "time"

type Jwt struct {
	// jwt签名
	SigningKey string `mapstructure:"signingKey" json:"signingKey" yaml:"signing-key"`
	// 签名方法
	SigningMethod string `mapstructure:"signingMethod" json:"signingMethod" yaml:"signing-method"`
	// 过期时间
	Expires int `mapstructure:"expires" json:"expires" yaml:"expires"`
	// 缓冲时间
	BufferTime time.Time `mapstructure:"bufferTime" json:"bufferTime" yaml:"buffer-time"`
	// 签发者
	Issuer string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
}
