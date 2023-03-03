package httputil

import (
	"time"
)

// Options contains configuration options for http client
type Options struct {
	Timeout    time.Duration
	SkipVerify bool
	UserAgent  string
	Host       string
	Headers    []Header
}

type Header struct {
	Key   string
	Value string
}

// DefaultOptions is the default configuration options for the client
var DefaultOptions = &Options{
	Timeout:    30 * time.Second,
	SkipVerify: true,
}
