package index

import (
	"github.com/mujhtech/agent-cli/pkg/views/index"
	"github.com/spf13/cobra"
)

func RegisterIndexCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "index",
		Short: "Index the current workspace",
		RunE: func(cmd *cobra.Command, args []string) error {

			err := index.Run()

			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
