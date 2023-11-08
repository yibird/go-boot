package config

// 限流结构体
type RateLimit struct {
	// 每秒允许通过的请求速率
	Limit float64 `mapstructure:"limit" json:"limit" yaml:"limit"`
	// 令牌桶的容量
	Burst int `mapstructure:"burst" json:"burst" yaml:"burst"`
}

// 加密结构体
type ResponseEncrypt struct {
	// 是否启用api响应加密
	Enabled bool
	// 加密方式
	Mode string
}

type System struct {
	Env             string          `mapstructure:"env" json:"env" yaml:"env"`
	Addr            string          `mapstructure:"addr" json:"addr" yaml:"addr"`
	RouterPrefix    string          `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	DbType          string          `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	RateLimit       RateLimit       `mapstructure:"rate-limit" json:"rate-limit" yaml:"rate-limit"`
	ResponseEncrypt ResponseEncrypt `mapstructure:"response-encrypt" json:"response-encrypt" yaml:"response-encrypt"`
}
