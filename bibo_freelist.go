package kit

import (
	"math/bits"

	"github.com/mattgonewild/kit/internal/help"
)

// TODO: complete me

type BiboFreelist7[T any] struct {
	queue   [help.WL7]uint
	head    uint
	length  uint
	bitmap  [help.WL7]uint
	element [help.L7]T
}

func NewBiboFreelist7[T any](factory func(uint) T) *BiboFreelist7[T] {
	freelist := new(BiboFreelist7[T])

	for index := range help.L7 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL7 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL7
	return freelist
}

func InitBiboFreelist7[T any](freelist *BiboFreelist7[T], factory func(uint) T) {
	for index := range help.L7 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL7 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL7
}

func (this *BiboFreelist7[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM7
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist7[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM7
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist7[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL7 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist7[T]) Cap() int { return int(help.L7) }

type BiboFreelist8[T any] struct {
	queue   [help.WL8]uint
	head    uint
	length  uint
	bitmap  [help.WL8]uint
	element [help.L8]T
}

func NewBiboFreelist8[T any](factory func(uint) T) *BiboFreelist8[T] {
	freelist := new(BiboFreelist8[T])

	for index := range help.L8 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL8 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL8
	return freelist
}

func InitBiboFreelist8[T any](freelist *BiboFreelist8[T], factory func(uint) T) {
	for index := range help.L8 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL8 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL8
}

func (this *BiboFreelist8[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM8
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist8[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM8
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist8[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL8 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist8[T]) Cap() int { return int(help.L8) }

type BiboFreelist9[T any] struct {
	queue   [help.WL9]uint
	head    uint
	length  uint
	bitmap  [help.WL9]uint
	element [help.L9]T
}

func NewBiboFreelist9[T any](factory func(uint) T) *BiboFreelist9[T] {
	freelist := new(BiboFreelist9[T])

	for index := range help.L9 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL9 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL9
	return freelist
}

func InitBiboFreelist9[T any](freelist *BiboFreelist9[T], factory func(uint) T) {
	for index := range help.L9 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL9 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL9
}

func (this *BiboFreelist9[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM9
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist9[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM9
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist9[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL9 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist9[T]) Cap() int { return int(help.L9) }

type BiboFreelist10[T any] struct {
	queue   [help.WL10]uint
	head    uint
	length  uint
	bitmap  [help.WL10]uint
	element [help.L10]T
}

func NewBiboFreelist10[T any](factory func(uint) T) *BiboFreelist10[T] {
	freelist := new(BiboFreelist10[T])

	for index := range help.L10 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL10 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL10
	return freelist
}

func InitBiboFreelist10[T any](freelist *BiboFreelist10[T], factory func(uint) T) {
	for index := range help.L10 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL10 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL10
}

func (this *BiboFreelist10[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM10
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist10[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM10
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist10[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL10 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist10[T]) Cap() int { return int(help.L10) }

type BiboFreelist11[T any] struct {
	queue   [help.WL11]uint
	head    uint
	length  uint
	bitmap  [help.WL11]uint
	element [help.L11]T
}

func NewBiboFreelist11[T any](factory func(uint) T) *BiboFreelist11[T] {
	freelist := new(BiboFreelist11[T])

	for index := range help.L11 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL11 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL11
	return freelist
}

func InitBiboFreelist11[T any](freelist *BiboFreelist11[T], factory func(uint) T) {
	for index := range help.L11 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL11 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL11
}

func (this *BiboFreelist11[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM11
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist11[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM11
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist11[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL11 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist11[T]) Cap() int { return int(help.L11) }

type BiboFreelist12[T any] struct {
	queue   [help.WL12]uint
	head    uint
	length  uint
	bitmap  [help.WL12]uint
	element [help.L12]T
}

func NewBiboFreelist12[T any](factory func(uint) T) *BiboFreelist12[T] {
	freelist := new(BiboFreelist12[T])

	for index := range help.L12 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL12 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL12
	return freelist
}

func InitBiboFreelist12[T any](freelist *BiboFreelist12[T], factory func(uint) T) {
	for index := range help.L12 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL12 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL12
}

func (this *BiboFreelist12[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM12
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist12[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM12
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist12[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL12 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist12[T]) Cap() int { return int(help.L12) }

type BiboFreelist13[T any] struct {
	queue   [help.WL13]uint
	head    uint
	length  uint
	bitmap  [help.WL13]uint
	element [help.L13]T
}

func NewBiboFreelist13[T any](factory func(uint) T) *BiboFreelist13[T] {
	freelist := new(BiboFreelist13[T])

	for index := range help.L13 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL13 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL13
	return freelist
}

func InitBiboFreelist13[T any](freelist *BiboFreelist13[T], factory func(uint) T) {
	for index := range help.L13 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL13 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL13
}

func (this *BiboFreelist13[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM13
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist13[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM13
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist13[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL13 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist13[T]) Cap() int { return int(help.L13) }

type BiboFreelist14[T any] struct {
	queue   [help.WL14]uint
	head    uint
	length  uint
	bitmap  [help.WL14]uint
	element [help.L14]T
}

func NewBiboFreelist14[T any](factory func(uint) T) *BiboFreelist14[T] {
	freelist := new(BiboFreelist14[T])

	for index := range help.L14 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL14 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL14
	return freelist
}

func InitBiboFreelist14[T any](freelist *BiboFreelist14[T], factory func(uint) T) {
	for index := range help.L14 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL14 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL14
}

func (this *BiboFreelist14[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM14
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist14[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM14
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist14[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL14 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist14[T]) Cap() int { return int(help.L14) }

type BiboFreelist15[T any] struct {
	queue   [help.WL15]uint
	head    uint
	length  uint
	bitmap  [help.WL15]uint
	element [help.L15]T
}

func NewBiboFreelist15[T any](factory func(uint) T) *BiboFreelist15[T] {
	freelist := new(BiboFreelist15[T])

	for index := range help.L15 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL15 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL15
	return freelist
}

func InitBiboFreelist15[T any](freelist *BiboFreelist15[T], factory func(uint) T) {
	for index := range help.L15 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL15 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL15
}

func (this *BiboFreelist15[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM15
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist15[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM15
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist15[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL15 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist15[T]) Cap() int { return int(help.L15) }

type BiboFreelist16[T any] struct {
	queue   [help.WL16]uint
	head    uint
	length  uint
	bitmap  [help.WL16]uint
	element [help.L16]T
}

func NewBiboFreelist16[T any](factory func(uint) T) *BiboFreelist16[T] {
	freelist := new(BiboFreelist16[T])

	for index := range help.L16 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL16 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL16
	return freelist
}

func InitBiboFreelist16[T any](freelist *BiboFreelist16[T], factory func(uint) T) {
	for index := range help.L16 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL16 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL16
}

func (this *BiboFreelist16[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM16
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist16[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM16
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist16[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL16 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist16[T]) Cap() int { return int(help.L16) }

type BiboFreelist17[T any] struct {
	queue   [help.WL17]uint
	head    uint
	length  uint
	bitmap  [help.WL17]uint
	element [help.L17]T
}

func NewBiboFreelist17[T any](factory func(uint) T) *BiboFreelist17[T] {
	freelist := new(BiboFreelist17[T])

	for index := range help.L17 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL17 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL17
	return freelist
}

func InitBiboFreelist17[T any](freelist *BiboFreelist17[T], factory func(uint) T) {
	for index := range help.L17 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL17 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL17
}

func (this *BiboFreelist17[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM17
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist17[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM17
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist17[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL17 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist17[T]) Cap() int { return int(help.L17) }

type BiboFreelist18[T any] struct {
	queue   [help.WL18]uint
	head    uint
	length  uint
	bitmap  [help.WL18]uint
	element [help.L18]T
}

func NewBiboFreelist18[T any](factory func(uint) T) *BiboFreelist18[T] {
	freelist := new(BiboFreelist18[T])

	for index := range help.L18 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL18 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL18
	return freelist
}

func InitBiboFreelist18[T any](freelist *BiboFreelist18[T], factory func(uint) T) {
	for index := range help.L18 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL18 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL18
}

func (this *BiboFreelist18[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM18
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist18[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM18
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist18[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL18 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist18[T]) Cap() int { return int(help.L18) }

type BiboFreelist19[T any] struct {
	queue   [help.WL19]uint
	head    uint
	length  uint
	bitmap  [help.WL19]uint
	element [help.L19]T
}

func NewBiboFreelist19[T any](factory func(uint) T) *BiboFreelist19[T] {
	freelist := new(BiboFreelist19[T])

	for index := range help.L19 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL19 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL19
	return freelist
}

func InitBiboFreelist19[T any](freelist *BiboFreelist19[T], factory func(uint) T) {
	for index := range help.L19 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL19 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL19
}

func (this *BiboFreelist19[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM19
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist19[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM19
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist19[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL19 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist19[T]) Cap() int { return int(help.L19) }

type BiboFreelist20[T any] struct {
	queue   [help.WL20]uint
	head    uint
	length  uint
	bitmap  [help.WL20]uint
	element [help.L20]T
}

func NewBiboFreelist20[T any](factory func(uint) T) *BiboFreelist20[T] {
	freelist := new(BiboFreelist20[T])

	for index := range help.L20 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL20 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL20
	return freelist
}

func InitBiboFreelist20[T any](freelist *BiboFreelist20[T], factory func(uint) T) {
	for index := range help.L20 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL20 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL20
}

func (this *BiboFreelist20[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM20
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist20[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM20
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist20[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL20 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist20[T]) Cap() int { return int(help.L20) }

type BiboFreelist21[T any] struct {
	queue   [help.WL21]uint
	head    uint
	length  uint
	bitmap  [help.WL21]uint
	element [help.L21]T
}

func NewBiboFreelist21[T any](factory func(uint) T) *BiboFreelist21[T] {
	freelist := new(BiboFreelist21[T])

	for index := range help.L21 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL21 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL21
	return freelist
}

func InitBiboFreelist21[T any](freelist *BiboFreelist21[T], factory func(uint) T) {
	for index := range help.L21 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL21 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL21
}

func (this *BiboFreelist21[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM21
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist21[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM21
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist21[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL21 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist21[T]) Cap() int { return int(help.L21) }

type BiboFreelist22[T any] struct {
	queue   [help.WL22]uint
	head    uint
	length  uint
	bitmap  [help.WL22]uint
	element [help.L22]T
}

func NewBiboFreelist22[T any](factory func(uint) T) *BiboFreelist22[T] {
	freelist := new(BiboFreelist22[T])

	for index := range help.L22 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL22 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL22
	return freelist
}

func InitBiboFreelist22[T any](freelist *BiboFreelist22[T], factory func(uint) T) {
	for index := range help.L22 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL22 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL22
}

func (this *BiboFreelist22[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM22
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist22[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM22
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist22[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL22 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist22[T]) Cap() int { return int(help.L22) }

type BiboFreelist23[T any] struct {
	queue   [help.WL23]uint
	head    uint
	length  uint
	bitmap  [help.WL23]uint
	element [help.L23]T
}

func NewBiboFreelist23[T any](factory func(uint) T) *BiboFreelist23[T] {
	freelist := new(BiboFreelist23[T])

	for index := range help.L23 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL23 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL23
	return freelist
}

func InitBiboFreelist23[T any](freelist *BiboFreelist23[T], factory func(uint) T) {
	for index := range help.L23 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL23 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL23
}

func (this *BiboFreelist23[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM23
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist23[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM23
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist23[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL23 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist23[T]) Cap() int { return int(help.L23) }

type BiboFreelist24[T any] struct {
	queue   [help.WL24]uint
	head    uint
	length  uint
	bitmap  [help.WL24]uint
	element [help.L24]T
}

func NewBiboFreelist24[T any](factory func(uint) T) *BiboFreelist24[T] {
	freelist := new(BiboFreelist24[T])

	for index := range help.L24 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL24 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL24
	return freelist
}

func InitBiboFreelist24[T any](freelist *BiboFreelist24[T], factory func(uint) T) {
	for index := range help.L24 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL24 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL24
}

func (this *BiboFreelist24[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM24
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist24[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM24
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist24[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL24 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist24[T]) Cap() int { return int(help.L24) }

type BiboFreelist25[T any] struct {
	queue   [help.WL25]uint
	head    uint
	length  uint
	bitmap  [help.WL25]uint
	element [help.L25]T
}

func NewBiboFreelist25[T any](factory func(uint) T) *BiboFreelist25[T] {
	freelist := new(BiboFreelist25[T])

	for index := range help.L25 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL25 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL25
	return freelist
}

func InitBiboFreelist25[T any](freelist *BiboFreelist25[T], factory func(uint) T) {
	for index := range help.L25 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL25 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL25
}

func (this *BiboFreelist25[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM25
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist25[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM25
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist25[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL25 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist25[T]) Cap() int { return int(help.L25) }

type BiboFreelist26[T any] struct {
	queue   [help.WL26]uint
	head    uint
	length  uint
	bitmap  [help.WL26]uint
	element [help.L26]T
}

func NewBiboFreelist26[T any](factory func(uint) T) *BiboFreelist26[T] {
	freelist := new(BiboFreelist26[T])

	for index := range help.L26 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL26 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL26
	return freelist
}

func InitBiboFreelist26[T any](freelist *BiboFreelist26[T], factory func(uint) T) {
	for index := range help.L26 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL26 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL26
}

func (this *BiboFreelist26[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM26
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist26[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM26
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist26[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL26 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist26[T]) Cap() int { return int(help.L26) }

type BiboFreelist27[T any] struct {
	queue   [help.WL27]uint
	head    uint
	length  uint
	bitmap  [help.WL27]uint
	element [help.L27]T
}

func NewBiboFreelist27[T any](factory func(uint) T) *BiboFreelist27[T] {
	freelist := new(BiboFreelist27[T])

	for index := range help.L27 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL27 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL27
	return freelist
}

func InitBiboFreelist27[T any](freelist *BiboFreelist27[T], factory func(uint) T) {
	for index := range help.L27 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL27 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL27
}

func (this *BiboFreelist27[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM27
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist27[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM27
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist27[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL27 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist27[T]) Cap() int { return int(help.L27) }

type BiboFreelist28[T any] struct {
	queue   [help.WL28]uint
	head    uint
	length  uint
	bitmap  [help.WL28]uint
	element [help.L28]T
}

func NewBiboFreelist28[T any](factory func(uint) T) *BiboFreelist28[T] {
	freelist := new(BiboFreelist28[T])

	for index := range help.L28 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL28 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL28
	return freelist
}

func InitBiboFreelist28[T any](freelist *BiboFreelist28[T], factory func(uint) T) {
	for index := range help.L28 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL28 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL28
}

func (this *BiboFreelist28[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM28
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist28[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM28
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist28[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL28 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist28[T]) Cap() int { return int(help.L28) }

type BiboFreelist29[T any] struct {
	queue   [help.WL29]uint
	head    uint
	length  uint
	bitmap  [help.WL29]uint
	element [help.L29]T
}

func NewBiboFreelist29[T any](factory func(uint) T) *BiboFreelist29[T] {
	freelist := new(BiboFreelist29[T])

	for index := range help.L29 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL29 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL29
	return freelist
}

func InitBiboFreelist29[T any](freelist *BiboFreelist29[T], factory func(uint) T) {
	for index := range help.L29 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL29 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL29
}

func (this *BiboFreelist29[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM29
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist29[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM29
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist29[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL29 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist29[T]) Cap() int { return int(help.L29) }

type BiboFreelist30[T any] struct {
	queue   [help.WL30]uint
	head    uint
	length  uint
	bitmap  [help.WL30]uint
	element [help.L30]T
}

func NewBiboFreelist30[T any](factory func(uint) T) *BiboFreelist30[T] {
	freelist := new(BiboFreelist30[T])

	for index := range help.L30 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL30 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL30
	return freelist
}

func InitBiboFreelist30[T any](freelist *BiboFreelist30[T], factory func(uint) T) {
	for index := range help.L30 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL30 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL30
}

func (this *BiboFreelist30[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM30
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist30[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM30
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist30[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL30 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist30[T]) Cap() int { return int(help.L30) }

type BiboFreelist31[T any] struct {
	queue   [help.WL31]uint
	head    uint
	length  uint
	bitmap  [help.WL31]uint
	element [help.L31]T
}

func NewBiboFreelist31[T any](factory func(uint) T) *BiboFreelist31[T] {
	freelist := new(BiboFreelist31[T])

	for index := range help.L31 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL31 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL31
	return freelist
}

func InitBiboFreelist31[T any](freelist *BiboFreelist31[T], factory func(uint) T) {
	for index := range help.L31 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL31 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL31
}

func (this *BiboFreelist31[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM31
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist31[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM31
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist31[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL31 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist31[T]) Cap() int { return int(help.L31) }

type BiboFreelist32[T any] struct {
	queue   [help.WL32]uint
	head    uint
	length  uint
	bitmap  [help.WL32]uint
	element [help.L32]T
}

func NewBiboFreelist32[T any](factory func(uint) T) *BiboFreelist32[T] {
	freelist := new(BiboFreelist32[T])

	for index := range help.L32 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL32 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL32
	return freelist
}

func InitBiboFreelist32[T any](freelist *BiboFreelist32[T], factory func(uint) T) {
	for index := range help.L32 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL32 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL32
}

func (this *BiboFreelist32[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM32
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist32[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM32
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist32[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL32 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist32[T]) Cap() int { return int(help.L32) }

type BiboFreelist33[T any] struct {
	queue   [help.WL33]uint
	head    uint
	length  uint
	bitmap  [help.WL33]uint
	element [help.L33]T
}

func NewBiboFreelist33[T any](factory func(uint) T) *BiboFreelist33[T] {
	freelist := new(BiboFreelist33[T])

	for index := range help.L33 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL33 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL33
	return freelist
}

func InitBiboFreelist33[T any](freelist *BiboFreelist33[T], factory func(uint) T) {
	for index := range help.L33 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL33 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL33
}

func (this *BiboFreelist33[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM33
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist33[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM33
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist33[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL33 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist33[T]) Cap() int { return int(help.L33) }

type BiboFreelist34[T any] struct {
	queue   [help.WL34]uint
	head    uint
	length  uint
	bitmap  [help.WL34]uint
	element [help.L34]T
}

func NewBiboFreelist34[T any](factory func(uint) T) *BiboFreelist34[T] {
	freelist := new(BiboFreelist34[T])

	for index := range help.L34 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL34 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL34
	return freelist
}

func InitBiboFreelist34[T any](freelist *BiboFreelist34[T], factory func(uint) T) {
	for index := range help.L34 {
		freelist.element[index] = factory(index)
	}

	for index := range help.WL34 {
		freelist.queue[index] = index
	}

	freelist.length = help.WL34
}

func (this *BiboFreelist34[T]) Get() (*T, bool) {
	if this.length == help.ZeroUint {
		return nil, false
	}

	var (
		head      = this.head
		wordIndex = *help.GetPtr(&this.queue[help.ZeroInt], int(head))
		word      = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		inverse   = ^word
		low       = inverse & -inverse
	)

	word |= low
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word
	if word == help.MaxUint {
		this.head = (head + help.OneUint) & help.WRM34
		this.length--
	}

	var (
		bitIndex  = uint(bits.TrailingZeros(low))
		elemIndex = (wordIndex << help.WordShift) | bitIndex
	)

	return help.GetPtr(&this.element[help.ZeroInt], int(elemIndex)), true
}

func (this *BiboFreelist34[T]) Release(id uint) bool {
	var (
		wordIndex   = id >> help.WordShift
		word        = *help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex))
		wordWasFull = word == help.MaxUint
	)

	word &^= help.OneUint << (id & help.WordRMask)
	*help.GetPtr(&this.bitmap[help.ZeroInt], int(wordIndex)) = word

	if wordWasFull {
		tail := (this.head + this.length) & help.WRM34
		*help.GetPtr(&this.queue[help.ZeroInt], int(tail)) = wordIndex
		this.length++
	}

	return true
}

func (this *BiboFreelist34[T]) Len() int {
	length := help.ZeroInt
	for index := range help.WL34 {
		length += bits.OnesCount(this.bitmap[index])
	}

	return length
}

func (this *BiboFreelist34[T]) Cap() int { return int(help.L34) }
