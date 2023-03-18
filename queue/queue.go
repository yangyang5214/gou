package queue

import (
	"container/list"
	"sync"
)

// Taken from https://stackoverflow.com/a/64641330/9546749

type Queue struct {
	ll    *list.List
	mutex sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		ll:    list.New(),
		mutex: sync.Mutex{},
	}
}

func (s *Queue) Push(v any) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ll.PushFront(v)
}

// Peek method returns the element at the front the container.
// It does not delete the element in the container.
func (s *Queue) Peek() any {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	front := s.ll.Front()
	if front != nil {
		return front.Value
	}
	return nil
}

// Len returns len of the Queue
func (s *Queue) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.ll.Len()
}

func (s *Queue) Pop() any {
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
