package list

import (
	"testing"
)

type People struct {
	name string
}

func TestName(t *testing.T) {
	l := New[*People]()
	l.PushFront(&People{
		name: "beer",
	})

	var p *People
	p = l.Back().Value

	t.Log(p.name)
}
