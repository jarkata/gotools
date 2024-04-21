package date

import (
	"strings"
	"time"

	"github.com/jarkata/gotools/logger"
)

const (
	yyyyMMddHHmmss  = "20060102150405"
	IsoDateTime     = "2006-01-02 15:04:05"
	IsoDateTime_1   = "2006-01-02T15:04:05"
	IsoDateTime_2_1 = "2006-01-02T15:04:05.999"
	IsoDateTime_2_2 = "2006-01-02 15:04:05.999"
	IsoDateTime_3_1 = "2006-01-02T15:04:05.999Z"
	IsoDateTime_3_2 = "2006-01-02 15:04:05.999Z"
	BasicDate       = "20060102"
)

/*
*
解析以下格式的数据

	yyyyMMddHHmmss  = "20060102150405"
	IsoDateTime     = "2006-01-02 15:04:05"
	IsoDateTime_1   = "2006-01-02T15:04:05"
	IsoDateTime_2_1 = "2006-01-02T15:04:05.999"
	IsoDateTime_2_2 = "2006-01-02 15:04:05.999"
	IsoDateTime_3_1 = "2006-01-02T15:04:05.999Z"
	IsoDateTime_3_2 = "2006-01-02 15:04:05.999Z"
	BasicDate       = "20060102"
*/
func ParseDateTime(date string) (time.Time, error) {
	tmpValue := date
	valueLen := len(tmpValue)
	if valueLen == 8 {
		return time.Parse(BasicDate, tmpValue)
	} else if valueLen == 10 {
		return time.Parse(time.DateOnly, tmpValue)
	} else if valueLen == 14 {
		return time.Parse(yyyyMMddHHmmss, tmpValue)
	} else if valueLen == 19 {
		index := strings.IndexByte(tmpValue, 'T')
		if index > 0 {
			return time.Parse(IsoDateTime_1, tmpValue)
		} else {
			return time.Parse(time.DateTime, tmpValue)
		}
	} else if valueLen == 23 {
		index := strings.IndexByte(tmpValue, 'T')
		if index > 0 {
			return time.Parse(IsoDateTime_2_1, tmpValue)
		} else {
			return time.Parse(IsoDateTime_2_2, tmpValue)
		}
	} else if valueLen == 24 {
		index := strings.IndexByte(tmpValue, 'T')
		if index > 0 {
			return time.Parse(IsoDateTime_3_1, tmpValue)
		} else {
			return time.Parse(IsoDateTime_3_2, tmpValue)
		}
	}
	return time.Time{}, nil
}

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

/*
*
可解析日期为yyyyMMddHHmmss
*/
func ParseStdDateTime(date string) *time.Time {
	t, err := time.Parse(yyyyMMddHHmmss, date)
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
