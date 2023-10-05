package date

import (
	"time"

	"gitee.com/tasphere/gotools/logger"
)

const (
	yyyyMMddHHmmss = "20060102150405"
	IsoDateTime    = "2006-01-02 15:04:05"
	BasicDate      = "20060102"
)

/*
*
可解析格式为:yyyyMMdd的日期
*
*/
func ParseBasicDate(date string) *time.Time {
	t, err := time.Parse(BasicDate, date)
	if err != nil {
		logger.Info("ERROR", err)
		return nil
	}
	return &t
}

func ParseISODateTime(date string) *time.Time {
	t, err := time.Parse(IsoDateTime, date)
	if err != nil {
		logger.Info("ERROR", err)
		return nil
	}
	return &t
}

/*
*
可解析日期为yyyy-MM-dd HH:mm:ss
*/
func ParseIsoDateTime(datetime string) *time.Time {
	t, err := time.Parse(IsoDateTime, datetime)
	if err != nil {
		logger.Info("ERROR", err)
		return nil
	}
	return &t
}

/*
*
return date style : yyyyMMdd
*
*/
func FormatBasicDate(date time.Time) string {
	return date.Format(BasicDate)
}

/*
*
return date style: yyyy-MM-dd HH:mm:ss
*
*/
func FormatIsoDateTime(datetime time.Time) string {
	return datetime.Format(IsoDateTime)
}
