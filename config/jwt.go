package config

type Jwt struct {
	// jwt签名
	SigningKey string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`
	// 过期时间
	ExpiresTime string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time`
	// 缓冲时间
	BufferTime string `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`
	// 签发者
	Issuer string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
}
