package logger_test

import (
	"testing"

	"github.com/jarkata/gotools/logger"
)

func TestInfo(t *testing.T) {
	logger.Info("teest", "fsdfs")
}
