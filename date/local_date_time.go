package date

import (
	"fmt"
	"time"
)

type LocalTime struct {
	time.Time
}

func New(t time.Time) LocalTime {
	return LocalTime{t}
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	//格式化秒
	return []byte(fmt.Sprintf("\"%v\"", t.Time.Format("2006-01-02 15:04:05"))), nil
}
