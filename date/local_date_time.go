package date

import (
	"errors"
	"fmt"
	"time"
)

type LocalDateTime struct {
	time.Time
}

func NewLocalDateTime(t time.Time) LocalDateTime {
	return LocalDateTime{t}
}

func (t LocalDateTime) ToTime() time.Time {
	return t.Time
}

/*
*
格式化日期
yyyy-MM-dd HH:mm:ss
*/
func (t LocalDateTime) MarshalJSON() ([]byte, error) {
	//格式化秒
	return []byte(fmt.Sprintf("\"%v\"", FormatIsoDateTime(t.Time))), nil
}

/** 解析JSON */
func (t *LocalDateTime) UnmarshalJSON(data []byte) error {
	if len(data) < 1 {
		return errors.New("param invalid")
	}
	tmpValue := string(data[1 : len(data)-1])
	value, err := ParseDateTime(tmpValue)
	if err != nil {
		return err
	}
	t.Time = value
	return nil
}
