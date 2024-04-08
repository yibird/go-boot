package config

type DsnProvider interface {
	Dsn() string
}

// GeneralDB 也被 Pgsql 和 Mysql 原样使用
type GeneralDB struct {
	// 服务器地址:端口
	Path string `mapstructure:"path" json:"path" yaml:"path"`
	// 端口
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	// 高级配置
	Config string `mapstructure:"config" json:"config" yaml:"config"`
	// 数据库名
	Dbname string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	// 数据库用户名
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	// 数据库密码
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	// 全局表前缀，单独定义TableName则不生效
	Prefix string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	// 是否开启全局禁用复数，true表示开启
	Singular bool `mapstructure:"singular" json:"singular" yaml:"singular"`
	// 数据库引擎，默认InnoDB
	Engine string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`
	// 空闲中的最大连接数
	MaxIdleConns int `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	// 打开到数据库的最大连接数
	MaxOpenConns int `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	// 连接最大生存时间(单位秒)
	ConnMaxLifetime int `mapstructure:"conn-max-lifetime" json:"conn-max-lifetime" yaml:"conn-max-lifetime"`
	// 是否开启Gorm全局日志
	LogMode string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
	// 是否通过zap写入日志文件
	LogZap bool `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`
}

type SpecializedDB struct {
	Disable   bool   `mapstructure:"disable" json:"disable" yaml:"disable"`
	Type      string `mapstructure:"_type" json:"_type" yaml:"_type"`
	AliasName string `mapstructure:"alias-name" json:"alias-name" yaml:"alias-name"`
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}
