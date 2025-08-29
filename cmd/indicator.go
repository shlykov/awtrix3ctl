package cmd

import (
	"awtrix3ctl/client"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func NewIndicatorCommand(clt *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "indicator <1|2|3> <#RRGGBB> [<fade> <duration_ms>|<blink> <duration_ms>]",
		GroupID: "interaction",
		Short:   "Set an indicator LED color, add optional effect",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 && len(args) != 4 {
				return fmt.Errorf("accepts 2 or 4 arg(s), received %d", len(args))
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var effect string
			var duration int
			var err error
			if len(args) == 4 {
				effect = args[2]
				duration, err = strconv.Atoi(args[3])
				if err != nil {
					return err
				}
			}
			if err := clt.SetIndicator(args[0], args[1], effect, duration); err != nil {
				return err
			}
			fmt.Println("Indicator updated")
			return nil
		},
	}
	return cmd
}
