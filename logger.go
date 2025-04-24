package lop

import (
	"context"
	"github.com/De-Santos/lop/lopcore"
)

type logger struct {
	core lopcore.Core
}

func New(core lopcore.Core) lopcore.Logger {
	return &logger{core: core}
}

func (l *logger) Log(level lopcore.LogLevel, msg string, fields ...lopcore.Field) {
	l.core.Log(level, msg, fields...)
}

func (l *logger) Info(msg string, fields ...lopcore.Field) {
	l.Log(lopcore.LevelInfo, msg, fields...)
}

func (l *logger) Warn(msg string, fields ...lopcore.Field) {
	l.Log(lopcore.LevelWarn, msg, fields...)
}

func (l *logger) Error(msg string, fields ...lopcore.Field) {
	l.Log(lopcore.LevelError, msg, fields...)
}

func (l *logger) Fatal(msg string, fields ...lopcore.Field) {
	l.Log(lopcore.LevelFatal, msg, fields...)
}

func (l *logger) Debug(msg string, fields ...lopcore.Field) {
	l.Log(lopcore.LevelDebug, msg, fields...)
}

func (l *logger) Trace(msg string, fields ...lopcore.Field) {
	l.Log(lopcore.LevelTrace, msg, fields...)
}

func (l *logger) With(fields ...lopcore.Field) lopcore.Logger {
	return &logger{core: l.core.With(fields...)}
}

func (l *logger) WithContext(ctx context.Context) lopcore.Logger {
	return &logger{core: l.core.WithContext(ctx)}
}
