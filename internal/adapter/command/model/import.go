package model

import (
	util "adg/internal/adapter/command"
	"adg/internal/application/inputport"
	domain "adg/internal/domain/config"
	"fmt"

	"github.com/spf13/cobra"
)

func NewImportCommand(input inputport.ModelImport, config domain.ConfigService) *cobra.Command {
	var modelPath, sourcePath, idFilter, titlePattern string
	var tagFilters, statusFilters []string

	cmd := &cobra.Command{
		Use:   "import",
		Short: "Imports a decision model into an existing model",
		Long: `Import all decisions from a source model into an existing target model.

IDs will be renumbered automatically to avoid conflicts.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if sourcePath == "" {
				return fmt.Errorf("--target must both be provided")
			}

			modelPath, err := util.ResolveModelPathOrDefault(modelPath, config)
			if err != nil {
				return err
			}

			filters := make(map[string][]string)
			if len(tagFilters) > 0 {
				filters["tag"] = tagFilters
			}
			if len(statusFilters) > 0 {
				filters["status"] = statusFilters
			}
			if titlePattern != "" {
				filters["title"] = []string{titlePattern}
			}
			if idFilter != "" {
				filters["id"] = []string{idFilter}
			}

			return input.Import(sourcePath, modelPath, filters)
		},
	}

	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the target model (optional if configured)")
	cmd.Flags().StringVar(&sourcePath, "source", "", "Path to the source model which is imported to the target model (required)")
	cmd.Flags().StringArrayVar(&tagFilters, "tag", []string{}, "Filter by tag (repeatable)")
	cmd.Flags().StringArrayVar(&statusFilters, "status", []string{}, "Filter by status (repeatable)")
	cmd.Flags().StringVar(&titlePattern, "title", "", "Regex pattern to match titles")
	cmd.Flags().StringVar(&idFilter, "id", "", "Match specific IDs or ranges (e.g. 0001,0003-0005)")

	return cmd
}
