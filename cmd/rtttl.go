package cmd

import (
	"awtrix3ctl/client"

	"github.com/spf13/cobra"
)

func NewRtttlCommand(clt *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rtttl <rtttl_string>",
		GroupID: "manage",
		Short:   "Play a RTTTL sound from a given RTTTL string.",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return clt.PlayRtttl(args[0])
		},
	}
	return cmd
}
