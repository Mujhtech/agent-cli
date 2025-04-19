package code

import "github.com/spf13/cobra"

func RegisterCodeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "code",
		Short: "Code generation",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	return cmd
}
