package model

import (
	util "adg/internal/adapter/command"
	"adg/internal/application/inputport"
	domain "adg/internal/domain/config"
	"fmt"

	"github.com/spf13/cobra"
)

func NewCopyCommand(input inputport.ModelCopy, config domain.ConfigService) *cobra.Command {
	var modelPath, targetPath, idFilter, titlePattern string
	var tagFilters, statusFilters []string

	cmd := &cobra.Command{
		Use:   "copy",
		Short: "Copies a model, optionally a subset based on filters",
		RunE: func(cmd *cobra.Command, args []string) error {
			if targetPath == "" {
				return fmt.Errorf("--target is required")
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

			return input.Copy(modelPath, targetPath, filters)
		},
	}

	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the source model (optional if configured)")
	cmd.Flags().StringVar(&targetPath, "target", "", "Destination path where the source model is copied to (required)")
	cmd.Flags().StringArrayVar(&tagFilters, "tag", []string{}, "Filter by tag (repeatable)")
	cmd.Flags().StringArrayVar(&statusFilters, "status", []string{}, "Filter by status (repeatable)")
	cmd.Flags().StringVar(&titlePattern, "title", "", "Regex pattern to match titles")
	cmd.Flags().StringVar(&idFilter, "id", "", "Match specific IDs or ranges (e.g. 0001,0003-0005)")

	return cmd
}
