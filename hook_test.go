package golanglogging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SimpleHook struct {
}

func (s *SimpleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SimpleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Simple hook", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SimpleHook{})
	logger.Info("hello info")
	logger.Warn("hello warn")
	logger.Error("hello error")
}
