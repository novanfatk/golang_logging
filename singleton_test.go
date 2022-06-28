package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("hello world")
	logrus.Warn("hello world")

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("hello world")
	logrus.Warn("hello world")
}
