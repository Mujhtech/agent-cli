package create

import "github.com/spf13/cobra"

func RegisterCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new project",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	return cmd
}
