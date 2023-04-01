package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSyncSet(t *testing.T) {
	s := NewSyncSet()

	go func() {
		for i := 0; i < 1000; i++ {
			s.Add(i)
		}
	}()

	go func() {
		for i := 500; i < 1500; i++ {
			s.Remove(i)
		}
	}()

	time.Sleep(time.Second * 3)
	assert.Equal(t, 1000, s.Size())
}
