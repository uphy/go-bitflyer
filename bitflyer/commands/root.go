package commands

import (
	"github.com/spf13/cobra"
	bitflyer "github.com/uphy/go-bitflyer"
)

var (
	client = bitflyer.New()
	root   = &cobra.Command{
		Use:   "bitflyer",
		Short: "bitflyer parent command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
)

func Execute() error {
	return root.Execute()
}
