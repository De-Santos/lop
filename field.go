package lop

import (
	"fmt"
	"github.com/De-Santos/lop/lopcore"
	"time"
)

func String(key, value string) lopcore.Field {
	return lopcore.Field{Key: key, Value: value, Type: lopcore.FieldString}
}

func Int(key string, value int) lopcore.Field {
	return lopcore.Field{Key: key, Value: value, Type: lopcore.FieldInt}
}

func Int64(key string, value int64) lopcore.Field {
	return lopcore.Field{Key: key, Value: value, Type: lopcore.FieldInt64}
}

func Bool(key string, value bool) lopcore.Field {
	return lopcore.Field{Key: key, Value: value, Type: lopcore.FieldBool}
}

func Float64(key string, value float64) lopcore.Field {
	return lopcore.Field{Key: key, Value: value, Type: lopcore.FieldFloat64}
}

func Complex128(key string, value complex128) lopcore.Field {
	return lopcore.Field{Key: key, Value: value, Type: lopcore.FieldComplex128}
}

func TimeField(key string, value time.Time) lopcore.Field {
	return lopcore.Field{Key: key, Value: value, Type: lopcore.FieldTime}
}

func Any(key string, value any) lopcore.Field {
	return lopcore.Field{Key: key, Value: value, Type: lopcore.FieldAny}
}

func Array(key string, value any) lopcore.Field {
	return lopcore.Field{Key: key, Value: value, Type: lopcore.FieldArray}
}

func Map(key string, value any) lopcore.Field {
	return lopcore.Field{Key: key, Value: value, Type: lopcore.FieldMap}
}

func ErrorField(err error) lopcore.Field {
	return lopcore.Field{Key: "error", Value: fmt.Sprintf("%v", err), Type: lopcore.FieldError}
}
