package utils

import "encoding/binary"

func ToBytes(val int16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(val))
	return buf
}
