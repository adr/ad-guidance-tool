package decision

import (
	util "github.com/adr/ad-guidance-tool/internal/adapter/command"
	"github.com/adr/ad-guidance-tool/internal/application/inputport"
	domain "github.com/adr/ad-guidance-tool/internal/domain/config"

	"github.com/spf13/cobra"
)

func NewRuleCommand(input inputport.DecisionRule, config domain.ConfigService) *cobra.Command {
	var id string
	var title string
	var modelPath string
	var outputPath string
	var err error

	cmd := &cobra.Command{
		Use:   "rule",
		Short: "Generate a .rule file template for an ADR",
		Long: `Generate a .rule file template based on an existing architectural decision record.
The rule file uses the ADR-DSL format and can be used to define enforceable architectural rules.
Specify either --id or --title to identify the ADR.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			modelPath, err = util.ResolveModelPathOrDefault(modelPath, config)
			if err != nil {
				return err
			}

			return input.Rule(modelPath, id, title, outputPath)
		},
	}

	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the decision model (optional if configured)")
	cmd.Flags().StringVar(&id, "id", "", "ID of the decision to generate a rule for")
	cmd.Flags().StringVar(&title, "title", "", "Title of the decision to generate a rule for")
	cmd.Flags().StringVar(&outputPath, "output", "", "Output path for the rule file (defaults to ADR directory)")

	return cmd
}
