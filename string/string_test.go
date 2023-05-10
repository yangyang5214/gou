package stringutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	stringutil "github.com/yangyang5214/gou/string"
)

func TestStringToBytes(t *testing.T) {
	r := stringutil.StringToBytes("beer")
	t.Log(r)
}

func TestName(t *testing.T) {
	r := stringutil.StringToBytes("beer")
	assert.Equal(t, stringutil.BytesToString(r), "beer")
}
