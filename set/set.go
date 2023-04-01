package set

// Set is a collection of unique elements.
type Set[T comparable] struct {
	data map[T]bool
}

// NewSet creates a new Set instance.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]bool),
	}
}

// Add adds an element to the set.
func (s *Set[T]) Add(element T) {
	s.data[element] = true
}

// Remove removes an element from the set.
func (s *Set[T]) Remove(element T) {
	delete(s.data, element)
}

// Contains checks if an element exists in the set.
func (s *Set[T]) Contains(element T) bool {
	_, exists := s.data[element]
	return exists
}

// Size returns the number of elements in the set.
func (s *Set[T]) Size() int {
	return len(s.data)
}

// Clear removes all elements from the set.
func (s *Set[T]) Clear() {
	s.data = make(map[T]bool)
}

// Elements returns a slice containing all elements in the set.
func (s *Set[T]) Elements() []T {
	elements := make([]T, 0, len(s.data))
	for element := range s.data {
		elements = append(elements, element)
	}
	return elements
}

// Union returns the union of two sets.
func Union[T comparable](a, b *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range a.data {
		result.Add(element)
	}
	for element := range b.data {
		result.Add(element)
	}
	return result
}

// Intersection returns the intersection of two sets.
func Intersection[T comparable](a, b *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range a.data {
		if b.Contains(element) {
			result.Add(element)
		}
	}
	return result
}

// Difference returns the difference of two sets (elements in a but not in b).
func Difference[T comparable](a, b *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range a.data {
		if !b.Contains(element) {
			result.Add(element)
		}
	}
	return result
}
