package kit

import (
	"errors"
	"iter"
	"sync"

	"github.com/mattgonewild/common"
	"github.com/mattgonewild/kit/internal/help"
)

type doublyLinkedIndexCache struct {
	previous     *doublyLinkedIndexCache
	next         *doublyLinkedIndexCache
	length       int
	indexInCache int
	id           uint
	index        [help.SixtyFourInt]int
}

func (this *doublyLinkedIndexCache) delete(offset int) {
	ib := &this.index[help.ZeroInt]
	this.length--

	hole := offset
	end := this.length

	for hole < end {
		*help.GetPtr(ib, hole) = *help.GetPtr(ib, hole+help.OneInt)
		hole++
	}
}

func (this *doublyLinkedIndexCache) split(right *doublyLinkedIndexCache) {
	mid := help.ThirtyTwoInt
	ib, rb := &this.index[help.ZeroInt], &right.index[help.ZeroInt]

	for index := range mid {
		*help.GetPtr(rb, index) = *help.GetPtr(ib, mid+index)
	}

	this.length = help.ThirtyTwoInt
	right.length = help.ThirtyTwoInt
	right.previous = this
	right.next = this.next
	if this.next != nil {
		this.next.previous = right
	}
	this.next = right
}

func (this *doublyLinkedIndexCache) append(index int) {
	*help.GetPtr(&this.index[help.ZeroInt], this.length) = index
	this.length++
}

func (this *doublyLinkedIndexCache) shiftRightSet(offset, index int) {
	ib := &this.index[help.ZeroInt]

	{
		hole := this.length
		for hole > offset {
			*help.GetPtr(ib, hole) = *help.GetPtr(ib, hole-help.OneInt)
			hole--
		}
	}

	*help.GetPtr(ib, offset) = index
	this.length++
}

var (
	ErrSetNotFound = errors.New("matt::kit::set: not found")
	ErrSetFull     = errors.New("matt::kit::set: full")
)

// TODO: complete me

type CoarseSortedSet7[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet7[T]
}

