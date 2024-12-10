package logger

import (
	"os"

	"strings"

	"github.com/rs/zerolog"
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
)

type Logger struct {
	log            zerolog.Logger
	showCallerInfo bool
}

const LoggerId = interfaces.LoggerInterfaceId

func NewLogger(logLevel string) interfaces.LoggerInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	instance := &Logger{
		log:            log,
		showCallerInfo: true,
	}

	instance.SetLevel(logLevel)

	dependencyContainer.AddDependency(LoggerId, instance)

	return instance
}

func (l *Logger) SetLevel(logLevel string) {
	switch strings.ToLower(logLevel) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}
}

func (l *Logger) DisableCallerInfo() {
	l.showCallerInfo = false
}

func (l *Logger) EnableCallerInfo() {
	l.showCallerInfo = true
}

func (l *Logger) addCallerInfo(log *zerolog.Event) *zerolog.Event {
	if l.showCallerInfo {
		callerInfo := NewCallerInfo(3)
		log.Str("caller", callerInfo.Caller)
		log.Str("file", callerInfo.File)
		log.Int("line", callerInfo.Line)
	}

	return log
}

func (l *Logger) Debug(message string, keysAndValues ...interface{}) {
	log := l.log.Debug()

	l.addCallerInfo(log)

	log.Fields(keysAndValues).Msg(message)
}

func (l *Logger) Info(message string, keysAndValues ...interface{}) {
	log := l.log.Info()

	l.addCallerInfo(log)

	log.Fields(keysAndValues).Msg(message)
}

func (l *Logger) Warn(message string, keysAndValues ...interface{}) {
	log := l.log.Warn()

	l.addCallerInfo(log)

	log.Fields(keysAndValues).Msg(message)
}

func (l *Logger) Error(message string, keysAndValues ...interface{}) {
	log := l.log.Error()

	l.addCallerInfo(log)

	log.Fields(keysAndValues).Msg(message)
}

func (l *Logger) Fatal(message string, keysAndValues ...interface{}) {
	log := l.log.Fatal()

	l.addCallerInfo(log)

	log.Fields(keysAndValues).Msg(message)
}
