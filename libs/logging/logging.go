package logging

import (
	"os"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logLevelDebug = "debug"
)

type syncer interface {
	Sync() error
}

// NewFlusher creates a new syncer from given syncer that log a error message if failed to sync.
func NewFlusher(s syncer) func() {
	return func() {
		// ignore the error as the sync function will always fail in Linux
		// https://github.com/uber-go/zap/issues/370
		_ = s.Sync()
	}
}

// NewSugaredLogger creates a new sugared logger and a flush function. The flush function should be
// called by consumer before quitting application.
// This function should be use most of the time unless
// the application requires extensive performance, in this case use NewLogger.
func NewSugaredLogger() (*zap.SugaredLogger, func(), error) {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = SyslogTimeEncoder
	encoder := zapcore.NewConsoleEncoder(encoderCfg)

	var atom zap.AtomicLevel
	logLevel := viper.GetString("METROGALAXY_API_LOG_LEVEL")

	switch logLevel {
	case logLevelDebug:
		atom = zap.NewAtomicLevelAt(zap.DebugLevel)
	default:
		atom = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	logger := zap.New(zapcore.NewCore(
		encoder,
		zapcore.Lock(os.Stdout),
		atom,
	), zap.AddCaller())

	sugar := logger.Sugar()
	return sugar, NewFlusher(logger), nil
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
