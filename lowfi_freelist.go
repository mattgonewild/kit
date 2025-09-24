package kit

import (
	"math/bits"

	"github.com/ringboundio/kit/internal/help"
)

func NewLowfiAs() bool { return true }

type LowfiFreelist0[T any] struct {
	bitmap  uint
	element [help.L0]T
}

func NewLowfiFreelist0[T any](factory func(uint) T) *LowfiFreelist0[T] {
	freelist := new(LowfiFreelist0[T])
	for index := range help.L0 {
		freelist.element[index] = factory(index)
	}

	freelist.bitmap |= ^help.BM0
	return freelist
}

func InitLowfiFreelist0[T any](freelist *LowfiFreelist0[T], factory func(uint) T) {
	for index := range help.L0 {
		freelist.element[index] = factory(index)
	}

	freelist.bitmap |= ^help.BM0
}

func (this *LowfiFreelist0[T]) Get() (*T, bool) {
	bitmap := this.bitmap
	if bitmap == help.MaxUint {
		return nil, false
	}
	zi := uint(bits.TrailingZeros(^bitmap))

	this.bitmap = bitmap | (help.OneUint << zi)
	return help.GetPtr(&this.element[help.ZeroInt], int(zi)), true
}

func (this *LowfiFreelist0[T]) Release(id uint) bool {
	this.bitmap &^= help.OneUint << (id & help.WordRMask)
	return true
}

func (this *LowfiFreelist0[T]) Len() int {
	return bits.OnesCount(this.bitmap & help.BM0)
}

func (this *LowfiFreelist0[T]) Cap() int { return int(help.L0) }

type LowfiFreelist1[T any] struct {
	bitmap  uint
	element [help.L1]T
}

func NewLowfiFreelist1[T any](factory func(uint) T) *LowfiFreelist1[T] {
	freelist := new(LowfiFreelist1[T])
	for index := range help.L1 {
		freelist.element[index] = factory(index)
	}

	freelist.bitmap |= ^help.BM1
	return freelist
}

func InitLowfiFreelist1[T any](freelist *LowfiFreelist1[T], factory func(uint) T) {
	for index := range help.L1 {
		freelist.element[index] = factory(index)
	}

	freelist.bitmap |= ^help.BM1
}

func (this *LowfiFreelist1[T]) Get() (*T, bool) {
	bitmap := this.bitmap
	if bitmap == help.MaxUint {
		return nil, false
	}
	zi := uint(bits.TrailingZeros(^bitmap))

	this.bitmap = bitmap | (help.OneUint << zi)
	return help.GetPtr(&this.element[help.ZeroInt], int(zi)), true
}

func (this *LowfiFreelist1[T]) Release(id uint) bool {
	this.bitmap &^= help.OneUint << (id & help.WordRMask)
	return true
}

func (this *LowfiFreelist1[T]) Len() int {
	return bits.OnesCount(this.bitmap & help.BM1)
}

func (this *LowfiFreelist1[T]) Cap() int { return int(help.L1) }

type LowfiFreelist2[T any] struct {
	bitmap  uint
	element [help.L2]T
}

func NewLowfiFreelist2[T any](factory func(uint) T) *LowfiFreelist2[T] {
	freelist := new(LowfiFreelist2[T])
	for index := range help.L2 {
		freelist.element[index] = factory(index)
	}

	freelist.bitmap |= ^help.BM2
	return freelist
}

func InitLowfiFreelist2[T any](freelist *LowfiFreelist2[T], factory func(uint) T) {
	for index := range help.L2 {
		freelist.element[index] = factory(index)
	}

	freelist.bitmap |= ^help.BM2
}

func (this *LowfiFreelist2[T]) Get() (*T, bool) {
	bitmap := this.bitmap
	if bitmap == help.MaxUint {
		return nil, false
	}
	zi := uint(bits.TrailingZeros(^bitmap))

	this.bitmap = bitmap | (help.OneUint << zi)
	return help.GetPtr(&this.element[help.ZeroInt], int(zi)), true
}

