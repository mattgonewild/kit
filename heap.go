// SPDX-License-Identifier: Unlicense

// Special thanks to the Go Authors for their original work on container/heap,
// which inspired and informed this implementation. While the concept draws from
// their work, the core algorithms and structure have been redesigned to meet our specific needs.

package kit

import (
	"github.com/ringboundio/common"
	"github.com/ringboundio/kit/internal/help"
)

func NewHeapAs() bool { return true }

type heapLowBitSet[T common.Standard[T, uint]] struct {
	heap     []T
	indexMap []int

	length int
	mask   uint
}

func NewHeapLowBitSet[T common.Standard[T, uint]](lowBitsUsed uint) common.Queue[T] {
	var (
		capacity = help.OneUint << lowBitsUsed
		heap     = make([]T, capacity)
		indexMap = make([]int, capacity)
		mask     = capacity - help.OneUint
	)

	for index := range indexMap {
		indexMap[index] = help.NegativeOneInt
	}

	return &heapLowBitSet[T]{
		heap:     heap,
		indexMap: indexMap,
		mask:     mask,
	}
}

func (this *heapLowBitSet[T]) Push(node T) bool {
	var mapIndex = this.getMapIndex(node)

	if this.indexMap[mapIndex] == help.NegativeOneInt {
		heapIndex := this.Len()
		this.heap[heapIndex] = node
		this.length++
		this.siftUp(heapIndex)

		return true
	}

	var heapIndex = this.indexMap[mapIndex]

	this.heap[heapIndex] = node
	this.fix(heapIndex)

	return true
}

func (this *heapLowBitSet[T]) siftUp(index int) {
	var (
		nodeToSift = this.heap[index]
		holeIndex  = index
	)

	for holeIndex > help.ZeroInt {
		parentIndex := (holeIndex - help.OneInt) >> help.ThreeUint
		if this.heap[parentIndex].Before(nodeToSift) {
			break
		}

		this.heap[holeIndex] = this.heap[parentIndex]
		this.setMapIndex(this.heap[holeIndex], holeIndex)

		holeIndex = parentIndex
	}

	this.heap[holeIndex] = nodeToSift
	this.setMapIndex(nodeToSift, holeIndex)
}

func (this *heapLowBitSet[T]) Pop() (T, bool) {
	var length = this.Len()

	if length == help.ZeroInt {
		return *new(T), false
	}

	var (
		root = this.heap[help.ZeroInt]
		high = length - help.OneInt
	)

	this.setMapIndex(root, help.NegativeOneInt)
	this.heap[help.ZeroInt] = this.heap[high]
	this.length--
	this.siftDown(help.ZeroInt)

	return root, true
}

func (this *heapLowBitSet[T]) siftDown(index int) bool {
	var (
		nodeToSift = this.heap[index]
		holeIndex  = index
		length     = this.Len()
	)

	for low := lowChild(holeIndex); low < length; low = lowChild(holeIndex) {
		var (
			high         = min((low + help.SevenInt), (length - help.OneInt))
			minNodeIndex = minimum(this.heap[low:high+help.OneInt]) + low
		)

		if nodeToSift.Before(this.heap[minNodeIndex]) {
			break
		}

		this.heap[holeIndex] = this.heap[minNodeIndex]
		this.setMapIndex(this.heap[holeIndex], holeIndex)
		holeIndex = minNodeIndex
	}

	this.heap[holeIndex] = nodeToSift
	this.setMapIndex(nodeToSift, holeIndex)

	return index != holeIndex
}

func lowChild(index int) int { return (index << help.ThreeUint) + help.OneInt }

func minimum[T common.Standard[T, uint]](group []T) int {
	var minimumIndex = help.ZeroInt

	for index, node := range group {
		if node.Before(group[minimumIndex]) {
			minimumIndex = index
		}
	}

	return minimumIndex
}

func (this *heapLowBitSet[T]) Peek() (T, bool) {
	if this.Len() == help.ZeroInt {
		return *new(T), false
	}

	return this.heap[help.ZeroInt], true
}

