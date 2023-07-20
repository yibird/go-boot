package internal

import (
	"fmt"
	"go-boot/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

type _zap struct{}

var Zap = new(_zap)

// CustomTimeEncoder 自定义日志输出时间格式
func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.CONFIG.Zap.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

// GetEncoderConfig 获取zap编码器配置文件,返回 zapcore.EncoderConfig 实例
func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    global.CONFIG.Zap.ZapEncodeLevel(),
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

// GetEncoder 获取 zapcore.Encoder
func (z *_zap) GetEncoder() zapcore.Encoder {
	// 获取配置文件中Zap的格式化方式
	format := global.CONFIG.Zap.Format

	formatMap := map[string]zapcore.Encoder{
		"json":    zapcore.NewJSONEncoder(z.GetEncoderConfig()),
		"console": zapcore.NewConsoleEncoder(z.GetEncoderConfig()),
	}
	if encoder, exists := formatMap[format]; exists {
		return encoder
	}
	return formatMap["console"]
}

// GetLevelPriority 根据日志级别获取 zap.LevelEnablerFunc(级别启用函数),LevelEnablerFunc
// 用于根据日志级别来判断是否记录日志
func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	maps := map[zapcore.Level]func(level zapcore.Level) bool{
		// 调试级别
		zapcore.DebugLevel: func(level zapcore.Level) bool {
			return level == zap.DebugLevel
		},
		// 日志级别
		zapcore.InfoLevel: func(level zapcore.Level) bool {
			return level == zap.InfoLevel
		},
		// 警告级别
		zapcore.WarnLevel: func(level zapcore.Level) bool {
			return level == zap.WarnLevel
		},
		// 错误级别
		zapcore.ErrorLevel: func(level zapcore.Level) bool {
			return level == zap.ErrorLevel
		},
		// dpanic级别
		zapcore.DPanicLevel: func(level zapcore.Level) bool {
			return level == zap.DPanicLevel
		},
		// panic级别
		zapcore.PanicLevel: func(level zapcore.Level) bool {
			return level == zap.PanicLevel
		},
		// 终止级别
		zapcore.FatalLevel: func(level zapcore.Level) bool {
			return level == zap.FatalLevel
		},
	}
	if fn, exists := maps[level]; exists {
		return fn
	}
	return maps[zapcore.DebugLevel]
}

func getWriter(level string) zapcore.WriteSyncer {
	director := global.CONFIG.Zap.Director
	filePrefix := global.CONFIG.Zap.FilePrefix
	filename := fmt.Sprintf("%s/%s-%s.log", director, filePrefix, level)

	// 使用lumberjack进行日志分割
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    global.CONFIG.Zap.MaxSize,
		MaxBackups: global.CONFIG.Zap.MaxBackups,
		MaxAge:     global.CONFIG.Zap.MaxAge,
		Compress:   global.CONFIG.Zap.Compress,
	}
	return zapcore.AddSync(lumberjackLogger)
}

// GetEncoderCore 获取Encoder的 zapcore.Core
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	return zapcore.NewCore(z.GetEncoder(), getWriter(l.String()), level)
}

// GetZapCores 根据配置文件的Level获取 []zapcore.Core
func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := global.CONFIG.Zap.ZapLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}