func (this *LowfiFreelist2[T]) Release(id uint) bool {
	this.bitmap &^= help.OneUint << (id & help.WordRMask)
	return true
}

func (this *LowfiFreelist2[T]) Len() int {
	return bits.OnesCount(this.bitmap & help.BM2)
}

func (this *LowfiFreelist2[T]) Cap() int { return int(help.L2) }

type LowfiFreelist3[T any] struct {
	bitmap  uint
	element [help.L3]T
}

func NewLowfiFreelist3[T any](factory func(uint) T) *LowfiFreelist3[T] {
	freelist := new(LowfiFreelist3[T])
	for index := range help.L3 {
		freelist.element[index] = factory(index)
	}

	freelist.bitmap |= ^help.BM3
	return freelist
}

func InitLowfiFreelist3[T any](freelist *LowfiFreelist3[T], factory func(uint) T) {
	for index := range help.L3 {
		freelist.element[index] = factory(index)
	}

	freelist.bitmap |= ^help.BM3
}

func (this *LowfiFreelist3[T]) Get() (*T, bool) {
	bitmap := this.bitmap
	if bitmap == help.MaxUint {
		return nil, false
	}
	zi := uint(bits.TrailingZeros(^bitmap))

	this.bitmap = bitmap | (help.OneUint << zi)
	return help.GetPtr(&this.element[help.ZeroInt], int(zi)), true
}

func (this *LowfiFreelist3[T]) Release(id uint) bool {
	this.bitmap &^= help.OneUint << (id & help.WordRMask)
	return true
}

func (this *LowfiFreelist3[T]) Len() int {
	return bits.OnesCount(this.bitmap & help.BM3)
}

func (this *LowfiFreelist3[T]) Cap() int { return int(help.L3) }

type LowfiFreelist4[T any] struct {
	bitmap  uint
	element [help.L4]T
}

func NewLowfiFreelist4[T any](factory func(uint) T) *LowfiFreelist4[T] {
	freelist := new(LowfiFreelist4[T])
	for index := range help.L4 {
		freelist.element[index] = factory(index)
	}

	freelist.bitmap |= ^help.BM4
	return freelist
}

func InitLowfiFreelist4[T any](freelist *LowfiFreelist4[T], factory func(uint) T) {
	for index := range help.L4 {
		freelist.element[index] = factory(index)
	}

	freelist.bitmap |= ^help.BM4
}

func (this *LowfiFreelist4[T]) Get() (*T, bool) {
	bitmap := this.bitmap
	if bitmap == help.MaxUint {
		return nil, false
	}
	zi := uint(bits.TrailingZeros(^bitmap))

	this.bitmap = bitmap | (help.OneUint << zi)
	return help.GetPtr(&this.element[help.ZeroInt], int(zi)), true
}

func (this *LowfiFreelist4[T]) Release(id uint) bool {
	this.bitmap &^= help.OneUint << (id & help.WordRMask)
	return true
}

func (this *LowfiFreelist4[T]) Len() int {
	return bits.OnesCount(this.bitmap & help.BM4)
}

func (this *LowfiFreelist4[T]) Cap() int { return int(help.L4) }

type LowfiFreelist5[T any] struct {
	bitmap  uint
	element [help.L5]T
}

func NewLowfiFreelist5[T any](factory func(uint) T) *LowfiFreelist5[T] {
	freelist := new(LowfiFreelist5[T])
	for index := range help.L5 {
		freelist.element[index] = factory(index)
	}

	freelist.bitmap |= ^help.BM5
	return freelist
}

func InitLowfiFreelist5[T any](freelist *LowfiFreelist5[T], factory func(uint) T) {
	for index := range help.L5 {
		freelist.element[index] = factory(index)
	}

	freelist.bitmap |= ^help.BM5
}

func (this *LowfiFreelist5[T]) Get() (*T, bool) {
	bitmap := this.bitmap
	if bitmap == help.MaxUint {
		return nil, false
	}
	zi := uint(bits.TrailingZeros(^bitmap))

	this.bitmap = bitmap | (help.OneUint << zi)
	return help.GetPtr(&this.element[help.ZeroInt], int(zi)), true
}

