package client

import (
	"io"
	"net/http"
)

var HttpClient = &http.Client{}

func ForwardRequest(method, url string, body io.Reader, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header = headers.Clone()
	return HttpClient.Do(req)
}
