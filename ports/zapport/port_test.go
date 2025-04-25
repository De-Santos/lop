package zapport_test

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/De-Santos/lop"
	"github.com/De-Santos/lop/lopcore"
	"github.com/De-Santos/lop/ports/zapport"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapLogger(buf *bytes.Buffer) *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "message"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(buf),
		zapcore.DebugLevel,
	)
	return zap.New(core)
}

func parseZapLog(t *testing.T, buf *bytes.Buffer) map[string]any {
	decoder := json.NewDecoder(buf)
	decoder.UseNumber()
	var out map[string]any
	err := decoder.Decode(&out)
	assert.NoError(t, err)
	return out
}

func TestZapPort_LogBasicTypes(t *testing.T) {
	buf := &bytes.Buffer{}
	zlogger := newZapLogger(buf)
	core := zapport.NewCore(zlogger)
	logger := lop.New(core)

	now := time.Now().UTC()
	logger.Info("testing",
		lop.String("str", "value"),
		lop.Int("int", 42),
		lop.Int64("int64", 4200000000),
		lop.Bool("bool", true),
		lop.Float64("float", 3.14),
		lop.Complex128("complex", 2+3i),
		lop.TimeField("time", now),
		lop.Map("map", map[string]any{"k": "v"}),
		lop.Array("slice", []string{"a", "b"}),
		lop.ErrorField(assert.AnError),
	)

	data := parseZapLog(t, buf)
	assert.Equal(t, "testing", data["message"])
	assert.Equal(t, "value", data["str"])
	assert.Equal(t, json.Number("42"), data["int"])
	assert.Equal(t, json.Number("4200000000"), data["int64"])
	assert.Equal(t, true, data["bool"])
	fnum, err := data["float"].(json.Number).Float64()
	assert.NoError(t, err)
	assert.Equal(t, 3.14, fnum)
	assert.Equal(t, "(2.000000 + 3.000000i)", data["complex"])
	assert.Equal(t, assert.AnError.Error(), data["error"])
	assert.IsType(t, map[string]any{}, data["map"])
	assert.IsType(t, []any{}, data["slice"])
}

func TestZapPort_WithContext_NoCrash(t *testing.T) {
	buf := &bytes.Buffer{}
	core := zapport.NewCore(newZapLogger(buf))
	ctx := context.Background()
	core = core.WithContext(ctx)
	core.Log(lopcore.LevelInfo, "hello context")
	output := buf.String()
	assert.True(t, strings.Contains(output, "hello context"))
}
