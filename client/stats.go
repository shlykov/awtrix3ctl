package client

import "encoding/json"

type Stats struct {
	Bat        int    `json:"bat"`
	BatRaw     int    `json:"bat_raw"`
	Type       int    `json:"type"`
	Lux        int    `json:"lux"`
	LdrRaw     int    `json:"ldr_raw"`
	Ram        int    `json:"ram"`
	Bri        int    `json:"bri"`
	Temp       int    `json:"temp"`
	Hum        int    `json:"hum"`
	Uptime     int    `json:"uptime"`
	WifiSignal int    `json:"wifi_signal"`
	Messages   int    `json:"messages"`
	Version    string `json:"version"`
	Indicator1 bool   `json:"indicator1"`
	Indicator2 bool   `json:"indicator2"`
	Indicator3 bool   `json:"indicator3"`
	App        string `json:"app"`
	UID        string `json:"uid"`
	Matrix     bool   `json:"matrix"`
	IPAddress  string `json:"ip_address"`
}

func (srs *Stats) GetDisplayState() string {
	if srs.Matrix {
		return "on"
	}
	return "off"
}

func (c *Client) GetStats() (*Stats, error) {
	body, err := c.doRequest("GET", "stats", nil)
	if err != nil {
		return nil, err
	}

	var stats Stats
	err = json.Unmarshal(body, &stats)
	return &stats, err
}
