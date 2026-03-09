package model

import (
	util "github.com/adr/ad-guidance-tool/internal/adapter/command"
	"github.com/adr/ad-guidance-tool/internal/application/inputport"
	domain "github.com/adr/ad-guidance-tool/internal/domain/config"

	"github.com/spf13/cobra"
)

func NewValidateCommand(input inputport.ModelValidate, config domain.ConfigService) *cobra.Command {
	var modelPath string

	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate the models decisions by checking if the files match the index file",
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedPath, err := util.ResolveModelPathOrDefault(modelPath, config)
			if err != nil {
				return err
			}

			return input.Validate(resolvedPath)
		},
	}

	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the decision model directory (optional if configured)")

	return cmd
}
