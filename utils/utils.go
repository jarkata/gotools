package utils

import (
	"encoding/binary"
	"log/slog"
	"strconv"
	"strings"
)

func ToBytes(val int16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(val))
	return buf
}

func IsBlank(val string) bool {
	return len(val) <= 0
}

func IsNull(val interface{}) bool {
	return val == nil
}

func ArrayToInt64(value []string) int64 {
	amount := strings.Join(value, "")
	if len(amount) <= 0 {
		return 0
	}
	real, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		slog.Error("类型转换异常:", "ERR", err, "OriVal", amount)
		return -1
	}

	return real
}

func Int64ToStr(value int64) string {
	return strconv.FormatInt(value, 10)
}

func StrToInt64(value string) int64 {
	if len(value) <= 0 {
		return 0
	}
	real, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		slog.Error("类型转换异常:", "ERR", err, "OriVal", value)
		return -1
	}
	return real
}
