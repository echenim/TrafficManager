package logger

import (
	"github.com/sirupsen/logrus"
)

func Setup(level string) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel)
		return
	}
	logrus.SetLevel(lvl)
}
