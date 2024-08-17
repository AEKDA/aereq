package logger

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// global глобальный экземпляр логгера.
	global       *zap.SugaredLogger
	defaultLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
)

func init() {
	SetLogger(NewStdOut(defaultLevel))
}

// New создает экземпляр *zap.SugaredLogger cо стандартным json выводом.
// Если уровень логгирования не передан - будет использоваться уровень
// по умолчанию (zap.ErrorLevel)
func NewStdOut(level zapcore.LevelEnabler, options ...zap.Option) *zap.SugaredLogger {
	return NewWithSink(level, os.Stdout, options...)
}

// NewWithSink создает экземпляр *zap.SugaredLogger cо стандартным json выводом.
// Если уровень логгирования не передан - будет использоваться уровень
// по умолчанию (zap.ErrorLevel). Sink используется для вывода логов.
func NewWithSink(level zapcore.LevelEnabler, sink io.Writer, options ...zap.Option) *zap.SugaredLogger {
	if level == nil {
		level = defaultLevel
	}

	core := newZapCore(level, sink, false)

	return zap.New(core, options...).Sugar()
}

func newZapCore(level zapcore.LevelEnabler, sink io.Writer, isStructure bool) zapcore.Core {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var encoder zapcore.Encoder
	if isStructure {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder,
		zapcore.AddSync(sink),
		level,
	)
}

// SetLogger устанавливает глобальный логгер. Функция непотокобезопасна.
func SetLogger(l *zap.SugaredLogger) {
	global = l
}

// SetLevel устанавливает уровень логгирования глобального логгера.
func SetLevel(l zapcore.Level) {
	defaultLevel.SetLevel(l)
}

// Level возвращает текущий уровень логгирования глобального логгера.
func Level() zapcore.Level {
	return defaultLevel.Level()
}

// Logger возвращает глобальный логгер.
func Logger() *zap.SugaredLogger {
	return global
}
