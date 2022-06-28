package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("usernamme", "codino").Info("hello world")
	logger.WithField("username", "novan").WithField("name", "Novan Java").Info("hello world")
	logger.Info("hello world")
}

func TestFileds(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithFields(logrus.Fields{
		"username": "novan",
		"name":     "Novan Java",
	}).Info("hello world")

}
