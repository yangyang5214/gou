package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	q := NewStack()
	q.Push("11")
	q.Push(22)
	q.Push("77")

	assert.Equal(t, q.Pop(), "77")
	assert.Equal(t, q.Pop(), 22)
	assert.Equal(t, q.Pop(), "11")
	assert.Equal(t, q.Pop(), nil)
}
