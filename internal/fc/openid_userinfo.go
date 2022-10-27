package fc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	gdhttp "github.com/tyrm/godent/internal/http"
)

type openidUserinfo struct {
	Sub string `json:"sub"`
}

func (c *Client) OpenidUserinfo(ctx context.Context, matrixServer, accessToken string) (*openidUserinfo, error) {
	homeServer := c.getHomeServer(ctx, matrixServer)

	query := url.Values{}
	query.Set(gdhttp.QueryAccessToken, accessToken)

	userinfoURL := url.URL{
		Scheme:   "https",
		Host:     homeServer,
		Path:     "/_matrix/federation/v1/openid/userinfo",
		RawQuery: query.Encode(),
	}

	resp, err := c.get(ctx, userinfoURL.String())
	if err != nil {
		return nil, fmt.Errorf("get: %s", err.Error())
	}
	if resp.StatusCode != 200 {
		return nil, ErrNotOKStatus
	}

	var data openidUserinfo
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("json: %s", err.Error())
	}

	return &data, nil
}
