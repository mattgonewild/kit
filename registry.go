package kit

import (
	"errors"
	"sync"
)

func NewRegistryAs() bool { return true }

type registryEntry[V any] struct {
	claimed bool
	value   V
}

type CoarseRegistry[K comparable, V any] struct {
	mu    sync.RWMutex
	entry map[K]registryEntry[V]
}

var ErrRegistryNotFound = errors.New("matt::kit::registry: not found")

func NewCoarseRegistry[K comparable, V any](capacity int) *CoarseRegistry[K, V] {
	return &CoarseRegistry[K, V]{
		entry: make(map[K]registryEntry[V], capacity),
	}
}

func InitCoarseRegistry[K comparable, V any](registry *CoarseRegistry[K, V], capacity int) {
	registry.entry = make(map[K]registryEntry[V], capacity)
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

func (this *CoarseRegistry[K, V]) Release(key K) bool {
	this.mu.Lock()
	entry, ok := this.entry[key]
	if !ok {
		this.mu.Unlock()
		return false
	}

	entry.claimed = false
	this.entry[key] = entry
	this.mu.Unlock()
	return true
}

func (this *CoarseRegistry[K, V]) Claim(key K) bool {
	this.mu.Lock()
	entry, ok := this.entry[key]
	if !ok {
		this.mu.Unlock()
		return false
	}

	if entry.claimed {
		this.mu.Unlock()
		return false
	}

	entry.claimed = true
	this.entry[key] = entry
	this.mu.Unlock()
	return true
}

func (this *CoarseRegistry[K, V]) ClaimOrRegister(key K, new func() V) (V, bool) {
	this.mu.Lock()
	entry, ok := this.entry[key]
	if !ok {
		entry = registryEntry[V]{claimed: true, value: new()}
		this.entry[key] = entry
		this.mu.Unlock()
		return entry.value, true
	}

	if entry.claimed {
		this.mu.Unlock()
		return entry.value, false
	}

	entry.claimed = true
	this.entry[key] = entry
	this.mu.Unlock()
	return entry.value, true
}
