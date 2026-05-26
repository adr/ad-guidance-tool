package cmd

import (
	util "github.com/adr/ad-guidance-tool/internal/adapter/command"
	adgmcp "github.com/adr/ad-guidance-tool/internal/adapter/mcp"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newMCPCommand())
}

func newMCPCommand() *cobra.Command {
	var modelPath string

	c := &cobra.Command{
		Use:    "mcp",
		Short:  "Start the ADG MCP server for AI tool integration",
		Hidden: true,
		Long: `Start the ADG MCP server for AI tool integration.

The server communicates over stdio using the Model Context Protocol (MCP) and
provides AI assistants (e.g. VS Code Copilot, Claude Desktop) with tools to
read ADRs, access the ADE rule DSL reference, and browse existing rule files.

Configure it in your AI tool's MCP settings, for example in .vscode/mcp.json:

  {
    "servers": {
      "adg": {
        "command": "adg",
        "args": ["mcp", "--model", "./docs/adr"]
      }
    }
  }

If --model is omitted, the default model path from the adg config is used.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedPath, err := util.ResolveModelPathOrDefault(modelPath, configSvc)
			if err != nil {
				return err
			}
			return adgmcp.Serve(resolvedPath, decisionSvc)
		},
	}

	c.Flags().StringVar(&modelPath, "model", "", "Path to the decision model (optional if configured)")
	return c
}
