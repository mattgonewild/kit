package kit

import (
	"encoding/binary"
	"time"

	"github.com/ringboundio/kit/internal/help"
)

func Pack16(low, high uint64) (array [help.SixteenInt]byte) {
	binary.LittleEndian.PutUint64(array[help.ZeroInt:help.EightInt], low)
	binary.LittleEndian.PutUint64(array[help.EightInt:help.SixteenInt], high)
	return array
}

func Unpack16(array [help.SixteenInt]byte) (low, high uint64) {
	low = binary.LittleEndian.Uint64(array[help.ZeroInt:help.EightInt])
	high = binary.LittleEndian.Uint64(array[help.EightInt:help.SixteenInt])
	return low, high
}

func UnixNano() int64          { return time.Now().UnixNano() }
func BoolToInt(true bool) int  { return help.BoolToInt(true) }
func Close[T any](ch chan<- T) { close(ch) }
