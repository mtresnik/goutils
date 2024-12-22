package goutils

import (
	"sync"
)

type SyncList[T any] struct {
	mutex sync.Mutex
	items []T
}

func NewSyncList[T any]() *SyncList[T] {
	return &SyncList[T]{
		items: make([]T, 0),
	}
}

func (s *SyncList[T]) Add(item T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.items = append(s.items, item)
}

func (l *SyncList[T]) Remove(item T) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	for i, v := range l.items {
		if any(v) == any(item) {
			l.items = append(l.items[:i], l.items[i+1:]...)
			return true
		}
	}
	return false
}

func (l *SyncList[T]) Contains(item T) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	for _, v := range l.items {
		if any(v) == any(item) {
			return true
		}
	}
	return false
}

func (l *SyncList[T]) Len() int {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return len(l.items)
}

func (l *SyncList[T]) Get(index int) (T, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if index < 0 || index >= len(l.items) {
		var zero T
		return zero, false
	}
	return l.items[index], true
}

func (l *SyncList[T]) Clear() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.items = make([]T, 0)
}

func (l *SyncList[T]) ToSlice() []T {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	result := make([]T, len(l.items))
	copy(result, l.items)
	return result
}
