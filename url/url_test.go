package urlutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHtmlTitle(t *testing.T) {
	assert.Equal(t, GetHtmlTitle("<title>beer</title>"), "beer")
	assert.Equal(t, GetHtmlTitle("<title>beer"), "")
}

func TestAppendSchema(t *testing.T) {
	assert.Equal(t, AppendSchema("www.baidu.com", HTTP), "http://www.baidu.com")
	assert.Equal(t, AppendSchema("www.baidu.com", HTTPS), "https://www.baidu.com")
}

func TestGetUrlSuffix(t *testing.T) {
	assert.Equal(t, GetUrlSuffix("www.baidu.com/111.png"), "png")
}

func TestRemoveQuery(t *testing.T) {
	assert.Equal(t, RemoveQuery("https://www.baidu.com/111.png?1=1"), "https://www.baidu.com/111.png")
}

func TestGetUrlHost(t *testing.T) {
	assert.Equal(t, GetUrlHost("111"), "111")
	assert.Equal(t, GetUrlHost("https://www.baidu.com/111.png?1=1"), "www.baidu.com")
	assert.Equal(t, GetUrlHost("baidu.com/111.png?1=1"), "baidu.com")
	assert.Equal(t, GetUrlHost("www.baidu.com/111.png?1=1"), "www.baidu.com")
	assert.Equal(t, GetUrlHost("111"), "111")
}
