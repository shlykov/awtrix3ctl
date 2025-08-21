package client

import "encoding/json"

func (c *Client) GetSettings() (*Settings, error) {
	body, err := c.doRequest("GET", "settings", nil)
	if err != nil {
		return nil, err
	}

	var settings Settings
	err = json.Unmarshal(body, &settings)
	return &settings, err
}

type Settings struct {
	// Duration an app is displayed in seconds. Value Range: Positive integer. Default: 7
	AppDisplayDuration int `json:"ATIME"`

	// Choose between app transition effects. Value Range: 0–10. Default: 1
	TransitionEffect int `json:"TEFF"`

	// Time taken for the transition to the next app in milliseconds. Value Range: Positive integer. Default: 500
	TransitionSpeed int `json:"TSPEED"`

	// Global text color. Value Range: RGB array or hex color. Default: N/A
	GlobalTextColor int `json:"TCOL"`

	// Changes the time app style. Value Range: 0–6. Default: 1
	TimeAppStyle int `json:"TMODE"`

	// Calendar header color of the time app. Value Range: RGB array or hex color. Default: #FF0000
	CalendarHeaderColor int `json:"CHCOL"`

	// Calendar body color of the time app. Value Range: RGB array or hex color. Default: #FFFFFF
	CalendarBodyColor int `json:"CBCOL"`

	// Calendar text color in the time app. Value Range: RGB array or hex color. Default: #000000
	CalendarTextColor int `json:"CTCOL"`

	// Enable or disable the weekday display. Value Range: true/false. Default: true
	WeekdayDisplay bool `json:"WD"`

	// Active weekday color. Value Range: RGB array or hex color. Default: N/A
	ActiveWeekdayColor int `json:"WDCA"`

	// Inactive weekday color. Value Range: RGB array or hex color. Default: N/A
	InactiveWeekdayColor int `json:"WDCI"`

	// Matrix brightness. Value Range: 0–255. Default: N/A
	MatrixBrightness int `json:"BRI"`

	// Automatic brightness control. Value Range: true/false. Default: N/A
	AutoBrightnessControl bool `json:"ABRI"`

	// Automatic switching to the next app. Value Range: true/false. Default: N/A
	AutoAppTransition bool `json:"ATRANS"`

	// Color correction for the matrix. Value Range: RGB array. Default: N/A
	ColorCorrection string `json:"CCORRECTION"`

	// Color temperature for the matrix. Value Range: RGB array. Default: N/A
	ColorTemperature string `json:"CTEMP"`

	// Time format for the TimeApp. Value Range: Varies (see documentation). Default: N/A
	TimeFormat string `json:"TFORMAT"`

	// Date format for the DateApp. Value Range: Varies (see documentation). Default: N/A
	DateFormat string `json:"DFORMAT"`

	// Start the week on Monday. Value Range: true/false. Default: true
	StartWeekOnMonday bool `json:"SOM"`

	// Shows the temperature in Celsius (Fahrenheit when false). Value Range: true/false. Default: true
	ShowTemperatureInCelsius bool `json:"CEL"`

	// Block physical navigation keys (still sends input to MQTT). Value Range: true/false. Default: false
	BlockPhysicalNavigationKeys bool `json:"BLOCKN"`

	// Display text in uppercase. Value Range: true/false. Default: true
	DisplayTextInUppercase bool `json:"UPPERCASE"`

	// Text color of the time app. Use 0 for global text color. Value Range: RGB array or hex color. Default: N/A
	TimeAppTextColor int `json:"TIME_COL"`

	// Text color of the date app. Use 0 for global text color. Value Range: RGB array or hex color. Default: N/A
	DateAppTextColor int `json:"DATE_COL"`

	// Text color of the temperature app. Use 0 for global text color. Value Range: RGB array or hex color. Default: N/A
	TemperatureAppTextColor int `json:"TEMP_COL"`

	// Text color of the humidity app. Use 0 for global text color. Value Range: RGB array or hex color. Default: N/A
	HumidityAppTextColor int `json:"HUM_COL"`

	// Text color of the battery app. Use 0 for global text color. Value Range: RGB array or hex color. Default: N/A
	BatteryAppTextColor int `json:"BAT_COL"`

	// Scroll speed modification. Value Range: Percentage of original scroll speed. Default: 100
	ScrollSpeed int `json:"SSPEED"`

	// Enable or disable the native time app (requires reboot). Value Range: true/false. Default: true
	EnableTimeApp bool `json:"TIM"`

	// Enable or disable the native date app (requires reboot). Value Range: true/false. Default: true
	EnableDateApp bool `json:"DAT"`

	// Enable or disable the native humidity app (requires reboot). Value Range: true/false. Default: true
	EnableHumidityApp bool `json:"HUM"`

	// Enable or disable the native temperature app (requires reboot). Value Range: true/false. Default: true
	EnableTemperatureApp bool `json:"TEMP"`

	// Enable or disable the native battery app (requires reboot). Value Range: true/false. Default: true
	EnableBatteryApp bool `json:"BAT"`

	// Enable or disable the matrix. Similar to power endpoint but without the animation. Value Range: true/false. Default: true
	EnableMatrix bool `json:"MATP"`

	// Allows to set the volume of the buzzer and DFplayer. Value Range: 0–30. Default: true
	BuzzerVolume int `json:"VOL"`

	// Sets a global effect overlay (cannot be used with app specific overlays). Value Range: Varies (see documentation). Default: N/A
	GlobalEffectOverlay string `json:"OVERLAY"`

	GAMMA float32 `json:"GAMMA"`
	MAT   int     `json:"MAT"`
	SOUND bool    `json:"SOUND"`
}
