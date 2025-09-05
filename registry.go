package kit

import (
	"errors"
	"iter"
	"sync"
)

type registryEntry[V any] struct {
	claimed bool
	value   V
}

type CoarseRegistry[K comparable, V any] struct {
	mu    sync.RWMutex
	entry map[K]registryEntry[V]
}

var (
	ErrRegistryConflict = errors.New("matt::kit::registry: key conflict")
	ErrRegistryNotFound = errors.New("matt::kit::registry: not found")
	ErrRegistryLocked   = errors.New("matt::kit::registry: locked")
)

func NewCoarseRegistry[K comparable, V any](capacity int) *CoarseRegistry[K, V] {
	return &CoarseRegistry[K, V]{
		entry: make(map[K]registryEntry[V], capacity),
	}
}

func InitCoarseRegistry[K comparable, V any](registry *CoarseRegistry[K, V], capacity int) {
	registry.entry = make(map[K]registryEntry[V], capacity)
}

func (this *CoarseRegistry[K, V]) Claim(key K) (V, bool) {
	this.mu.Lock()
	entry, ok := this.entry[key]
	if !ok || entry.claimed {
		this.mu.Unlock()
		return entry.value, false
	}

	entry.claimed = true
	this.entry[key] = entry
	this.mu.Unlock()
	return entry.value, true
}

func (this *CoarseRegistry[K, V]) Register(key K, value V) error {
	this.mu.Lock()
	_, ok := this.entry[key]
	if ok {
		this.mu.Unlock()
		return ErrRegistryConflict
	}

	this.entry[key] = registryEntry[V]{claimed: true, value: value}
	this.mu.Unlock()
	return nil
}

func (this *CoarseRegistry[K, V]) Release(key K) error {
	this.mu.Lock()
	entry, ok := this.entry[key]
	if !ok {
		this.mu.Unlock()
		return ErrRegistryNotFound
	}

	entry.claimed = false
	this.entry[key] = entry
	this.mu.Unlock()
	return nil
}

func (this *CoarseRegistry[K, V]) Get(key K) (V, error) {
	this.mu.RLock()
	entry, ok := this.entry[key]
	this.mu.RUnlock()
	if !ok {
		return entry.value, ErrRegistryNotFound
	}
	return entry.value, nil
}

func (this *CoarseRegistry[K, V]) Delete(key K) error {
	this.mu.Lock()
	entry, ok := this.entry[key]
	if !ok {
		this.mu.Unlock()
		return ErrRegistryNotFound
	}

	if entry.claimed {
		this.mu.Unlock()
		return ErrRegistryLocked
	}

	delete(this.entry, key)
	this.mu.Unlock()
	return nil
}

func (this *CoarseRegistry[K, V]) ForEach(yield func(K, V) bool) {
	this.mu.RLock()
	for key, entry := range this.entry {
		if !yield(key, entry.value) {
			break
		}
	}

	this.mu.RUnlock()
}

func (this *CoarseRegistry[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		this.mu.RLock()
		for key, entry := range this.entry {
			if !yield(key, entry.value) {
				break
			}
		}

		this.mu.RUnlock()
	}
}
