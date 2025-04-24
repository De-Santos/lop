package zerologport

import (
	"context"
	"fmt"
	"time"

	"github.com/De-Santos/lop/lopcore"
	"github.com/rs/zerolog"
)

type ZeroCore struct {
	logger zerolog.Logger
}

func NewCore(logger zerolog.Logger) lopcore.Core {
	return &ZeroCore{logger: logger}
}

func (z *ZeroCore) Log(level lopcore.LogLevel, msg string, fields ...lopcore.Field) {
	event := z.entry(level)
	if event == nil {
		return
	}

	for _, f := range fields {
		switch f.Type {
		case lopcore.FieldString:
			event.Str(f.Key, f.Value.(string))
		case lopcore.FieldInt:
			event.Int(f.Key, f.Value.(int))
		case lopcore.FieldInt64:
			event.Int64(f.Key, f.Value.(int64))
		case lopcore.FieldBool:
			event.Bool(f.Key, f.Value.(bool))
		case lopcore.FieldFloat64:
			event.Float64(f.Key, f.Value.(float64))
		case lopcore.FieldComplex128:
			event.Str(f.Key, formatComplex(f.Value.(complex128)))
		case lopcore.FieldTime:
			event.Time(f.Key, f.Value.(time.Time))
		case lopcore.FieldError:
			event.Str(f.Key, f.Value.(string))
		case lopcore.FieldArray, lopcore.FieldMap, lopcore.FieldAny:
			fallthrough
		default:
			event.Interface(f.Key, f.Value)
		}
	}

	event.Msg(msg)
}

func (z *ZeroCore) With(fields ...lopcore.Field) lopcore.Core {
	l := z.logger.With()
	for _, f := range fields {
		l = l.Interface(f.Key, f.Value)
	}
	return &ZeroCore{logger: l.Logger()}
}

func (z *ZeroCore) WithContext(ctx context.Context) lopcore.Core {
	// zerolog allows attaching context with log.Ctx(ctx),
	// but here we keep it simple and return the existing logger
	return z
}

func (z *ZeroCore) entry(level lopcore.LogLevel) *zerolog.Event {
	switch level {
	case lopcore.LevelDebug:
		return z.logger.Debug()
	case lopcore.LevelInfo:
		return z.logger.Info()
	case lopcore.LevelWarn:
		return z.logger.Warn()
	case lopcore.LevelError:
		return z.logger.Error()
	case lopcore.LevelFatal:
		return z.logger.Fatal()
	case lopcore.LevelTrace:
		return z.logger.Trace()
	default:
		return z.logger.Info()
	}
}

func formatComplex(c complex128) string {
	return fmt.Sprintf("(%f + %fi)", real(c), imag(c))
}
