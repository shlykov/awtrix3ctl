package cmd

import (
	"awtrix3ctl/client"
	"fmt"

	"github.com/spf13/cobra"
)

func NewStatsCommand(clt *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stats",
		Short: "Show device statistics",
		RunE: func(cmd *cobra.Command, args []string) error {
			stats, err := clt.GetStats()
			if err != nil {
				return err
			}
			fmt.Printf("Device Status:\n")
			fmt.Printf("  Battery: %d%%\n", stats.Bat)
			fmt.Printf("  Battery Raw: %d\n", stats.BatRaw)
			fmt.Printf("  Type: %d\n", stats.Type)
			fmt.Printf("  Lux: %d\n", stats.Lux)
			fmt.Printf("  LDR Raw: %d\n", stats.LdrRaw)
			fmt.Printf("  Free RAM: %d bytes\n", stats.Ram)
			fmt.Printf("  Brightness: %d\n", stats.Bri)
			fmt.Printf("  Temperature: %d °C\n", stats.Temp)
			fmt.Printf("  Humidity: %d %%\n", stats.Hum)
			fmt.Printf("  Uptime: %d seconds\n", stats.Uptime)
			fmt.Printf("  WiFi Signal: %d dBm\n", stats.WifiSignal)
			fmt.Printf("  Messages: %d\n", stats.Messages)
			fmt.Printf("  Version: %s\n", stats.Version)
			fmt.Printf("  Indicator1: %v\n", stats.Indicator1)
			fmt.Printf("  Indicator2: %v\n", stats.Indicator2)
			fmt.Printf("  Indicator3: %v\n", stats.Indicator3)
			fmt.Printf("  Active app: %s\n", stats.App)
			fmt.Printf("  UID: %s\n", stats.UID)
			fmt.Printf("  Display: %v\n", stats.GetDisplayState())
			fmt.Printf("  IP Address: %s\n", stats.IPAddress)
			return nil
		},
	}
	return cmd
}
