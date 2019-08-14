package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

type (
	// MyLogger struct
	MyLogger struct {
		*logrus.Logger
	}
)

var logger *MyLogger

func init() {
	logger = &MyLogger{
		logrus.New(),
	}
}

// Logger return MyLogger instance
func Logger() *MyLogger {
	return logger
}

// InitLog set logger
func InitLog(level string) {
	logger.SetLevel(getLogLevel(level))
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetOutput(newLogFile())
}

func newLogFile() *os.File {
	filename := "log/" + time.Now().Format("2006-01-02") + ".log"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}

func getLogLevel(level string) logrus.Level {
	lv := strings.ToUpper(level)
	switch lv {
	case "PANIC":
		return logrus.PanicLevel
	case "FATAL":
		return logrus.FatalLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "WARN":
		return logrus.WarnLevel
	case "INFO":
		return logrus.InfoLevel
	case "DEBUG":
		return logrus.DebugLevel
	default:
		return logrus.DebugLevel
	}
}

// Print logger.Print
func Print(i ...interface{}) {
	logger.Print(i...)
}

// Printf logger.Printf
func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

// Debug logger.Debug
func Debug(i ...interface{}) {
	logger.Debug(i...)
}

// Debugf logger.Debugf
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Info logger.Info
func Info(i ...interface{}) {
	logger.Info(i...)
}

// Infof logger.Infof
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warn logger.Warn
func Warn(i ...interface{}) {
	logger.Warn(i...)
}

// Warnf logger.Warnf
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Error logger.Error
func Error(i ...interface{}) {
	logger.Error(i...)
}

// Errorf logger.Errorf
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Fatal logger.Fatal
func Fatal(i ...interface{}) {
	logger.Fatal(i...)
}

// Fatalf logger.Fatalf
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Panic logger.Panic
func Panic(i ...interface{}) {
	logger.Panic(i...)
}

// Panicf logger.Panicf
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}
