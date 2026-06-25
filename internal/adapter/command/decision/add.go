package decision

import (
	"fmt"

	util "github.com/adr/ad-guidance-tool/internal/adapter/command"
	"github.com/adr/ad-guidance-tool/internal/application/inputport"
	domain "github.com/adr/ad-guidance-tool/internal/domain/config"

	"github.com/spf13/cobra"
)

func NewAddCommand(input inputport.DecisionAdd, config domain.ConfigService) *cobra.Command {
	var titles []string
	var modelPath string
	var err error

	cmd := &cobra.Command{
		Use:   "add [title...]",
		Short: "Adds one or more decision points to a model",
		Long: `Adds one or more decision points to a model.

You can provide the title either as positional arguments or via the --title flag.

Examples:
  adg add My Decision Title
  adg add --title "First Decision" --title "Second Decision"
  adg add --model my-model My Decision Title`,
		RunE: func(cmd *cobra.Command, args []string) error {
			modelPath, err = util.ResolveModelPathOrDefault(modelPath, config)
			if err != nil {
				return err
			}

			// If no --title flags provided, use positional arguments
			if len(titles) == 0 && len(args) > 0 {
				// Join all args as a single title
				titles = []string{joinArgs(args)}
			}

			if len(titles) == 0 {
				return fmt.Errorf("at least one title must be provided (via arguments or --title flag)")
			}

			return input.Add(modelPath, titles)
		},
	}

	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the decision model (optional if configured)")
	cmd.Flags().StringSliceVar(&titles, "title", nil, "One or more titles for new decisions (optional if using positional arguments)")

	return cmd
}

// joinArgs joins command arguments with spaces, preserving the original spacing
func joinArgs(args []string) string {
	result := ""
	for i, arg := range args {
		if i > 0 {
			result += " "
		}
		result += arg
	}
	return result
}
