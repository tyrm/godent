package fc

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
)

type serverWellKnown struct {
	MatrixServer string `json:"m.server"`
}

func (c *Client) fetchServerWellKnown(ctx context.Context, serverName string) (string, error) {
	wellKnowURL := url.URL{
		Scheme: "https",
		Host:   serverName,
		Path:   "/.well-known/matrix/server",
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, wellKnowURL.String(), nil)
	if err != nil {
		return "", fmt.Errorf("req: %s", err.Error())
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return "", fmt.Errorf("get: %s", err.Error())
	}
	if resp.StatusCode != 200 {
		return "", ErrNotOKStatus
	}

	var data serverWellKnown
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", fmt.Errorf("decode: %s", err.Error())
	}
	if data.MatrixServer == "" {
		return "", ErrHomeServerNotFound
	}

	return data.MatrixServer, nil
}

func (c *Client) fetchServerSRV(serverName string) (string, error) {
	_, srvs, err := net.LookupSRV("matrix", "tcp", serverName)
	if err != nil {
		return "", fmt.Errorf("lookup: %s", err.Error())
	}
	if len(srvs) == 0 {
		return "", ErrHomeServerNotFound
	}

	return fmt.Sprintf("%s:%d", strings.TrimSuffix(srvs[0].Target, "."), srvs[0].Port), nil
}
