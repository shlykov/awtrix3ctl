package client

func (c *Client) DisplayOn() error {
	_, err := c.doRequest("POST", "power", map[string]bool{"power": true})
	return err
}

func (c *Client) DisplayOff() error {
	_, err := c.doRequest("POST", "power", map[string]bool{"power": false})
	return err
}
