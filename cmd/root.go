package cmd

import (
	decisiondomain "adg/internal/domain/decision"
	modeldomain "adg/internal/domain/model"
	configinfra "adg/internal/infrastructure/config"
	decisioninfra "adg/internal/infrastructure/decision"
	modelinfra "adg/internal/infrastructure/model"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "adg",
	Short: "Architectural Decision Guidance CLI",
	Long:  "CLI tool for managing architectural decision records and models",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true, // hides completion cmd from help text but it is still available
	},
}

var configSvc, err = configinfra.NewConfigService()
var decisionRepo = decisioninfra.NewFileDecisionRepository(configSvc)
var modelRepo = modelinfra.NewFileModelRepository()
var modelSvc = modeldomain.NewModelService(modelRepo, decisionRepo)
var decisionSvc = decisiondomain.NewDecisionService(decisionRepo)

func Execute() error {
	if err != nil {
		log.Fatalf("failed to initialize config service: %v", err)
	}

	// todo: check if index needs to be rebuilt

	return rootCmd.Execute()
}
