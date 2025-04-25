# Lop - Log Portable

**Lop** is a simple, fast, and portable Go logging library designed with clean architecture principles.
It provides a consistent logging interface (`Logger` and `Core`) and supports pluggable backends like Zap and Zerolog.

## ✨ Features
- Lightweight
- Supports all basic Go types: `string`, `int`, `int64`, `bool`, `float64`, `complex128`, `time`, arrays, and maps
- Structured logging with typed fields
- Pluggable backends (ports): Zap, Zerolog
- Easy to extend with other logging systems (e.g., slog, logrus)

## 📦 Installation

```bash
go get github.com/De-Santos/lop
```

## 🚀 Usage

### Basic Example

```go
import (
    "github.com/De-Santos/lop"
    "github.com/De-Santos/lop/lopcore"
    "github.com/De-Santos/lop/ports/zerologport"
    "github.com/rs/zerolog"
    "os"
)

func main() {
    core := zerologport.NewCore(zerolog.New(os.Stdout).With().Timestamp().Logger())
    logger := lop.New(core)

    logger.Info("User created",
        lop.String("username", "alice"),
        lop.Int("age", 30),
        lop.Bool("verified", true),
    )
}
```

### Supported Fields
- `String(key, value string)`
- `Int(key string, value int)`
- `Int64(key string, value int64)`
- `Bool(key string, value bool)`
- `Float64(key string, value float64)`
- `Complex128(key string, value complex128)`
- `TimeField(key string, value time.Time)`
- `Array(key string, value any)`
- `Map(key string, value any)`
- `ErrorField(err error)`

### Logging Levels
- `Debug`
- `Info`
- `Warn`
- `Error`
- `Fatal`
- `Trace`

## 🔌 Available Ports

- [Zerolog Port](ports/zerologport)
- [Zap Port](ports/zapport)

## 🛠 Future Plans
- Add support for all go types
- Middleware integration (Fiber, Echo, Gin)
- Automatic structured tracing fields (trace-id, span-id)

## 🤝 Contribution

Contributions, issues and feature requests are welcome!

---

by De-Santos :3

