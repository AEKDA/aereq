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
	SetLogger(New(defaultLevel))
}

// New создает экземпляр *zap.SugaredLogger cо стандартным json выводом.
// Если уровень логгирования не передан - будет использоваться уровень
// по умолчанию (zap.ErrorLevel)
func New(level zapcore.LevelEnabler, options ...zap.Option) *zap.SugaredLogger {
	return NewWithSink(level, os.Stdout, options...)
}

// NewWithSink создает экземпляр *zap.SugaredLogger cо стандартным json выводом.
// Если уровень логгирования не передан - будет использоваться уровень
// по умолчанию (zap.ErrorLevel). Sink используется для вывода логов.
func NewWithSink(level zapcore.LevelEnabler, sink io.Writer, options ...zap.Option) *zap.SugaredLogger {
	if level == nil {
		level = defaultLevel
	}

	core := newZapCore(level, sink)

	return zap.New(core, options...).Sugar()
}

func newZapCore(level zapcore.LevelEnabler, sink io.Writer) zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
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
		}),
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

func Debug(args ...interface{}) {
	global.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	global.Debugf(format, args...)
}

func Info(args ...interface{}) {
	global.Info(args...)
}

func Infof(format string, args ...interface{}) {
	global.Infof(format, args...)
}

func Warn(args ...interface{}) {
	global.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	global.Warnf(format, args...)
}

func Error(args ...interface{}) {
	global.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	global.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	global.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	global.Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	global.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	global.Panicf(format, args...)
}
