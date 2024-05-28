package main

import (
	"fmt"
	"sync"
)

type Synced[K comparable, V any] struct { // Generic types for key and value of map
	data map[K]V
	mu   sync.RWMutex // Mutex locks for others to read or write or both.
}

func (s *Synced[K, V]) Set(key K, val V) { // For changing something in map, should use Lock for write lock.
	// Lock locks for reading and writing.
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = val
}
func (s *Synced[K, V]) Get(key K) (V, bool) { // For reading it's good to read lock.
	// Rlock doesn't lock for reading but locks for writing.
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	return val, ok
}
func (s *Synced[K, V]) Delete(key K) { // For deleting value by key, and setting presence variable to false.
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

func New[K comparable, V any]() *Synced[K, V] {
	return &Synced[K, V]{
		data: make(map[K]V),
	}
}

func main() {
	wg := sync.WaitGroup{}
	cache := New[string, int]()

	cache.Set("One", 1)
	fmt.Println(cache.data)

	wg.Add(2)
	go func() {
		defer wg.Done()
		cache.Set("One", 2)
	}()
	go func() {
		defer wg.Done()
		cache.Set("Three", 3)
	}()
	wg.Wait() // Waiting for all goroutines to end.

	fmt.Println(cache.data)
}
