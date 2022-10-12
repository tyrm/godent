package fc

import "context"

func (c *Client) getHomeServer(ctx context.Context, serverName string) string {
	// try to get http
	homeServer, err := c.fetchServerWellKnown(ctx, serverName)
	if err == nil {
		return homeServer
	}

	// try to get dns
	homeServer, err = c.fetchServerSRV(ctx, serverName)
	if err == nil {
		return homeServer
	}

	return serverName
}
