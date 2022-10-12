package fc

func (c *Client) getHomeServer(serverName string) string {
	// try to get http
	homeServer, err := c.fetchServerWellKnown(serverName)
	if err == nil {
		return homeServer
	}

	// try to get dns
	homeServer, err = c.fetchServerSRV(serverName)
	if err == nil {
		return homeServer
	}

	return serverName
}
