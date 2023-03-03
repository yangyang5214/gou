package urlutil_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yangyang5214/gou/url"
	"testing"
)

func TestGetHtmlTitle(t *testing.T) {
	assert.Equal(t, urlutil.GetHtmlTitle("<title>beer</title>"), "beer")
	assert.Equal(t, urlutil.GetHtmlTitle("<title>beer"), "")
}

func TestAppendSchema(t *testing.T) {
	assert.Equal(t, urlutil.AppendSchema("www.baidu.com", urlutil.HTTP), "http://www.baidu.com")
	assert.Equal(t, urlutil.AppendSchema("www.baidu.com", urlutil.HTTPS), "https://www.baidu.com")
}

func TestGetUrlSuffix(t *testing.T) {
	assert.Equal(t, urlutil.GetUrlSuffix("www.baidu.com/111.png"), "png")
}

func TestRemoveQuery(t *testing.T) {
	assert.Equal(t, urlutil.RemoveQuery("https://www.baidu.com/111.png?1=1"), "https://www.baidu.com/111.png")
}

func TestGetUrlHost(t *testing.T) {
	assert.Equal(t, urlutil.GetUrlHost("111"), "111")
	assert.Equal(t, urlutil.GetUrlHost("https://www.baidu.com/111.png?1=1"), "www.baidu.com")
	assert.Equal(t, urlutil.GetUrlHost("baidu.com/111.png?1=1"), "baidu.com")
	assert.Equal(t, urlutil.GetUrlHost("www.baidu.com/111.png?1=1"), "www.baidu.com")
	assert.Equal(t, urlutil.GetUrlHost("111"), "111")
}
