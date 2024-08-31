package zlog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Zlog = NewSLog()

func NewSLog() *zap.Logger {
	// 创建一个 lumberjack.Logger 实例
	logger := &lumberjack.Logger{
		Filename:   "app.log", // 日志文件名
		MaxSize:    10,        // 每个日志文件的最大大小（MB）
		MaxBackups: 3,         // 保留旧文件的最大个数
		MaxAge:     28,        // 保留旧文件的最大天数
		Compress:   true,      // 是否压缩/归档旧文件
	}

	// 创建 zap logger
	zapLogger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), // 日志格式
		zapcore.AddSync(logger),                                  // 使用 lumberjack 作为输出
		zapcore.InfoLevel,                                        // 日志级别
	))

	defer zapLogger.Sync() // 确保日志被写入
	return zapLogger
}
