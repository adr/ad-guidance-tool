package decision

import (
	util "adg/internal/adapter/command"
	"adg/internal/application/inputport"
	domain "adg/internal/domain/config"

	"github.com/spf13/cobra"
)

func NewListCommand(input inputport.DecisionList, config domain.ConfigService) *cobra.Command {
	var tags []string
	var statuses []string
	var format, titlePattern, idFilter string
	var modelPath string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists decisions in the model, optionally filtering by tag, status, title, or ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			modelPath, err := util.ResolveModelPathOrDefault(modelPath, config)
			if err != nil {
				return err
			}

			filters := make(map[string][]string)

			if len(tags) > 0 {
				filters["tag"] = tags
			}
			if len(statuses) > 0 {
				filters["status"] = statuses
			}
			if titlePattern != "" {
				filters["title"] = []string{titlePattern}
			}
			if idFilter != "" {
				filters["id"] = []string{idFilter}
			}

			return input.ListDecisions(modelPath, filters, format)
		},
	}

	cmd.Flags().StringSliceVar(&tags, "tag", nil, "Filter decisions by one or more tags")
	cmd.Flags().StringSliceVar(&statuses, "status", nil, "Filter decisions by one or more statuses")
	cmd.Flags().StringVar(&format, "format", "simple", "Output format: simple, yaml, json, or md")
	cmd.Flags().StringVar(&titlePattern, "title", "", "Regex pattern to match titles")
	cmd.Flags().StringVar(&idFilter, "id", "", "Match specific IDs or ranges (e.g. 0002,0004-0006)")
	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the decision model (overrides config)")

	return cmd
}
