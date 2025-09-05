package kit

import (
	"errors"
	"iter"
	"sync"
)

type CoarseMap[K comparable, V any] struct {
	mu      sync.RWMutex
	element map[K]V
}

var ErrMapNotFound = errors.New("matt::kit::map: not found")

func NewCoarseMap[K comparable, V any](capacity int) *CoarseMap[K, V] {
	return &CoarseMap[K, V]{
		element: make(map[K]V, capacity),
	}
}

func InitCoarseMap[K comparable, V any](m *CoarseMap[K, V], capacity int) {
	m.element = make(map[K]V, capacity)
}

func (this *CoarseMap[K, V]) Set(key K, value V) error {
	this.mu.Lock()
	this.element[key] = value
	this.mu.Unlock()
	return nil
}

func (this *CoarseMap[K, V]) Get(key K) (V, error) {
	this.mu.RLock()
	value, ok := this.element[key]
	this.mu.RUnlock()
	if !ok {
		return value, ErrMapNotFound
	}
	return value, nil
}

func (this *CoarseMap[K, V]) Delete(key K) error {
	this.mu.Lock()
	delete(this.element, key)
	this.mu.Unlock()
	return nil
}

func (this *CoarseMap[K, V]) ForEach(yield func(K, V) bool) {
	this.mu.RLock()
	for key, value := range this.element {
		if !yield(key, value) {
			break
		}
	}

	this.mu.RUnlock()
}

func (this *CoarseMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		this.mu.RLock()
		for key, value := range this.element {
			if !yield(key, value) {
				break
			}
		}

		this.mu.RUnlock()
	}
}
