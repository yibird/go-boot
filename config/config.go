package config

// Config 服务端配置结构体
type Config struct {
	JWT         Jwt         `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap         Zap         `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis       Redis       `mapstructure:"redis" json:"redis" yaml:"redis"`
	Cors        CORS        `mapstructure:"cors" json:"cors" yaml:"cors"`
	System      System      `mapstructure:"system" json:"system" yaml:"system"`
	Swagger     Swagger     `mapstructure:"swagger" json:"swagger" yaml:"swagger"`
	Application Application `mapstructure:"application" json:"application" yaml:"application"`

	// ----------------- DB
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
