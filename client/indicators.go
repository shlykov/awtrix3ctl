package client

import (
	"fmt"
)

func (c *Client) SetIndicator(number, color, effect string, duration int) error {
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

	if effect != "blink" && effect != "fade" && effect != "" {
		return fmt.Errorf("invalid effect. Expected blink or fade, got %s", effect)
	}

	payload := map[string]any{"color": color}
	if effect != "" {
		payload[effect] = duration
	}
	_, err := c.doRequest("POST", "indicator"+number, payload)
	if err != nil {
		return err
	}
	return nil
}
