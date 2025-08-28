package cmd

import (
	"awtrix3ctl/client"
	"fmt"

	"github.com/spf13/cobra"
)

func NewIndicatorCommand(clt *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "indicator <1|2|3> <#RRGGBB>",
		Short: "Set an indicator LED color",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := clt.SetIndicator(args[0], args[1]); err != nil {
				return err
			}
			fmt.Println("Indicator updated")
			return nil
		},
	}
	return cmd
}
