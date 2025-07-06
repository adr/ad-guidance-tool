package decision

import (
	util "adg/internal/adapter/command"
	"adg/internal/application/inputport"
	domain "adg/internal/domain/config"
	"fmt"

	"github.com/spf13/cobra"
)

func NewEditCommand(input inputport.DecisionEdit, config domain.ConfigService) *cobra.Command {
	var modelPath, idOrTitle, id, title string
	var question, criteria string
	var options []string
	var err error

	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a decision file",
		RunE: func(cmd *cobra.Command, args []string) error {
			modelPath, err = util.ResolveModelPathOrDefault(modelPath, config)
			if err != nil {
				return err
			}

			err := util.ResolveIdOrTitle(idOrTitle, &id, &title)
			if err != nil {
				return err
			}

			// validate: must be editing something
			if question == "" && criteria == "" && len(options) == 0 {
				return fmt.Errorf("at least one of --question, --option, or --criteria must be provided")
			}

			// prepare pointer args
			var qPtr, cPtr *string
			var oPtr *[]string

			if question != "" {
				qPtr = &question
			}
			if criteria != "" {
				cPtr = &criteria
			}
			if len(options) > 0 {
				oPtr = &options
			}

			return input.Edit(modelPath, id, title, qPtr, oPtr, cPtr)
		},
	}

	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the decision model (optional if set in config)")
	cmd.Flags().StringVar(&idOrTitle, "id", "", "ID or title of the decision to edit (e.g. 0001, 'my-decision')")
	cmd.Flags().StringVar(&question, "question", "", "Edit the Question section")
	cmd.Flags().StringArrayVar(&options, "option", nil, "Add one or more options (use multiple --option flags to specify multiple options at once or repeat command)")
	cmd.Flags().StringVar(&criteria, "criteria", "", "Edit the Criterion section")

	return cmd
}
