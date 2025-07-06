package model

import (
	"adg/internal/application/inputport"

	"github.com/spf13/cobra"
)

func NewInitCommand(input inputport.ModelInit) *cobra.Command {
	return &cobra.Command{
		Use:   "init <model-path>",
		Short: "Initializes a new model",
		Long:  `Creates a directory using the provided path and initializes an empty index file`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			modelPath := args[0]
			return input.Init(modelPath)
		},
	}
}
