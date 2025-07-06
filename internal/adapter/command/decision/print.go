package decision

import (
	util "adg/internal/adapter/command"
	"adg/internal/application/inputport"
	domain "adg/internal/domain/config"
	"fmt"

	"github.com/spf13/cobra"
)

// todo: rename all related functions and files to view
func NewPrintCommand(input inputport.DecisionPrint, config domain.ConfigService) *cobra.Command {
	var err error
	var modelPath string
	var idsOrTitles, ids, titles []string
	var printQuestion, printOptions, printCriteria, printComments, printOutcome bool

	cmd := &cobra.Command{
		Use:   "view",
		Short: "Show the full or partial content of one or more decision files",
		RunE: func(cmd *cobra.Command, args []string) error {
			modelPath, err = util.ResolveModelPathOrDefault(modelPath, config)
			if err != nil {
				return err
			}

			for _, value := range idsOrTitles {
				var id, title string
				if err := util.ResolveIdOrTitle(value, &id, &title); err != nil {
					return err
				}
				if id != "" {
					ids = append(ids, id)
				} else {
					titles = append(titles, title)
				}
			}

			if len(ids) == 0 && len(titles) == 0 {
				return fmt.Errorf("at least one --id or --title must be provided")
			}

			if !printQuestion && !printOptions && !printCriteria && !printComments && !printOutcome {
				printQuestion = true
				printOptions = true
				printCriteria = true
				printComments = true
				printOutcome = true
			}

			sections := map[string]bool{
				"question": printQuestion,
				"options":  printOptions,
				"criteria": printCriteria,
				"comments": printComments,
				"outcome":  printOutcome,
			}

			return input.Print(modelPath, ids, titles, sections)
		},
	}

	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the decision model directory")
	cmd.Flags().StringSliceVar(&idsOrTitles, "id", nil, "IDs or titles of the decisions to print (e.g. 0001, 'my-decision') (can be repeated)")

	cmd.Flags().BoolVar(&printQuestion, "question", false, "Print the Question section")
	cmd.Flags().BoolVar(&printOptions, "options", false, "Print the Options section")
	cmd.Flags().BoolVar(&printCriteria, "criteria", false, "Print the Criteria section")
	cmd.Flags().BoolVar(&printComments, "comments", false, "Print the Comments section")
	cmd.Flags().BoolVar(&printOutcome, "outcome", false, "Print the Outcome section")

	return cmd
}
