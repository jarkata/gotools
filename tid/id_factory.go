package tid

import (
	"time"
)

var lastTime time.Time

/*
计算uuid
1 + 41 +5 + 5 + 12
*/
var sequece int = 0

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

func New(nodeId int64, workId int64) int64 {
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
	var uuid = (0 | milliTime<<22 | int64(nodeId)<<17 | int64(workId)<<12 | int64(sequece))
	return uuid
}
