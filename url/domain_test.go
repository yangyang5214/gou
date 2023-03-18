package urlutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	r := GetIpsByDomain("https://www.baidu.com/")
	assert.True(t, len(r) != 0)

	r = GetIpsByDomain("/www.baidu.com/")
	assert.True(t, len(r) == 0)
}
