package config

type System struct {
	Env          string `mapstructure:"env" json:"env" yaml:"env"`
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
}