func (this *LowfiFreelist5[T]) Release(id uint) bool {
	this.bitmap &^= help.OneUint << (id & help.WordRMask)
	return true
}

func (this *LowfiFreelist5[T]) Len() int {
	return bits.OnesCount(this.bitmap & help.BM5)
}

func (this *LowfiFreelist5[T]) Cap() int { return int(help.L5) }

type LowfiFreelist6[T any] struct {
	bitmap  uint
	element [help.L6]T
}

func NewLowfiFreelist6[T any](factory func(uint) T) *LowfiFreelist6[T] {
	freelist := new(LowfiFreelist6[T])
	for index := range help.L6 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitLowfiFreelist6[T any](freelist *LowfiFreelist6[T], factory func(uint) T) {
	for index := range help.L6 {
		freelist.element[index] = factory(index)
	}
}

func (this *LowfiFreelist6[T]) Get() (*T, bool) {
	bitmap := this.bitmap
	if bitmap == help.MaxUint {
		return nil, false
	}
	zi := uint(bits.TrailingZeros(^bitmap))

	this.bitmap = bitmap | (help.OneUint << zi)
	return help.GetPtr(&this.element[help.ZeroInt], int(zi)), true
}

func (this *LowfiFreelist6[T]) Release(id uint) bool {
	this.bitmap &^= help.OneUint << (id & help.WordRMask)
	return true
}

func (this *LowfiFreelist6[T]) Len() int {
	return bits.OnesCount(this.bitmap)
}

func (this *LowfiFreelist6[T]) Cap() int { return int(help.L6) }

type LowfiFreelist7[T any] struct {
	summary [help.SL7]uint
	bitmap  [help.WL7]uint
	element [help.L7]T
	cursor  uint
}

func NewLowfiFreelist7[T any](factory func(uint) T) *LowfiFreelist7[T] {
	freelist := new(LowfiFreelist7[T])
	for index := range help.L7 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM7] |= ^help.SBM7
	return freelist
}

func InitLowfiFreelist7[T any](freelist *LowfiFreelist7[T], factory func(uint) T) {
	for index := range help.L7 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM7] |= ^help.SBM7
}

