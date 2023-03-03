package urlutil_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yangyang5214/gou/url"
	"testing"
)

func Test(t *testing.T) {
	r := urlutil.GetIpsByDomain("https://www.baidu.com/")
	assert.True(t, len(r) != 0)

	r = urlutil.GetIpsByDomain("/www.baidu.com/")
	assert.True(t, len(r) == 0)
}
