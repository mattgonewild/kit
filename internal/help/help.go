package help

import (
	"encoding/binary"
	"math/bits"
	"time"
	"unsafe"

	"github.com/ringboundio/common"
)

func BoolToInt(true bool) int {
	if true {
		return OneInt
	}

	return ZeroInt
}

func LittleEndianDurationToByte(duration time.Duration) []byte {
	buf := make([]byte, EightInt)
	binary.LittleEndian.PutUint64(buf, uint64(duration))
	return buf
}

func CeilPowerOfTwo(value uint) uint {
	if value < TwoUint {
		return OneUint
	}

	return OneUint << (uint(bits.Len(value - OneUint)))
}

func CeilPowerOfTwoWithShift(value uint) (uint, uint) {
	if value < TwoUint {
		return OneUint, ZeroUint
	}

	var shift = uint(bits.Len(value - OneUint))
	return OneUint << shift, shift
}

func CeilPowerOfTwoShiftOnly(value uint) uint {
	if value < TwoUint {
		return ZeroUint
	}

	return uint(bits.Len(value - OneUint))
}

// AlignDownPowerOfTwo aligns value down to the nearest multiple of alignment.
// Assumes alignment is a power of two. If value is already aligned, it remains unchanged.
// Example: AlignDownPowerOfTwo(10, 8) = 8, AlignDownPowerOfTwo(16, 8) = 16
func AlignDownPowerOfTwo(value, alignment uint) uint {
	return value & -alignment
}

// AlignUpPowerOfTwo aligns value up to the nearest multiple of alignment.
// Assumes alignment is a power of two. If value is already aligned, it remains unchanged.
// Example: AlignUpPowerOfTwo(10, 8) = 16, AlignUpPowerOfTwo(16, 8) = 16
func AlignUpPowerOfTwo(value, alignment uint) uint {
	return (value + alignment - OneUint) & -alignment
}

func UnixAt[T common.UnixTimestamped](base *T, index int) int64 {
	return (*GetPtr(base, index)).UnixNano()
}

func GetPtr[T any](base *T, index int) *T {
	var (
		length  = unsafe.Sizeof(*base)
		offset  = uintptr(index) * length
		pointer = unsafe.Add(unsafe.Pointer(base), offset)
	)

	return (*T)(pointer)
}

func Select1(word uint, rank uint) int {
	var (
		index = ZeroUint
		sum   = ZeroUint
	)

	sum = uint(bits.OnesCount(word & ((OneUint << ThirtyTwoUint) - OneUint)))
	if rank >= sum {
		index += ThirtyTwoUint
		rank -= sum
		word >>= ThirtyTwoUint
	}

	sum = uint(bits.OnesCount(word & ((OneUint << SixteenUint) - OneUint)))
	if rank >= sum {
		index += SixteenUint
		rank -= sum
		word >>= SixteenUint
	}

	sum = uint(bits.OnesCount(word & ((OneUint << EightUint) - OneUint)))
	if rank >= sum {
		index += EightUint
		rank -= sum
		word >>= EightUint
	}

	sum = uint(bits.OnesCount(word & ((OneUint << FourUint) - OneUint)))
	if rank >= sum {
		index += FourUint
		rank -= sum
		word >>= FourUint
	}

	sum = uint(bits.OnesCount(word & ((OneUint << TwoUint) - OneUint)))
	if rank >= sum {
		index += TwoUint
		rank -= sum
		word >>= TwoUint
	}

	if rank >= (word & OneUint) {
		index++
	}

	return int(index)
}
