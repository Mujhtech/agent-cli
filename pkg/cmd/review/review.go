package review

import "github.com/spf13/cobra"

func RegisterReviewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "review",
		Short: "Review code",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	return cmd
}
