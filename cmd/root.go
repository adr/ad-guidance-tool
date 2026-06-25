package cmd

import (
	"log"

	decisiondomain "github.com/adr/ad-guidance-tool/internal/domain/decision"
	modeldomain "github.com/adr/ad-guidance-tool/internal/domain/model"
	configinfra "github.com/adr/ad-guidance-tool/internal/infrastructure/config"
	decisioninfra "github.com/adr/ad-guidance-tool/internal/infrastructure/decision"
	modelinfra "github.com/adr/ad-guidance-tool/internal/infrastructure/model"

	"github.com/spf13/cobra"
)

// Version is set at build time via -ldflags.
var Version = "1.0.3-dev"

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

	rootCmd.Version = Version

	// todo: check if index needs to be rebuilt

	return rootCmd.Execute()
}
