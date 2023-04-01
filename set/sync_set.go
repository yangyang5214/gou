package set

import "sync"

// SyncSet is a collection of unique elements.
type SyncSet struct {
	data sync.Map
}

// NewSyncSet creates a new SyncSet instance.
func NewSyncSet() *SyncSet {
	return &SyncSet{
		data: sync.Map{},
	}
}

// Add adds an element to the SyncSet.
func (s *SyncSet) Add(element interface{}) {
	s.data.Store(element, true)
}

// Remove removes an element from the SyncSet.
func (s *SyncSet) Remove(element interface{}) {
	s.data.Delete(element)
}

// Contains checks if an element exists in the SyncSet.
func (s *SyncSet) Contains(element interface{}) bool {
	_, exists := s.data.Load(element)
	return exists
}

// Size returns the number of elements in the SyncSet.
func (s *SyncSet) Size() int {
	size := 0
	s.data.Range(func(_, _ interface{}) bool {
		size++
		return true
	})
	return size
}

// Clear removes all elements from the SyncSet.
func (s *SyncSet) Clear() {
	s.data = sync.Map{}
}

// Elements returns a slice containing all elements in the SyncSet.
func (s *SyncSet) Elements() []interface{} {
	elements := make([]interface{}, 0, s.Size())
	s.data.Range(func(key, _ interface{}) bool {
		elements = append(elements, key)
		return true
	})
	return elements
}
