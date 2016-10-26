package log

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"io"
	"os"
)

var (
	logger = &logrus.Logger{
		Out:       initLogOut(),
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}
)

func EnableTimestamps() {
	// Logrus always enable timestamps
}

func EnableDebug() {
	logger.Level = logrus.DebugLevel
}

func Debug(v ...interface{}) {
	logger.Debug(fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
	logger.Debug(fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	logger.Info(fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	logger.Info(fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	logger.Error(fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	logger.Error(fmt.Sprintf(format, v...))
}

func Warning(v ...interface{}) {
	logger.Warn(fmt.Sprint(v...))
}

func Warningf(format string, v ...interface{}) {
	logger.Warn(fmt.Sprintf(format, v...))
}

func Fatal(v ...interface{}) {
	logger.Fatal(fmt.Sprint(v...))
}

func Fatalf(format string, v ...interface{}) {
	logger.Fatal(fmt.Sprintf(format, v...))
}

func initLogOut() io.Writer {
	var result io.Writer
	logPath := os.Getenv("DEX_LOG_PATH")
	if logPath != "" {
		logfile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return os.Stderr
		}
		result = logfile
	} else {
		result = os.Stdout
	}

	return result
}

type logWriter string

func (l logWriter) Write(p []byte) (int, error) {
	logger.Info(string(p))
	return len(p), nil
}

func InfoWriter() io.Writer {
	return logWriter("INFO")
}
