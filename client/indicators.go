package client

import (
	"fmt"
)

func (c *Client) SetIndicator(number, color string) error {
	if number != "1" && number != "2" && number != "3" {
		return fmt.Errorf("invalid indicator number. 1,2 or 3 expected, got %d", number)
	}
	if len(color) != 7 || color[0] != '#' {
		return fmt.Errorf("invalid color format. Expected RGB hex format like #RRGGBB, got %s", color)
	}
	for _, char := range color[1:] {
		if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'f') || (char >= 'A' && char <= 'F')) {
			return fmt.Errorf("invalid color format. Expected RGB hex format like #RRGGBB, got %s", color)
		}
	}

	_, err := c.doRequest("POST", "indicator"+number, map[string]string{"color": color})
	if err != nil {
		return err
	}
	return nil
}
