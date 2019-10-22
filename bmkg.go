package bmkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://data.bmkg.go.id"

// Client is used for communicating with BMKG data
type Client interface {
	GetXMLBytes(url string) ([]byte, error)
}

type client struct {
	httpClient *http.Client
}

// Config ...
type Config struct {
	BaseURL string
}

// Bmkg ...
type Bmkg struct {
	client Client
	config *Config
}

// NewBmkg returns client object
func NewBmkg(httpClient *http.Client) *Bmkg {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := new(client)
	client.httpClient = httpClient

	config := new(Config)
	config.BaseURL = baseURL

	return &Bmkg{
		client: client,
		config: config,
	}
}

// GetXMLBytes ...
func (c *client) GetXMLBytes(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error http bad request: %v", http.StatusBadRequest)
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
