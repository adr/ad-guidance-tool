package model

import (
	"adg/internal/application/inputport"
	"fmt"

	"github.com/spf13/cobra"
)

func NewMergeModelsCommand(input inputport.ModelMerge) *cobra.Command {
	var modelAPath, modelBPath, targetPath, idFilter, titlePattern string
	var tagFilters, statusFilters []string

	cmd := &cobra.Command{
		Use:   "merge",
		Short: "Merges two decision models into a new target model",
		Long: `Merge two existing decision models by combining all decisions into a new target model directory.

IDs will be renumbered if needed to ensure uniqueness.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Validate flags
			if modelAPath == "" || modelBPath == "" || targetPath == "" {
				return fmt.Errorf("--model1, --model2 and --target must all be provided")
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

			return input.Merge(modelAPath, modelBPath, targetPath, filters)
		},
	}

	cmd.Flags().StringVar(&modelAPath, "model1", "", "Path to the first model (required)")
	cmd.Flags().StringVar(&modelBPath, "model2", "", "Path to the second model (required)")
	cmd.Flags().StringVar(&targetPath, "target", "", "Path where the merged model will be created (required)")
	cmd.Flags().StringArrayVar(&tagFilters, "tag", []string{}, "Filter by tag (repeatable)")
	cmd.Flags().StringArrayVar(&statusFilters, "status", []string{}, "Filter by status (repeatable)")
	cmd.Flags().StringVar(&titlePattern, "title", "", "Regex pattern to match titles")
	cmd.Flags().StringVar(&idFilter, "id", "", "Match specific IDs or ranges (e.g. 0001,0003-0005)")

	return cmd
}
