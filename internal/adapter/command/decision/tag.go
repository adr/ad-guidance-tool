package decision

import (
	util "adg/internal/adapter/command"
	"adg/internal/application/inputport"
	domain "adg/internal/domain/config"
	"fmt"

	"github.com/spf13/cobra"
)

func NewTagCommand(input inputport.DecisionTag, config domain.ConfigService) *cobra.Command {
	var modelPath, idOrTitle, id, title string
	var tags []string
	var err error

	cmd := &cobra.Command{
		Use:   "tag",
		Short: "Categorizes a decision by adding one or more tags to its metadata",
		RunE: func(cmd *cobra.Command, args []string) error {
			modelPath, err = util.ResolveModelPathOrDefault(modelPath, config)
			if err != nil {
				return err
			}

			err := util.ResolveIdOrTitle(idOrTitle, &id, &title)
			if err != nil {
				return err
			}

			if len(tags) == 0 {
				return fmt.Errorf("at least one tag must be specified using --tag")
			}

			return input.Tag(modelPath, id, title, tags)
		},
	}

	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the decision model (optional if set in config)")
	cmd.Flags().StringVar(&idOrTitle, "id", "", "ID or title of the decision to tag (e.g. 0001, 'my-decision')")
	cmd.Flags().StringArrayVar(&tags, "tag", nil, "Tag(s) to add to the decision (can be repeated)")

	return cmd
}
