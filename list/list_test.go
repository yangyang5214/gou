package list

import (
	"testing"
)

type People struct {
	name string
}

type NewType[T int | *People] struct {
	value T
}

func TestName(t *testing.T) {
	l := New[*People]()
	l.PushFront(&People{
		name: "beer",
	})

	var p *People
	p = l.Back().Value

	t.Log(p.name)

	l2 := New[*NewType[int]]()

	l2.PushBack(&NewType[int]{
		value: 200,
	})
	l2.PushBack(&NewType[int]{
		value: 100,
	})
	t.Log(l2.Len())
	t.Log(l2.Back().Value) // 100
	l2.Remove(l2.Back())
	t.Log(l2.Back().Value) //200
}
