package decision

import (
	util "adg/internal/adapter/command"
	"adg/internal/application/inputport"
	domain "adg/internal/domain/config"
	"fmt"

	"github.com/spf13/cobra"
)

func NewDecideCommand(input inputport.DecisionDecide, config domain.ConfigService) *cobra.Command {
	var modelPath, idOrTitle, id, title, option, reason, author string
	var enforce bool
	var err error

	cmd := &cobra.Command{
		Use:   "decide",
		Short: "Marks a decision as decided by selecting one of its options",
		Long: `Decide finalizes a decision by selecting a specific option and marking the decision as decided.
You must provide --id to identify the decision.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			modelPath, err = util.ResolveModelPathOrDefault(modelPath, config)
			if err != nil {
				return err
			}

			err := util.ResolveIdOrTitle(idOrTitle, &id, &title)
			if err != nil {
				return err
			}

			if author == "" {
				author = config.GetAuthor()
			}

			if option == "" {
				return fmt.Errorf("--option must be provided (either its name or a positive integer (1-based index)")
			}

			return input.Decide(modelPath, id, title, option, reason, author, enforce)
		},
	}

	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the decision model (optional if configured)")
	cmd.Flags().StringVar(&idOrTitle, "id", "", "ID or title of the decision to decide, e.g., 0001, 'my-decision'")
	cmd.Flags().StringVar(&option, "option", "", "Name or the number of the option being selected, e.g., 'first-option' or '1' (required)")
	cmd.Flags().StringVar(&reason, "rationale", "", "Optional rationale or explanation for the selected option")
	cmd.Flags().StringVar(&author, "author", "", "Name of the person deciding (overrides config)")
	cmd.Flags().BoolVarP(&enforce, "force", "f", false, "If an option name is provided which does not exist in the decision, using --force will automatically add it as an option and use it for the decision.")

	return cmd
}
