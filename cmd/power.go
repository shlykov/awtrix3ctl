package cmd

import (
	"awtrix3ctl/client"
	"fmt"

	"github.com/spf13/cobra"
)

func NewPowerCommand(clt *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "power [on|off]",
		GroupID: "manage",
		Short:   "Set power state",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "on":
				if err := clt.PowerOn(); err != nil {
					return err
				}
				fmt.Println("Power turned on")
				return nil
			case "off":
				if err := clt.PowerOff(); err != nil {
					return err
				}
				fmt.Println("Power turned off")
				return nil
			default:
				return fmt.Errorf("unknown power state: %s", args[0])
			}
		},
	}
	return cmd
}
