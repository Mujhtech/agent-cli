package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "dev"

func RegisterVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version and exit",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Agent Cli Version: %s\n", Version)
			return nil
		},
	}

	return cmd
}
