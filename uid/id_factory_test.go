package uid_test

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/jarkata/gotools/uid"
)

func TestGettime(t *testing.T) {
	fmt.Printf("tid.GetInitTime(): %v\n", uid.GetInitTime())
	nodeId := uint64(1)
	workId := uint64(2)
	id := uid.New(nodeId, workId)
	slog.Info("", "", id)
}
