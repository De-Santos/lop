package zapport

import (
	"context"
	"fmt"
	"time"

	"github.com/De-Santos/lop/lopcore"
	"go.uber.org/zap"
)

type ZapCore struct {
	logger *zap.Logger
}

func NewCore(logger *zap.Logger) lopcore.Core {
	return &ZapCore{logger: logger}
}

func (z *ZapCore) Log(level lopcore.LogLevel, msg string, fields ...lopcore.Field) {
	zapFields := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		switch f.Type {
		case lopcore.FieldString:
			zapFields = append(zapFields, zap.String(f.Key, f.Value.(string)))
		case lopcore.FieldInt:
			zapFields = append(zapFields, zap.Int(f.Key, f.Value.(int)))
		case lopcore.FieldInt64:
			zapFields = append(zapFields, zap.Int64(f.Key, f.Value.(int64)))
		case lopcore.FieldBool:
			zapFields = append(zapFields, zap.Bool(f.Key, f.Value.(bool)))
		case lopcore.FieldFloat64:
			zapFields = append(zapFields, zap.Float64(f.Key, f.Value.(float64)))
		case lopcore.FieldComplex128:
			zapFields = append(zapFields, zap.String(f.Key, formatComplex(f.Value.(complex128))))
		case lopcore.FieldTime:
			zapFields = append(zapFields, zap.Time(f.Key, f.Value.(time.Time)))
		case lopcore.FieldArray, lopcore.FieldMap, lopcore.FieldAny:
			fallthrough
		default:
			zapFields = append(zapFields, zap.Any(f.Key, f.Value))
		}
	}

	switch level {
	case lopcore.LevelDebug:
		z.logger.Debug(msg, zapFields...)
	case lopcore.LevelInfo:
		z.logger.Info(msg, zapFields...)
	case lopcore.LevelWarn:
		z.logger.Warn(msg, zapFields...)
	case lopcore.LevelError:
		z.logger.Error(msg, zapFields...)
	case lopcore.LevelFatal:
		z.logger.Fatal(msg, zapFields...)
	case lopcore.LevelTrace:
		z.logger.Debug(msg, zapFields...) // Zap doesn't have Trace; use Debug
	default:
		z.logger.Info(msg, zapFields...)
	}
}

func (z *ZapCore) With(fields ...lopcore.Field) lopcore.Core {
	zapFields := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		zapFields = append(zapFields, zap.Any(f.Key, f.Value))
	}
	return &ZapCore{logger: z.logger.With(zapFields...)}
}

func (z *ZapCore) WithContext(ctx context.Context) lopcore.Core {
	// Zap does not natively use context.
	return z
}

func formatComplex(c complex128) string {
	return fmt.Sprintf("(%f + %fi)", real(c), imag(c))
}
