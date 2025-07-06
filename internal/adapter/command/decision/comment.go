package decision

import (
	util "adg/internal/adapter/command"
	"adg/internal/application/inputport"
	domain "adg/internal/domain/config"
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommentCommand(input inputport.DecisionComment, config domain.ConfigService) *cobra.Command {
	var modelPath, idOrTitle, id, title, text, authorFlag string

	cmd := &cobra.Command{
		Use:   "comment",
		Short: "Add a comment to a decision",
		Long: `Adds a comment to the specified decision.
You must provide --id to identify the decision and a --text for the comment.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if text == "" {
				return fmt.Errorf("--text is required to provide the comment")
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
	cmd.Flags().StringVar(&text, "text", "", "Text content of the comment (required)")
	cmd.Flags().StringVar(&authorFlag, "author", "", "Name of the commenter (overrides config)")

	return cmd
}
