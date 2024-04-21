package tid_test

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/jarkata/gotools/tid"
)

func TestGettime(t *testing.T) {
	fmt.Printf("tid.GetInitTime(): %v\n", tid.GetInitTime())
	nodeId := int64(1)
	workId := int64(2)
	id := tid.New(nodeId, workId)
	slog.Info("", "", id)
}
