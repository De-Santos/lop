package zerologport_test

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/De-Santos/lop"
	"github.com/De-Santos/lop/lopcore"
	"github.com/De-Santos/lop/ports/zerologport"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func parseLog(t *testing.T, buf *bytes.Buffer) map[string]any {
	decoder := json.NewDecoder(buf)
	decoder.UseNumber()
	var out map[string]any
	err := decoder.Decode(&out)
	assert.NoError(t, err)
	return out
}

func TestZerologPort_LogBasicTypes(t *testing.T) {
	buf := &bytes.Buffer{}
	zlogger := zerolog.New(buf).With().Timestamp().Logger()
	core := zerologport.NewCore(zlogger)
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

	data := parseLog(t, buf)
	assert.Equal(t, "testing", data["message"])
	assert.Equal(t, "value", data["str"])
	assert.Equal(t, json.Number("42"), data["int"])
	assert.Equal(t, json.Number("4200000000"), data["int64"])
	assert.Equal(t, true, data["bool"])
	assert.Equal(t, 3.14, data["float"])
	assert.Equal(t, "(2.000000 + 3.000000i)", data["complex"])
	assert.Equal(t, assert.AnError.Error(), data["error"])
	assert.IsType(t, map[string]any{}, data["map"])
	assert.IsType(t, []any{}, data["slice"])
}

func TestZerologPort_WithContext_NoCrash(t *testing.T) {
	buf := &bytes.Buffer{}
	core := zerologport.NewCore(zerolog.New(buf))
	ctx := context.Background()
	core = core.WithContext(ctx)
	core.Log(lopcore.LevelInfo, "hello context")
	output := buf.String()
	assert.True(t, strings.Contains(output, "hello context"))
}
