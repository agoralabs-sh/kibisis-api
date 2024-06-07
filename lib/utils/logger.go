package utils

import (
	"github.com/fatih/color"
	"log"
	"os"
	"runtime"
)

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelError
	LogLevelInfo
	LogLevelSilent
	LogLevelWarn
)

type Logger struct {
	Level LogLevel
}

// NewLogger Creates a new logger using the "ENVIRONMENT" env var. It defaults to error level logs.
//
//	"development" = debug level - all types of logs are printed
//	"test" = silent level - no logs are printed
func NewLogger() *Logger {
	logLevel := LogLevelError

	switch os.Getenv("ENVIRONMENT") {
	case "development":
		logLevel = LogLevelDebug
		break
	case "test":
		logLevel = LogLevelSilent
		break
	default:
		logLevel = LogLevelError
		break
	}

	return &Logger{
		Level: logLevel,
	}
}

// NewLoggerWithLevel Creates a new logger with a specific level.
func NewLoggerWithLevel(level LogLevel) *Logger {
	return &Logger{
		Level: level,
	}
}

func (logger *Logger) canLog(allowedLevel LogLevel) bool {
	switch logger.Level {
	case LogLevelError:
		return allowedLevel == LogLevelError
	case LogLevelWarn:
		return allowedLevel == LogLevelError || allowedLevel == LogLevelInfo
	case LogLevelInfo:
		return allowedLevel == LogLevelError || allowedLevel == LogLevelInfo || allowedLevel == LogLevelWarn
	case LogLevelDebug:
		return true
	default: // silent
		return false
	}
}

func (logger *Logger) Debug(message interface{}) {
	if logger.canLog(LogLevelDebug) {
		// use 1 to get the caller of this function
		pc, filename, line, _ := runtime.Caller(1)
		prefix := color.New(color.FgBlue).SprintFunc()

		log.Printf("%s in %s[%s:%d] %v", prefix("[debug]"), runtime.FuncForPC(pc).Name(), filename, line, message)
	}
}

func (logger *Logger) Error(message interface{}) {
	if logger.canLog(LogLevelError) {
		// use 1 to get the caller of this function
		pc, filename, line, _ := runtime.Caller(1)
		prefix := color.New(color.FgRed).SprintFunc()

		log.Printf("%s in %s[%s:%d] %v", prefix("[error]"), runtime.FuncForPC(pc).Name(), filename, line, message)
	}
}

func (logger *Logger) Info(message interface{}) {
	if logger.canLog(LogLevelInfo) {
		// use 1 to get the caller of this function
		pc, filename, line, _ := runtime.Caller(1)
		prefix := color.New(color.FgWhite).SprintFunc()

		log.Printf("%s in %s[%s:%d] %v", prefix("[info]"), runtime.FuncForPC(pc).Name(), filename, line, message)
	}
}

func (logger *Logger) Warn(message interface{}) {
	if logger.canLog(LogLevelWarn) {
		// use 1 to get the caller of this function
		pc, filename, line, _ := runtime.Caller(1)
		prefix := color.New(color.FgYellow).SprintFunc()

		log.Printf("%s in %s[%s:%d] %v", prefix("[warn]"), runtime.FuncForPC(pc).Name(), filename, line, message)
	}
}
