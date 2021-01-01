package main

import "net/http"

type client struct {
	*http.Client
}

func (c *client) Get(url string, headers mapping) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return c.Do(req)
}
