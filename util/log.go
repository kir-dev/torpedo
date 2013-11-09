package util

import (
	"bytes"
	"log"
)

const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARNING"
	ERROR = "ERROR"
	FATAL = "FATAL"
)

func LogDebug(format string, args ...interface{}) {
	// debug messages only printed in dev mode
	if !IsDev() {
		return
	}

	logMessage(DEBUG, format, args...)
}

// logs a message with info prefix
func LogInfo(format string, args ...interface{}) {
	logMessage(INFO, format, args...)
}

// logs a message with warn prefix
func LogWarn(format string, args ...interface{}) {
	logMessage(WARN, format, args...)
}

// logs a message with error prefix
func LogError(format string, args ...interface{}) {
	logMessage(ERROR, format, args...)
}

// logs a message with fatal prefix and exists
func LogFatal(format string, args ...interface{}) {
	log.Fatalf(prepareLogMessage(FATAL, format), args)
}

func logMessage(level, format string, args ...interface{}) {
	log.Printf(prepareLogMessage(level, format), args...)
}

func prepareLogMessage(level, format string) string {
	buffer := bytes.Buffer{}

	buffer.WriteRune('[')
	buffer.WriteString(level)
	buffer.WriteRune(']')
	buffer.WriteString(" - ")
	buffer.WriteString(format)

	return buffer.String()
}
