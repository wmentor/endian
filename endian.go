package endian

import (
	"encoding/binary"
	"unsafe"
)

type (
	ByteOrder int
	Codec     = binary.ByteOrder
)

const (
	BigEndian ByteOrder = iota
	LittleEndian
)

const (
	intSize = int(unsafe.Sizeof(int(0)))
)

var (
	order ByteOrder
)

func init() {
	i := int(1)
	if (*[intSize]byte)(unsafe.Pointer(&i))[0] == 1 {
		order = LittleEndian
	} else {
		order = BigEndian
	}
}

func GetOrder() ByteOrder {
	return order
}

func GetCodec() Codec {
	if order == LittleEndian {
		return binary.LittleEndian
	}
	return binary.BigEndian
}
