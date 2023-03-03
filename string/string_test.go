package stringutil_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yangyang5214/gou/string"
	"testing"
)

func TestStringToBytes(t *testing.T) {
	r := stringutil.StringToBytes("beer")
	t.Log(r)
}

func TestName(t *testing.T) {
	r := stringutil.StringToBytes("beer")
	assert.Equal(t, stringutil.BytesToString(r), "beer")
}
