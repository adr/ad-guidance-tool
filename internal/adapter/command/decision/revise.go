package decision

import (
	util "adg/internal/adapter/command"
	"adg/internal/application/inputport"
	domain "adg/internal/domain/config"

	"github.com/spf13/cobra"
)

func NewReviseCommand(input inputport.DecisionRevise, config domain.ConfigService) *cobra.Command {
	var modelPath, idOrTitle, id, title string
	var err error

	cmd := &cobra.Command{
		Use:   "revise",
		Short: "Creates a copy of a decision and resets its status to 'open' (if not already)",
		Long: `Revise creates a new version of an existing decision by duplicating it and setting its status to 'open'.
You must provide either --id or --title to identify the original decision.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			modelPath, err = util.ResolveModelPathOrDefault(modelPath, config)
			if err != nil {
				return err
			}

			err := util.ResolveIdOrTitle(idOrTitle, &id, &title)
			if err != nil {
				return err
			}

			return input.ReviseDecision(modelPath, id, title)
		},
	}

	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the decision model (optional if set in config)")
	cmd.Flags().StringVar(&idOrTitle, "id", "", "ID or title of the decision to revise (e.g. 0001, 'my-decision')")

	return cmd
}
