package kit

import (
	"errors"
	"math/bits"
	"time"
	"unsafe"

	"github.com/mattgonewild/common"
	"github.com/mattgonewild/kit/internal/help"
)

func NewLogAs() bool { return true }

type linkedSegBucket[T common.UnixTimestamped] struct {
	element []T

	base *T
	next *linkedSegBucket[T]
}

func (this *linkedSegBucket[T]) getLowerOffsetMinusOnePtr(unixTime int64) (int, *T, bool) {
	var (
		low  = help.ZeroInt
		base = this.base
		high = len(this.element)
	)

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		midUnixTime := help.UnixAt(base, mid)

		if midUnixTime < unixTime {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	if low == help.ZeroInt {
		return help.ZeroInt, nil, false
	}

	return low, help.GetPtr(base, low-help.OneInt), true
}

func (this *linkedSegBucket[T]) getUpperOffsetMinusOnePtr(unixTime int64) (int, *T, bool) {
	var (
		low  = help.ZeroInt
		base = this.base
		high = len(this.element)
	)

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		midUnixTime := help.UnixAt(base, mid)

		if midUnixTime > unixTime {
			high = mid
		} else {
			low = mid + help.OneInt
		}
	}

	if low == help.ZeroInt {
		return help.ZeroInt, nil, false
	}

	return low, help.GetPtr(base, low-help.OneInt), true
}

func (this *linkedSegBucket[T]) getUpperOffsetPtr(unixTime int64) (int, *T) {
	var (
		low  = help.ZeroInt
		base = this.base
		high = len(this.element)
	)

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		midUnixTime := help.UnixAt(base, mid)

		if midUnixTime > unixTime {
			high = mid
		} else {
			low = mid + help.OneInt
		}
	}

	return low, help.GetPtr(base, low)
}

func (this *linkedSegBucket[T]) getLowerOffsetPtr(unixTime int64) (int, *T, bool) {
	var (
		low    = help.ZeroInt
		base   = this.base
		length = len(this.element)
		high   = length
	)

	for low < high {
		mid := low + ((high - low) >> help.OneInt)
		midUnixTime := help.UnixAt(base, mid)

		if midUnixTime < unixTime {
			low = mid + help.OneInt
		} else {
			high = mid
		}
	}

	if low == length {
		return length, nil, false
	}

	return low, help.GetPtr(base, low), true
}

func (this *linkedSegBucket[T]) getLengthHighUnixTime() (int, int64) {
	var (
		length = len(this.element)
		high   = length - help.OneInt
	)

	return length, help.UnixAt(this.base, high)
}

func (this *linkedSegBucket[T]) getHighUnixTime() int64 {
	return help.UnixAt(this.base, len(this.element)-help.OneInt)
}

type logSegment[T common.UnixTimestamped] struct {
	head linkedSegBucket[T]
	tail *linkedSegBucket[T]
	step uint
}

func (this *logSegment[T]) rotateToWith(step uint, element T) {
	this.head.element = append(this.head.element[:help.ZeroInt], element)
	this.head.next = nil
	this.tail = &this.head
	this.step = step
}

func (this *logSegment[T]) append(element T) {
	var (
		tail     = this.tail
		bucket   = tail.element
		length   = len(bucket)
		capacity = cap(bucket)
	)

	if length == capacity {
		var (
			bucket = make([]T, help.OneInt, capacity)
			next   = &linkedSegBucket[T]{element: bucket}
		)

		next.base = &bucket[help.ZeroInt]
		bucket[help.ZeroInt] = element

		tail.next = next
		this.tail = next

		return
	}

	tail.element = append(bucket, element)
}

func (this *logSegment[T]) getHigh() (element T, ok bool) {
	var (
		bucket = this.tail.element
		length = len(bucket)
		high   = length - help.OneInt
	)

	if high < help.ZeroInt {
		return
	}

	return bucket[high], true
}

func (this *logSegment[T]) getHighOffsetPtr() (int, *T) {
	var (
		offset = help.ZeroInt
		bucket = &this.head
		base   *T
		length = help.ZeroInt
	)

	for bucket != nil {
		base = bucket.base
		length = len(bucket.element)
		offset += length
		bucket = bucket.next
	}

	offset--
	high := length - help.OneInt
	return offset, help.GetPtr(base, high)
}

func (this *logSegment[T]) getLowPtr() *T { return this.head.base }

