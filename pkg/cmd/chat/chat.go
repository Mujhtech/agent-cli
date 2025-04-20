package chat

import (
	"github.com/mujhtech/agent-cli/pkg/views/chat"
	"github.com/spf13/cobra"
)

func RegisterChatCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chat",
		Short: "Chat in the current workspace",
		RunE: func(cmd *cobra.Command, args []string) error {

			_, err := chat.Run()

			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
