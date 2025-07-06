package config

import (
	"fmt"

	util "adg/internal/adapter/command"
	domain "adg/internal/domain/config"

	"github.com/spf13/cobra"
)

func NewSetCommand(config domain.ConfigService) *cobra.Command {
	var template, question, criteria, options, comments, outcome, author, modelPath, configPath string

	cmd := &cobra.Command{
		Use:   "set-config",
		Short: "Set persistent configuration values",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cmd.Flags().NFlag() == 0 {
				return fmt.Errorf("at least one configuration flag must be provided")
			}

			if configPath != "" {
				if err := config.SetConfigPath(configPath); err != nil {
					return err
				}
				fmt.Printf("Config path set to: %s\n", configPath)

			}

			if template != "" {
				preset, err := util.GetTemplateSections(template)
				if err != nil {
					return err
				}

				if question == "" {
					question = preset["question"]
				}
				if criteria == "" {
					criteria = preset["criteria"]
				}
				if options == "" {
					options = preset["options"]
				}
				if comments == "" {
					comments = preset["comments"]
				}
				if outcome == "" {
					outcome = preset["outcome"]
				}
			}

			if question != "" || criteria != "" || options != "" || comments != "" || outcome != "" || author != "" || modelPath != "" {
				path, err := config.Save(question, criteria, options, comments, outcome, author, modelPath)
				if err != nil {
					return err
				}

				fmt.Printf("Config saved to: %s\n", path)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&template, "template", "", "Apply a template to auto-fill section headers (available: Nygard, MADR). Use reset-config --template to go back to the default.")
	cmd.Flags().StringVar(&question, "question", "", "Define section header for decision question (default: Question)")
	cmd.Flags().StringVar(&criteria, "criteria", "", "Define section header for decision criteria (default: Criteria)")
	cmd.Flags().StringVar(&options, "options", "", "Define section header for decision options (default: Options)")
	cmd.Flags().StringVar(&comments, "comments", "", "Define section header for decision comments (default: Comments)")
	cmd.Flags().StringVar(&outcome, "outcome", "", "Define section header for decision outcome (default: Outcome)")
	cmd.Flags().StringVar(&author, "author", "", "Default author name to use")
	cmd.Flags().StringVar(&modelPath, "model", "", "Default model path")
	cmd.Flags().StringVar(&configPath, "config-path", "", "Optional path to store the configuration file (default: $HOME/.adgconfig.yaml)")

	return cmd
}
