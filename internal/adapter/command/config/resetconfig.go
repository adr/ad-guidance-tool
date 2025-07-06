package config

import (
	"fmt"

	domain "adg/internal/domain/config"

	"github.com/spf13/cobra"
)

func NewResetCommand(config domain.ConfigService) *cobra.Command {
	var onlyTemplate bool

	cmd := &cobra.Command{
		Use:   "reset-config",
		Short: "Reset all configuration (or only template headers with --template)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if onlyTemplate {
				if err := config.ResetTemplateHeaders(); err != nil {
					return err
				}
				fmt.Println("Template header configuration reset.")
			} else {
				if err := config.ResetAll(); err != nil {
					return err
				}
				fmt.Println("Full configuration reset completed.")
			}
			return nil
		},
	}

	cmd.Flags().BoolVar(&onlyTemplate, "template", false, "Reset only template-related headers (question, criteria, options, comments, outcome)")

	return cmd
}
