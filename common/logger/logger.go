package logger

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
)

func init() {

	// Enable logrus to DebugLevel. Debug level messages will only show on console.
	logrus.SetLevel(logrus.DebugLevel)

	// Set text formatter with time format to millisecond without date.
	logrus.SetFormatter(&prefixed.TextFormatter{
		TimestampFormat: "15:04:05.000",
		ForceColors:     true,
		SpacePadding:    15,
	})
}

type Logger struct {
	logger *logrus.Entry
}

type Fields map[string]interface{}

func NewLogger(name string) *Logger {

	with_fields := make(logrus.Fields)

	with_fields["prefix"] = fmt.Sprintf("%25s", name)

	internal_log := logrus.WithFields(with_fields)

	logger := &Logger{
		logger: internal_log,
	}

	return logger
}

func (this *Logger) WithFields(fields Fields) *Logger {

	logger := &Logger{
		logger: this.logger.WithFields(logrus.Fields(fields)),
	}

	return logger
}

// PanicLevel level, highest level of severity.
// Logs and then calls panic with the message passed to Debug, Info, ...
func (this *Logger) PANIC(format string, args ...interface{}) {
	this.logger.Panicf(format, args...)
}

// ErrorLevel level. Logs. Used for errors that should definitely be noted.
// Commonly used for hooks to send errors to an error tracking service.
func (this *Logger) ERROR(format string, args ...interface{}) {
	this.logger.Errorf(format, args...)
}

// WarnLevel level. Non-critical entries that deserve eyes.
func (this *Logger) WARN(format string, args ...interface{}) {
	this.logger.Warnf(format, args...)
}

// InfoLevel level. General operational entries about what's going on inside the application.
func (this *Logger) INFO(format string, args ...interface{}) {
	this.logger.Infof(format, args...)
}

// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
func (this *Logger) DEBUG(format string, args ...interface{}) {
	this.logger.Debugf(format, args...)
}
