package kit

import (
	"errors"
	"iter"
	"sync"
)

type CoarseMap[K comparable, V any] struct {
	sync.RWMutex
	element map[K]V
}

var ErrMapNotFound = errors.New("matt::kit::map: not found")

func NewCoarseMap[K comparable, V any](capacity int) *CoarseMap[K, V] {
	return &CoarseMap[K, V]{
		element: make(map[K]V, capacity),
	}
}

func (this *CoarseMap[K, V]) Set(key K, value V) error {
	this.Lock()
	this.element[key] = value
	this.Unlock()
	return nil
}

func (this *CoarseMap[K, V]) Get(key K) (V, error) {
	this.RLock()
	value, ok := this.element[key]
	this.RUnlock()
	if !ok {
		return value, ErrMapNotFound
	}
	return value, nil
}

func (this *CoarseMap[K, V]) Delete(key K) error {
	this.Lock()
	delete(this.element, key)
	this.Unlock()
	return nil
}

func (this *CoarseMap[K, V]) ForEach(yield func(K, V) bool) {
	this.RLock()
	for key, value := range this.element {
		if !yield(key, value) {
			break
		}
	}

	this.RUnlock()
}

func (this *CoarseMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		this.RLock()
		for key, value := range this.element {
			if !yield(key, value) {
				break
			}
		}

		this.RUnlock()
	}
}
