package kit

import (
	"github.com/mattgonewild/common"
	"github.com/mattgonewild/kit/internal/help"
)

func NewRingAs() bool { return true }

type Ring1[T common.Comparable[T]] struct {
	element            [help.L1]T
	head, tail, length uint
}

func (this *Ring1[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM1
	this.head = (this.head + (this.length >> help.E1)) & help.RM1
	this.length += help.OneUint - (this.length >> help.E1)
	return true
}

func (this *Ring1[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM1) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring1[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring1[T]) Remove(T) bool { return false }
func (this *Ring1[T]) Len() int      { return int(this.length) }
func (this *Ring1[T]) Cap() int      { return int(help.L1) }

type Ring2[T common.Comparable[T]] struct {
	element            [help.L2]T
	head, tail, length uint
}

func (this *Ring2[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM2
	this.head = (this.head + (this.length >> help.E2)) & help.RM2
	this.length += help.OneUint - (this.length >> help.E2)
	return true
}

func (this *Ring2[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM2) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring2[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring2[T]) Remove(T) bool { return false }
func (this *Ring2[T]) Len() int      { return int(this.length) }
func (this *Ring2[T]) Cap() int      { return int(help.L2) }

type Ring3[T common.Comparable[T]] struct {
	element            [help.L3]T
	head, tail, length uint
}

func (this *Ring3[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM3
	this.head = (this.head + (this.length >> help.E3)) & help.RM3
	this.length += help.OneUint - (this.length >> help.E3)
	return true
}

func (this *Ring3[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM3) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring3[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring3[T]) Remove(T) bool { return false }
func (this *Ring3[T]) Len() int      { return int(this.length) }
func (this *Ring3[T]) Cap() int      { return int(help.L3) }

type Ring4[T common.Comparable[T]] struct {
	element            [help.L4]T
	head, tail, length uint
}

func (this *Ring4[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM4
	this.head = (this.head + (this.length >> help.E4)) & help.RM4
	this.length += help.OneUint - (this.length >> help.E4)
	return true
}

func (this *Ring4[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM4) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring4[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring4[T]) Remove(T) bool { return false }
func (this *Ring4[T]) Len() int      { return int(this.length) }
func (this *Ring4[T]) Cap() int      { return int(help.L4) }

type Ring5[T common.Comparable[T]] struct {
	element            [help.L5]T
	head, tail, length uint
}

func (this *Ring5[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM5
	this.head = (this.head + (this.length >> help.E5)) & help.RM5
	this.length += help.OneUint - (this.length >> help.E5)
	return true
}

func (this *Ring5[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM5) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring5[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring5[T]) Remove(T) bool { return false }
func (this *Ring5[T]) Len() int      { return int(this.length) }
func (this *Ring5[T]) Cap() int      { return int(help.L5) }

type Ring6[T common.Comparable[T]] struct {
	element            [help.L6]T
	head, tail, length uint
}

func (this *Ring6[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM6
	this.head = (this.head + (this.length >> help.E6)) & help.RM6
	this.length += help.OneUint - (this.length >> help.E6)
	return true
}

func (this *Ring6[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM6) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring6[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring6[T]) Remove(T) bool { return false }
func (this *Ring6[T]) Len() int      { return int(this.length) }
func (this *Ring6[T]) Cap() int      { return int(help.L6) }

type Ring7[T common.Comparable[T]] struct {
	element            [help.L7]T
	head, tail, length uint
}

func (this *Ring7[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM7
	this.head = (this.head + (this.length >> help.E7)) & help.RM7
	this.length += help.OneUint - (this.length >> help.E7)
	return true
}

func (this *Ring7[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM7) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring7[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring7[T]) Remove(T) bool { return false }
func (this *Ring7[T]) Len() int      { return int(this.length) }
func (this *Ring7[T]) Cap() int      { return int(help.L7) }

type Ring8[T common.Comparable[T]] struct {
	element            [help.L8]T
	head, tail, length uint
}

func (this *Ring8[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM8
	this.head = (this.head + (this.length >> help.E8)) & help.RM8
	this.length += help.OneUint - (this.length >> help.E8)
	return true
}

func (this *Ring8[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM8) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring8[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring8[T]) Remove(T) bool { return false }
func (this *Ring8[T]) Len() int      { return int(this.length) }
func (this *Ring8[T]) Cap() int      { return int(help.L8) }

type Ring9[T common.Comparable[T]] struct {
	element            [help.L9]T
	head, tail, length uint
}

func (this *Ring9[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM9
	this.head = (this.head + (this.length >> help.E9)) & help.RM9
	this.length += help.OneUint - (this.length >> help.E9)
	return true
}

func (this *Ring9[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM9) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring9[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring9[T]) Remove(T) bool { return false }
func (this *Ring9[T]) Len() int      { return int(this.length) }
func (this *Ring9[T]) Cap() int      { return int(help.L9) }

type Ring10[T common.Comparable[T]] struct {
	element            [help.L10]T
	head, tail, length uint
}

func (this *Ring10[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM10
	this.head = (this.head + (this.length >> help.E10)) & help.RM10
	this.length += help.OneUint - (this.length >> help.E10)
	return true
}

func (this *Ring10[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM10) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring10[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring10[T]) Remove(T) bool { return false }
func (this *Ring10[T]) Len() int      { return int(this.length) }
func (this *Ring10[T]) Cap() int      { return int(help.L10) }

type Ring11[T common.Comparable[T]] struct {
	element            [help.L11]T
	head, tail, length uint
}

func (this *Ring11[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM11
	this.head = (this.head + (this.length >> help.E11)) & help.RM11
	this.length += help.OneUint - (this.length >> help.E11)
	return true
}

func (this *Ring11[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM11) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring11[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring11[T]) Remove(T) bool { return false }
func (this *Ring11[T]) Len() int      { return int(this.length) }
func (this *Ring11[T]) Cap() int      { return int(help.L11) }

type Ring12[T common.Comparable[T]] struct {
	element            [help.L12]T
	head, tail, length uint
}

func (this *Ring12[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM12
	this.head = (this.head + (this.length >> help.E12)) & help.RM12
	this.length += help.OneUint - (this.length >> help.E12)
	return true
}

func (this *Ring12[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM12) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring12[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring12[T]) Remove(T) bool { return false }
func (this *Ring12[T]) Len() int      { return int(this.length) }
func (this *Ring12[T]) Cap() int      { return int(help.L12) }

type Ring13[T common.Comparable[T]] struct {
	element            [help.L13]T
	head, tail, length uint
}

func (this *Ring13[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM13
	this.head = (this.head + (this.length >> help.E13)) & help.RM13
	this.length += help.OneUint - (this.length >> help.E13)
	return true
}

func (this *Ring13[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM13) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring13[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring13[T]) Remove(T) bool { return false }
func (this *Ring13[T]) Len() int      { return int(this.length) }
func (this *Ring13[T]) Cap() int      { return int(help.L13) }

type Ring14[T common.Comparable[T]] struct {
	element            [help.L14]T
	head, tail, length uint
}

func (this *Ring14[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM14
	this.head = (this.head + (this.length >> help.E14)) & help.RM14
	this.length += help.OneUint - (this.length >> help.E14)
	return true
}

func (this *Ring14[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM14) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring14[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring14[T]) Remove(T) bool { return false }
func (this *Ring14[T]) Len() int      { return int(this.length) }
func (this *Ring14[T]) Cap() int      { return int(help.L14) }

type Ring15[T common.Comparable[T]] struct {
	element            [help.L15]T
	head, tail, length uint
}

func (this *Ring15[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM15
	this.head = (this.head + (this.length >> help.E15)) & help.RM15
	this.length += help.OneUint - (this.length >> help.E15)
	return true
}

func (this *Ring15[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM15) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring15[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring15[T]) Remove(T) bool { return false }
func (this *Ring15[T]) Len() int      { return int(this.length) }
func (this *Ring15[T]) Cap() int      { return int(help.L15) }

type Ring16[T common.Comparable[T]] struct {
	element            [help.L16]T
	head, tail, length uint
}

func (this *Ring16[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM16
	this.head = (this.head + (this.length >> help.E16)) & help.RM16
	this.length += help.OneUint - (this.length >> help.E16)
	return true
}

func (this *Ring16[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM16) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring16[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring16[T]) Remove(T) bool { return false }
func (this *Ring16[T]) Len() int      { return int(this.length) }
func (this *Ring16[T]) Cap() int      { return int(help.L16) }

type Ring17[T common.Comparable[T]] struct {
	element            [help.L17]T
	head, tail, length uint
}

func (this *Ring17[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM17
	this.head = (this.head + (this.length >> help.E17)) & help.RM17
	this.length += help.OneUint - (this.length >> help.E17)
	return true
}

func (this *Ring17[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM17) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring17[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring17[T]) Remove(T) bool { return false }
func (this *Ring17[T]) Len() int      { return int(this.length) }
func (this *Ring17[T]) Cap() int      { return int(help.L17) }

type Ring18[T common.Comparable[T]] struct {
	element            [help.L18]T
	head, tail, length uint
}

func (this *Ring18[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM18
	this.head = (this.head + (this.length >> help.E18)) & help.RM18
	this.length += help.OneUint - (this.length >> help.E18)
	return true
}

func (this *Ring18[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM18) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring18[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring18[T]) Remove(T) bool { return false }
func (this *Ring18[T]) Len() int      { return int(this.length) }
func (this *Ring18[T]) Cap() int      { return int(help.L18) }

type Ring19[T common.Comparable[T]] struct {
	element            [help.L19]T
	head, tail, length uint
}

func (this *Ring19[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM19
	this.head = (this.head + (this.length >> help.E19)) & help.RM19
	this.length += help.OneUint - (this.length >> help.E19)
	return true
}

func (this *Ring19[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM19) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring19[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring19[T]) Remove(T) bool { return false }
func (this *Ring19[T]) Len() int      { return int(this.length) }
func (this *Ring19[T]) Cap() int      { return int(help.L19) }

type Ring20[T common.Comparable[T]] struct {
	element            [help.L20]T
	head, tail, length uint
}

func (this *Ring20[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM20
	this.head = (this.head + (this.length >> help.E20)) & help.RM20
	this.length += help.OneUint - (this.length >> help.E20)
	return true
}

func (this *Ring20[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM20) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring20[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring20[T]) Remove(T) bool { return false }
func (this *Ring20[T]) Len() int      { return int(this.length) }
func (this *Ring20[T]) Cap() int      { return int(help.L20) }

type Ring21[T common.Comparable[T]] struct {
	element            [help.L21]T
	head, tail, length uint
}

func (this *Ring21[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM21
	this.head = (this.head + (this.length >> help.E21)) & help.RM21
	this.length += help.OneUint - (this.length >> help.E21)
	return true
}

func (this *Ring21[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM21) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring21[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring21[T]) Remove(T) bool { return false }
func (this *Ring21[T]) Len() int      { return int(this.length) }
func (this *Ring21[T]) Cap() int      { return int(help.L21) }

type Ring22[T common.Comparable[T]] struct {
	element            [help.L22]T
	head, tail, length uint
}

func (this *Ring22[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM22
	this.head = (this.head + (this.length >> help.E22)) & help.RM22
	this.length += help.OneUint - (this.length >> help.E22)
	return true
}

func (this *Ring22[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM22) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring22[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring22[T]) Remove(T) bool { return false }
func (this *Ring22[T]) Len() int      { return int(this.length) }
func (this *Ring22[T]) Cap() int      { return int(help.L22) }

type Ring23[T common.Comparable[T]] struct {
	element            [help.L23]T
	head, tail, length uint
}

func (this *Ring23[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM23
	this.head = (this.head + (this.length >> help.E23)) & help.RM23
	this.length += help.OneUint - (this.length >> help.E23)
	return true
}

func (this *Ring23[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM23) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring23[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring23[T]) Remove(T) bool { return false }
func (this *Ring23[T]) Len() int      { return int(this.length) }
func (this *Ring23[T]) Cap() int      { return int(help.L23) }

type Ring24[T common.Comparable[T]] struct {
	element            [help.L24]T
	head, tail, length uint
}

func (this *Ring24[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM24
	this.head = (this.head + (this.length >> help.E24)) & help.RM24
	this.length += help.OneUint - (this.length >> help.E24)
	return true
}

func (this *Ring24[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM24) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring24[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring24[T]) Remove(T) bool { return false }
func (this *Ring24[T]) Len() int      { return int(this.length) }
func (this *Ring24[T]) Cap() int      { return int(help.L24) }

type Ring25[T common.Comparable[T]] struct {
	element            [help.L25]T
	head, tail, length uint
}

func (this *Ring25[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM25
	this.head = (this.head + (this.length >> help.E25)) & help.RM25
	this.length += help.OneUint - (this.length >> help.E25)
	return true
}

func (this *Ring25[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM25) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring25[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring25[T]) Remove(T) bool { return false }
func (this *Ring25[T]) Len() int      { return int(this.length) }
func (this *Ring25[T]) Cap() int      { return int(help.L25) }

type Ring26[T common.Comparable[T]] struct {
	element            [help.L26]T
	head, tail, length uint
}

func (this *Ring26[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM26
	this.head = (this.head + (this.length >> help.E26)) & help.RM26
	this.length += help.OneUint - (this.length >> help.E26)
	return true
}

func (this *Ring26[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM26) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring26[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring26[T]) Remove(T) bool { return false }
func (this *Ring26[T]) Len() int      { return int(this.length) }
func (this *Ring26[T]) Cap() int      { return int(help.L26) }

type Ring27[T common.Comparable[T]] struct {
	element            [help.L27]T
	head, tail, length uint
}

func (this *Ring27[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM27
	this.head = (this.head + (this.length >> help.E27)) & help.RM27
	this.length += help.OneUint - (this.length >> help.E27)
	return true
}

func (this *Ring27[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM27) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring27[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring27[T]) Remove(T) bool { return false }
func (this *Ring27[T]) Len() int      { return int(this.length) }
func (this *Ring27[T]) Cap() int      { return int(help.L27) }

type Ring28[T common.Comparable[T]] struct {
	element            [help.L28]T
	head, tail, length uint
}

func (this *Ring28[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM28
	this.head = (this.head + (this.length >> help.E28)) & help.RM28
	this.length += help.OneUint - (this.length >> help.E28)
	return true
}

func (this *Ring28[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM28) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring28[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring28[T]) Remove(T) bool { return false }
func (this *Ring28[T]) Len() int      { return int(this.length) }
func (this *Ring28[T]) Cap() int      { return int(help.L28) }

type Ring29[T common.Comparable[T]] struct {
	element            [help.L29]T
	head, tail, length uint
}

func (this *Ring29[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM29
	this.head = (this.head + (this.length >> help.E29)) & help.RM29
	this.length += help.OneUint - (this.length >> help.E29)
	return true
}

func (this *Ring29[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM29) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}
func (this *Ring29[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring29[T]) Remove(T) bool { return false }
func (this *Ring29[T]) Len() int      { return int(this.length) }
func (this *Ring29[T]) Cap() int      { return int(help.L29) }

type Ring30[T common.Comparable[T]] struct {
	element            [help.L30]T
	head, tail, length uint
}

func (this *Ring30[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM30
	this.head = (this.head + (this.length >> help.E30)) & help.RM30
	this.length += help.OneUint - (this.length >> help.E30)
	return true
}

func (this *Ring30[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM30) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring30[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring30[T]) Remove(T) bool { return false }
func (this *Ring30[T]) Len() int      { return int(this.length) }
func (this *Ring30[T]) Cap() int      { return int(help.L30) }

type Ring31[T common.Comparable[T]] struct {
	element            [help.L31]T
	head, tail, length uint
}

func (this *Ring31[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM31
	this.head = (this.head + (this.length >> help.E31)) & help.RM31
	this.length += help.OneUint - (this.length >> help.E31)
	return true
}

func (this *Ring31[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM31) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring31[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring31[T]) Remove(T) bool { return false }
func (this *Ring31[T]) Len() int      { return int(this.length) }
func (this *Ring31[T]) Cap() int      { return int(help.L31) }

type Ring32[T common.Comparable[T]] struct {
	element            [help.L32]T
	head, tail, length uint
}

func (this *Ring32[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM32
	this.head = (this.head + (this.length >> help.E32)) & help.RM32
	this.length += help.OneUint - (this.length >> help.E32)
	return true
}

func (this *Ring32[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM32) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring32[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring32[T]) Remove(T) bool { return false }
func (this *Ring32[T]) Len() int      { return int(this.length) }
func (this *Ring32[T]) Cap() int      { return int(help.L32) }

type Ring33[T common.Comparable[T]] struct {
	element            [help.L33]T
	head, tail, length uint
}

func (this *Ring33[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM33
	this.head = (this.head + (this.length >> help.E33)) & help.RM33
	this.length += help.OneUint - (this.length >> help.E33)
	return true
}

func (this *Ring33[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM33) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring33[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring33[T]) Remove(T) bool { return false }
func (this *Ring33[T]) Len() int      { return int(this.length) }
func (this *Ring33[T]) Cap() int      { return int(help.L33) }

type Ring34[T common.Comparable[T]] struct {
	element            [help.L34]T
	head, tail, length uint
}

func (this *Ring34[T]) Push(element T) bool {
	*help.GetPtr(&this.element[help.ZeroInt], int(this.tail)) = element
	this.tail = (this.tail + help.OneUint) & help.RM34
	this.head = (this.head + (this.length >> help.E34)) & help.RM34
	this.length += help.OneUint - (this.length >> help.E34)
	return true
}

func (this *Ring34[T]) Pop() (T, bool) {
	element := *help.GetPtr(&this.element[help.ZeroInt], int(this.head))
	nz := (this.length | -this.length) >> help.WordRMask
	this.head = ((((this.head + help.OneUint) & help.RM34) & -nz) | (this.head &^ (-nz)))
	this.length -= nz
	return element, nz == help.OneUint
}

func (this *Ring34[T]) Peek() (T, bool) {
	return *help.GetPtr(&this.element[help.ZeroInt], int(this.head)), ((this.length | -this.length) >> help.WordRMask) == help.OneUint
}

func (this *Ring34[T]) Remove(T) bool { return false }
func (this *Ring34[T]) Len() int      { return int(this.length) }
func (this *Ring34[T]) Cap() int      { return int(help.L34) }
