package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Setup() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel) // Set to logrus.DebugLevel as needed
}
