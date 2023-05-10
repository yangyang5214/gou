package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	q := NewQueue()
	q.Push("11")
	q.Push(22)
	q.Push("77")

	assert.Equal(t, q.Pop(), "11")
	assert.Equal(t, q.Pop(), 22)

	assert.Equal(t, q.Peek(), "77")

	var a string
	a = q.Peek().(string)
	t.Log(a)

	assert.Equal(t, q.Pop(), "77")
}
