package logger_test

import (
	"testing"

	"gitee.com/tasphere/gotools/logger"
)

func TestInfo(t *testing.T) {
	logger.Info("teest", "fsdfs")
}
