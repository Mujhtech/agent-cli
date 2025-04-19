package index

import "github.com/spf13/cobra"

func RegisterIndexCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "index",
		Short: "Index the current workspace",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	return cmd
}
