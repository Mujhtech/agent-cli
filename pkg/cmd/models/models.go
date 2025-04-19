package model

import (
	"fmt"

	"github.com/mujhtech/agent-cli/pkg/agent"
	"github.com/mujhtech/agent-cli/pkg/views/models"
	"github.com/spf13/cobra"
)

var setDefaultModel string

func RegisterModelsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "models",
		Short: "List of available models",
		RunE: func(cmd *cobra.Command, args []string) error {

			// check if flag is set
			if setDefaultModel != "" {

				model, err := agent.GetModelCatalogByIndex(setDefaultModel)

				if err != nil {
					return err
				}

				fmt.Printf("%s (%s) set as default model", model.Name, model.Model)

			} else {
				_, err := models.Run()

				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&setDefaultModel, "set-default-model", "d", "", "Set default model")

	return cmd
}
