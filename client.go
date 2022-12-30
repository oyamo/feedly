package feedly

import (
	"errors"
	"io"
	"net/http"
	"time"
)

type ClientError struct {
	error
}

type RequestError struct {
	error
}

type Client struct {

	// Client is a wrapper around the http.Client struct.
	// It is used to make requests to the RSS feed.
	c *http.Client
}

// NewClient returns a new instance of Client.
// Set timeout to 10 seconds
func NewClient() *Client {

	c := &http.Client{}
	c.Timeout = 10 * time.Second
	return &Client{c}
}

// FetchFeed fetches the RSS feed from the given url.
// It returns the response code, the response body and any errors.
func (c Client) FetchFeed(url string) (respCode uint32, res []byte, err RequestError) {
	// MakeRequest makes a GET request to the RSS feed.
	// It returns the response and any errors.
	req, newReqErr := http.NewRequest("GET", url, nil)
	if newReqErr != nil {
		return 0, nil, RequestError{err}
	}

	// se content type to xml
	req.Header.Set("Content-Type", "application/xml")

	// set user agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")

	resp, doErr := c.c.Do(req)
	if doErr != nil {
		return 0, nil, RequestError{doErr}
	}

	res, rerr := io.ReadAll(resp.Body)
	if rerr != nil {
		return uint32(resp.StatusCode), nil, RequestError{rerr}
	}

	if resp.StatusCode != http.StatusOK {
		return uint32(resp.StatusCode), nil, RequestError{errors.New(resp.Status)}
	}

	return uint32(resp.StatusCode), res, RequestError{nil}
}
