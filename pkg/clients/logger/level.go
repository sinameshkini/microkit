package logger

import (
	"context"
	"go.uber.org/zap/zapcore"
)

type Level string

func (l Level) toZapCoreLevel() zapcore.Level {
	switch l {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case PanicLevel:
		return zapcore.DPanicLevel
	case FatalLevel:
		return zapcore.FatalLevel
	}

	return zapcore.InvalidLevel
}

const (
	DebugLevel Level = "debug"
	InfoLevel  Level = "info"
	WarnLevel  Level = "warn"
	ErrorLevel Level = "error"
	PanicLevel Level = "panic"
	FatalLevel Level = "fatal"
)

func (l *Logger) newLogs(ctx context.Context, lvl Level) *Logs {
	if ctx == nil {
		ctx = context.Background()
	}

	logs := &Logs{
		ctx:    ctx,
		level:  lvl,
		fields: make(map[Field]interface{}),
	}
	logs.prepare(l)
	return logs
}

func (l *Logger) Debug(ctx context.Context) *Logs {
	return l.newLogs(ctx, DebugLevel)
}

func (l *Logger) Info(ctx context.Context) *Logs {
	return l.newLogs(ctx, InfoLevel)
}

func (l *Logger) Warning(ctx context.Context) *Logs {
	return l.newLogs(ctx, WarnLevel)
}

func (l *Logger) Error(ctx context.Context) *Logs {
	return l.newLogs(ctx, ErrorLevel)
}

func (l *Logger) Panic(ctx context.Context) *Logs {
	return l.newLogs(ctx, PanicLevel)
}

func (l *Logger) Fatal(ctx context.Context) *Logs {
	return l.newLogs(ctx, FatalLevel)
}

func (l *Logs) prepare(logger *Logger) {
	logger.addFieldsTo(l)
	l.logContext(l.ctx)
	l.logger = logger.logger
}

func (l *Logs) SetContext() context.Context {
	for k, v := range l.fields {
		l.ctx = context.WithValue(l.ctx, k, v)
	}

	return l.ctx
}
