// Пакет lrucache предоставляет реализацию LRU кэша
package lrucache

import (
	"container/list"
	"sync"
)

type entry[K comparable, V any] struct {
	key K
	val V
}

type LRUCache[K comparable, V any] struct {
	mu       sync.Mutex
	capacity int
	list     *list.List
	cache    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	capacity = max(capacity, 1)

	return &LRUCache[K, V]{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[K]*list.Element, capacity),
	}
}

func (s *LRUCache[K, V]) Get(key K) (V, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if e, ok := s.cache[key]; ok {
		s.list.MoveToFront(e)
		return e.Value.(*entry[K, V]).val, true
	}
	var v V
	return v, false
}

func (s *LRUCache[K, V]) Put(key K, value V) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if e, ok := s.cache[key]; ok {
		e.Value.(*entry[K, V]).val = value
		s.list.MoveToFront(e)
		return
	}

	e := s.list.PushFront(&entry[K, V]{key, value})
	s.cache[key] = e
	if s.list.Len() > s.capacity {
		s.removeOldest()
	}
}

func (s *LRUCache[K, V]) Delete(key K) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.cache[key]
	if !ok {
		return
	}

	s.list.Remove(e)
	delete(s.cache, key)
	return
}

func (s *LRUCache[K, V]) removeOldest() {
	e := s.list.Back()
	if e == nil {
		return
	}
	s.list.Remove(e)
	delete(s.cache, e.Value.(*entry[K, V]).key)
}
