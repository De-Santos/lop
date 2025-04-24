package lopcore

import "context"

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelTrace
)

func (l LogLevel) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelTrace:
		return "trace"
	default:
		return "unknown"
	}
}

type FieldType int

const (
	FieldAny FieldType = iota
	FieldString
	FieldInt
	FieldInt64
	FieldBool
	FieldFloat64
	FieldComplex128
	FieldTime
	FieldArray
	FieldMap
	FieldError
)

type Field struct {
	Key   string
	Value any
	Type  FieldType
}

type Logger interface {
	Log(level LogLevel, msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Trace(msg string, fields ...Field)
	With(fields ...Field) Logger
	WithContext(ctx context.Context) Logger
}

type Core interface {
	Log(level LogLevel, msg string, fields ...Field)
	With(fields ...Field) Core
	WithContext(ctx context.Context) Core
}
