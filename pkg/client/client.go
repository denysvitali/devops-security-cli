package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const UserAgent = "devops-security-cli "
const ApiEndpoint = "https://api.devops.security"
var Version = "develop"

type Client struct {
	httpClient *http.Client
	apiUrl *url.URL
}

func mustGetPath(u *url.URL, path string) string {
	finalUrl, err := u.Parse(path)
	if err != nil {
		panic(err)
	}
	return finalUrl.String()
}

func Get[T any](c *Client, path string, body io.Reader) (*T, error) {
	req, err := http.NewRequest(http.MethodGet, mustGetPath(c.apiUrl, path), body)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %v", err)
	}
	req.Header.Add("Accept", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to perform GET request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code %d", res.StatusCode)
	}

	dec := json.NewDecoder(res.Body)
	var decoded T
	err = dec.Decode(&decoded)
	return &decoded, err
}

func New(token string) (*Client, error) {
	c := Client{
		httpClient: http.DefaultClient,
	}

	var err error
	c.apiUrl, err = url.Parse(ApiEndpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to parse API Endpoint: %v", err)
	}

	c.httpClient.Transport = CustomTransport{
		Token: token,
	}

	return &c, nil
}