func (this *logSegment[T]) getLowerOffsetMinusOnePtr(unixTime int64) (int, *T, bool) {
	var (
		offset = help.ZeroInt
		bucket = &this.head

		pb *T = nil
		pl    = help.ZeroInt
	)

	for bucket != nil {
		length, highUnixTime := bucket.getLengthHighUnixTime()

		if highUnixTime >= unixTime {
			suffix, ptr, ok := bucket.getLowerOffsetMinusOnePtr(unixTime)

			if !ok && pb != nil {
				return offset, help.GetPtr(pb, pl-help.OneInt), true
			}

			return offset + suffix, ptr, ok
		}

		offset += length
		pb = bucket.base
		pl = length
		bucket = bucket.next
	}

	return offset, help.GetPtr(pb, pl-help.OneInt), true
}

func (this *logSegment[T]) getUpperOffsetMinusOnePtr(unixTime int64) (int, *T, bool) {
	var (
		offset = help.ZeroInt
		bucket = &this.head

		pb *T = nil
		pl    = help.ZeroInt
	)

	for bucket != nil {
		length, highUnixTime := bucket.getLengthHighUnixTime()

		if highUnixTime > unixTime {
			suffix, ptr, ok := bucket.getUpperOffsetMinusOnePtr(unixTime)

			if !ok && pb != nil {
				return offset, help.GetPtr(pb, pl-help.OneInt), true
			}

			return offset + suffix, ptr, ok
		}

		offset += length
		pb = bucket.base
		pl = length
		bucket = bucket.next
	}

	return offset, help.GetPtr(pb, pl-help.OneInt), true
}

func (this *logSegment[T]) getUpperOffsetPtr(unixTime int64) (int, *T) {
	var (
		offset = help.ZeroInt
		bucket = &this.head
	)

	for bucket != nil {
		length, highUnixTime := bucket.getLengthHighUnixTime()

		if highUnixTime > unixTime {
			suffix, ptr := bucket.getUpperOffsetPtr(unixTime)
			return offset + suffix, ptr
		}

		offset += length
		bucket = bucket.next
	}

	return offset, nil
}

func (this *logSegment[T]) getLowerOffsetPtr(unixTime int64) (int, *T, bool) {
	var (
		offset = help.ZeroInt
		bucket = &this.head
	)

	for bucket != nil {
		length, highUnixTime := bucket.getLengthHighUnixTime()

		if highUnixTime >= unixTime {
			suffix, ptr, ok := bucket.getLowerOffsetPtr(unixTime)
			return offset + suffix, ptr, ok
		}

		offset += length
		bucket = bucket.next
	}

	return offset, nil, false
}

func (this *logSegment[T]) getOffsetPtr(offset int) *T {
	for bucket := &this.head; bucket != nil; bucket = bucket.next {
		length := len(bucket.element)

		if offset < length {
			return help.GetPtr(bucket.base, offset)
		}

		offset -= length
	}

	return nil
}

type slidingLog[T common.UnixTimestamped] struct {
	segment []logSegment[T]

	base, tail             *logSegment[T]
	shift, mask, stepShift uint
}

// TODO: element timestamps must be > than logWindow; bucket's element caps must remain constant
func NewLog[T common.UnixTimestamped](stepWindow, retention time.Duration, capPerLinkedNode int) common.Log[T] {
	var (
		stepNano      = uint(stepWindow.Nanoseconds())
		retentionNano = uint(retention.Nanoseconds())

		shift      = help.CeilPowerOfTwoShiftOnly(stepNano)
		logWindow  = help.CeilPowerOfTwo(retentionNano)
		windowStep = logWindow >> shift
		mask       = windowStep - help.OneUint
		segment    = make([]logSegment[T], help.OneInt, windowStep) // TODO: the smallest windowStep can be is one

		slidingLog = &slidingLog[T]{
			segment:   segment,
			shift:     shift,
			mask:      mask,
			stepShift: uint(bits.Len(mask)),
		}
	)

	slidingLog.base = &segment[help.ZeroInt]
	slidingLog.tail = slidingLog.base
	slidingLog.segment = segment[:windowStep]

	var (
		minCap = int(windowStep) * capPerLinkedNode

		element = make([]T, help.OneInt, minCap)
		base    = &element[help.ZeroInt]
	)

	for index := range slidingLog.segment {
		var (
			low     = index * capPerLinkedNode
			high    = low
			max     = low + capPerLinkedNode
			segment = slidingLog.getSegPtr(index)

			base    = help.GetPtr(base, low)
			element = unsafe.Slice(base, max-low)
		)

		segment.head.base = base
		segment.head.element = element[: high-low : max-low]

		segment.tail = &segment.head
		segment.step = uint(index)
	}

	return slidingLog
}

