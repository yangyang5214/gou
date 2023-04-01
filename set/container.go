package set

type Container[T comparable] interface {
	Add(T)
	Remove(T)
	Contains(T) bool
	Size() int
	Clear()
	Elements() []T
}
