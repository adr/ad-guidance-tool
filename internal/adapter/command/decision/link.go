package decision

import (
	util "adg/internal/adapter/command"
	"adg/internal/application/inputport"
	domain "adg/internal/domain/config"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func NewLinkCommand(
	input inputport.DecisionLink,
	config domain.ConfigService,
) *cobra.Command {
	var modelPath, fromIdOrTitle, fromID, fromTitle, toIdOrTitle, toID, toTitle string
	var tag, reverseTag string
	var err error

	cmd := &cobra.Command{
		Use:   "link",
		Short: "Link two decisions using optional custom tags or default precedes/succeeds logic",
		Long: `Links a source decision (--from) to a target decision (--to).
	
Default behavior:
  - If no --tag or --reverse-tag is provided, the link is treated as "precedes" → "succeeds".
  - A check is performed when using the default tags as to not create a cycle. 

Custom tag behavior:
  - You may provide --tag (and optionally --reverse-tag) to use a custom relationship.
  - You may not use "precedes" or "succeeds" explicitly as --tag or --reverse-tag.
    These are reserved for the default implicit mode.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			modelPath, err = util.ResolveModelPathOrDefault(modelPath, config)
			if err != nil {
				return err
			}

			if err := util.ResolveIdOrTitle(fromIdOrTitle, &fromID, &fromTitle); err != nil {
				return fmt.Errorf("you must specify the decisions via --from by either providing the numbered id (e.g., 0001) or the name of the decision (e.g, 'my-decision')")
			}

			if err := util.ResolveIdOrTitle(toIdOrTitle, &toID, &toTitle); err != nil {
				return fmt.Errorf("you must specify the decisions via --to by either providing the numbered id (e.g., 0001) or the name of the decision (e.g, 'my-decision')")
			}

			if (fromID == "" && fromTitle == "") || (toID == "" && toTitle == "") {
				return fmt.Errorf("must provide both --from and --to with either ID or title")
			}

			// Reject reserved tags explicitly
			if strings.EqualFold(tag, "precedes") || strings.EqualFold(tag, "succeeds") ||
				strings.EqualFold(reverseTag, "precedes") || strings.EqualFold(reverseTag, "succeeds") {
				return fmt.Errorf(`you cannot use "precedes" or "succeeds" as custom tags; omit --tag and --reverse-tag to use them implicitly`)
			}

			// Tag logic
			finalTag := tag
			finalReverseTag := reverseTag

			if tag == "" && reverseTag == "" {
				finalTag = "precedes"
				finalReverseTag = "succeeds"
			} else if tag != "" && reverseTag == "" {
				finalReverseTag = tag
			}

			return input.Link(modelPath, fromID, fromTitle, toID, toTitle, finalTag, finalReverseTag)
		},
	}

	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the decision model (optional if set in config)")
	cmd.Flags().StringVar(&fromIdOrTitle, "from", "", "ID or title of the source decision (e.g. 0001, 'my-decision')")
	cmd.Flags().StringVar(&toIdOrTitle, "to", "", "ID or title of the target decision (e.g. 0002, 'other-decision')")
	cmd.Flags().StringVar(&tag, "tag", "", `Custom link tag (e.g. "invalidated by"). Cannot be "precedes" or "succeeds".`)
	cmd.Flags().StringVar(&reverseTag, "reverse-tag", "", `Optional reverse tag (e.g. "invalidates"). Cannot be "precedes" or "succeeds".`)

	return cmd
}
