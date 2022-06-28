package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.Info("hello logging")
	logger.WithField("useername", "novan").Info("hello logging")

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "novan")
	entry.Info("hello entry")
}
