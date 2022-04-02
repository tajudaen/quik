package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"quik/config"
)

var (
	Log *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel(config.C.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}

	Log = &logrus.Logger{
		Level:     level,
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
	}
}

func Debug(msg string, tags map[string]interface{}) {
	if Log.Level < logrus.DebugLevel {
		return
	}
	Log.WithFields(logrus.Fields(tags)).Debug(msg)
}

func Info(msg string, tags map[string]interface{}) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(logrus.Fields(tags)).Info(msg)
}

func Error(msg string, err error, tags map[string]interface{}) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	msg = fmt.Sprintf("%s - ERROR - %v", msg, err)
	Log.WithFields(logrus.Fields(tags)).Error(msg)
}
