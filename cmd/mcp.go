package cmd

import (
	"fmt"
	"net"

	util "github.com/adr/ad-guidance-tool/internal/adapter/command"
	adgmcp "github.com/adr/ad-guidance-tool/internal/adapter/mcp"
	"github.com/spf13/cobra"
)

func init() {
	mcpCmd := newMCPCommand()
	mcpCmd.AddCommand(newMCPRunCommand())
	rootCmd.AddCommand(mcpCmd)
}

func newMCPCommand() *cobra.Command {
	var modelPath string

	c := &cobra.Command{
		Use:   "mcp",
		Short: "MCP server setup for AI tool integration",
		Long:  `Configure the ADG MCP server for use with VS Code Copilot or other MCP-compatible AI assistants.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			displayPath := modelPath
			if displayPath == "" {
				if configSvc.IsLoaded() && configSvc.GetDefaultModelPath() != "" {
					displayPath = configSvc.GetDefaultModelPath()
				} else {
					displayPath = "<path-to-model>"
				}
			}

			fmt.Printf(`ADG MCP server - AI tool integration

When running, the server provides AI assistants with tools to read ADRs,
access the ADE rule DSL reference, browse existing rule files, and validate
generated rule content.

Add to .vscode/mcp.json in your project:

  {
    "servers": {
      "adg": {
        "command": "adg",
        "args": ["mcp", "run", "--model", "%s"]
      }
    }
  }

Or run "MCP: Add Server" in VS Code and select "Command (stdio)" to
configure this interactively.

Other MCP-compatible AI tools use the same command and args but may require
a different config file and structure.

`, displayPath)
			return nil
		},
	}

	c.Flags().StringVar(&modelPath, "model", "", "Path to the ADR model (used in the generated config snippet)")
	return c
}

func newMCPRunCommand() *cobra.Command {
	var modelPath string
	var httpAddr string

	c := &cobra.Command{
		Use:    "run",
		Short:  "Start the ADG MCP server over stdio, or over Streamable HTTP with --http",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedPath, err := util.ResolveModelPathOrDefault(modelPath, configSvc)
			if err != nil {
				return err
			}
			if httpAddr == "" {
				return adgmcp.Serve(resolvedPath, decisionSvc)
			}

			ln, err := net.Listen("tcp", httpAddr)
			if err != nil {
				return fmt.Errorf("failed to listen on %s: %w", httpAddr, err)
			}
			// The bound address, not --http: port 0 or a hostname resolve only here.
			cmd.PrintErrf("ADG MCP server listening on http://%s%s (model: %s)\n", ln.Addr(), adgmcp.EndpointPath, resolvedPath)
			return adgmcp.ServeStreamableHTTP(ln, resolvedPath, decisionSvc)
		},
	}

	c.Flags().StringVar(&modelPath, "model", "", "Path to the decision model (optional if configured)")
	c.Flags().StringVar(&httpAddr, "http", "", "Serve over Streamable HTTP at this address (e.g. 127.0.0.1:8080) instead of stdio; the MCP endpoint is /mcp")
	return c
}
