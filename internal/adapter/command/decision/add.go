package decision

import (
	util "adg/internal/adapter/command"
	"adg/internal/application/inputport"
	domain "adg/internal/domain/config"
	"fmt"

	"github.com/spf13/cobra"
)

func NewAddCommand(input inputport.DecisionAdd, config domain.ConfigService) *cobra.Command {
	var titles []string
	var modelPath string
	var err error

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Adds one or more decision points to a model",
		RunE: func(cmd *cobra.Command, args []string) error {
			modelPath, err = util.ResolveModelPathOrDefault(modelPath, config)
			if err != nil {
				return err
			}

			if len(titles) == 0 {
				return fmt.Errorf("at least one --title must be provided")
			}

			return input.Add(modelPath, titles)
		},
	}

	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the decision model (optional if configured)")
	cmd.Flags().StringSliceVar(&titles, "title", nil, "One or more titles for new decisions (required)")

	return cmd
}
