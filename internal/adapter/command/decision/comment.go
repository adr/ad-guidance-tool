package decision

import (
	"fmt"

	util "github.com/adr/ad-guidance-tool/internal/adapter/command"
	"github.com/adr/ad-guidance-tool/internal/application/inputport"
	domain "github.com/adr/ad-guidance-tool/internal/domain/config"

	"github.com/spf13/cobra"
)

func NewCommentCommand(input inputport.DecisionComment, config domain.ConfigService) *cobra.Command {
	var modelPath, idOrTitle, id, title, text, authorFlag string

	cmd := &cobra.Command{
		Use:   "comment [comment-text...]",
		Short: "Add a comment to a decision",
		Long: `Adds a comment to the specified decision.

You can provide the comment text either as positional arguments or via the --text flag.

Examples:
  adg comment --id 0001 This is my comment text
  adg comment --id 0001 --text "This is my comment text"
  adg comment --id my-decision Great decision about architecture`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// If no --text flag provided, use positional arguments
			if text == "" && len(args) > 0 {
				text = joinArgs(args)
			}

			if text == "" {
				return fmt.Errorf("comment text must be provided (via arguments or --text flag)")
			}

			err := util.ResolveIdOrTitle(idOrTitle, &id, &title)
			if err != nil {
				return err
			}

			modelPath, err := util.ResolveModelPathOrDefault(modelPath, config)
			if err != nil {
				return err
			}

			author := authorFlag
			if author == "" {
				author = config.GetAuthor()
			}
			if author == "" {
				return fmt.Errorf("author must be provided using --author or set in config")
			}

			return input.Comment(modelPath, id, title, author, text)
		},
	}

	cmd.Flags().StringVar(&modelPath, "model", "", "Path to the model directory (optional if configured)")
	cmd.Flags().StringVar(&idOrTitle, "id", "", "ID or title of the decision to comment on (e.g. 0001, 'my-decision')")
	cmd.Flags().StringVar(&text, "text", "", "Text content of the comment (optional if using positional arguments)")
	cmd.Flags().StringVar(&authorFlag, "author", "", "Name of the commenter (overrides config)")

	return cmd
}