func (this *heapLowBitSet[T]) Remove(node T) bool {
	var (
		mapIndex  = this.getMapIndex(node)
		heapIndex = this.indexMap[mapIndex]
	)

	if heapIndex == help.NegativeOneInt {
		return false
	}

	var (
		nodeToRemove = this.heap[heapIndex]
		high         = this.Len() - help.OneInt
		nodeToSift   = this.heap[high]
	)

	this.setMapIndex(nodeToRemove, help.NegativeOneInt)
	this.heap[heapIndex] = nodeToSift
	this.length--
	this.fix(heapIndex)

	return true
}

func (this *heapLowBitSet[T]) fix(index int) {
	if !this.siftDown(index) {
		this.siftUp(index)
	}
}

func (this *heapLowBitSet[T]) setMapIndex(node T, index int) {
	this.indexMap[this.getMapIndex(node)] = index
}

func (this *heapLowBitSet[T]) getMapIndex(node T) int { return int(node.ID() & this.mask) }

func (this *heapLowBitSet[T]) Len() int { return this.length }
func (this *heapLowBitSet[T]) Cap() int { return cap(this.heap) }

type LeakyHeap[T common.Comparable[T]] struct {
	heap [help.TenInt]T

	length, poleIndex int
}

func (this *LeakyHeap[T]) Push(node T) {
	var length = this.length
	switch length {
	case help.ZeroInt:
		this.heap[help.ZeroInt] = node
		this.length = help.OneInt
		this.poleIndex = help.ZeroInt
		return
	case help.TenInt:
		var pole = this.heap[this.poleIndex]
		if pole.Equal(node) {
			this.heap[this.poleIndex] = node
			return
		}

		if !node.Before(pole) {
			return
		}
	}

	for index, current := range this.heap[:length] {
		if current.Equal(node) {
			this.heap[index] = node
			return
		}
	}

	if length < help.TenInt {
		this.heap[length] = node
		this.length++
		this.siftUp(length)
		this.updatePole()
		return
	}

	this.heap[this.poleIndex] = node
	this.fix(this.poleIndex)
	this.updatePole()
}

func (this *LeakyHeap[T]) Peek() (T, bool) {
	if this.length == help.ZeroInt {
		return this.heap[help.ZeroInt], false
	}

	return this.heap[help.ZeroInt], true
}

func (this *LeakyHeap[T]) Remove(node T) {
	if this.length == help.ZeroInt || this.heap[this.poleIndex].Before(node) {
		return
	}

	for index, current := range this.heap[:this.length] {
		if current.Equal(node) {
			high := this.length - help.OneInt
			if index == high {
				this.length--
				this.updatePole()
				return
			}

			this.heap[index] = this.heap[high]
			this.length--
			this.fix(index)
			this.updatePole()
			return
		}
	}
}

func (this *LeakyHeap[T]) fix(index int) {
	if !this.siftDown(index) {
		this.siftUp(index)
	}
}

func (this *LeakyHeap[T]) siftDown(index int) bool {
	var (
		nodeToSift = this.heap[index]
		holeIndex  = index
	)

	for low := (holeIndex << help.OneUint) + help.OneInt; low < this.length; low = (holeIndex << help.OneUint) + help.OneInt {
		var (
			high         = low + help.OneInt
			minNodeIndex = low
		)

		if high < this.length && this.heap[high].Before(this.heap[low]) {
			minNodeIndex = high
		}

		if nodeToSift.Before(this.heap[minNodeIndex]) {
			break
		}

		this.heap[holeIndex] = this.heap[minNodeIndex]
		holeIndex = minNodeIndex
	}

	this.heap[holeIndex] = nodeToSift
	return index != holeIndex
}

func (this *LeakyHeap[T]) siftUp(index int) {
	var (
		nodeToSift = this.heap[index]
		holeIndex  = index
	)

	for holeIndex > help.ZeroInt {
		parentIndex := (holeIndex - help.OneInt) >> help.OneUint
		if this.heap[parentIndex].Before(nodeToSift) {
			break
		}

		this.heap[holeIndex] = this.heap[parentIndex]

		holeIndex = parentIndex
	}

	this.heap[holeIndex] = nodeToSift
}

func (this *LeakyHeap[T]) updatePole() {
	var poleIndex int
	for index, current := range this.heap[:this.length] {
		if this.heap[poleIndex].Before(current) {
			poleIndex = index
		}
	}

	this.poleIndex = poleIndex
}

func (this *LeakyHeap[T]) Len() int { return this.length }
func (this *LeakyHeap[T]) Cap() int { return help.TenInt }
