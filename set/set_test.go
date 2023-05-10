package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	assert.Equal(t, 2, s.Size())
}