func (this *LowfiFreelist7[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL7; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM7
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist7[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist7[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL7 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist7[T]) Cap() int { return int(help.L7) }

type LowfiFreelist8[T any] struct {
	summary [help.SL8]uint
	bitmap  [help.WL8]uint
	element [help.L8]T
	cursor  uint
}

func NewLowfiFreelist8[T any](factory func(uint) T) *LowfiFreelist8[T] {
	freelist := new(LowfiFreelist8[T])
	for index := range help.L8 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM8] |= ^help.SBM8
	return freelist
}

func InitLowfiFreelist8[T any](freelist *LowfiFreelist8[T], factory func(uint) T) {
	for index := range help.L8 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM8] |= ^help.SBM8
}

func (this *LowfiFreelist8[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL8; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM8
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist8[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist8[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL8 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist8[T]) Cap() int { return int(help.L8) }

type LowfiFreelist9[T any] struct {
	summary [help.SL9]uint
	bitmap  [help.WL9]uint
	element [help.L9]T
	cursor  uint
}

func NewLowfiFreelist9[T any](factory func(uint) T) *LowfiFreelist9[T] {
	freelist := new(LowfiFreelist9[T])
	for index := range help.L9 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM9] |= ^help.SBM9
	return freelist
}

func InitLowfiFreelist9[T any](freelist *LowfiFreelist9[T], factory func(uint) T) {
	for index := range help.L9 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM9] |= ^help.SBM9
}

func (this *LowfiFreelist9[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL9; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM9
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist9[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist9[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL9 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist9[T]) Cap() int { return int(help.L9) }

type LowfiFreelist10[T any] struct {
	summary [help.SL10]uint
	bitmap  [help.WL10]uint
	element [help.L10]T
	cursor  uint
}

func NewLowfiFreelist10[T any](factory func(uint) T) *LowfiFreelist10[T] {
	freelist := new(LowfiFreelist10[T])
	for index := range help.L10 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM10] |= ^help.SBM10
	return freelist
}

func InitLowfiFreelist10[T any](freelist *LowfiFreelist10[T], factory func(uint) T) {
	for index := range help.L10 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM10] |= ^help.SBM10
}

func (this *LowfiFreelist10[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL10; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM10
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist10[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist10[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL10 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist10[T]) Cap() int { return int(help.L10) }

type LowfiFreelist11[T any] struct {
	summary [help.SL11]uint
	bitmap  [help.WL11]uint
	element [help.L11]T
	cursor  uint
}

func NewLowfiFreelist11[T any](factory func(uint) T) *LowfiFreelist11[T] {
	freelist := new(LowfiFreelist11[T])
	for index := range help.L11 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM11] |= ^help.SBM11
	return freelist
}

func InitLowfiFreelist11[T any](freelist *LowfiFreelist11[T], factory func(uint) T) {
	for index := range help.L11 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM11] |= ^help.SBM11
}

func (this *LowfiFreelist11[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL11; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM11
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist11[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist11[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL11 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist11[T]) Cap() int { return int(help.L11) }

type LowfiFreelist12[T any] struct {
	summary [help.SL12]uint
	bitmap  [help.WL12]uint
	element [help.L12]T
	cursor  uint
}

func NewLowfiFreelist12[T any](factory func(uint) T) *LowfiFreelist12[T] {
	freelist := new(LowfiFreelist12[T])
	for index := range help.L12 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM12] |= ^help.SBM12
	return freelist
}

func InitLowfiFreelist12[T any](freelist *LowfiFreelist12[T], factory func(uint) T) {
	for index := range help.L12 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM12] |= ^help.SBM12
}

func (this *LowfiFreelist12[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL12; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM12
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist12[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist12[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL12 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist12[T]) Cap() int { return int(help.L12) }

type LowfiFreelist13[T any] struct {
	summary [help.SL13]uint
	bitmap  [help.WL13]uint
	element [help.L13]T
	cursor  uint
}

func NewLowfiFreelist13[T any](factory func(uint) T) *LowfiFreelist13[T] {
	freelist := new(LowfiFreelist13[T])
	for index := range help.L13 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM13] |= ^help.SBM13
	return freelist
}

func InitLowfiFreelist13[T any](freelist *LowfiFreelist13[T], factory func(uint) T) {
	for index := range help.L13 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM13] |= ^help.SBM13
}

func (this *LowfiFreelist13[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL13; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM13
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist13[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist13[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL13 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist13[T]) Cap() int { return int(help.L13) }

type LowfiFreelist14[T any] struct {
	summary [help.SL14]uint
	bitmap  [help.WL14]uint
	element [help.L14]T
	cursor  uint
}

func NewLowfiFreelist14[T any](factory func(uint) T) *LowfiFreelist14[T] {
	freelist := new(LowfiFreelist14[T])
	for index := range help.L14 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM14] |= ^help.SBM14
	return freelist
}

func InitLowfiFreelist14[T any](freelist *LowfiFreelist14[T], factory func(uint) T) {
	for index := range help.L14 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM14] |= ^help.SBM14
}

func (this *LowfiFreelist14[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL14; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM14
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist14[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist14[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL14 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist14[T]) Cap() int { return int(help.L14) }

type LowfiFreelist15[T any] struct {
	summary [help.SL15]uint
	bitmap  [help.WL15]uint
	element [help.L15]T
	cursor  uint
}

func NewLowfiFreelist15[T any](factory func(uint) T) *LowfiFreelist15[T] {
	freelist := new(LowfiFreelist15[T])
	for index := range help.L15 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM15] |= ^help.SBM15
	return freelist
}

func InitLowfiFreelist15[T any](freelist *LowfiFreelist15[T], factory func(uint) T) {
	for index := range help.L15 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM15] |= ^help.SBM15
}

func (this *LowfiFreelist15[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL15; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM15
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist15[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist15[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL15 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist15[T]) Cap() int { return int(help.L15) }

type LowfiFreelist16[T any] struct {
	summary [help.SL16]uint
	bitmap  [help.WL16]uint
	element [help.L16]T
	cursor  uint
}

func NewLowfiFreelist16[T any](factory func(uint) T) *LowfiFreelist16[T] {
	freelist := new(LowfiFreelist16[T])
	for index := range help.L16 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM16] |= ^help.SBM16
	return freelist
}

func InitLowfiFreelist16[T any](freelist *LowfiFreelist16[T], factory func(uint) T) {
	for index := range help.L16 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM16] |= ^help.SBM16
}

func (this *LowfiFreelist16[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL16; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM16
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist16[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist16[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL16 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist16[T]) Cap() int { return int(help.L16) }

type LowfiFreelist17[T any] struct {
	summary [help.SL17]uint
	bitmap  [help.WL17]uint
	element [help.L17]T
	cursor  uint
}

func NewLowfiFreelist17[T any](factory func(uint) T) *LowfiFreelist17[T] {
	freelist := new(LowfiFreelist17[T])
	for index := range help.L17 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM17] |= ^help.SBM17
	return freelist
}

func InitLowfiFreelist17[T any](freelist *LowfiFreelist17[T], factory func(uint) T) {
	for index := range help.L17 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM17] |= ^help.SBM17
}

func (this *LowfiFreelist17[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL17; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM17
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist17[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist17[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL17 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist17[T]) Cap() int { return int(help.L17) }

type LowfiFreelist18[T any] struct {
	summary [help.SL18]uint
	bitmap  [help.WL18]uint
	element [help.L18]T
	cursor  uint
}

func NewLowfiFreelist18[T any](factory func(uint) T) *LowfiFreelist18[T] {
	freelist := new(LowfiFreelist18[T])
	for index := range help.L18 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM18] |= ^help.SBM18
	return freelist
}

func InitLowfiFreelist18[T any](freelist *LowfiFreelist18[T], factory func(uint) T) {
	for index := range help.L18 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM18] |= ^help.SBM18
}

func (this *LowfiFreelist18[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL18; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM18
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist18[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist18[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL18 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist18[T]) Cap() int { return int(help.L18) }

type LowfiFreelist19[T any] struct {
	summary [help.SL19]uint
	bitmap  [help.WL19]uint
	element [help.L19]T
	cursor  uint
}

func NewLowfiFreelist19[T any](factory func(uint) T) *LowfiFreelist19[T] {
	freelist := new(LowfiFreelist19[T])
	for index := range help.L19 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM19] |= ^help.SBM19
	return freelist
}

func InitLowfiFreelist19[T any](freelist *LowfiFreelist19[T], factory func(uint) T) {
	for index := range help.L19 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM19] |= ^help.SBM19
}

func (this *LowfiFreelist19[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL19; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM19
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist19[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist19[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL19 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist19[T]) Cap() int { return int(help.L19) }

type LowfiFreelist20[T any] struct {
	summary [help.SL20]uint
	bitmap  [help.WL20]uint
	element [help.L20]T
	cursor  uint
}

func NewLowfiFreelist20[T any](factory func(uint) T) *LowfiFreelist20[T] {
	freelist := new(LowfiFreelist20[T])
	for index := range help.L20 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM20] |= ^help.SBM20
	return freelist
}

func InitLowfiFreelist20[T any](freelist *LowfiFreelist20[T], factory func(uint) T) {
	for index := range help.L20 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM20] |= ^help.SBM20
}

func (this *LowfiFreelist20[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL20; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM20
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist20[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist20[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL20 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist20[T]) Cap() int { return int(help.L20) }

type LowfiFreelist21[T any] struct {
	summary [help.SL21]uint
	bitmap  [help.WL21]uint
	element [help.L21]T
	cursor  uint
}

func NewLowfiFreelist21[T any](factory func(uint) T) *LowfiFreelist21[T] {
	freelist := new(LowfiFreelist21[T])
	for index := range help.L21 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM21] |= ^help.SBM21
	return freelist
}

func InitLowfiFreelist21[T any](freelist *LowfiFreelist21[T], factory func(uint) T) {
	for index := range help.L21 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM21] |= ^help.SBM21
}

func (this *LowfiFreelist21[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL21; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM21
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist21[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist21[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL21 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist21[T]) Cap() int { return int(help.L21) }

type LowfiFreelist22[T any] struct {
	summary [help.SL22]uint
	bitmap  [help.WL22]uint
	element [help.L22]T
	cursor  uint
}

func NewLowfiFreelist22[T any](factory func(uint) T) *LowfiFreelist22[T] {
	freelist := new(LowfiFreelist22[T])
	for index := range help.L22 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM22] |= ^help.SBM22
	return freelist
}

func InitLowfiFreelist22[T any](freelist *LowfiFreelist22[T], factory func(uint) T) {
	for index := range help.L22 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM22] |= ^help.SBM22
}

func (this *LowfiFreelist22[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL22; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM22
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist22[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist22[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL22 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist22[T]) Cap() int { return int(help.L22) }

type LowfiFreelist23[T any] struct {
	summary [help.SL23]uint
	bitmap  [help.WL23]uint
	element [help.L23]T
	cursor  uint
}

func NewLowfiFreelist23[T any](factory func(uint) T) *LowfiFreelist23[T] {
	freelist := new(LowfiFreelist23[T])
	for index := range help.L23 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM23] |= ^help.SBM23
	return freelist
}

func InitLowfiFreelist23[T any](freelist *LowfiFreelist23[T], factory func(uint) T) {
	for index := range help.L23 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM23] |= ^help.SBM23
}

func (this *LowfiFreelist23[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL23; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM23
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist23[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist23[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL23 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist23[T]) Cap() int { return int(help.L23) }

type LowfiFreelist24[T any] struct {
	summary [help.SL24]uint
	bitmap  [help.WL24]uint
	element [help.L24]T
	cursor  uint
}

func NewLowfiFreelist24[T any](factory func(uint) T) *LowfiFreelist24[T] {
	freelist := new(LowfiFreelist24[T])
	for index := range help.L24 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM24] |= ^help.SBM24
	return freelist
}

func InitLowfiFreelist24[T any](freelist *LowfiFreelist24[T], factory func(uint) T) {
	for index := range help.L24 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM24] |= ^help.SBM24
}

func (this *LowfiFreelist24[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL24; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM24
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist24[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist24[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL24 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist24[T]) Cap() int { return int(help.L24) }

type LowfiFreelist25[T any] struct {
	summary [help.SL25]uint
	bitmap  [help.WL25]uint
	element [help.L25]T
	cursor  uint
}

func NewLowfiFreelist25[T any](factory func(uint) T) *LowfiFreelist25[T] {
	freelist := new(LowfiFreelist25[T])
	for index := range help.L25 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM25] |= ^help.SBM25
	return freelist
}

func InitLowfiFreelist25[T any](freelist *LowfiFreelist25[T], factory func(uint) T) {
	for index := range help.L25 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM25] |= ^help.SBM25
}

func (this *LowfiFreelist25[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL25; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM25
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist25[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist25[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL25 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist25[T]) Cap() int { return int(help.L25) }

type LowfiFreelist26[T any] struct {
	summary [help.SL26]uint
	bitmap  [help.WL26]uint
	element [help.L26]T
	cursor  uint
}

func NewLowfiFreelist26[T any](factory func(uint) T) *LowfiFreelist26[T] {
	freelist := new(LowfiFreelist26[T])
	for index := range help.L26 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM26] |= ^help.SBM26
	return freelist
}

func InitLowfiFreelist26[T any](freelist *LowfiFreelist26[T], factory func(uint) T) {
	for index := range help.L26 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM26] |= ^help.SBM26
}

func (this *LowfiFreelist26[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL26; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM26
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist26[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist26[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL26 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist26[T]) Cap() int { return int(help.L26) }

type LowfiFreelist27[T any] struct {
	summary [help.SL27]uint
	bitmap  [help.WL27]uint
	element [help.L27]T
	cursor  uint
}

func NewLowfiFreelist27[T any](factory func(uint) T) *LowfiFreelist27[T] {
	freelist := new(LowfiFreelist27[T])
	for index := range help.L27 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM27] |= ^help.SBM27
	return freelist
}

func InitLowfiFreelist27[T any](freelist *LowfiFreelist27[T], factory func(uint) T) {
	for index := range help.L27 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM27] |= ^help.SBM27
}

func (this *LowfiFreelist27[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL27; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM27
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist27[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist27[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL27 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist27[T]) Cap() int { return int(help.L27) }

type LowfiFreelist28[T any] struct {
	summary [help.SL28]uint
	bitmap  [help.WL28]uint
	element [help.L28]T
	cursor  uint
}

func NewLowfiFreelist28[T any](factory func(uint) T) *LowfiFreelist28[T] {
	freelist := new(LowfiFreelist28[T])
	for index := range help.L28 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM28] |= ^help.SBM28
	return freelist
}

func InitLowfiFreelist28[T any](freelist *LowfiFreelist28[T], factory func(uint) T) {
	for index := range help.L28 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM28] |= ^help.SBM28
}

func (this *LowfiFreelist28[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL28; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM28
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist28[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist28[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL28 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist28[T]) Cap() int { return int(help.L28) }

type LowfiFreelist29[T any] struct {
	summary [help.SL29]uint
	bitmap  [help.WL29]uint
	element [help.L29]T
	cursor  uint
}

func NewLowfiFreelist29[T any](factory func(uint) T) *LowfiFreelist29[T] {
	freelist := new(LowfiFreelist29[T])
	for index := range help.L29 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM29] |= ^help.SBM29
	return freelist
}

func InitLowfiFreelist29[T any](freelist *LowfiFreelist29[T], factory func(uint) T) {
	for index := range help.L29 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM29] |= ^help.SBM29
}

func (this *LowfiFreelist29[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL29; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM29
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist29[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist29[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL29 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist29[T]) Cap() int { return int(help.L29) }

type LowfiFreelist30[T any] struct {
	summary [help.SL30]uint
	bitmap  [help.WL30]uint
	element [help.L30]T
	cursor  uint
}

func NewLowfiFreelist30[T any](factory func(uint) T) *LowfiFreelist30[T] {
	freelist := new(LowfiFreelist30[T])
	for index := range help.L30 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM30] |= ^help.SBM30
	return freelist
}

func InitLowfiFreelist30[T any](freelist *LowfiFreelist30[T], factory func(uint) T) {
	for index := range help.L30 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM30] |= ^help.SBM30
}

func (this *LowfiFreelist30[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL30; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM30
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist30[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist30[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL30 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist30[T]) Cap() int { return int(help.L30) }

type LowfiFreelist31[T any] struct {
	summary [help.SL31]uint
	bitmap  [help.WL31]uint
	element [help.L31]T
	cursor  uint
}

func NewLowfiFreelist31[T any](factory func(uint) T) *LowfiFreelist31[T] {
	freelist := new(LowfiFreelist31[T])
	for index := range help.L31 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM31] |= ^help.SBM31
	return freelist
}

func InitLowfiFreelist31[T any](freelist *LowfiFreelist31[T], factory func(uint) T) {
	for index := range help.L31 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM31] |= ^help.SBM31
}

func (this *LowfiFreelist31[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL31; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM31
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist31[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist31[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL31 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist31[T]) Cap() int { return int(help.L31) }

type LowfiFreelist32[T any] struct {
	summary [help.SL32]uint
	bitmap  [help.WL32]uint
	element [help.L32]T
	cursor  uint
}

func NewLowfiFreelist32[T any](factory func(uint) T) *LowfiFreelist32[T] {
	freelist := new(LowfiFreelist32[T])
	for index := range help.L32 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM32] |= ^help.SBM32
	return freelist
}

func InitLowfiFreelist32[T any](freelist *LowfiFreelist32[T], factory func(uint) T) {
	for index := range help.L32 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM32] |= ^help.SBM32
}

func (this *LowfiFreelist32[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL32; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM32
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist32[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist32[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL32 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist32[T]) Cap() int { return int(help.L32) }

type LowfiFreelist33[T any] struct {
	summary [help.SL33]uint
	bitmap  [help.WL33]uint
	element [help.L33]T
	cursor  uint
}

func NewLowfiFreelist33[T any](factory func(uint) T) *LowfiFreelist33[T] {
	freelist := new(LowfiFreelist33[T])
	for index := range help.L33 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM33] |= ^help.SBM33
	return freelist
}

func InitLowfiFreelist33[T any](freelist *LowfiFreelist33[T], factory func(uint) T) {
	for index := range help.L33 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM33] |= ^help.SBM33
}

func (this *LowfiFreelist33[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL33; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM33
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist33[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist33[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL33 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist33[T]) Cap() int { return int(help.L33) }

type LowfiFreelist34[T any] struct {
	summary [help.SL34]uint
	bitmap  [help.WL34]uint
	element [help.L34]T
	cursor  uint
}

func NewLowfiFreelist34[T any](factory func(uint) T) *LowfiFreelist34[T] {
	freelist := new(LowfiFreelist34[T])
	for index := range help.L34 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM34] |= ^help.SBM34
	return freelist
}

func InitLowfiFreelist34[T any](freelist *LowfiFreelist34[T], factory func(uint) T) {
	for index := range help.L34 {
		freelist.element[index] = factory(index)
	}

	freelist.summary[help.SRM34] |= ^help.SBM34
}

func (this *LowfiFreelist34[T]) Get() (*T, bool) {
	var (
		sb = &this.summary[help.ZeroInt]
		bb = &this.bitmap[help.ZeroInt]
		wi = this.cursor
		w  = *help.GetPtr(bb, int(wi))
	)

	var b uint
	if w == help.MaxUint {
		swi := wi >> help.WordShift
		b = ^(*help.GetPtr(sb, int(swi))) & (help.MaxUint << ((wi & help.WordRMask) + help.OneUint))

		if b == help.ZeroUint {
			for swi := swi + help.OneUint; swi < help.SL34; swi++ {
				b = ^(*help.GetPtr(sb, int(swi)))
				if b != help.ZeroUint {
					wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
					w = *help.GetPtr(bb, int(wi))
					goto mark
				}
			}

			return nil, false
		}

		wi = (swi << help.WordShift) + uint(bits.TrailingZeros(b))
		w = *help.GetPtr(bb, int(wi))
	}

mark:
	b = ^w & (w + help.OneUint)
	w |= b
	*help.GetPtr(bb, int(wi)) = w
	this.cursor = wi

	if w == help.MaxUint {
		*help.GetPtr(sb, int(wi>>help.WordShift)) |= help.OneUint << (wi & help.WordRMask)
		this.cursor = (wi + help.OneUint) & help.WRM34
	}

	return help.GetPtr(&this.element[help.ZeroInt], int((wi<<help.WordShift)|uint(bits.TrailingZeros(b)))), true
}

func (this *LowfiFreelist34[T]) Release(id uint) bool {
	wi := id >> help.WordShift

	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wi)) &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.summary[help.ZeroInt], int(wi>>help.WordShift)) &^= help.OneUint << (wi & help.WordRMask)

	if wi < this.cursor {
		this.cursor = wi
	}

	return true
}

func (this *LowfiFreelist34[T]) Len() int {
	var (
		base   = &this.bitmap[help.ZeroInt]
		length = help.ZeroInt
	)

	for index := range help.WL34 {
		length += bits.OnesCount(*help.GetPtr(base, int(index)))
	}

	return length
}

func (this *LowfiFreelist34[T]) Cap() int { return int(help.L34) }
