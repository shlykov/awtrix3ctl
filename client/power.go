package client

func (c *Client) PowerOff() error {
	_, err := c.doRequest("POST", "power", map[string]bool{"power": false})
	if err != nil {
		return err
	}
	return nil
}
func (c *Client) PowerOn() error {
	_, err := c.doRequest("POST", "power", map[string]bool{"power": true})
	if err != nil {
		return err
	}
	return nil
}
