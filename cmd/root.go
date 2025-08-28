package cmd

import (
	"awtrix3ctl/client"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Context struct {
	*client.Client
}

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "awtrix3ctl",
		Short: "Command-line client for AWTRIX3 device",
		Long:  `A small CLI to interact with an AWTRIX3 device via its HTTP API.`,
	}
	viper.SetEnvPrefix("AWTRIX3")
	viper.AutomaticEnv()

	var urlFlag string
	rootCmd.PersistentFlags().StringVar(&urlFlag, "url", "", "Awtrix3 device base URL (can also use AWTRIX3_URL env var)")
	_ = viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))

	clt := client.NewClient()

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		url := viper.GetString("url")
		if url == "" {
			return fmt.Errorf("device URL is required. Provide --url or AWTRIX3_URL env variable")
		}
		return clt.SetDeviceURL(url)
	}

	//Register sub commands
	rootCmd.AddCommand(NewStatsCommand(clt))
	rootCmd.AddCommand(NewDisplayCommand(clt))
	rootCmd.AddCommand(NewPowerCommand(clt))
	rootCmd.AddCommand(NewSleepCommand(clt))
	rootCmd.AddCommand(NewIndicatorCommand(clt))
	rootCmd.AddCommand(NewRebootCommand(clt))
	rootCmd.AddCommand(NewSettingsCommand(clt))
	return rootCmd
}
