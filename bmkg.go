package bmkg

import "net/http"

const baseURL = "http://data.bmkg.go.id"

// Client is used for communicating with BMKG data
type Client struct {
	client  *http.Client
	BaseURL string
}

// NewClient returns client object
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		client:  httpClient,
		BaseURL: baseURL,
	}
}
