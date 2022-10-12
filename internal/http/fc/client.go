package fc

import "net/http"

func New(httpClient *http.Client) *Client {
	return &Client{
		http: httpClient,
	}
}

type Client struct {
	http *http.Client
}
