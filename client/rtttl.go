package client

import (
	"fmt"
	"regexp"
)

var rtttlRegex = regexp.MustCompile(`^([A-Za-z0-9_]{1,15}):(?:(?:d=(?:1|2|4|8|16|32)|o=[4-7]|b=\d{1,3})(?:,(?:d=(?:1|2|4|8|16|32)|o=[4-7]|b=\d{1,3})){0,2}:)?((?:(?:1|2|4|8|16|32)?[PpAaBbCcDdEeFfGg]#?\.?[4-7]?)(?:,(?:1|2|4|8|16|32)?[PpAaBbCcDdEeFfGg]#?\.?[4-7]?)*?)$`)

func (c *Client) PlayRtttl(rtttl string) error {
	if !rtttlRegex.MatchString(rtttl) {
		return fmt.Errorf("invalid RTTTL format")
	}

	_, err := c.doRequest("POST", "rtttl", map[string]string{"rtttl": rtttl})
	if err != nil {
		return err
	}
	return nil
}
