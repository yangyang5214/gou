package urlutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	r := GetIpsByDomain("https://www.baidu.com/")
	assert.True(t, len(r) != 0)

	r = GetIpsByDomain("/www.baidu.com/")
	assert.True(t, len(r) == 0)

	r = GetIpsByDomain("http/www.baidu.com/")
	assert.True(t, len(r) == 0)
}
