package cmd

import (
	"awtrix3ctl/client"
	"fmt"

	"github.com/spf13/cobra"
)

func NewDisplayCommand(clt *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "display [on|off]",
		GroupID: "interaction",
		Short:   "Get or set display power state",
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				stats, err := clt.GetStats()
				if err != nil {
					return fmt.Errorf("can't check display status: %w", err)
				}
				fmt.Printf("Display: %s\n", stats.GetDisplayState())
				return nil
			}
			switch args[0] {
			case "on":
				if err := clt.DisplayOn(); err != nil {
					return err
				}
				fmt.Println("Display turned on")
				return nil
			case "off":
				if err := clt.DisplayOff(); err != nil {
					return err
				}
				fmt.Println("Display turned off")
				return nil
			default:
				return fmt.Errorf("unknown display state: %s", args[0])
			}
		},
	}
	return cmd
}
