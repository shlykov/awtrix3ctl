package cmd

import (
	"awtrix3ctl/client"
	"fmt"

	"github.com/spf13/cobra"
)

func NewSettingsCommand(clt *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settings",
		Short: "Show device settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			settings, err := clt.GetSettings()
			if err != nil {
				return fmt.Errorf("can't check settings: %w", err)
			}
			fmt.Println(settings)
			return nil
		},
	}
	return cmd
}
