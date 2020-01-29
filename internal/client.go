package internal

import "net/http"

// Client is a client for interacting with the YouTube Data API.
type Client struct {
	Endpoint   string
	HTTPClient *http.Client
	Token      string
}

// ClientOption are functions that are passed into NewClient to modify the behaviour of the Client.
type ClientOption func(*Client)

// NewClient makes a new Client capable of making API requests.
func NewClient(token string, opts ...ClientOption) *Client {
	c := &Client{
		Endpoint: defaultEndpoint,
		Token:    token,
	}

	for _, optionFunc := range opts {
		optionFunc(c)
	}

	if c.HTTPClient == nil {
		c.HTTPClient = http.DefaultClient
	}

	return c
}
