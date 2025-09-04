package kit

import (
	"github.com/mattgonewild/kit/internal/help"
)

type LifoFreelist0[T any] struct {
	element [help.L0]T
	stack   [help.L0]uint
	top     int
}

func NewLifoFreelist0[T any](factory func(uint) T) *LifoFreelist0[T] {
	freelist := &LifoFreelist0[T]{
		top: int(help.L0),
	}

	for index := range help.L0 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist0[T any](freelist *LifoFreelist0[T], factory func(uint) T) {
	for index := range help.L0 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L0)
}

func (this *LifoFreelist0[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist0[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist0[T]) Len() int { return int(help.L0) - this.top }
func (this *LifoFreelist0[T]) Cap() int { return int(help.L0) }

type LifoFreelist1[T any] struct {
	element [help.L1]T
	stack   [help.L1]uint
	top     int
}

func NewLifoFreelist1[T any](factory func(uint) T) *LifoFreelist1[T] {
	freelist := &LifoFreelist1[T]{
		top: int(help.L1),
	}

	for index := range help.L1 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist1[T any](freelist *LifoFreelist1[T], factory func(uint) T) {
	for index := range help.L1 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L1)
}

func (this *LifoFreelist1[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist1[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist1[T]) Len() int { return int(help.L1) - this.top }
func (this *LifoFreelist1[T]) Cap() int { return int(help.L1) }

type LifoFreelist2[T any] struct {
	element [help.L2]T
	stack   [help.L2]uint
	top     int
}

func NewLifoFreelist2[T any](factory func(uint) T) *LifoFreelist2[T] {
	freelist := &LifoFreelist2[T]{
		top: int(help.L2),
	}

	for index := range help.L2 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist2[T any](freelist *LifoFreelist2[T], factory func(uint) T) {
	for index := range help.L2 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L2)
}

func (this *LifoFreelist2[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist2[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist2[T]) Len() int { return int(help.L2) - this.top }
func (this *LifoFreelist2[T]) Cap() int { return int(help.L2) }

type LifoFreelist3[T any] struct {
	element [help.L3]T
	stack   [help.L3]uint
	top     int
}

func NewLifoFreelist3[T any](factory func(uint) T) *LifoFreelist3[T] {
	freelist := &LifoFreelist3[T]{
		top: int(help.L3),
	}

	for index := range help.L3 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist3[T any](freelist *LifoFreelist3[T], factory func(uint) T) {
	for index := range help.L3 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L3)
}

func (this *LifoFreelist3[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist3[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist3[T]) Len() int { return int(help.L3) - this.top }
func (this *LifoFreelist3[T]) Cap() int { return int(help.L3) }

type LifoFreelist4[T any] struct {
	element [help.L4]T
	stack   [help.L4]uint
	top     int
}

func NewLifoFreelist4[T any](factory func(uint) T) *LifoFreelist4[T] {
	freelist := &LifoFreelist4[T]{
		top: int(help.L4),
	}

	for index := range help.L4 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist4[T any](freelist *LifoFreelist4[T], factory func(uint) T) {
	for index := range help.L4 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L4)
}

func (this *LifoFreelist4[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist4[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist4[T]) Len() int { return int(help.L4) - this.top }
func (this *LifoFreelist4[T]) Cap() int { return int(help.L4) }

type LifoFreelist5[T any] struct {
	element [help.L5]T
	stack   [help.L5]uint
	top     int
}

func NewLifoFreelist5[T any](factory func(uint) T) *LifoFreelist5[T] {
	freelist := &LifoFreelist5[T]{
		top: int(help.L5),
	}

	for index := range help.L5 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist5[T any](freelist *LifoFreelist5[T], factory func(uint) T) {
	for index := range help.L5 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L5)
}

func (this *LifoFreelist5[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist5[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist5[T]) Len() int { return int(help.L5) - this.top }
func (this *LifoFreelist5[T]) Cap() int { return int(help.L5) }

type LifoFreelist6[T any] struct {
	element [help.L6]T
	stack   [help.L6]uint
	top     int
}

func NewLifoFreelist6[T any](factory func(uint) T) *LifoFreelist6[T] {
	freelist := &LifoFreelist6[T]{
		top: int(help.L6),
	}

	for index := range help.L6 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist6[T any](freelist *LifoFreelist6[T], factory func(uint) T) {
	for index := range help.L6 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L6)
}

func (this *LifoFreelist6[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist6[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist6[T]) Len() int { return int(help.L6) - this.top }
func (this *LifoFreelist6[T]) Cap() int { return int(help.L6) }

type LifoFreelist7[T any] struct {
	element [help.L7]T
	stack   [help.L7]uint
	top     int
}

func NewLifoFreelist7[T any](factory func(uint) T) *LifoFreelist7[T] {
	freelist := &LifoFreelist7[T]{
		top: int(help.L7),
	}

	for index := range help.L7 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist7[T any](freelist *LifoFreelist7[T], factory func(uint) T) {
	for index := range help.L7 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L7)
}

func (this *LifoFreelist7[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist7[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist7[T]) Len() int { return int(help.L7) - this.top }
func (this *LifoFreelist7[T]) Cap() int { return int(help.L7) }

type LifoFreelist8[T any] struct {
	element [help.L8]T
	stack   [help.L8]uint
	top     int
}

func NewLifoFreelist8[T any](factory func(uint) T) *LifoFreelist8[T] {
	freelist := &LifoFreelist8[T]{
		top: int(help.L8),
	}

	for index := range help.L8 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist8[T any](freelist *LifoFreelist8[T], factory func(uint) T) {
	for index := range help.L8 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L8)
}

func (this *LifoFreelist8[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist8[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist8[T]) Len() int { return int(help.L8) - this.top }
func (this *LifoFreelist8[T]) Cap() int { return int(help.L8) }

type LifoFreelist9[T any] struct {
	element [help.L9]T
	stack   [help.L9]uint
	top     int
}

func NewLifoFreelist9[T any](factory func(uint) T) *LifoFreelist9[T] {
	freelist := &LifoFreelist9[T]{
		top: int(help.L9),
	}

	for index := range help.L9 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist9[T any](freelist *LifoFreelist9[T], factory func(uint) T) {
	for index := range help.L9 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L9)
}

func (this *LifoFreelist9[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist9[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist9[T]) Len() int { return int(help.L9) - this.top }
func (this *LifoFreelist9[T]) Cap() int { return int(help.L9) }

type LifoFreelist10[T any] struct {
	element [help.L10]T
	stack   [help.L10]uint
	top     int
}

func NewLifoFreelist10[T any](factory func(uint) T) *LifoFreelist10[T] {
	freelist := &LifoFreelist10[T]{
		top: int(help.L10),
	}

	for index := range help.L10 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist10[T any](freelist *LifoFreelist10[T], factory func(uint) T) {
	for index := range help.L10 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L10)
}

func (this *LifoFreelist10[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist10[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist10[T]) Len() int { return int(help.L10) - this.top }
func (this *LifoFreelist10[T]) Cap() int { return int(help.L10) }

type LifoFreelist11[T any] struct {
	element [help.L11]T
	stack   [help.L11]uint
	top     int
}

func NewLifoFreelist11[T any](factory func(uint) T) *LifoFreelist11[T] {
	freelist := &LifoFreelist11[T]{
		top: int(help.L11),
	}

	for index := range help.L11 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist11[T any](freelist *LifoFreelist11[T], factory func(uint) T) {
	for index := range help.L11 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L11)
}

func (this *LifoFreelist11[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist11[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist11[T]) Len() int { return int(help.L11) - this.top }
func (this *LifoFreelist11[T]) Cap() int { return int(help.L11) }

type LifoFreelist12[T any] struct {
	element [help.L12]T
	stack   [help.L12]uint
	top     int
}

func NewLifoFreelist12[T any](factory func(uint) T) *LifoFreelist12[T] {
	freelist := &LifoFreelist12[T]{
		top: int(help.L12),
	}

	for index := range help.L12 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist12[T any](freelist *LifoFreelist12[T], factory func(uint) T) {
	for index := range help.L12 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L12)
}

func (this *LifoFreelist12[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist12[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist12[T]) Len() int { return int(help.L12) - this.top }
func (this *LifoFreelist12[T]) Cap() int { return int(help.L12) }

type LifoFreelist13[T any] struct {
	element [help.L13]T
	stack   [help.L13]uint
	top     int
}

func NewLifoFreelist13[T any](factory func(uint) T) *LifoFreelist13[T] {
	freelist := &LifoFreelist13[T]{
		top: int(help.L13),
	}

	for index := range help.L13 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist13[T any](freelist *LifoFreelist13[T], factory func(uint) T) {
	for index := range help.L13 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L13)
}

func (this *LifoFreelist13[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist13[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist13[T]) Len() int { return int(help.L13) - this.top }
func (this *LifoFreelist13[T]) Cap() int { return int(help.L13) }

type LifoFreelist14[T any] struct {
	element [help.L14]T
	stack   [help.L14]uint
	top     int
}

func NewLifoFreelist14[T any](factory func(uint) T) *LifoFreelist14[T] {
	freelist := &LifoFreelist14[T]{
		top: int(help.L14),
	}

	for index := range help.L14 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist14[T any](freelist *LifoFreelist14[T], factory func(uint) T) {
	for index := range help.L14 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L14)
}

func (this *LifoFreelist14[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist14[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist14[T]) Len() int { return int(help.L14) - this.top }
func (this *LifoFreelist14[T]) Cap() int { return int(help.L14) }

type LifoFreelist15[T any] struct {
	element [help.L15]T
	stack   [help.L15]uint
	top     int
}

func NewLifoFreelist15[T any](factory func(uint) T) *LifoFreelist15[T] {
	freelist := &LifoFreelist15[T]{
		top: int(help.L15),
	}

	for index := range help.L15 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist15[T any](freelist *LifoFreelist15[T], factory func(uint) T) {
	for index := range help.L15 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L15)
}

func (this *LifoFreelist15[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist15[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist15[T]) Len() int { return int(help.L15) - this.top }
func (this *LifoFreelist15[T]) Cap() int { return int(help.L15) }

type LifoFreelist16[T any] struct {
	element [help.L16]T
	stack   [help.L16]uint
	top     int
}

func NewLifoFreelist16[T any](factory func(uint) T) *LifoFreelist16[T] {
	freelist := &LifoFreelist16[T]{
		top: int(help.L16),
	}

	for index := range help.L16 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist16[T any](freelist *LifoFreelist16[T], factory func(uint) T) {
	for index := range help.L16 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L16)
}

func (this *LifoFreelist16[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist16[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist16[T]) Len() int { return int(help.L16) - this.top }
func (this *LifoFreelist16[T]) Cap() int { return int(help.L16) }

type LifoFreelist17[T any] struct {
	element [help.L17]T
	stack   [help.L17]uint
	top     int
}

func NewLifoFreelist17[T any](factory func(uint) T) *LifoFreelist17[T] {
	freelist := &LifoFreelist17[T]{
		top: int(help.L17),
	}

	for index := range help.L17 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist17[T any](freelist *LifoFreelist17[T], factory func(uint) T) {
	for index := range help.L17 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L17)
}

func (this *LifoFreelist17[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist17[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist17[T]) Len() int { return int(help.L17) - this.top }
func (this *LifoFreelist17[T]) Cap() int { return int(help.L17) }

type LifoFreelist18[T any] struct {
	element [help.L18]T
	stack   [help.L18]uint
	top     int
}

func NewLifoFreelist18[T any](factory func(uint) T) *LifoFreelist18[T] {
	freelist := &LifoFreelist18[T]{
		top: int(help.L18),
	}

	for index := range help.L18 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist18[T any](freelist *LifoFreelist18[T], factory func(uint) T) {
	for index := range help.L18 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L18)
}

func (this *LifoFreelist18[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist18[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist18[T]) Len() int { return int(help.L18) - this.top }
func (this *LifoFreelist18[T]) Cap() int { return int(help.L18) }

type LifoFreelist19[T any] struct {
	element [help.L19]T
	stack   [help.L19]uint
	top     int
}

func NewLifoFreelist19[T any](factory func(uint) T) *LifoFreelist19[T] {
	freelist := &LifoFreelist19[T]{
		top: int(help.L19),
	}

	for index := range help.L19 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist19[T any](freelist *LifoFreelist19[T], factory func(uint) T) {
	for index := range help.L19 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L19)
}

func (this *LifoFreelist19[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist19[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist19[T]) Len() int { return int(help.L19) - this.top }
func (this *LifoFreelist19[T]) Cap() int { return int(help.L19) }

type LifoFreelist20[T any] struct {
	element [help.L20]T
	stack   [help.L20]uint
	top     int
}

func NewLifoFreelist20[T any](factory func(uint) T) *LifoFreelist20[T] {
	freelist := &LifoFreelist20[T]{
		top: int(help.L20),
	}

	for index := range help.L20 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist20[T any](freelist *LifoFreelist20[T], factory func(uint) T) {
	for index := range help.L20 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L20)
}

func (this *LifoFreelist20[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist20[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist20[T]) Len() int { return int(help.L20) - this.top }
func (this *LifoFreelist20[T]) Cap() int { return int(help.L20) }

type LifoFreelist21[T any] struct {
	element [help.L21]T
	stack   [help.L21]uint
	top     int
}

func NewLifoFreelist21[T any](factory func(uint) T) *LifoFreelist21[T] {
	freelist := &LifoFreelist21[T]{
		top: int(help.L21),
	}

	for index := range help.L21 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist21[T any](freelist *LifoFreelist21[T], factory func(uint) T) {
	for index := range help.L21 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L21)
}

func (this *LifoFreelist21[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist21[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist21[T]) Len() int { return int(help.L21) - this.top }
func (this *LifoFreelist21[T]) Cap() int { return int(help.L21) }

type LifoFreelist22[T any] struct {
	element [help.L22]T
	stack   [help.L22]uint
	top     int
}

func NewLifoFreelist22[T any](factory func(uint) T) *LifoFreelist22[T] {
	freelist := &LifoFreelist22[T]{
		top: int(help.L22),
	}

	for index := range help.L22 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist22[T any](freelist *LifoFreelist22[T], factory func(uint) T) {
	for index := range help.L22 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L22)
}

func (this *LifoFreelist22[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist22[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist22[T]) Len() int { return int(help.L22) - this.top }
func (this *LifoFreelist22[T]) Cap() int { return int(help.L22) }

type LifoFreelist23[T any] struct {
	element [help.L23]T
	stack   [help.L23]uint
	top     int
}

func NewLifoFreelist23[T any](factory func(uint) T) *LifoFreelist23[T] {
	freelist := &LifoFreelist23[T]{
		top: int(help.L23),
	}

	for index := range help.L23 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist23[T any](freelist *LifoFreelist23[T], factory func(uint) T) {
	for index := range help.L23 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L23)
}

func (this *LifoFreelist23[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist23[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist23[T]) Len() int { return int(help.L23) - this.top }
func (this *LifoFreelist23[T]) Cap() int { return int(help.L23) }

type LifoFreelist24[T any] struct {
	element [help.L24]T
	stack   [help.L24]uint
	top     int
}

func NewLifoFreelist24[T any](factory func(uint) T) *LifoFreelist24[T] {
	freelist := &LifoFreelist24[T]{
		top: int(help.L24),
	}

	for index := range help.L24 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist24[T any](freelist *LifoFreelist24[T], factory func(uint) T) {
	for index := range help.L24 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L24)
}

func (this *LifoFreelist24[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist24[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist24[T]) Len() int { return int(help.L24) - this.top }
func (this *LifoFreelist24[T]) Cap() int { return int(help.L24) }

type LifoFreelist25[T any] struct {
	element [help.L25]T
	stack   [help.L25]uint
	top     int
}

func NewLifoFreelist25[T any](factory func(uint) T) *LifoFreelist25[T] {
	freelist := &LifoFreelist25[T]{
		top: int(help.L25),
	}

	for index := range help.L25 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist25[T any](freelist *LifoFreelist25[T], factory func(uint) T) {
	for index := range help.L25 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L25)
}

func (this *LifoFreelist25[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist25[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist25[T]) Len() int { return int(help.L25) - this.top }
func (this *LifoFreelist25[T]) Cap() int { return int(help.L25) }

type LifoFreelist26[T any] struct {
	element [help.L26]T
	stack   [help.L26]uint
	top     int
}

func NewLifoFreelist26[T any](factory func(uint) T) *LifoFreelist26[T] {
	freelist := &LifoFreelist26[T]{
		top: int(help.L26),
	}

	for index := range help.L26 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist26[T any](freelist *LifoFreelist26[T], factory func(uint) T) {
	for index := range help.L26 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L26)
}

func (this *LifoFreelist26[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist26[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist26[T]) Len() int { return int(help.L26) - this.top }
func (this *LifoFreelist26[T]) Cap() int { return int(help.L26) }

type LifoFreelist27[T any] struct {
	element [help.L27]T
	stack   [help.L27]uint
	top     int
}

func NewLifoFreelist27[T any](factory func(uint) T) *LifoFreelist27[T] {
	freelist := &LifoFreelist27[T]{
		top: int(help.L27),
	}

	for index := range help.L27 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist27[T any](freelist *LifoFreelist27[T], factory func(uint) T) {
	for index := range help.L27 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L27)
}

func (this *LifoFreelist27[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist27[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist27[T]) Len() int { return int(help.L27) - this.top }
func (this *LifoFreelist27[T]) Cap() int { return int(help.L27) }

type LifoFreelist28[T any] struct {
	element [help.L28]T
	stack   [help.L28]uint
	top     int
}

func NewLifoFreelist28[T any](factory func(uint) T) *LifoFreelist28[T] {
	freelist := &LifoFreelist28[T]{
		top: int(help.L28),
	}

	for index := range help.L28 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist28[T any](freelist *LifoFreelist28[T], factory func(uint) T) {
	for index := range help.L28 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L28)
}

func (this *LifoFreelist28[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist28[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist28[T]) Len() int { return int(help.L28) - this.top }
func (this *LifoFreelist28[T]) Cap() int { return int(help.L28) }

type LifoFreelist29[T any] struct {
	element [help.L29]T
	stack   [help.L29]uint
	top     int
}

func NewLifoFreelist29[T any](factory func(uint) T) *LifoFreelist29[T] {
	freelist := &LifoFreelist29[T]{
		top: int(help.L29),
	}

	for index := range help.L29 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist29[T any](freelist *LifoFreelist29[T], factory func(uint) T) {
	for index := range help.L29 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L29)
}

func (this *LifoFreelist29[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist29[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist29[T]) Len() int { return int(help.L29) - this.top }
func (this *LifoFreelist29[T]) Cap() int { return int(help.L29) }

type LifoFreelist30[T any] struct {
	element [help.L30]T
	stack   [help.L30]uint
	top     int
}

func NewLifoFreelist30[T any](factory func(uint) T) *LifoFreelist30[T] {
	freelist := &LifoFreelist30[T]{
		top: int(help.L30),
	}

	for index := range help.L30 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist30[T any](freelist *LifoFreelist30[T], factory func(uint) T) {
	for index := range help.L30 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L30)
}

func (this *LifoFreelist30[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist30[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist30[T]) Len() int { return int(help.L30) - this.top }
func (this *LifoFreelist30[T]) Cap() int { return int(help.L30) }

type LifoFreelist31[T any] struct {
	element [help.L31]T
	stack   [help.L31]uint
	top     int
}

func NewLifoFreelist31[T any](factory func(uint) T) *LifoFreelist31[T] {
	freelist := &LifoFreelist31[T]{
		top: int(help.L31),
	}

	for index := range help.L31 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist31[T any](freelist *LifoFreelist31[T], factory func(uint) T) {
	for index := range help.L31 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L31)
}

func (this *LifoFreelist31[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist31[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist31[T]) Len() int { return int(help.L31) - this.top }
func (this *LifoFreelist31[T]) Cap() int { return int(help.L31) }

type LifoFreelist32[T any] struct {
	element [help.L32]T
	stack   [help.L32]uint
	top     int
}

func NewLifoFreelist32[T any](factory func(uint) T) *LifoFreelist32[T] {
	freelist := &LifoFreelist32[T]{
		top: int(help.L32),
	}

	for index := range help.L32 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist32[T any](freelist *LifoFreelist32[T], factory func(uint) T) {
	for index := range help.L32 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L32)
}

func (this *LifoFreelist32[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist32[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist32[T]) Len() int { return int(help.L32) - this.top }
func (this *LifoFreelist32[T]) Cap() int { return int(help.L32) }

type LifoFreelist33[T any] struct {
	element [help.L33]T
	stack   [help.L33]uint
	top     int
}

func NewLifoFreelist33[T any](factory func(uint) T) *LifoFreelist33[T] {
	freelist := &LifoFreelist33[T]{
		top: int(help.L33),
	}

	for index := range help.L33 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist33[T any](freelist *LifoFreelist33[T], factory func(uint) T) {
	for index := range help.L33 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L33)
}

func (this *LifoFreelist33[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist33[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist33[T]) Len() int { return int(help.L33) - this.top }
func (this *LifoFreelist33[T]) Cap() int { return int(help.L33) }

type LifoFreelist34[T any] struct {
	element [help.L34]T
	stack   [help.L34]uint
	top     int
}

func NewLifoFreelist34[T any](factory func(uint) T) *LifoFreelist34[T] {
	freelist := &LifoFreelist34[T]{
		top: int(help.L34),
	}

	for index := range help.L34 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	return freelist
}

func InitLifoFreelist34[T any](freelist *LifoFreelist34[T], factory func(uint) T) {
	for index := range help.L34 {
		freelist.element[index] = factory(index)
		freelist.stack[index] = index
	}

	freelist.top = int(help.L34)
}

func (this *LifoFreelist34[T]) Get() (*T, bool) {
	if this.top == help.ZeroInt {
		return nil, false
	}
	this.top--

	return help.GetPtr(&this.element[help.ZeroInt], int(*help.GetPtr(&this.stack[help.ZeroInt], this.top))), true
}

func (this *LifoFreelist34[T]) Release(id uint) bool {
	*help.GetPtr(&this.stack[help.ZeroInt], this.top) = id

	this.top++
	return true
}

func (this *LifoFreelist34[T]) Len() int { return int(help.L34) - this.top }
func (this *LifoFreelist34[T]) Cap() int { return int(help.L34) }
