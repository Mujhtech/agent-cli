package initial

import (
	model "github.com/mujhtech/agent-cli/pkg/cmd/models"
	"github.com/mujhtech/agent-cli/pkg/views/initals"
	"github.com/spf13/cobra"
)

func RunInitialCommand(cmd *cobra.Command, args []string) error {

	command, err := initals.Run()

	switch command {
	case "models":
		modelsCmd := model.RegisterModelsCommand()
		return modelsCmd.RunE(cmd, []string{})
	case "help":
		return cmd.Help()
	}

	return err
}
