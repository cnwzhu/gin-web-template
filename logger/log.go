package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *logrus.Logger

type Config struct {
	Filename string
	MaxSize  int
	MaxAge   int
	Compress bool
}

func Init(config *Config, profile string) {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	if "dev" != profile {
		logger.SetOutput(&lumberjack.Logger{
			Filename: config.Filename,
			MaxSize:  config.MaxSize,
			MaxAge:   config.MaxAge,
			Compress: config.Compress,
		})
	} else {
		logger.SetReportCaller(true)
	}
	Logger = logger
}
