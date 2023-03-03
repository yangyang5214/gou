package httputil

import (
	"crypto/tls"
	"github.com/yangyang5214/gou/type"
	"io"
	"net/http"
)

type HttpClient struct {
	options *Options
	client  *http.Client
}

func NewClient(options *Options) *HttpClient {
	var tr http.Transport
	if options.SkipVerify {
		tr = http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	return &HttpClient{
		client: &http.Client{
			Transport: &tr,
			Timeout:   options.Timeout,
		},
		options: options,
	}
}

func (client *HttpClient) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client.buildReq(req)
	return client.client.Do(req)
}

func (client *HttpClient) buildReq(req *http.Request) {
	req.Host = client.options.Host
	req.Header.Set("User-Agent", client.options.UserAgent)
	for _, item := range client.options.Headers {
		req.Header.Set(item.Key, item.Value)
	}
}

func (client *HttpClient) Post(url, contentType typeutil.ContentType, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	client.buildReq(req)
	req.Header.Set("Content-Type", contentType)
	return client.client.Do(req)
}

func (client *HttpClient) ReadBody(resp *http.Response) (data []byte, err error) {
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
