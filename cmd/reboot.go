package cmd

import (
	"awtrix3ctl/client"

	"github.com/spf13/cobra"
)

func NewRebootCommand(clt *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reboot",
		Short: "Reboot the device",
		RunE: func(cmd *cobra.Command, args []string) error {
			return clt.Reboot()
		},
	}
	return cmd
}
