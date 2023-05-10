package slices

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	a := []string{"demo", "demo1"}
	assert.True(t, Contains(a, "demo"))
}

func TestIndex(t *testing.T) {
	a := []string{"demo", "demo1"}
	assert.Equal(t, Index(a, "demo"), 0)
	assert.Equal(t, Index(a, "demo1"), 1)
}

func TestConvertSlice(t *testing.T) {
	ages := []int{1, 2}
	r := ConvertSlice[int, string](ages, func(item int) string {
		return strconv.Itoa(item)
	})
	assert.Equal(t, r[0], "1")
	assert.Equal(t, r[1], strconv.Itoa(2))
}

type TestPeople struct {
	name string
	age  int32
}

func TestForeachFind(t *testing.T) {
	peoples := []TestPeople{
		{
			name: "beef",
			age:  20,
		},
		{
			name: "beer",
			age:  1,
		},
	}
	ok, v := ForeachFind(peoples, func(v TestPeople) (bool, TestPeople) {
		if v.age < 10 && strings.Contains(v.name, "bee") {
			return true, v
		}
		return false, TestPeople{}
	})
	assert.True(t, ok, true)
	assert.Equal(t, v.name, "beer")
	assert.Equal(t, v.age, int32(1))
}
