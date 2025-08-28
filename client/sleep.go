package client

func (c *Client) Sleep(sec int) error {
	_, err := c.doRequest("POST", "sleep", map[string]int{"sleep": sec})
	if err != nil {
		return err
	}
	return nil
}