func (this *slidingLog[T]) getSegPtr(index int) *logSegment[T] { return help.GetPtr(this.base, index) }

func (this *slidingLog[T]) Append(element T) error {
	var (
		elemUnixTime = element.UnixNano()
		step         = uint(elemUnixTime) >> this.shift
		segment      = this.getSegPtr(int(step & this.mask))
	)

	if segment.step != step {
		segment.rotateToWith(step, element)
		this.tail = segment
		return nil
	}

	segment.append(element)
	return nil
}

func (this *slidingLog[T]) Tail() (element T, ok bool) { return this.tail.getHigh() }

func (this *slidingLog[T]) NewCursorBefore(endUnixTime int64) (common.Cursor[T], error) { // offline
	return newSlidingLogCursorBefore(this, endUnixTime)
}

func (this *slidingLog[T]) NewCursorAt(floorUnixTime int64) (common.Cursor[T], error) { // event
	return newSlidingLogCursorAt(this, floorUnixTime)
}

func (this *slidingLog[T]) NewCursorAfter(startUnixTime int64) (common.Cursor[T], error) { // online
	return newSlidingLogCursorAfter(this, startUnixTime)
}

type logPoint struct {
	step   uint
	offset int
}

type slidingLogCursor[T common.UnixTimestamped] struct {
	log         *slidingLog[T]
	head        logPoint
	endUnixTime int64

	cursor T
}

var (
	ErrLogUnsupported = errors.New("matt::kit::log: unsupported")
	ErrLogCorruption  = errors.New("matt::kit::log: corruption")
	ErrLogObsolete    = errors.New("matt::kit::log: obsolete")
	ErrLogNoData      = errors.New("matt::kit::log: no data")
)

func newSlidingLogCursorBefore[T common.UnixTimestamped](log *slidingLog[T], endUnixTime int64) (common.Cursor[T], error) {
	var (
		shift      = log.shift
		mask       = log.mask
		windowStep = help.OneUint << log.stepShift
	)

	{
		minSupportedUnixTime := int64((windowStep << help.OneUint) << shift)
		if endUnixTime < minSupportedUnixTime {
			return nil, ErrLogUnsupported
		}
	}

	var (
		step = uint(endUnixTime) >> shift
		end  = step - windowStep
	)

	for step > end {
		var (
			segment = log.getSegPtr(int(step & mask))
			got     = segment.step
		)

		if got > step {
			return nil, ErrLogCorruption
		}

		if got < step {
			step--
			continue
		}

		offset, ptr, ok := segment.getLowerOffsetMinusOnePtr(endUnixTime)
		if !ok {
			if segment.step > got {
				return nil, ErrLogObsolete
			}

			step--
			continue
		}

		element, ok := snapshotIfStepUnsafe(ptr, shift, step)
		if !ok {
			return nil, ErrLogObsolete
		}

		return &slidingLogCursor[T]{
			log: log, head: logPoint{step: step, offset: (offset - help.OneInt)}, cursor: element,
			endUnixTime: endUnixTime,
		}, nil
	}

	return nil, ErrLogNoData
}

func newSlidingLogCursorAt[T common.UnixTimestamped](log *slidingLog[T], floorUnixTime int64) (common.Cursor[T], error) {
	var (
		shift      = log.shift
		mask       = log.mask
		windowStep = help.OneUint << log.stepShift
	)

	{
		minSupportedUnixTime := int64((windowStep << help.OneUint) << shift)
		if floorUnixTime < minSupportedUnixTime {
			return nil, ErrLogUnsupported
		}
	}

	var (
		step = uint(floorUnixTime) >> shift
		end  = step - windowStep
	)

	for step > end {
		var (
			segment = log.getSegPtr(int(step & mask))
			got     = segment.step
		)

		if got > step {
			return nil, ErrLogCorruption
		}

		if got < step {
			step--
			continue
		}

		offset, ptr, ok := segment.getUpperOffsetMinusOnePtr(floorUnixTime)
		if !ok {
			if segment.step > got {
				return nil, ErrLogObsolete
			}

			step--
			continue
		}

		element, ok := snapshotIfStepUnsafe(ptr, shift, step)
		if !ok {
			return nil, ErrLogObsolete
		}

		return &slidingLogCursor[T]{
			log: log, head: logPoint{step: step, offset: (offset - help.OneInt)}, cursor: element,
			endUnixTime: max(floorUnixTime, (floorUnixTime + help.OneInt64)),
		}, nil
	}

	return nil, ErrLogNoData
}

