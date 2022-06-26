package config

import (
	"github.com/sirupsen/logrus"
)

var timeFormat = "02/Jan/2006:15:04:05 -0700"

func Logger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: timeFormat,
		PrettyPrint:     false,
	})

	err := logger.Level.UnmarshalText([]byte(LogLevel()))
	if err != nil {
		logger.SetLevel(logrus.InfoLevel)
		logger.Info("Failed to parse log level, defaulting 'info' level logs")
	}

	return logger
}