func NewCoarseSortedSet7[T common.Comparable[T]]() *CoarseSortedSet7[T] {
	set := new(CoarseSortedSet7[T])

	InitLowfiFreelist7(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist1(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet7[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet7.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet7[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet7.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet7[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet7.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet7[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet7.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet7[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet7[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L1]int
	indexCache [help.L1]*doublyLinkedIndexCache

	free    LowfiFreelist7[uint]
	arena   LowfiFreelist1[doublyLinkedIndexCache]
	element [help.L7]T
}

func NewSortedSet7[T common.Comparable[T]]() *SortedSet7[T] {
	set := new(SortedSet7[T])

	InitLowfiFreelist7(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist1(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet7[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet7[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet7[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet7[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet7[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet7[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet7[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet7[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet7[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet7[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet7[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet7[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet7[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet7[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet8[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet8[T]
}

func NewCoarseSortedSet8[T common.Comparable[T]]() *CoarseSortedSet8[T] {
	set := new(CoarseSortedSet8[T])

	InitLowfiFreelist8(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist2(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet8[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet8.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet8[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet8.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet8[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet8.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet8[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet8.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet8[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet8[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L2]int
	indexCache [help.L2]*doublyLinkedIndexCache

	free    LowfiFreelist8[uint]
	arena   LowfiFreelist2[doublyLinkedIndexCache]
	element [help.L8]T
}

func NewSortedSet8[T common.Comparable[T]]() *SortedSet8[T] {
	set := new(SortedSet8[T])

	InitLowfiFreelist8(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist2(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet8[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet8[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet8[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet8[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet8[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet8[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet8[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet8[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet8[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet8[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet8[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet8[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet8[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet8[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet9[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet9[T]
}

func NewCoarseSortedSet9[T common.Comparable[T]]() *CoarseSortedSet9[T] {
	set := new(CoarseSortedSet9[T])

	InitLowfiFreelist9(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist3(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet9[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet9.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet9[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet9.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet9[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet9.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet9[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet9.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet9[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet9[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L3]int
	indexCache [help.L3]*doublyLinkedIndexCache

	free    LowfiFreelist9[uint]
	arena   LowfiFreelist3[doublyLinkedIndexCache]
	element [help.L9]T
}

func NewSortedSet9[T common.Comparable[T]]() *SortedSet9[T] {
	set := new(SortedSet9[T])

	InitLowfiFreelist9(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist3(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet9[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet9[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet9[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet9[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet9[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet9[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet9[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet9[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet9[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet9[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet9[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet9[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet9[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet9[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet10[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet10[T]
}

func NewCoarseSortedSet10[T common.Comparable[T]]() *CoarseSortedSet10[T] {
	set := new(CoarseSortedSet10[T])

	InitLowfiFreelist10(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist4(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet10[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet10.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet10[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet10.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet10[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet10.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet10[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet10.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet10[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet10[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L4]int
	indexCache [help.L4]*doublyLinkedIndexCache

	free    LowfiFreelist10[uint]
	arena   LowfiFreelist4[doublyLinkedIndexCache]
	element [help.L10]T
}

func NewSortedSet10[T common.Comparable[T]]() *SortedSet10[T] {
	set := new(SortedSet10[T])

	InitLowfiFreelist10(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist4(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet10[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet10[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet10[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet10[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet10[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet10[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet10[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet10[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet10[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet10[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet10[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet10[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet10[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet10[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet11[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet11[T]
}

func NewCoarseSortedSet11[T common.Comparable[T]]() *CoarseSortedSet11[T] {
	set := new(CoarseSortedSet11[T])

	InitLowfiFreelist11(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist5(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet11[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet11.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet11[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet11.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet11[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet11.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet11[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet11.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet11[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet11[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L5]int
	indexCache [help.L5]*doublyLinkedIndexCache

	free    LowfiFreelist11[uint]
	arena   LowfiFreelist5[doublyLinkedIndexCache]
	element [help.L11]T
}

func NewSortedSet11[T common.Comparable[T]]() *SortedSet11[T] {
	set := new(SortedSet11[T])

	InitLowfiFreelist11(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist5(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet11[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet11[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet11[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet11[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet11[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet11[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet11[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet11[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet11[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet11[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet11[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet11[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet11[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet11[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet12[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet12[T]
}

func NewCoarseSortedSet12[T common.Comparable[T]]() *CoarseSortedSet12[T] {
	set := new(CoarseSortedSet12[T])

	InitLowfiFreelist12(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist6(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet12[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet12.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet12[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet12.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet12[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet12.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet12[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet12.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet12[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet12[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L6]int
	indexCache [help.L6]*doublyLinkedIndexCache

	free    LowfiFreelist12[uint]
	arena   LowfiFreelist6[doublyLinkedIndexCache]
	element [help.L12]T
}

func NewSortedSet12[T common.Comparable[T]]() *SortedSet12[T] {
	set := new(SortedSet12[T])

	InitLowfiFreelist12(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist6(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet12[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet12[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet12[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet12[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet12[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet12[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet12[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet12[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet12[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet12[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet12[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet12[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet12[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet12[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet13[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet13[T]
}

func NewCoarseSortedSet13[T common.Comparable[T]]() *CoarseSortedSet13[T] {
	set := new(CoarseSortedSet13[T])

	InitLowfiFreelist13(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist7(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet13[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet13.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet13[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet13.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet13[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet13.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet13[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet13.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet13[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet13[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L7]int
	indexCache [help.L7]*doublyLinkedIndexCache

	free    LowfiFreelist13[uint]
	arena   LowfiFreelist7[doublyLinkedIndexCache]
	element [help.L13]T
}

func NewSortedSet13[T common.Comparable[T]]() *SortedSet13[T] {
	set := new(SortedSet13[T])

	InitLowfiFreelist13(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist7(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet13[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet13[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet13[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet13[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet13[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet13[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet13[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet13[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet13[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet13[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet13[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet13[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet13[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet13[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet14[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet14[T]
}

func NewCoarseSortedSet14[T common.Comparable[T]]() *CoarseSortedSet14[T] {
	set := new(CoarseSortedSet14[T])

	InitLowfiFreelist14(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist8(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet14[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet14.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet14[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet14.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet14[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet14.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet14[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet14.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet14[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet14[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L8]int
	indexCache [help.L8]*doublyLinkedIndexCache

	free    LowfiFreelist14[uint]
	arena   LowfiFreelist8[doublyLinkedIndexCache]
	element [help.L14]T
}

func NewSortedSet14[T common.Comparable[T]]() *SortedSet14[T] {
	set := new(SortedSet14[T])

	InitLowfiFreelist14(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist8(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet14[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet14[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet14[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet14[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet14[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet14[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet14[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet14[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet14[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet14[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet14[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet14[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet14[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet14[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet15[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet15[T]
}

func NewCoarseSortedSet15[T common.Comparable[T]]() *CoarseSortedSet15[T] {
	set := new(CoarseSortedSet15[T])

	InitLowfiFreelist15(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist9(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet15[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet15.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet15[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet15.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet15[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet15.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet15[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet15.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet15[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet15[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L9]int
	indexCache [help.L9]*doublyLinkedIndexCache

	free    LowfiFreelist15[uint]
	arena   LowfiFreelist9[doublyLinkedIndexCache]
	element [help.L15]T
}

func NewSortedSet15[T common.Comparable[T]]() *SortedSet15[T] {
	set := new(SortedSet15[T])

	InitLowfiFreelist15(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist9(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet15[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet15[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet15[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet15[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet15[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet15[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet15[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet15[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet15[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet15[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet15[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet15[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet15[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet15[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet16[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet16[T]
}

func NewCoarseSortedSet16[T common.Comparable[T]]() *CoarseSortedSet16[T] {
	set := new(CoarseSortedSet16[T])

	InitLowfiFreelist16(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist10(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet16[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet16.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet16[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet16.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet16[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet16.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet16[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet16.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet16[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet16[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L10]int
	indexCache [help.L10]*doublyLinkedIndexCache

	free    LowfiFreelist16[uint]
	arena   LowfiFreelist10[doublyLinkedIndexCache]
	element [help.L16]T
}

func NewSortedSet16[T common.Comparable[T]]() *SortedSet16[T] {
	set := new(SortedSet16[T])

	InitLowfiFreelist16(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist10(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet16[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet16[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet16[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet16[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet16[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet16[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet16[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet16[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet16[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet16[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet16[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet16[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet16[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet16[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet17[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet17[T]
}

func NewCoarseSortedSet17[T common.Comparable[T]]() *CoarseSortedSet17[T] {
	set := new(CoarseSortedSet17[T])

	InitLowfiFreelist17(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist11(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet17[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet17.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet17[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet17.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet17[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet17.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet17[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet17.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet17[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet17[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L11]int
	indexCache [help.L11]*doublyLinkedIndexCache

	free    LowfiFreelist17[uint]
	arena   LowfiFreelist11[doublyLinkedIndexCache]
	element [help.L17]T
}

func NewSortedSet17[T common.Comparable[T]]() *SortedSet17[T] {
	set := new(SortedSet17[T])

	InitLowfiFreelist17(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist11(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet17[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet17[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet17[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet17[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet17[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet17[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet17[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet17[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet17[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet17[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet17[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet17[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet17[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet17[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet18[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet18[T]
}

func NewCoarseSortedSet18[T common.Comparable[T]]() *CoarseSortedSet18[T] {
	set := new(CoarseSortedSet18[T])

	InitLowfiFreelist18(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist12(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet18[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet18.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet18[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet18.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet18[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet18.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet18[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet18.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet18[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet18[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L12]int
	indexCache [help.L12]*doublyLinkedIndexCache

	free    LowfiFreelist18[uint]
	arena   LowfiFreelist12[doublyLinkedIndexCache]
	element [help.L18]T
}

func NewSortedSet18[T common.Comparable[T]]() *SortedSet18[T] {
	set := new(SortedSet18[T])

	InitLowfiFreelist18(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist12(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet18[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet18[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet18[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet18[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet18[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet18[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet18[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet18[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet18[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet18[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet18[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet18[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet18[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet18[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet19[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet19[T]
}

func NewCoarseSortedSet19[T common.Comparable[T]]() *CoarseSortedSet19[T] {
	set := new(CoarseSortedSet19[T])

	InitLowfiFreelist19(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist13(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet19[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet19.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet19[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet19.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet19[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet19.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet19[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet19.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet19[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet19[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L13]int
	indexCache [help.L13]*doublyLinkedIndexCache

	free    LowfiFreelist19[uint]
	arena   LowfiFreelist13[doublyLinkedIndexCache]
	element [help.L19]T
}

func NewSortedSet19[T common.Comparable[T]]() *SortedSet19[T] {
	set := new(SortedSet19[T])

	InitLowfiFreelist19(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist13(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet19[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet19[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet19[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet19[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet19[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet19[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet19[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet19[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet19[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet19[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet19[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet19[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet19[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet19[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet20[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet20[T]
}

func NewCoarseSortedSet20[T common.Comparable[T]]() *CoarseSortedSet20[T] {
	set := new(CoarseSortedSet20[T])

	InitLowfiFreelist20(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist14(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet20[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet20.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet20[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet20.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet20[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet20.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet20[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet20.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet20[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet20[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L14]int
	indexCache [help.L14]*doublyLinkedIndexCache

	free    LowfiFreelist20[uint]
	arena   LowfiFreelist14[doublyLinkedIndexCache]
	element [help.L20]T
}

func NewSortedSet20[T common.Comparable[T]]() *SortedSet20[T] {
	set := new(SortedSet20[T])

	InitLowfiFreelist20(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist14(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet20[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet20[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet20[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet20[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet20[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet20[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet20[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet20[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet20[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet20[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet20[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet20[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet20[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet20[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet21[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet21[T]
}

func NewCoarseSortedSet21[T common.Comparable[T]]() *CoarseSortedSet21[T] {
	set := new(CoarseSortedSet21[T])

	InitLowfiFreelist21(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist15(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet21[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet21.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet21[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet21.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet21[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet21.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet21[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet21.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet21[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet21[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L15]int
	indexCache [help.L15]*doublyLinkedIndexCache

	free    LowfiFreelist21[uint]
	arena   LowfiFreelist15[doublyLinkedIndexCache]
	element [help.L21]T
}

func NewSortedSet21[T common.Comparable[T]]() *SortedSet21[T] {
	set := new(SortedSet21[T])

	InitLowfiFreelist21(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist15(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet21[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet21[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet21[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet21[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet21[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet21[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet21[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet21[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet21[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet21[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet21[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet21[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet21[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet21[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet22[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet22[T]
}

func NewCoarseSortedSet22[T common.Comparable[T]]() *CoarseSortedSet22[T] {
	set := new(CoarseSortedSet22[T])

	InitLowfiFreelist22(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist16(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet22[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet22.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet22[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet22.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet22[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet22.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet22[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet22.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet22[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet22[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L16]int
	indexCache [help.L16]*doublyLinkedIndexCache

	free    LowfiFreelist22[uint]
	arena   LowfiFreelist16[doublyLinkedIndexCache]
	element [help.L22]T
}

func NewSortedSet22[T common.Comparable[T]]() *SortedSet22[T] {
	set := new(SortedSet22[T])

	InitLowfiFreelist22(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist16(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet22[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet22[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet22[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet22[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet22[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet22[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet22[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet22[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet22[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet22[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet22[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet22[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet22[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet22[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet23[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet23[T]
}

func NewCoarseSortedSet23[T common.Comparable[T]]() *CoarseSortedSet23[T] {
	set := new(CoarseSortedSet23[T])

	InitLowfiFreelist23(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist17(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet23[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet23.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet23[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet23.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet23[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet23.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet23[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet23.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet23[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet23[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L17]int
	indexCache [help.L17]*doublyLinkedIndexCache

	free    LowfiFreelist23[uint]
	arena   LowfiFreelist17[doublyLinkedIndexCache]
	element [help.L23]T
}

func NewSortedSet23[T common.Comparable[T]]() *SortedSet23[T] {
	set := new(SortedSet23[T])

	InitLowfiFreelist23(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist17(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet23[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet23[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet23[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet23[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet23[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet23[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet23[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet23[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet23[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet23[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet23[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet23[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet23[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet23[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet24[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet24[T]
}

func NewCoarseSortedSet24[T common.Comparable[T]]() *CoarseSortedSet24[T] {
	set := new(CoarseSortedSet24[T])

	InitLowfiFreelist24(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist18(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet24[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet24.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet24[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet24.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet24[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet24.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet24[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet24.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet24[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet24[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L18]int
	indexCache [help.L18]*doublyLinkedIndexCache

	free    LowfiFreelist24[uint]
	arena   LowfiFreelist18[doublyLinkedIndexCache]
	element [help.L24]T
}

func NewSortedSet24[T common.Comparable[T]]() *SortedSet24[T] {
	set := new(SortedSet24[T])

	InitLowfiFreelist24(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist18(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet24[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet24[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet24[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet24[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet24[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet24[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet24[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet24[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet24[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet24[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet24[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet24[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet24[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet24[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet25[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet25[T]
}

func NewCoarseSortedSet25[T common.Comparable[T]]() *CoarseSortedSet25[T] {
	set := new(CoarseSortedSet25[T])

	InitLowfiFreelist25(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist19(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet25[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet25.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet25[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet25.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet25[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet25.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet25[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet25.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet25[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet25[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L19]int
	indexCache [help.L19]*doublyLinkedIndexCache

	free    LowfiFreelist25[uint]
	arena   LowfiFreelist19[doublyLinkedIndexCache]
	element [help.L25]T
}

func NewSortedSet25[T common.Comparable[T]]() *SortedSet25[T] {
	set := new(SortedSet25[T])

	InitLowfiFreelist25(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist19(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet25[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet25[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet25[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet25[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet25[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet25[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet25[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet25[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet25[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet25[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet25[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet25[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet25[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet25[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet26[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet26[T]
}

func NewCoarseSortedSet26[T common.Comparable[T]]() *CoarseSortedSet26[T] {
	set := new(CoarseSortedSet26[T])

	InitLowfiFreelist26(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist20(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet26[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet26.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet26[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet26.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet26[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet26.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet26[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet26.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet26[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet26[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L20]int
	indexCache [help.L20]*doublyLinkedIndexCache

	free    LowfiFreelist26[uint]
	arena   LowfiFreelist20[doublyLinkedIndexCache]
	element [help.L26]T
}

func NewSortedSet26[T common.Comparable[T]]() *SortedSet26[T] {
	set := new(SortedSet26[T])

	InitLowfiFreelist26(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist20(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet26[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet26[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet26[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet26[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet26[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet26[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet26[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet26[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet26[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet26[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet26[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet26[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet26[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet26[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet27[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet27[T]
}

func NewCoarseSortedSet27[T common.Comparable[T]]() *CoarseSortedSet27[T] {
	set := new(CoarseSortedSet27[T])

	InitLowfiFreelist27(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist21(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet27[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet27.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet27[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet27.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet27[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet27.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet27[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet27.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet27[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet27[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L21]int
	indexCache [help.L21]*doublyLinkedIndexCache

	free    LowfiFreelist27[uint]
	arena   LowfiFreelist21[doublyLinkedIndexCache]
	element [help.L27]T
}

func NewSortedSet27[T common.Comparable[T]]() *SortedSet27[T] {
	set := new(SortedSet27[T])

	InitLowfiFreelist27(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist21(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet27[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet27[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet27[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet27[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet27[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet27[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet27[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet27[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet27[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet27[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet27[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet27[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet27[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet27[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet28[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet28[T]
}

func NewCoarseSortedSet28[T common.Comparable[T]]() *CoarseSortedSet28[T] {
	set := new(CoarseSortedSet28[T])

	InitLowfiFreelist28(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist22(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet28[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet28.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet28[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet28.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet28[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet28.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet28[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet28.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet28[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet28[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L22]int
	indexCache [help.L22]*doublyLinkedIndexCache

	free    LowfiFreelist28[uint]
	arena   LowfiFreelist22[doublyLinkedIndexCache]
	element [help.L28]T
}

func NewSortedSet28[T common.Comparable[T]]() *SortedSet28[T] {
	set := new(SortedSet28[T])

	InitLowfiFreelist28(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist22(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet28[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet28[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet28[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet28[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet28[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet28[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet28[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet28[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet28[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet28[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet28[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet28[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet28[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet28[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet29[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet29[T]
}

func NewCoarseSortedSet29[T common.Comparable[T]]() *CoarseSortedSet29[T] {
	set := new(CoarseSortedSet29[T])

	InitLowfiFreelist29(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist23(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet29[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet29.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet29[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet29.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet29[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet29.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet29[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet29.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet29[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet29[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L23]int
	indexCache [help.L23]*doublyLinkedIndexCache

	free    LowfiFreelist29[uint]
	arena   LowfiFreelist23[doublyLinkedIndexCache]
	element [help.L29]T
}

func NewSortedSet29[T common.Comparable[T]]() *SortedSet29[T] {
	set := new(SortedSet29[T])

	InitLowfiFreelist29(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist23(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet29[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet29[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet29[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet29[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet29[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet29[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet29[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet29[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet29[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet29[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet29[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet29[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet29[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet29[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet30[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet30[T]
}

func NewCoarseSortedSet30[T common.Comparable[T]]() *CoarseSortedSet30[T] {
	set := new(CoarseSortedSet30[T])

	InitLowfiFreelist30(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist24(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet30[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet30.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet30[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet30.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet30[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet30.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet30[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet30.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet30[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet30[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L24]int
	indexCache [help.L24]*doublyLinkedIndexCache

	free    LowfiFreelist30[uint]
	arena   LowfiFreelist24[doublyLinkedIndexCache]
	element [help.L30]T
}

func NewSortedSet30[T common.Comparable[T]]() *SortedSet30[T] {
	set := new(SortedSet30[T])

	InitLowfiFreelist30(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist24(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet30[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet30[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet30[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet30[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet30[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet30[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet30[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet30[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet30[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet30[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet30[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet30[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet30[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet30[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet31[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet31[T]
}

func NewCoarseSortedSet31[T common.Comparable[T]]() *CoarseSortedSet31[T] {
	set := new(CoarseSortedSet31[T])

	InitLowfiFreelist31(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist25(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet31[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet31.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet31[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet31.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet31[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet31.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet31[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet31.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet31[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet31[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L25]int
	indexCache [help.L25]*doublyLinkedIndexCache

	free    LowfiFreelist31[uint]
	arena   LowfiFreelist25[doublyLinkedIndexCache]
	element [help.L31]T
}

func NewSortedSet31[T common.Comparable[T]]() *SortedSet31[T] {
	set := new(SortedSet31[T])

	InitLowfiFreelist31(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist25(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet31[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet31[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet31[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet31[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet31[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet31[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet31[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet31[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet31[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet31[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet31[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet31[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet31[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet31[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet32[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet32[T]
}

func NewCoarseSortedSet32[T common.Comparable[T]]() *CoarseSortedSet32[T] {
	set := new(CoarseSortedSet32[T])

	InitLowfiFreelist32(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist26(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet32[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet32.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet32[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet32.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet32[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet32.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet32[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet32.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet32[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet32[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L26]int
	indexCache [help.L26]*doublyLinkedIndexCache

	free    LowfiFreelist32[uint]
	arena   LowfiFreelist26[doublyLinkedIndexCache]
	element [help.L32]T
}

func NewSortedSet32[T common.Comparable[T]]() *SortedSet32[T] {
	set := new(SortedSet32[T])

	InitLowfiFreelist32(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist26(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet32[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet32[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet32[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet32[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet32[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet32[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet32[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet32[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet32[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet32[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet32[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet32[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet32[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet32[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet33[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet33[T]
}

func NewCoarseSortedSet33[T common.Comparable[T]]() *CoarseSortedSet33[T] {
	set := new(CoarseSortedSet33[T])

	InitLowfiFreelist33(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist27(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet33[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet33.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet33[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet33.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet33[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet33.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet33[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet33.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet33[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet33[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L27]int
	indexCache [help.L27]*doublyLinkedIndexCache

	free    LowfiFreelist33[uint]
	arena   LowfiFreelist27[doublyLinkedIndexCache]
	element [help.L33]T
}

func NewSortedSet33[T common.Comparable[T]]() *SortedSet33[T] {
	set := new(SortedSet33[T])

	InitLowfiFreelist33(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist27(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet33[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet33[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet33[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet33[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet33[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet33[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet33[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet33[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet33[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet33[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet33[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet33[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet33[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet33[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}

type CoarseSortedSet34[T common.Comparable[T]] struct {
	mu sync.RWMutex
	SortedSet34[T]
}

func NewCoarseSortedSet34[T common.Comparable[T]]() *CoarseSortedSet34[T] {
	set := new(CoarseSortedSet34[T])

	InitLowfiFreelist34(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist28(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *CoarseSortedSet34[T]) Put(element T) error {
	this.mu.Lock()
	err := this.SortedSet34.Put(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet34[T]) Get(element T) (T, error) {
	this.mu.RLock()
	element, err := this.SortedSet34.Get(element)
	this.mu.RUnlock()
	return element, err
}

func (this *CoarseSortedSet34[T]) Delete(element T) error {
	this.mu.Lock()
	err := this.SortedSet34.Delete(element)
	this.mu.Unlock()
	return err
}

func (this *CoarseSortedSet34[T]) ForEach(yield func(T) bool) {
	this.mu.RLock()
	this.SortedSet34.ForEach(yield)
	this.mu.RUnlock()
}

func (this *CoarseSortedSet34[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		this.mu.RLock()

		if this.length == help.ZeroInt {
			this.mu.RUnlock()
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					this.mu.RUnlock()
					return
				}
			}

			node = node.next
		}

		this.mu.RUnlock()
	}
}

type SortedSet34[T common.Comparable[T]] struct {
	length     int
	maxima     [help.L28]int
	indexCache [help.L28]*doublyLinkedIndexCache

	free    LowfiFreelist34[uint]
	arena   LowfiFreelist28[doublyLinkedIndexCache]
	element [help.L34]T
}

func NewSortedSet34[T common.Comparable[T]]() *SortedSet34[T] {
	set := new(SortedSet34[T])

	InitLowfiFreelist34(&set.free, func(index uint) uint {
		return index
	})

	InitLowfiFreelist28(&set.arena, func(index uint) doublyLinkedIndexCache {
		return doublyLinkedIndexCache{id: index}
	})

	return set
}

func (this *SortedSet34[T]) Put(element T) error {
	if this.length == help.ZeroInt {
		head := this.getFreeNodeUnsafe()
		head.previous = nil
		head.next = nil
		head.indexInCache = help.ZeroInt
		index := this.getFreeIndexUnsafe()
		head.append(index)
		this.maxima[help.ZeroInt] = index
		this.indexCache[help.ZeroInt] = head
		*help.GetPtr(&this.element[help.ZeroInt], index) = element
		return nil
	}

	mb := &this.maxima[help.ZeroInt]
	ib := &this.indexCache[help.ZeroInt]
	eb := &this.element[help.ZeroInt]

	{
		head := this.indexCache[help.ZeroInt]

		e := *help.GetPtr(eb, head.index[help.ZeroInt])
		if e.After(element) {
			if head.length == help.SixtyFourInt {
				left, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftRight(head)
					head.shiftRightSet(help.ZeroInt, index)
					this.maxima[help.ZeroInt] = *help.GetPtr(&head.index[help.ZeroInt], help.SixtyThreeInt)
					*help.GetPtr(eb, index) = element
					return nil
				}

				left.previous = nil
				left.next = head
				left.indexInCache = help.ZeroInt
				index := this.getFreeIndexUnsafe()
				left.append(index)
				head.previous = left
				head.indexInCache = help.OneInt
				this.icShiftRightAdd63(head)
				this.maxima[help.ZeroInt] = index
				this.indexCache[help.ZeroInt] = left
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			head.shiftRightSet(help.ZeroInt, index)
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	{
		high := this.length - help.OneInt

		e := *help.GetPtr(eb, *help.GetPtr(mb, high))
		if e.Before(element) {
			tail := *help.GetPtr(ib, high)
			if tail.length == help.SixtyFourInt {
				right, ok := this.arena.Get()
				if !ok {
					index, ok := this.getFreeIndex()
					if !ok {
						return ErrSetFull
					}

					this.xShiftLeft(tail)
					tail.append(index)
					*help.GetPtr(mb, tail.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				right.previous = tail
				right.next = nil
				tail.next = right
				index := this.getFreeIndexUnsafe()
				right.append(index)
				high += help.OneInt
				right.indexInCache = high
				*help.GetPtr(mb, high) = index
				*help.GetPtr(ib, high) = right
				*help.GetPtr(eb, index) = element
				this.length++
				return nil
			}

			index := this.getFreeIndexUnsafe()
			tail.append(index)
			*help.GetPtr(mb, high) = index
			*help.GetPtr(eb, index) = element
			return nil
		}
	}

	// element is not greater than the current max
	node, offset := this.getInsertionPoint(element)

	{
		ptr := help.GetPtr(eb, *help.GetPtr(&node.index[help.ZeroInt], offset))
		e := *ptr

		if e.Equal(element) {
			*ptr = element
			return nil
		}
	}

	// element is not present; offset is > than element
	if node.length == help.SixtyFourInt {
		right, ok := this.arena.Get()
		if !ok {
			index, ok := this.getFreeIndex()
			if !ok {
				return ErrSetFull
			}

			if !this.xShiftRight(node) {
				if offset == help.ZeroInt {
					node = node.previous
					this.xShiftLeft(node)
					node.append(index)
					*help.GetPtr(mb, node.indexInCache) = index
					*help.GetPtr(eb, index) = element
					return nil
				}

				this.xShiftLeft(node)
				node.shiftRightSet(offset-help.OneInt, index)
				*help.GetPtr(eb, index) = element
				return nil
			}

			node.shiftRightSet(offset, index)
			*help.GetPtr(mb, node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
			*help.GetPtr(eb, index) = element
			return nil
		}

		node.split(right)
		if offset < help.ThirtyTwoInt {
			index := this.getFreeIndexUnsafe()
			node.shiftRightSet(offset, index)
			nodeIndexInCache := node.indexInCache
			right.indexInCache = nodeIndexInCache + help.OneInt
			this.icShiftRightAdd31(right)
			*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
			*help.GetPtr(eb, index) = element
			this.length++
			return nil
		}

		index := this.getFreeIndexUnsafe()
		right.shiftRightSet(offset-help.ThirtyTwoInt, index)
		nodeIndexInCache := node.indexInCache
		right.indexInCache = nodeIndexInCache + help.OneInt
		this.icShiftRightAdd32(right)
		*help.GetPtr(mb, nodeIndexInCache) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
		*help.GetPtr(eb, index) = element
		this.length++
		return nil
	}

	index := this.getFreeIndexUnsafe()
	node.shiftRightSet(offset, index)
	*help.GetPtr(eb, index) = element
	return nil
}

func (this *SortedSet34[T]) getFreeNodeUnsafe() *doublyLinkedIndexCache {
	node, _ := this.arena.Get()
	this.length++
	return node
}

func (this *SortedSet34[T]) getFreeIndex() (int, bool) {
	ptr, ok := this.free.Get()
	if !ok {
		return help.NegativeOneInt, false
	}

	return int(*ptr), true
}

func (this *SortedSet34[T]) getFreeIndexUnsafe() int {
	ptr, _ := this.free.Get()
	return int(*ptr)
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet34[T]) xShiftLeft(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.previous
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]
	for sink != start {
		next := sink.next
		low := *help.GetPtr(&next.index[help.ZeroInt], help.ZeroInt)
		sink.append(low)
		*help.GetPtr(mb, sink.indexInCache) = low
		next.delete(help.ZeroInt)
		sink = next
	}

	return true
}

// should not be called on a non-full node and start's maxima must be updated by caller
func (this *SortedSet34[T]) xShiftRight(start *doublyLinkedIndexCache) bool {
	sink := start
	for sink != nil && sink.length == help.SixtyFourInt {
		sink = sink.next
	}

	if sink == nil {
		return false
	}

	mb := &this.maxima[help.ZeroInt]

	{
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	for sink != start {
		previous := sink.previous
		sink.shiftRightSet(help.ZeroInt, *help.GetPtr(&previous.index[help.ZeroInt], help.SixtyThreeInt))
		*help.GetPtr(mb, sink.indexInCache) = *help.GetPtr(&sink.index[help.ZeroInt], help.SixtyThreeInt)
		previous.delete(help.SixtyThreeInt)
		sink = previous
	}

	return true
}

func (this *SortedSet34[T]) Get(element T) (T, error) {
	if this.length == help.ZeroInt {
		return element, ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	if node.length == offset {
		return element, ErrSetNotFound
	}

	e := *help.GetPtr(&this.element[help.ZeroInt], *help.GetPtr(&node.index[help.ZeroInt], offset))
	if e.Equal(element) {
		return e, nil
	}

	return element, ErrSetNotFound
}

func (this *SortedSet34[T]) Delete(element T) error {
	setNodeLen := this.length
	if setNodeLen == help.ZeroInt {
		return ErrSetNotFound
	}

	node, offset := this.getInsertionPoint(element)
	nodeIdxLen := node.length
	if nodeIdxLen == offset {
		return ErrSetNotFound
	}

	elemIndex := *help.GetPtr(&node.index[help.ZeroInt], offset)
	e := *help.GetPtr(&this.element[help.ZeroInt], elemIndex)
	if !e.Equal(element) {
		return ErrSetNotFound
	}

	if nodeIdxLen == help.OneInt {
		this.free.Release(uint(elemIndex))

		if node.previous != nil {
			node.previous.next = node.next
		}

		if node.next != nil {
			node.next.previous = node.previous
		}

		hole := node.indexInCache
		this.arena.Release(node.id)

		{
			maxima := this.maxima[hole:setNodeLen]
			copy(maxima, maxima[help.OneInt:])
		}

		ib := &this.indexCache[help.ZeroInt]
		setNodeLen--
		for hole < setNodeLen {
			node := *help.GetPtr(ib, hole+help.OneInt)
			*help.GetPtr(ib, hole) = node
			node.indexInCache = hole
			hole++
		}

		this.length = setNodeLen
		return nil
	}

	if offset == nodeIdxLen-help.OneInt {
		this.free.Release(uint(elemIndex))
		node.delete(offset)
		*help.GetPtr(&this.maxima[help.ZeroInt], node.indexInCache) = *help.GetPtr(&node.index[help.ZeroInt], offset-help.OneInt)
		return nil
	}

	this.free.Release(uint(elemIndex))
	node.delete(offset)
	return nil
}

// must not be called when this.length == zero
func (this *SortedSet34[T]) getInsertionPoint(element T) (*doublyLinkedIndexCache, int) {
	var (
		eb     = &this.element[help.ZeroInt]
		low    = help.ZeroInt
		length = this.length
		high   = length - help.OneInt // TODO: this is intentionally incorrect to avoid ever getting a nil pointer
	)

	{
		mb := &this.maxima[help.ZeroInt]

		for low < high {
			mid := low + ((high - low) >> help.OneInt)
			e := *help.GetPtr(eb, *help.GetPtr(mb, mid))

			if e.Before(element) {
				low = mid + help.OneInt
			} else {
				high = mid
			}
		}
	}

	var (
		node = *help.GetPtr(&this.indexCache[help.ZeroInt], low)
		nb   = &node.index[help.ZeroInt]
	)

	low, high = help.ZeroInt, node.length

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		e := *help.GetPtr(eb, *help.GetPtr(nb, mid))

		if e.Before(element) {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	return node, low
}

func (this *SortedSet34[T]) icShiftRightAdd63(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.SixtyThreeInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet34[T]) icShiftRightAdd31(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyOneInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet34[T]) icShiftRightAdd32(node *doublyLinkedIndexCache) {
	var (
		mb   = &this.maxima[help.ZeroInt]
		ib   = &this.indexCache[help.ZeroInt]
		hole = this.length
		mark = node.indexInCache
	)

	for hole > mark {
		*help.GetPtr(mb, hole) = *help.GetPtr(mb, hole-help.OneInt)
		left := *help.GetPtr(ib, hole-help.OneInt)
		*help.GetPtr(ib, hole) = left
		left.indexInCache = hole
		hole--
	}

	*help.GetPtr(mb, mark) = *help.GetPtr(&node.index[help.ZeroInt], help.ThirtyTwoInt)
	*help.GetPtr(ib, mark) = node
}

func (this *SortedSet34[T]) ForEach(yield func(T) bool) {
	// TODO: As common.Set[T], ForEach was roughly half the cost, so omit All from the interface.
	if this.length == help.ZeroInt {
		return
	}

	var (
		eb   = &this.element[help.ZeroInt]
		node = this.indexCache[help.ZeroInt]
	)

	for node != nil {
		nb := &node.index[help.ZeroInt]
		for index := range node.length {
			if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
				return
			}
		}

		node = node.next
	}
}

func (this *SortedSet34[T]) All() iter.Seq[T] {
	// TODO: 2/3 cost reduction when returning an anonymous function; keep both.
	return func(yield func(T) bool) {
		if this.length == help.ZeroInt {
			return
		}

		var (
			eb   = &this.element[help.ZeroInt]
			node = this.indexCache[help.ZeroInt]
		)

		for node != nil {
			nb := &node.index[help.ZeroInt]
			for index := range node.length {
				if !yield(*help.GetPtr(eb, *help.GetPtr(nb, index))) {
					return
				}
			}

			node = node.next
		}
	}
}