func newSlidingLogCursorAfter[T common.UnixTimestamped](log *slidingLog[T], startUnixTime int64) (common.Cursor[T], error) {
	var (
		shift      = log.shift
		mask       = log.mask
		windowStep = help.OneUint << log.stepShift
	)

	{
		minSupportedUnixTime := int64((windowStep << help.OneUint) << shift)
		if startUnixTime < minSupportedUnixTime {
			return nil, ErrLogUnsupported
		}
	}

	var (
		step = uint(startUnixTime) >> shift
		end  = step + windowStep
	)

	for step < end {
		var (
			segment = log.getSegPtr(int(step & mask))
			got     = segment.step
		)

		if got > step {
			return nil, ErrLogCorruption
		}

		if got < step {
			step++
			continue
		}

		highUnixTime := segment.tail.getHighUnixTime()
		if highUnixTime <= startUnixTime {
			step++
			continue
		}

		offset, ptr := segment.getUpperOffsetPtr(startUnixTime)
		element, ok := snapshotIfStepUnsafe(ptr, shift, step)
		if !ok {
			return nil, ErrLogObsolete
		}

		return &slidingLogCursor[T]{
			log: log, head: logPoint{step: step, offset: offset}, cursor: element,
			endUnixTime: help.MaxInt64,
		}, nil
	}

	return nil, ErrLogNoData
}

func snapshotIfStepUnsafe[T common.UnixTimestamped](ptr *T, shift, want uint) (element T, ok bool) {
	element = *ptr
	elemUnixTime := element.UnixNano()
	return element, (uint(elemUnixTime) >> shift) == want
}

func snapshotIfStep[T common.UnixTimestamped](ptr *T, shift, want uint) (element T, ok bool) {
	if ptr == nil {
		return
	}

	element = *ptr
	elemUnixTime := element.UnixNano()
	return element, (uint(elemUnixTime) >> shift) == want
}

func (this *slidingLogCursor[T]) Cursor() T { return this.cursor }

func (this *slidingLogCursor[T]) Seek(pointUnixTime int64) error {
	var (
		log     = this.log
		shift   = log.shift
		step    = uint(pointUnixTime) >> shift
		segment = log.getSegPtr(int(step & log.mask))
		got     = segment.step
	)

	if got > step {
		return ErrLogCorruption
	}

	if got < step {
		return ErrLogNoData
	}

	offset, ptr, ok := segment.getLowerOffsetPtr(pointUnixTime)
	if !ok {
		return ErrLogNoData
	}

	element, ok := snapshotIfStepUnsafe(ptr, shift, step)
	if !ok {
		return ErrLogObsolete
	}

	elemUnixTime := element.UnixNano()
	if elemUnixTime > pointUnixTime {
		return ErrLogNoData
	}

	this.cursor = element
	this.head = logPoint{step: step, offset: offset}
	return nil
}

func (this *slidingLogCursor[T]) Start() error {
	var (
		log        = this.log
		mask       = log.mask
		windowStep = help.OneUint << log.stepShift
		tail       = log.tail
		ts         = tail.step
		step       = (ts - windowStep) + (windowStep >> help.FiveUint)
	)

	for step < ts {
		var (
			segment = log.getSegPtr(int(step & mask))
			got     = segment.step
		)

		if got > step {
			return ErrLogCorruption
		}

		if got < step {
			step++
			continue
		}

		ptr := segment.getLowPtr()
		element, ok := snapshotIfStepUnsafe(ptr, log.shift, step)
		if !ok {
			return ErrLogObsolete
		}

		this.cursor = element
		this.head = logPoint{step: step, offset: help.ZeroInt}
		return nil
	}

	ptr := tail.getLowPtr()
	element, ok := snapshotIfStepUnsafe(ptr, log.shift, ts)
	if !ok {
		return ErrLogObsolete
	}

	this.cursor = element
	this.head = logPoint{step: ts, offset: help.ZeroInt}
	return nil
}

