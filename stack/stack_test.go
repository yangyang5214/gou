package stack

import (
	"testing"
)

func TestName(t *testing.T) {
	q := NewStack[string]()
	q.Push("11")
	q.Push("77")
}
