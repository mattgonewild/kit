package kit

import "github.com/ringboundio/kit/internal/help"

func NewFifoAs() bool { return true }

type FifoFreelist0[T any] struct {
	element [help.L0]T
	queue   [help.L0]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist0[T any](factory func(uint) T) *FifoFreelist0[T] {
	freelist := new(FifoFreelist0[T])
	for index := range help.L0 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist0[T any](freelist *FifoFreelist0[T], factory func(uint) T) {
	for index := range help.L0 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist0[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM0
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L0 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist0[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM0)) = id

	this.length++
	return true
}

func (this *FifoFreelist0[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist0[T]) Cap() int { return int(help.L0) }

type FifoFreelist1[T any] struct {
	element [help.L1]T
	queue   [help.L1]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist1[T any](factory func(uint) T) *FifoFreelist1[T] {
	freelist := new(FifoFreelist1[T])
	for index := range help.L1 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist1[T any](freelist *FifoFreelist1[T], factory func(uint) T) {
	for index := range help.L1 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist1[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM1
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L1 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist1[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM1)) = id

	this.length++
	return true
}

func (this *FifoFreelist1[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist1[T]) Cap() int { return int(help.L1) }

type FifoFreelist2[T any] struct {
	element [help.L2]T
	queue   [help.L2]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist2[T any](factory func(uint) T) *FifoFreelist2[T] {
	freelist := new(FifoFreelist2[T])
	for index := range help.L2 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist2[T any](freelist *FifoFreelist2[T], factory func(uint) T) {
	for index := range help.L2 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist2[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM2
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L2 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist2[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM2)) = id

	this.length++
	return true
}

func (this *FifoFreelist2[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist2[T]) Cap() int { return int(help.L2) }

type FifoFreelist3[T any] struct {
	element [help.L3]T
	queue   [help.L3]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist3[T any](factory func(uint) T) *FifoFreelist3[T] {
	freelist := new(FifoFreelist3[T])
	for index := range help.L3 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist3[T any](freelist *FifoFreelist3[T], factory func(uint) T) {
	for index := range help.L3 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist3[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM3
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L3 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist3[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM3)) = id

	this.length++
	return true
}

func (this *FifoFreelist3[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist3[T]) Cap() int { return int(help.L3) }

type FifoFreelist4[T any] struct {
	element [help.L4]T
	queue   [help.L4]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist4[T any](factory func(uint) T) *FifoFreelist4[T] {
	freelist := new(FifoFreelist4[T])
	for index := range help.L4 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist4[T any](freelist *FifoFreelist4[T], factory func(uint) T) {
	for index := range help.L4 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist4[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM4
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L4 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist4[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM4)) = id

	this.length++
	return true
}

func (this *FifoFreelist4[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist4[T]) Cap() int { return int(help.L4) }

type FifoFreelist5[T any] struct {
	element [help.L5]T
	queue   [help.L5]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist5[T any](factory func(uint) T) *FifoFreelist5[T] {
	freelist := new(FifoFreelist5[T])
	for index := range help.L5 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist5[T any](freelist *FifoFreelist5[T], factory func(uint) T) {
	for index := range help.L5 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist5[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM5
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L5 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist5[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM5)) = id

	this.length++
	return true
}

func (this *FifoFreelist5[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist5[T]) Cap() int { return int(help.L5) }

type FifoFreelist6[T any] struct {
	element [help.L6]T
	queue   [help.L6]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist6[T any](factory func(uint) T) *FifoFreelist6[T] {
	freelist := new(FifoFreelist6[T])
	for index := range help.L6 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist6[T any](freelist *FifoFreelist6[T], factory func(uint) T) {
	for index := range help.L6 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist6[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM6
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L6 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist6[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM6)) = id

	this.length++
	return true
}

func (this *FifoFreelist6[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist6[T]) Cap() int { return int(help.L6) }

type FifoFreelist7[T any] struct {
	element [help.L7]T
	queue   [help.L7]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist7[T any](factory func(uint) T) *FifoFreelist7[T] {
	freelist := new(FifoFreelist7[T])
	for index := range help.L7 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist7[T any](freelist *FifoFreelist7[T], factory func(uint) T) {
	for index := range help.L7 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist7[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM7
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L7 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist7[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM7)) = id

	this.length++
	return true
}

func (this *FifoFreelist7[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist7[T]) Cap() int { return int(help.L7) }

type FifoFreelist8[T any] struct {
	element [help.L8]T
	queue   [help.L8]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist8[T any](factory func(uint) T) *FifoFreelist8[T] {
	freelist := new(FifoFreelist8[T])
	for index := range help.L8 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist8[T any](freelist *FifoFreelist8[T], factory func(uint) T) {
	for index := range help.L8 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist8[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM8
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L8 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist8[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM8)) = id

	this.length++
	return true
}

func (this *FifoFreelist8[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist8[T]) Cap() int { return int(help.L8) }

type FifoFreelist9[T any] struct {
	element [help.L9]T
	queue   [help.L9]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist9[T any](factory func(uint) T) *FifoFreelist9[T] {
	freelist := new(FifoFreelist9[T])
	for index := range help.L9 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist9[T any](freelist *FifoFreelist9[T], factory func(uint) T) {
	for index := range help.L9 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist9[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM9
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L9 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist9[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM9)) = id

	this.length++
	return true
}

func (this *FifoFreelist9[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist9[T]) Cap() int { return int(help.L9) }

type FifoFreelist10[T any] struct {
	element [help.L10]T
	queue   [help.L10]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist10[T any](factory func(uint) T) *FifoFreelist10[T] {
	freelist := new(FifoFreelist10[T])
	for index := range help.L10 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist10[T any](freelist *FifoFreelist10[T], factory func(uint) T) {
	for index := range help.L10 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist10[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM10
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L10 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist10[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM10)) = id

	this.length++
	return true
}

func (this *FifoFreelist10[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist10[T]) Cap() int { return int(help.L10) }

type FifoFreelist11[T any] struct {
	element [help.L11]T
	queue   [help.L11]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist11[T any](factory func(uint) T) *FifoFreelist11[T] {
	freelist := new(FifoFreelist11[T])
	for index := range help.L11 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist11[T any](freelist *FifoFreelist11[T], factory func(uint) T) {
	for index := range help.L11 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist11[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM11
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L11 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist11[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM11)) = id

	this.length++
	return true
}

func (this *FifoFreelist11[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist11[T]) Cap() int { return int(help.L11) }

type FifoFreelist12[T any] struct {
	element [help.L12]T
	queue   [help.L12]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist12[T any](factory func(uint) T) *FifoFreelist12[T] {
	freelist := new(FifoFreelist12[T])
	for index := range help.L12 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist12[T any](freelist *FifoFreelist12[T], factory func(uint) T) {
	for index := range help.L12 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist12[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM12
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L12 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist12[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM12)) = id

	this.length++
	return true
}

func (this *FifoFreelist12[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist12[T]) Cap() int { return int(help.L12) }

type FifoFreelist13[T any] struct {
	element [help.L13]T
	queue   [help.L13]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist13[T any](factory func(uint) T) *FifoFreelist13[T] {
	freelist := new(FifoFreelist13[T])
	for index := range help.L13 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist13[T any](freelist *FifoFreelist13[T], factory func(uint) T) {
	for index := range help.L13 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist13[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM13
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L13 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist13[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM13)) = id

	this.length++
	return true
}

func (this *FifoFreelist13[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist13[T]) Cap() int { return int(help.L13) }

type FifoFreelist14[T any] struct {
	element [help.L14]T
	queue   [help.L14]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist14[T any](factory func(uint) T) *FifoFreelist14[T] {
	freelist := new(FifoFreelist14[T])
	for index := range help.L14 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist14[T any](freelist *FifoFreelist14[T], factory func(uint) T) {
	for index := range help.L14 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist14[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM14
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L14 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist14[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM14)) = id

	this.length++
	return true
}

func (this *FifoFreelist14[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist14[T]) Cap() int { return int(help.L14) }

type FifoFreelist15[T any] struct {
	element [help.L15]T
	queue   [help.L15]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist15[T any](factory func(uint) T) *FifoFreelist15[T] {
	freelist := new(FifoFreelist15[T])
	for index := range help.L15 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist15[T any](freelist *FifoFreelist15[T], factory func(uint) T) {
	for index := range help.L15 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist15[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM15
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L15 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist15[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM15)) = id

	this.length++
	return true
}

func (this *FifoFreelist15[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist15[T]) Cap() int { return int(help.L15) }

type FifoFreelist16[T any] struct {
	element [help.L16]T
	queue   [help.L16]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist16[T any](factory func(uint) T) *FifoFreelist16[T] {
	freelist := new(FifoFreelist16[T])
	for index := range help.L16 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist16[T any](freelist *FifoFreelist16[T], factory func(uint) T) {
	for index := range help.L16 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist16[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM16
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L16 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist16[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM16)) = id

	this.length++
	return true
}

func (this *FifoFreelist16[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist16[T]) Cap() int { return int(help.L16) }

type FifoFreelist17[T any] struct {
	element [help.L17]T
	queue   [help.L17]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist17[T any](factory func(uint) T) *FifoFreelist17[T] {
	freelist := new(FifoFreelist17[T])
	for index := range help.L17 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist17[T any](freelist *FifoFreelist17[T], factory func(uint) T) {
	for index := range help.L17 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist17[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM17
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L17 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist17[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM17)) = id

	this.length++
	return true
}

func (this *FifoFreelist17[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist17[T]) Cap() int { return int(help.L17) }

type FifoFreelist18[T any] struct {
	element [help.L18]T
	queue   [help.L18]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist18[T any](factory func(uint) T) *FifoFreelist18[T] {
	freelist := new(FifoFreelist18[T])
	for index := range help.L18 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist18[T any](freelist *FifoFreelist18[T], factory func(uint) T) {
	for index := range help.L18 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist18[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM18
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L18 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist18[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM18)) = id

	this.length++
	return true
}

func (this *FifoFreelist18[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist18[T]) Cap() int { return int(help.L18) }

type FifoFreelist19[T any] struct {
	element [help.L19]T
	queue   [help.L19]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist19[T any](factory func(uint) T) *FifoFreelist19[T] {
	freelist := new(FifoFreelist19[T])
	for index := range help.L19 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist19[T any](freelist *FifoFreelist19[T], factory func(uint) T) {
	for index := range help.L19 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist19[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM19
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L19 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist19[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM19)) = id

	this.length++
	return true
}

func (this *FifoFreelist19[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist19[T]) Cap() int { return int(help.L19) }

type FifoFreelist20[T any] struct {
	element [help.L20]T
	queue   [help.L20]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist20[T any](factory func(uint) T) *FifoFreelist20[T] {
	freelist := new(FifoFreelist20[T])
	for index := range help.L20 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist20[T any](freelist *FifoFreelist20[T], factory func(uint) T) {
	for index := range help.L20 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist20[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM20
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L20 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist20[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM20)) = id

	this.length++
	return true
}

func (this *FifoFreelist20[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist20[T]) Cap() int { return int(help.L20) }

type FifoFreelist21[T any] struct {
	element [help.L21]T
	queue   [help.L21]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist21[T any](factory func(uint) T) *FifoFreelist21[T] {
	freelist := new(FifoFreelist21[T])
	for index := range help.L21 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist21[T any](freelist *FifoFreelist21[T], factory func(uint) T) {
	for index := range help.L21 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist21[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM21
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L21 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist21[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM21)) = id

	this.length++
	return true
}

func (this *FifoFreelist21[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist21[T]) Cap() int { return int(help.L21) }

type FifoFreelist22[T any] struct {
	element [help.L22]T
	queue   [help.L22]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist22[T any](factory func(uint) T) *FifoFreelist22[T] {
	freelist := new(FifoFreelist22[T])
	for index := range help.L22 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist22[T any](freelist *FifoFreelist22[T], factory func(uint) T) {
	for index := range help.L22 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist22[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM22
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L22 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist22[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM22)) = id

	this.length++
	return true
}

func (this *FifoFreelist22[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist22[T]) Cap() int { return int(help.L22) }

type FifoFreelist23[T any] struct {
	element [help.L23]T
	queue   [help.L23]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist23[T any](factory func(uint) T) *FifoFreelist23[T] {
	freelist := new(FifoFreelist23[T])
	for index := range help.L23 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist23[T any](freelist *FifoFreelist23[T], factory func(uint) T) {
	for index := range help.L23 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist23[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM23
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L23 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist23[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM23)) = id

	this.length++
	return true
}

func (this *FifoFreelist23[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist23[T]) Cap() int { return int(help.L23) }

type FifoFreelist24[T any] struct {
	element [help.L24]T
	queue   [help.L24]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist24[T any](factory func(uint) T) *FifoFreelist24[T] {
	freelist := new(FifoFreelist24[T])
	for index := range help.L24 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist24[T any](freelist *FifoFreelist24[T], factory func(uint) T) {
	for index := range help.L24 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist24[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM24
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L24 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist24[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM24)) = id

	this.length++
	return true
}

func (this *FifoFreelist24[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist24[T]) Cap() int { return int(help.L24) }

type FifoFreelist25[T any] struct {
	element [help.L25]T
	queue   [help.L25]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist25[T any](factory func(uint) T) *FifoFreelist25[T] {
	freelist := new(FifoFreelist25[T])
	for index := range help.L25 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist25[T any](freelist *FifoFreelist25[T], factory func(uint) T) {
	for index := range help.L25 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist25[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM25
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L25 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist25[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM25)) = id

	this.length++
	return true
}

func (this *FifoFreelist25[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist25[T]) Cap() int { return int(help.L25) }

type FifoFreelist26[T any] struct {
	element [help.L26]T
	queue   [help.L26]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist26[T any](factory func(uint) T) *FifoFreelist26[T] {
	freelist := new(FifoFreelist26[T])
	for index := range help.L26 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist26[T any](freelist *FifoFreelist26[T], factory func(uint) T) {
	for index := range help.L26 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist26[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM26
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L26 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist26[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM26)) = id

	this.length++
	return true
}

func (this *FifoFreelist26[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist26[T]) Cap() int { return int(help.L26) }

type FifoFreelist27[T any] struct {
	element [help.L27]T
	queue   [help.L27]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist27[T any](factory func(uint) T) *FifoFreelist27[T] {
	freelist := new(FifoFreelist27[T])
	for index := range help.L27 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist27[T any](freelist *FifoFreelist27[T], factory func(uint) T) {
	for index := range help.L27 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist27[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM27
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L27 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist27[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM27)) = id

	this.length++
	return true
}

func (this *FifoFreelist27[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist27[T]) Cap() int { return int(help.L27) }

type FifoFreelist28[T any] struct {
	element [help.L28]T
	queue   [help.L28]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist28[T any](factory func(uint) T) *FifoFreelist28[T] {
	freelist := new(FifoFreelist28[T])
	for index := range help.L28 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist28[T any](freelist *FifoFreelist28[T], factory func(uint) T) {
	for index := range help.L28 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist28[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM28
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L28 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist28[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM28)) = id

	this.length++
	return true
}

func (this *FifoFreelist28[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist28[T]) Cap() int { return int(help.L28) }

type FifoFreelist29[T any] struct {
	element [help.L29]T
	queue   [help.L29]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist29[T any](factory func(uint) T) *FifoFreelist29[T] {
	freelist := new(FifoFreelist29[T])
	for index := range help.L29 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist29[T any](freelist *FifoFreelist29[T], factory func(uint) T) {
	for index := range help.L29 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist29[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM29
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L29 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist29[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM29)) = id

	this.length++
	return true
}

func (this *FifoFreelist29[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist29[T]) Cap() int { return int(help.L29) }

type FifoFreelist30[T any] struct {
	element [help.L30]T
	queue   [help.L30]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist30[T any](factory func(uint) T) *FifoFreelist30[T] {
	freelist := new(FifoFreelist30[T])
	for index := range help.L30 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist30[T any](freelist *FifoFreelist30[T], factory func(uint) T) {
	for index := range help.L30 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist30[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM30
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L30 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist30[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM30)) = id

	this.length++
	return true
}

func (this *FifoFreelist30[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist30[T]) Cap() int { return int(help.L30) }

type FifoFreelist31[T any] struct {
	element [help.L31]T
	queue   [help.L31]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist31[T any](factory func(uint) T) *FifoFreelist31[T] {
	freelist := new(FifoFreelist31[T])
	for index := range help.L31 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist31[T any](freelist *FifoFreelist31[T], factory func(uint) T) {
	for index := range help.L31 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist31[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM31
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L31 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist31[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM31)) = id

	this.length++
	return true
}

func (this *FifoFreelist31[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist31[T]) Cap() int { return int(help.L31) }

type FifoFreelist32[T any] struct {
	element [help.L32]T
	queue   [help.L32]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist32[T any](factory func(uint) T) *FifoFreelist32[T] {
	freelist := new(FifoFreelist32[T])
	for index := range help.L32 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist32[T any](freelist *FifoFreelist32[T], factory func(uint) T) {
	for index := range help.L32 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist32[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM32
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L32 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist32[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM32)) = id

	this.length++
	return true
}

func (this *FifoFreelist32[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist32[T]) Cap() int { return int(help.L32) }

type FifoFreelist33[T any] struct {
	element [help.L33]T
	queue   [help.L33]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist33[T any](factory func(uint) T) *FifoFreelist33[T] {
	freelist := new(FifoFreelist33[T])
	for index := range help.L33 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist33[T any](freelist *FifoFreelist33[T], factory func(uint) T) {
	for index := range help.L33 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist33[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM33
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L33 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist33[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM33)) = id

	this.length++
	return true
}

func (this *FifoFreelist33[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist33[T]) Cap() int { return int(help.L33) }

type FifoFreelist34[T any] struct {
	element [help.L34]T
	queue   [help.L34]uint
	head    uint
	length  uint
	next    uint
}

func NewFifoFreelist34[T any](factory func(uint) T) *FifoFreelist34[T] {
	freelist := new(FifoFreelist34[T])
	for index := range help.L34 {
		freelist.element[index] = factory(index)
	}

	return freelist
}

func InitFifoFreelist34[T any](freelist *FifoFreelist34[T], factory func(uint) T) {
	for index := range help.L34 {
		freelist.element[index] = factory(index)
	}
}

func (this *FifoFreelist34[T]) Get() (*T, bool) {
	if this.length > help.ZeroUint {
		head := this.head
		this.head = (head + help.OneUint) & help.RM34
		this.length--

		return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.queue[help.ZeroInt], int(head)))), true
	}

	next := this.next
	if next == help.L34 {
		return nil, false
	}
	this.next = next + help.OneUint

	return help.GetPtr(&this.element[help.ZeroInt], int(next)), true
}

func (this *FifoFreelist34[T]) Release(id uint) bool {
	*help.GetPtr(&this.queue[help.ZeroInt], int((this.head+this.length)&help.RM34)) = id

	this.length++
	return true
}

func (this *FifoFreelist34[T]) Len() int { return int(this.next - this.length) }
func (this *FifoFreelist34[T]) Cap() int { return int(help.L34) }