func (this *slidingLogCursor[T]) Previous() bool {
	var (
		log   = this.log
		shift = log.shift
		head  = this.head
	)

	if head.offset > help.ZeroInt {
		var (
			segment = log.getSegPtr(int(head.step & log.mask))
			offset  = head.offset - help.OneInt
		)

		ptr := segment.getOffsetPtr(offset)
		element, ok := snapshotIfStep(ptr, shift, head.step)
		if !ok {
			return false
		}

		this.cursor = element
		this.head.offset = offset
		return true
	}

	var (
		mask       = log.mask
		windowStep = help.OneUint << log.stepShift
		step       = head.step - help.OneUint
		end        = step - (windowStep - help.OneUint)
	)

	for step > end {
		var (
			segment = log.getSegPtr(int(step & mask))
			got     = segment.step
		)

		if got > step {
			return false
		}

		if got < step {
			step--
			continue
		}

		offset, ptr := segment.getHighOffsetPtr()
		element, ok := snapshotIfStepUnsafe(ptr, shift, step)
		if !ok {
			return false
		}

		this.cursor = element
		this.head = logPoint{step: step, offset: offset}
		return true
	}

	return false
}

func (this *slidingLogCursor[T]) Next() (bool, error) {
	var (
		log         = this.log
		shift       = log.shift
		mask        = log.mask
		head        = this.head
		endUnixTime = this.endUnixTime
	)

	{
		segment := log.getSegPtr(int(head.step & mask))
		if segment.step > head.step {
			return false, ErrLogCorruption
		}

		offset := head.offset + help.OneInt
		ptr := segment.getOffsetPtr(offset)
		if ptr != nil {
			element, ok := snapshotIfStepUnsafe(ptr, shift, head.step)
			if !ok {
				return false, ErrLogObsolete
			}

			elemUnixTime := element.UnixNano()
			if elemUnixTime >= endUnixTime {
				return false, ErrLogNoData
			}

			this.cursor = element
			this.head.offset = offset
			return true, nil
		}

		if segment.step > head.step {
			return false, ErrLogObsolete
		}
	}

	var (
		step = head.step + help.OneUint
		end  = min((log.tail.step + help.OneUint), ((uint(endUnixTime) >> shift) + help.OneUint))
	)

	for step < end {
		var (
			segment = log.getSegPtr(int(step & mask))
			got     = segment.step
		)

		if got > step {
			return false, ErrLogCorruption
		}

		if got < step {
			step++
			continue
		}

		ptr := segment.getLowPtr()
		element, ok := snapshotIfStepUnsafe(ptr, shift, step)
		if !ok {
			return false, ErrLogObsolete
		}

		elemUnixTime := element.UnixNano()
		if elemUnixTime >= endUnixTime {
			return false, ErrLogNoData
		}

		this.cursor = element
		this.head = logPoint{step: step, offset: help.ZeroInt}
		return true, nil
	}

	return false, nil
}

func (this *slidingLogCursor[T]) End() error {
	var (
		log         = this.log
		shift       = log.shift
		mask        = log.mask
		endUnixTime = this.endUnixTime
		step        = min(log.tail.step, (uint(endUnixTime) >> shift))
		max         = max(this.head.step, (step - (help.OneUint << log.stepShift) + help.OneUint))
	)

	for step >= max {
		var (
			segment = log.getSegPtr(int(step & mask))
			got     = segment.step
		)

		if got > step {
			return ErrLogCorruption
		}

		if got < step {
			step--
			continue
		}

		offset, ptr, ok := segment.getLowerOffsetMinusOnePtr(endUnixTime)
		if !ok {
			if segment.step > got {
				return ErrLogObsolete
			}

			step--
			continue
		}

		element, ok := snapshotIfStepUnsafe(ptr, shift, step)
		if !ok {
			return ErrLogObsolete
		}

		this.cursor = element
		this.head = logPoint{step: step, offset: (offset - help.OneInt)}
		return nil
	}

	return ErrLogObsolete // TODO: is this the correct error?
}

func (this *slidingLogCursor[T]) Point() int64 { return this.cursor.UnixNano() }
