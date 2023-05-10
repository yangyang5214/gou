package stack

import (
	"sync"

	"github.com/yangyang5214/gou/list"
)

// Taken from https://stackoverflow.com/a/64641330/9546749

type Stack[T any] struct {
	ll    *list.List[T]
	mutex sync.Mutex
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		ll:    list.New[T](),
		mutex: sync.Mutex{},
	}
}

func (s *Stack[T]) Push(x T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.ll.PushBack(x)
}

func (s *Stack[T]) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.ll.Len()
}

func (s *Stack[T]) Pop() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.ll.Len() == 0 {
		var zeroVal T // 使用 T 类型的零值变量
		return zeroVal, false
	}
	tail := s.ll.Back()
	val := tail.Value
	s.ll.Remove(tail)
	return val, true
}
