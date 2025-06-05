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
	// 配置 lumberjack 日志滚动
	rollingLogger := &lumberjack.Logger{
		Filename:   "test.log",
		MaxSize:    10,   // 每个日志文件最大10MB
		MaxBackups: 3,    // 保留最近的3个备份
		MaxAge:     28,   // 保留28天
		Compress:   true, // 压缩备份文件
	}

	// 创建一个 JSON 编码器
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
	// 创建 WriteSyncer
	writer := zapcore.AddSync(rollingLogger)

	// 创建 Core
	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)

	// 使用 zap 的选项创建 Logger
	logger := zap.New(core,
		zap.AddCaller(),                        // 启用调用者信息
		zap.Development(),                      // 开发模式，启用更详细的日志记录
		zap.Fields(zap.String("app", "myApp")), // 添加默认字段
	)

	// 使用 Logger 记录日志
	logger.Info("example : This is an info message")
	logger.Debug("example : This is a debug message")

	defer logger.Sync() // 确保日志被写入
	return logger
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

	logSugar := loggerWithFields.Sugar()
	logSugar.Infof("This is an info message with additional fields %d %d %d", 1, 2, 3)
	logSugar.Warn("This is an info message with additional fields ", 1, 2, 3, 4, 5)
}
