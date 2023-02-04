package client

import (
	"fmt"
	"net/http"
)

type CustomTransport struct {
	Token string
}

func (a CustomTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	if a.Token != "" {
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
	}
	request.Header.Add("User-Agent", fmt.Sprintf("%s - v%s", UserAgent, Version))
	return http.DefaultTransport.RoundTrip(request)
}

var _ http.RoundTripper = (*CustomTransport)(nil)
