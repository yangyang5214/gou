package stack

import (
	"container/list"
	"sync"
)

// Taken from https://stackoverflow.com/a/64641330/9546749

type Stack struct {
	ll    *list.List
	mutex sync.Mutex
}

func NewStack() *Stack {
	return &Stack{
		ll:    list.New(),
		mutex: sync.Mutex{},
	}
}

func (s *Stack) Push(x any) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.ll.PushBack(x)
}

func (s *Stack) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.ll.Len()
}

func (s *Stack) Pop() any {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.ll.Len() == 0 {
		return nil
	}
	tail := s.ll.Back()
	val := tail.Value
	s.ll.Remove(tail)
	return val
}
