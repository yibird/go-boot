package config

type Mysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// GetLoggerMode 获取日志模式
func (m *Mysql) GetLoggerMode() string {
	return m.LogMode
}

// Dsn 获取Mysql连接地址
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
