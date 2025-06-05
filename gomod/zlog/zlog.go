package zlog

import (
	"os"
	"time"

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

	zapcore.NewTee()

	defer zapLogger.Sync() // 确保日志被写入
	return zapLogger
}

// CustomCore 是一个自定义的 zapcore.Core 实现
type CustomCore struct {
	zapcore.LevelEnabler
	encoder zapcore.Encoder
	writer  zapcore.WriteSyncer
}

// Check 实现了 zapcore.Core 接口的 Check 方法
func (c *CustomCore) Check(entry zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(entry.Level) {
		return ce.AddCore(entry, c)
	}
	return ce
}

// Write 实现了 zapcore.Core 接口的 Write 方法
func (c *CustomCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	buf, err := c.encoder.EncodeEntry(entry, fields)
	if err != nil {
		return err
	}
	_, err = c.writer.Write(buf.Bytes())
	return err
}

// Sync 实现了 zapcore.Core 接口的 Sync 方法
func (c *CustomCore) Sync() error {
	return c.writer.Sync()
}

// With 实现了 zapcore.Core 接口的 With 方法
func (c *CustomCore) With(fields []zapcore.Field) zapcore.Core {
	// 创建一个新的 encoder，并添加额外的字段
	newEncoder := c.encoder.Clone()
	for _, field := range fields {
		field.AddTo(newEncoder)
	}
	return &CustomCore{
		LevelEnabler: c.LevelEnabler,
		encoder:      newEncoder,
		writer:       c.writer,
	}
}

func NewCustomZLog() {
	// 创建一个自定义的 encoder
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:          "time",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		MessageKey:       "msg",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeTime:       zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000000"),
		EncodeDuration:   zapcore.StringDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: " ",
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 创建一个自定义的 WriteSyncer
	writer := zapcore.AddSync(os.Stdout)

	// 创建一个自定义的 Core
	core := &CustomCore{
		LevelEnabler: zapcore.DebugLevel,
		encoder:      encoder,
		writer:       writer,
	}

	// 创建一个 logger
	logger := zap.New(core)

	// 使用 logger 记录日志
	logger.Info("This is an info message",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	logger.Warn("This is a warning message")

	logger.Error("This is an error message")

	// 使用 With 方法添加额外的字段
	loggerWithFields := logger.With(zap.String("request_id", "12345"), zap.String("user_id", "abcde"))

	// 使用带有额外字段的 logger 记录日志
	loggerWithFields.Info("This is an info message with additional fields", zap.String("status", "success"))
}
