package tid

import (
	"sync"
	"time"
)

var lastTime time.Time
var mux sync.Mutex

/*
计算uuid
1 + 41 +5 + 5 + 12
*/
var sequece uint64 = 0

const (
	yyyyMMddHHmmss      = "20060102150405"
	yyyy_MM_dd_HH_mm_ss = "2006-01-02 15:04:05"
)

func GetInitTime() *time.Time {
	t, err := time.Parse(yyyy_MM_dd_HH_mm_ss, "2020-01-01 00:00:00")
	if err != nil {
		return nil
	}
	return &t
}

func New(nodeId uint64, workId uint64) uint64 {
	mux.Lock()
	defer mux.Unlock()
	current := time.Now()
	if current.Before(lastTime) {
		current = time.Now()
	}
	if current.Equal(lastTime) {
		sequece++
		if sequece > 4096 {
			sequece = 0
		}
	} else {
		sequece = 0
	}
	milliTime := current.UnixMilli()
	lastTime = current
	//1 + 41 +5 + 5 + 12
	var uuid = (0 | uint64(milliTime)<<22 | nodeId<<17 | workId<<12 | sequece)
	return uuid
}
