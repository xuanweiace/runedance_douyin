package tools

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestLoggerInit(t *testing.T) {
	LoggerInit()
	log.Info("runedance")
}
