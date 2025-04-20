package initial

import (
	"github.com/mujhtech/agent-cli/pkg/cmd/chat"
	"github.com/mujhtech/agent-cli/pkg/cmd/create"
	"github.com/mujhtech/agent-cli/pkg/cmd/index"
	model "github.com/mujhtech/agent-cli/pkg/cmd/models"
	"github.com/mujhtech/agent-cli/pkg/cmd/review"
	"github.com/mujhtech/agent-cli/pkg/views/initals"
	"github.com/spf13/cobra"
)

func RunInitialCommand(cmd *cobra.Command, args []string) error {

	command, err := initals.Run()

	switch command {
	case "models":
		modelsCmd := model.RegisterModelsCommand()
		return modelsCmd.RunE(cmd, []string{})
	case "chat":
		chatCmd := chat.RegisterChatCommand()
		return chatCmd.RunE(cmd, []string{})
	case "index":
		indexCmd := index.RegisterIndexCommand()
		return indexCmd.RunE(cmd, []string{})
	case "review":
		reviewCmd := review.RegisterReviewCommand()
		return reviewCmd.RunE(cmd, []string{})
	case "create":
		createCmd := create.RegisterCreateCommand()
		return createCmd.RunE(cmd, []string{})
	case "help":
		return cmd.Help()
	}

	return err
}
