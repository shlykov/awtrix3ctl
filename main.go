package main

import (
	device "awtrix3-cli/client"
	"flag"
	"fmt"
	"os"
)

func main() {
	var awtrix3Url string
	awtrix3UrlFlag := flag.String("url", "", "Set Awtrix3 device url")
	flag.Parse()

	if *awtrix3UrlFlag != "" {
		awtrix3Url = *awtrix3UrlFlag
	} else {
		awtrix3Url = os.Getenv("AWTRIX3_URL")
	}

	if awtrix3Url == "" {
		fmt.Println("Please provide the Url of your Awtrix3 device. Use -url or set the environment variable AWTRIX3_URL.")
		os.Exit(1)
	}

	client, err := device.NewClient(awtrix3Url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = execCommand(flag.Args(), client)
	if err != nil {
		fmt.Println("can't execute command", err)
		os.Exit(1)
	}
}

func execCommand(args []string, client *device.Client) error {
	if len(args) == 0 {
		return fmt.Errorf("no command found. Use -help to show all commands")
	}

	switch args[0] {
	case "help":
		return cmdHelp()
	case "stats":
		return cmdStats(client)
	case "settings":
		return cmdSettings(client, args[1:])
	case "display":
		return cmdDisplay(client, args[1:])
	case "reboot":
		return client.Reboot()
	case "indicator":
		return cmdIndicator(client, args[1:])

	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func cmdIndicator(client *device.Client, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("expected indicator number as 1st argument")
	}
	if len(args) < 2 {
		return fmt.Errorf("expected indicator color as 2nd argument. use rgb hex color code")
	}

	return client.SetIndicator(args[0], args[1])
}

func cmdSettings(client *device.Client, args []string) error {
	if len(args) == 0 {
		settings, err := client.GetSettings()
		if err != nil {
			return fmt.Errorf("can't check display status: %w", err)
		}

		fmt.Println(settings)
		return nil
	}
	return nil
}

func cmdDisplay(client *device.Client, args []string) error {
	if len(args) == 0 {
		stats, err := client.GetStats()
		if err != nil {
			return fmt.Errorf("can't check display status: %w", err)
		}

		fmt.Printf("Display: %s\n", stats.GetDisplayState())
		return nil
	}
	switch args[0] {
	case "on":
		err := client.DisplayOn()
		if err == nil {
			fmt.Println("Display turned on")
		}
		return err
	case "off":
		err := client.DisplayOff()
		if err == nil {
			fmt.Println("Display turned off")
		}
	default:
		return fmt.Errorf("unknown display state: %s", args[0])
	}

	return nil
}

func cmdStats(client *device.Client) error {
	stats, err := client.GetStats()
	if err != nil {
		return fmt.Errorf("can't exec stats command: %w", err)
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
}

func cmdHelp() error {
	fmt.Println("Help")
	return nil
}
