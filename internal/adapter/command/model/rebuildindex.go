package model

import (
	util "adg/internal/adapter/command"
	"adg/internal/application/inputport"
	domain "adg/internal/domain/config"

	"github.com/spf13/cobra"
)

func NewRebuildIndexCommand(input inputport.ModelRebuildIndex, config domain.ConfigService) *cobra.Command {
	var modelPath string
	var err error

	cmd := &cobra.Command{
		Use:   "rebuild",
		Short: "Rebuilds the index file for the given model",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {

			modelPath, err = util.ResolveModelPathOrDefault(modelPath, config)
			if err != nil {
				return err
			}

			return input.RebuildIndex(modelPath)
		},
	}

	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the source model (optional if configured)")

	return cmd
}
