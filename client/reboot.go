package client

func (c *Client) Reboot() error {
	_, err := c.doRequest("POST", "reboot", nil)
	if err != nil {
		return err
	}
	return nil
}
