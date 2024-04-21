package date

import (
	"errors"
	"time"
)

type LocalTime struct {
	time.Time
}

func New(t time.Time) LocalTime {
	return LocalTime{t}
}

func (t LocalTime) ToTime() time.Time {
	return t.Time
}

/*
*
格式化日期
yyyy-MM-dd HH:mm:ss
*/
func (t LocalTime) MarshalJSON() ([]byte, error) {
	//格式化秒
	return []byte(FormatIsoDateTime(t.Time)), nil
}

/** 解析JSON */
func (t *LocalTime) UnmarshalJSON(data []byte) error {
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
