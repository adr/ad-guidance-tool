package ade

import (
	"fmt"
	"os"

	validateapp "github.com/adr/ad-guidance-tool/internal/ade/application/validate"
	"github.com/adr/ad-guidance-tool/internal/ade/domain"
	"github.com/spf13/cobra"
)

const (
	FLAG_VALIDATE_INPUT       = "input"
	FLAG_VALIDATE_INPUT_SHORT = "i"
	FLAG_VALIDATE_INPUT_USAGE = "input path(s) containing rule files (required, can be file or directory)"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate rule file(s) syntax without executing any plugin",
	Long: `Validate one or more rule files for syntax errors.
	
This command parses the rule files and reports any syntax or semantic errors
without actually executing any plugin.
Examples:
  ade validate -i rules/0001.rule
  ade validate -i rules/
  ade validate -i rules/0001.rule -i rules/0002.rule`,
	Run: validateCommand,
}

func init() {
	enforceCmd.AddCommand(validateCmd)

	validateCmd.Flags().StringArrayP(FLAG_VALIDATE_INPUT, FLAG_VALIDATE_INPUT_SHORT, []string{}, FLAG_VALIDATE_INPUT_USAGE)
	validateCmd.MarkFlagRequired(FLAG_VALIDATE_INPUT)
}

func validateCommand(cmd *cobra.Command, args []string) {
	inputs, err := cmd.Flags().GetStringArray(FLAG_VALIDATE_INPUT)
	domain.CheckFatalError(err, "reading input flag")

	if len(inputs) == 0 {
		fmt.Fprintf(os.Stderr, "Error: at least one input path required\n")
		os.Exit(1)
	}

	err = validateapp.Validate(validateapp.ValidateInput{
		Paths: inputs,
	})
	if err != nil {
		os.Exit(1)
	}
}
