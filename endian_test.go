package endian_test

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/wmentor/endian"
)

func TestOrder(t *testing.T) {
	t.Parallel()

	order := endian.GetOrder()
	val := uint32(0x87654321)
	data := (*[4]byte)(unsafe.Pointer(&val))[:]

	if order == endian.BigEndian {
		require.Equal(t, []byte{0x87, 0x65, 0x43, 0x21}, data)
	} else {
		require.Equal(t, []byte{0x21, 0x43, 0x65, 0x87}, data)
	}

	require.Equal(t, val, endian.GetCodec().Uint32(data))
}
