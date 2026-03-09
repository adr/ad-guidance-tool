package decision

import (
	"fmt"

	util "github.com/adr/ad-guidance-tool/internal/adapter/command"
	"github.com/adr/ad-guidance-tool/internal/application/inputport"
	domain "github.com/adr/ad-guidance-tool/internal/domain/config"

	"github.com/spf13/cobra"
)

func NewTagCommand(input inputport.DecisionTag, config domain.ConfigService) *cobra.Command {
	var modelPath, idOrTitle, id, title string
	var tags []string
	var err error

	cmd := &cobra.Command{
		Use:   "tag [tags...]",
		Short: "Categorizes a decision by adding one or more tags to its metadata",
		Long: `Categorizes a decision by adding one or more tags to its metadata.

You can provide tags either as positional arguments or via the --tag flag.

Examples:
  adg tag --id 0001 architecture urgent
  adg tag --id 0001 --tag architecture --tag urgent
  adg tag --id my-decision important`,
		RunE: func(cmd *cobra.Command, args []string) error {
			modelPath, err = util.ResolveModelPathOrDefault(modelPath, config)
			if err != nil {
				return err
			}

			err := util.ResolveIdOrTitle(idOrTitle, &id, &title)
			if err != nil {
				return err
			}

			// If no --tag flags provided, use positional arguments
			if len(tags) == 0 && len(args) > 0 {
				tags = args
			}

			if len(tags) == 0 {
				return fmt.Errorf("at least one tag must be specified (via arguments or --tag flag)")
			}

			return input.Tag(modelPath, id, title, tags)
		},
	}

	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the decision model (optional if set in config)")
	cmd.Flags().StringVar(&idOrTitle, "id", "", "ID or title of the decision to tag (e.g. 0001, 'my-decision')")
	cmd.Flags().StringArrayVar(&tags, "tag", nil, "Tag(s) to add to the decision (optional if using positional arguments)")

	return cmd
}
