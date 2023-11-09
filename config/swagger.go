package config

type Swagger struct {
	BasePath         string `mapstructure:"base-path" json:"base-path" yaml:"base-path"`
	Title            string `mapstructure:"title" json:"title" yaml:"title"`
	Description      string `mapstructure:"description" json:"description" yaml:"description"`
	Host             string `mapstructure:"host" json:"host" yaml:"host"`
	InfoInstanceName string `mapstructure:"info-instance-name" json:"info-instance-name" yaml:"info-instance-name"`
	Version          string `mapstructure:"version" json:"version" yaml:"version"`
}
