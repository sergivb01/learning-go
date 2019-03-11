package main

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// Client defines the http client
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new HTTP client for the API client
func NewClient(baseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	if baseURL == "" {
		baseURL = "https://httpbin.org"
	}

	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

// doRequest performs an HTTP request to an specified path using the given method
func (c Client) doRequest(method, path string) ([]byte, error) {
	if method == "" {
		return nil, errors.New("method is nil")
	}

	req, err := http.NewRequest(method, c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}

	if c.httpClient == nil {
		return nil, errors.New("httpClient is nil")
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("did not recive statusCode 200")
	}

	return body, nil
}
