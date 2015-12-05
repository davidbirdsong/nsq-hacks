package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/nsqio/nsq/internal/http_api"
	"github.com/nsqio/nsq/internal/version"
)

var httpclient *http.Client
var userAgent string

func init() {
	httpclient = &http.Client{
		Transport:     http_api.NewDeadlineTransport(*httpTimeout),
		CheckRedirect: boo,
	}
	userAgent = fmt.Sprintf("nsq_to_http v%s", version.Binary)
}
func boo(r *http.Request, v []*http.Request) error {
	log.Println("got redirect")
	return nil
}

func HTTPGet(host string, endpoint *url.URL) (*http.Response, error) {
	req, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	req.Host = host
	return httpclient.Do(req)
}

func HTTPPost(endpoint string, body *bytes.Buffer) (*http.Response, error) {
	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	//req.Header.Set("Content-Type", *contentType)
	return httpclient.Do(req)
}
