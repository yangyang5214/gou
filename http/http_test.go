package httputil_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yangyang5214/gou/http"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	httpClient := httputil.NewClient(httputil.DefaultOptions)
	r, err := httpClient.Get("https://github.com/")
	assert.True(t, err == nil)
	assert.True(t, r.StatusCode == http.StatusOK)

	r, err = httpClient.Get("https://10.0.90.50/api/alarm/global/count")
	assert.True(t, r.StatusCode == http.StatusUnauthorized)
}

func TestGetWithHeaders(t *testing.T) {
	httpClient := httputil.NewClient(&httputil.Options{
		SkipVerify: true,
		Headers: []httputil.Header{
			{
				Key:   "token",
				Value: "15f5e828a3ee4274c6066be4eeeae0ec",
			},
		},
	})
	r, err := httpClient.Get("https://10.0.90.50/api/alarm/global/count")
	if err != nil {
		t.Error(err)
	}
	assert.True(t, r.StatusCode == http.StatusOK)
}
