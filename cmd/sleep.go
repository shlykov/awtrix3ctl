package cmd

import (
	"awtrix3ctl/client"
	"strconv"

	"github.com/spf13/cobra"
)

func NewSleepCommand(clt *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sleep X",
		Short: "Send the board in deep sleep mode in X seconds(turns off the matrix as well), good for saving battery life",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			x, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			return clt.Sleep(x)
		},
	}
	return cmd
}
