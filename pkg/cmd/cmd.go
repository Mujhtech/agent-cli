package cmd

import (
	"github.com/mujhtech/agent-cli/pkg/cmd/chat"
	"github.com/mujhtech/agent-cli/pkg/cmd/code"
	"github.com/mujhtech/agent-cli/pkg/cmd/create"
	"github.com/mujhtech/agent-cli/pkg/cmd/index"
	"github.com/mujhtech/agent-cli/pkg/cmd/initial"
	model "github.com/mujhtech/agent-cli/pkg/cmd/models"
	"github.com/mujhtech/agent-cli/pkg/cmd/review"
	"github.com/mujhtech/agent-cli/pkg/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() error {
	cmd := &cobra.Command{
		Use:           "agent-cli",
		Short:         "An open-source alternative to OpenAI Codex & Claude Code",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          initial.RunInitialCommand,
	}

	cmd.AddCommand(version.RegisterVersionCommand())
	cmd.AddCommand(model.RegisterModelsCommand())
	cmd.AddCommand(code.RegisterCodeCommand())
	cmd.AddCommand(create.RegisterCreateCommand())
	cmd.AddCommand(review.RegisterReviewCommand())
	cmd.AddCommand(index.RegisterIndexCommand())
	cmd.AddCommand(chat.RegisterChatCommand())

	err := cmd.Execute()

	return err
}
