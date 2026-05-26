package mcp

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	aderule "github.com/phi42/ad-enforcement-tool/dslref"
	decisiondomain "github.com/adr/ad-guidance-tool/internal/domain/decision"
	mcplib "github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// Serve starts the ADG MCP server over stdio.
func Serve(modelPath string, decisionSvc decisiondomain.DecisionService) error {
	s := buildServer(modelPath, decisionSvc)
	return mcpserver.ServeStdio(s)
}

func buildServer(modelPath string, decisionSvc decisiondomain.DecisionService) *mcpserver.MCPServer {
	s := mcpserver.NewMCPServer(
		"adg",
		"1.0.0",
		mcpserver.WithInstructions(`ADG (Architectural Decision Guidance) MCP server.
Use these tools to read architectural decision records (ADRs) and generate .rule files that enforce them.
Always call get_dsl_reference before writing any .rule file content to ensure correct syntax.`),
	)

	s.AddTool(
		mcplib.NewTool("list_adrs",
			mcplib.WithDescription("List all architectural decision records (ADRs) in the decision model. Returns each ADR's ID, title, and status."),
			mcplib.WithReadOnlyHintAnnotation(true),
		),
		listADRsHandler(modelPath, decisionSvc),
	)

	s.AddTool(
		mcplib.NewTool("get_adr",
			mcplib.WithDescription("Get the full content of an ADR by its ID. Returns the markdown content and the path where the corresponding .rule file should be placed."),
			mcplib.WithString("id",
				mcplib.Required(),
				mcplib.Description(`The ADR ID, e.g. "0001"`),
			),
			mcplib.WithReadOnlyHintAnnotation(true),
		),
		getADRHandler(modelPath, decisionSvc),
	)

	s.AddTool(
		mcplib.NewTool("get_dsl_reference",
			mcplib.WithDescription("Get the full ADE rule DSL language reference. Always read this before writing .rule file content."),
			mcplib.WithReadOnlyHintAnnotation(true),
		),
		getDSLReferenceHandler(),
	)

	s.AddTool(
		mcplib.NewTool("list_rule_files",
			mcplib.WithDescription("List all existing .rule files in the decision model directory with their full content. Use these as examples when generating new rule files."),
			mcplib.WithReadOnlyHintAnnotation(true),
		),
		listRuleFilesHandler(modelPath),
	)

	s.AddTool(
		mcplib.NewTool("validate_rule",
			mcplib.WithDescription("Validate the syntax and semantics of rule file content. Call this after generating a .rule file to confirm it is correct before presenting it to the user."),
			mcplib.WithString("content",
				mcplib.Required(),
				mcplib.Description("The full .rule file content to validate."),
			),
		),
		validateRuleHandler(),
	)

	return s
}

func listADRsHandler(modelPath string, decisionSvc decisiondomain.DecisionService) mcpserver.ToolHandlerFunc {
	return func(ctx context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
		decisions, err := decisionSvc.GetAllDecisions(modelPath)
		if err != nil {
			return mcplib.NewToolResultError(fmt.Sprintf("failed to list ADRs: %v", err)), nil
		}

		if len(decisions) == 0 {
			return mcplib.NewToolResultText("No ADRs found in the decision model."), nil
		}

		var sb strings.Builder
		for _, d := range decisions {
			fmt.Fprintf(&sb, "- ID: %s | Title: %s | Status: %s\n", d.ID, d.Title, d.Status)
		}
		return mcplib.NewToolResultText(sb.String()), nil
	}
}

func getADRHandler(modelPath string, decisionSvc decisiondomain.DecisionService) mcpserver.ToolHandlerFunc {
	return func(ctx context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
		id, err := req.RequireString("id")
		if err != nil {
			return mcplib.NewToolResultError(`parameter "id" is required`), nil
		}

		decision, err := decisionSvc.GetDecisionByID(modelPath, id)
		if err != nil {
			return mcplib.NewToolResultError(fmt.Sprintf("ADR %q not found: %v", id, err)), nil
		}

		adrPath, err := decisionSvc.GetDecisionFilePath(modelPath, decision.ID)
		if err != nil {
			return mcplib.NewToolResultError(fmt.Sprintf("failed to resolve ADR file path: %v", err)), nil
		}

		content, err := os.ReadFile(adrPath)
		if err != nil {
			return mcplib.NewToolResultError(fmt.Sprintf("failed to read ADR file: %v", err)), nil
		}

		base := strings.TrimSuffix(filepath.Base(adrPath), filepath.Ext(adrPath))
		ruleFilePath := filepath.Join(filepath.Dir(adrPath), base+".rule")

		ruleFileStatus := "does not exist yet (create it at: " + ruleFilePath + ")"
		if _, statErr := os.Stat(ruleFilePath); statErr == nil {
			ruleFileStatus = "already exists at: " + ruleFilePath
		}

		var sb strings.Builder
		fmt.Fprintf(&sb, "ADR ID:     %s\n", decision.ID)
		fmt.Fprintf(&sb, "Title:      %s\n", decision.Title)
		fmt.Fprintf(&sb, "Status:     %s\n", decision.Status)
		fmt.Fprintf(&sb, "File:       %s\n", adrPath)
		fmt.Fprintf(&sb, "Rule file:  %s\n\n", ruleFileStatus)
		fmt.Fprintf(&sb, "=== Content ===\n%s", string(content))

		return mcplib.NewToolResultText(sb.String()), nil
	}
}

func getDSLReferenceHandler() mcpserver.ToolHandlerFunc {
	return func(ctx context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
		return mcplib.NewToolResultText(aderule.Reference), nil
	}
}

func validateRuleHandler() mcpserver.ToolHandlerFunc {
	return func(ctx context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
		content, err := req.RequireString("content")
		if err != nil {
			return mcplib.NewToolResultError(`parameter "content" is required`), nil
		}

		if err := aderule.Validate(content); err != nil {
			return mcplib.NewToolResultError(fmt.Sprintf("Validation failed:\n%v", err)), nil
		}
		return mcplib.NewToolResultText("\u2713 Rule is valid."), nil
	}
}

func listRuleFilesHandler(modelPath string) mcpserver.ToolHandlerFunc {
	return func(ctx context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
		entries, err := os.ReadDir(modelPath)
		if err != nil {
			return mcplib.NewToolResultError(fmt.Sprintf("failed to read model directory: %v", err)), nil
		}

		var sb strings.Builder
		count := 0
		for _, entry := range entries {
			if entry.IsDir() || filepath.Ext(entry.Name()) != ".rule" {
				continue
			}
			content, err := os.ReadFile(filepath.Join(modelPath, entry.Name()))
			if err != nil {
				continue
			}
			fmt.Fprintf(&sb, "=== %s ===\n%s\n", entry.Name(), string(content))
			count++
		}

		if count == 0 {
			return mcplib.NewToolResultText("No .rule files found in the model directory."), nil
		}
		return mcplib.NewToolResultText(sb.String()), nil
	}
}
