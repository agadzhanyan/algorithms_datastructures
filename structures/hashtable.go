package structures

import (
	"sync"
)

type HashTable struct {
	elements map[int]string
	sync sync.RWMutex
}

// хеширование по правилу Горнера
func (s *HashTable) hash(key string) int {
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31 * h + int(key[i])
	}
	return h
}

func (s *HashTable) add(key, value string) {
	hashedKey := s.hash(key)
	s.sync.Lock()
	defer s.sync.Unlock()
	s.elements[hashedKey] = value
}

func (s *HashTable) Get(key string) string {
	s.sync.RLock()
	defer s.sync.RUnlock()
	v, ok := s.elements[s.hash(key)]
	if !ok {
		return ""
	}
	return v
}

func (s *HashTable) Delete(key string) {
	s.sync.Lock()
	defer s.sync.Unlock()
	delete(s.elements, s.hash(key))
}

func (s *HashTable) Size() int {
	return len(s.elements)
}
