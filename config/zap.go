package config

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type Zap struct {
	// 日志级别
	Level string `mapstructure:"level" json:"level" yaml:"level"`
	/**
	 * 日志格式化。可选值 json、console。Zap支持zapcore.JSONEncoder和zapcore.ConsoleEncoder,
	 * zapcore.JSONEncoder将日志以JSON格式输出,由于JSON格式易于解析和处理,通常适用于生产环境。
	 * zapcore.ConsoleEncoder用于将日志以文本输出,可读好,通常适用于开发环境
	 */
	Format string `mapstructure:"format" json:"format" yaml:"format"`
	// 日志前缀
	Prefix string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	// 日志输出目录
	Director string `mapstructure:"director" json:"director"  yaml:"director"`
	// 日志文件前缀
	FilePrefix string `mapstructure:"file-prefix" json:"file-prefix"  yaml:"file-prefix"`
	// 是否显示文件的行号
	ShowLine bool `mapstructure:"show-line" json:"show-line" yaml:"show-line"`
	// 日志编码级别
	EncodeLevel string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`
	// 用于设置日志中的堆栈跟踪字段的键名
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"`
	// 是否将日志输出到控制台
	LogInConsole bool `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"`
	// 日志文件最大大小(MB),用于日志分割
	MaxSize int `mapstructure:"max-size" json:"max-size" yaml:"max-size"`
	// 保留的最大天数,用于日志分割
	MaxBackups int `mapstructure:"max-backups" json:"max-backups" yaml:"max-backups"`
	// 日志留存时间
	MaxAge int `mapstructure:"max-age" json:"max-age" yaml:"max-age"`
	// 日志留存时间
	Compress bool `mapstructure:"compress" json:"compress" yaml:"compress"`
}

// ZapLevel 根据Zap的Level返回对应的zapcore.Level
func (z *Zap) ZapLevel() zapcore.Level {
	levelMap := map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
	z.Level = strings.ToLower(z.Level)
	if level, exists := levelMap[z.Level]; exists {
		return level
	}
	return zapcore.DebugLevel
}

// ZapEncodeLevel 根据Zap的EncodeLevel返回对应的zapcore.LevelEncoder
func (z *Zap) ZapEncodeLevel() zapcore.LevelEncoder {
	encodeLevelMap := map[string]zapcore.LevelEncoder{
		// 小写编码器
		"LowercaseLevelEncoder": zapcore.LowercaseLevelEncoder,
		// 小写编码器带颜色
		"LowercaseColorLevelEncoder": zapcore.LowercaseColorLevelEncoder,
		// 大写编码器
		"CapitalLevelEncoder": zapcore.CapitalLevelEncoder,
		// 大写编码器带颜色
		"CapitalColorLevelEncoder": zapcore.LowercaseLevelEncoder,
	}
	if encode, exists := encodeLevelMap[z.EncodeLevel]; exists {
		return encode
	}
	return zapcore.LowercaseLevelEncoder
}
