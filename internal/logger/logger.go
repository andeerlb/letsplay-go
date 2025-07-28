package logger

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func NewLogger(level string) (*zap.Logger, error) {
    var zapLevel zapcore.Level
    if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
        zapLevel = zapcore.InfoLevel
    }

    cfg := zap.Config{
        Encoding:         "json",
        Level:            zap.NewAtomicLevelAt(zapLevel),
        OutputPaths:      []string{"stdout"},
        ErrorOutputPaths: []string{"stderr"},
        EncoderConfig: zapcore.EncoderConfig{
            TimeKey:        "ts",
            LevelKey:       "level",
            NameKey:        "logger",
            CallerKey:      "caller",
            MessageKey:     "msg",
            StacktraceKey:  "stacktrace",
            LineEnding:     zapcore.DefaultLineEnding,
            EncodeLevel:    zapcore.LowercaseLevelEncoder,
            EncodeTime:     zapcore.ISO8601TimeEncoder,
            EncodeDuration: zapcore.StringDurationEncoder,
            EncodeCaller:   zapcore.ShortCallerEncoder,
        },
    }

    return cfg.Build()
}