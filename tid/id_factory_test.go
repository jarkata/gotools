package tid_test

import (
	"fmt"
	"testing"

	"github.com/jarkata/gotools/tid"
)

func TestGettime(t *testing.T) {
	fmt.Printf("tid.GetInitTime(): %v\n", tid.GetInitTime())
}
