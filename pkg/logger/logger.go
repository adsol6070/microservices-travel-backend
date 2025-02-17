package logger

import (
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// InitLogger initializes the global logger instance
func InitLogger(logLevel string, logFilePath string) {
	level := getLogLevel(logLevel)

	// Console + File logging setup
	core := zapcore.NewTee(
		zapcore.NewCore(getEncoder(), getLogWriter(logFilePath), level), // File writer
		zapcore.NewCore(getEncoder(), zapcore.AddSync(os.Stdout), level), // Console writer
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer Logger.Sync()
}

// getLogLevel converts string log level to zapcore.Level
func getLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

// getEncoder sets up log format (JSON or Console)
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		MessageKey:     "message",
		CallerKey:      "caller",
		EncodeTime:     customTimeEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	return zapcore.NewJSONEncoder(encoderConfig) // Use JSON logs; change to NewConsoleEncoder for readable logs
}

// getLogWriter sets up file logging with rotation
func getLogWriter(logFilePath string) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    10,   // Max file size in MB
		MaxBackups: 5,    // Max number of old log files
		MaxAge:     30,   // Max days to keep old log files
		Compress:   true, // Enable log compression
	}

	return zapcore.AddSync(lumberjackLogger)
}

// customTimeEncoder formats the log timestamp
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// Debug logs a debug-level message
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

// Info logs an info-level message
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

// Warn logs a warning-level message
func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

// Error logs an error-level message
func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}
