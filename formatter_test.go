package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("hello logger")
	logger.Warn("hello logger")
	logger.Error("hello logger")
}
