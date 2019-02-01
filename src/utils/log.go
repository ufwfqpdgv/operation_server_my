package utils

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLog(c Log_info) (logger *zap.Logger) {
	hook := lumberjack.Logger{
		Filename:   c.Path_filename, // 日志文件路径
		MaxSize:    c.Max_size,      // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: c.Max_backups,   // 日志文件最多保存多少个备份
		MaxAge:     c.Max_age,       // 文件最多保存多少天
		Compress:   c.Compress,      // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     customkTimeEncoder,             // 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	switch c.Level {
	case "debug":
		atomicLevel.SetLevel(zap.DebugLevel)
	case "info":
		atomicLevel.SetLevel(zap.InfoLevel)
	case "warn":
		atomicLevel.SetLevel(zap.WarnLevel)
	case "error":
		atomicLevel.SetLevel(zap.ErrorLevel)
	case "panic":
		atomicLevel.SetLevel(zap.PanicLevel)
	case "fatal":
		atomicLevel.SetLevel(zap.FatalLevel)
	default:
		panic("input on of 'debug、info、warn、error、panic、fatal'")
	}

	var encoder zapcore.Encoder
	switch c.Encoding{
	case "console":
		encoder= zapcore.NewConsoleEncoder(encoderConfig)                                           // 编码器配置
	case "json":
		encoder= zapcore.NewJSONEncoder(encoderConfig)                                           // 编码器配置
	default:
		panic("input on of 'console、json'")
	}
	core := zapcore.NewCore(
		encoder,                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段，设置后每行都会带上
	// filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	logger = zap.New(core, caller, development, filed)

	logger.Info(

	return
}

func customkTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
