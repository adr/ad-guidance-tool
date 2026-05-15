package validate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adr/ad-guidance-tool/internal/ade/domain"
)

// ValidateInput contains paths to validate
type ValidateInput struct {
	Paths []string
}

// Validate validates rule files and reports errors
func Validate(input ValidateInput) error {
	// Collect all rule files
	var ruleFiles []string
	for _, path := range input.Paths {
		files, err := collectRuleFiles(path)
		if err != nil {
			return fmt.Errorf("collecting files from %q: %w", path, err)
		}
		ruleFiles = append(ruleFiles, files...)
	}

	if len(ruleFiles) == 0 {
		return fmt.Errorf("no rule files found")
	}

	// Validate each file
	hasErrors := false
	for _, file := range ruleFiles {
		if err := validateFile(file); err != nil {
			hasErrors = true
			fmt.Fprintf(os.Stderr, "X %s: %v\n", file, err)
		} else {
			fmt.Fprintf(os.Stdout, "✓ %s\n", file)
		}
	}

	if hasErrors {
		return fmt.Errorf("validation failed for one or more files")
	}

	return nil
}

// collectRuleFiles expands a path to all .rule files
func collectRuleFiles(path string) ([]string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		// Single file
		if strings.HasSuffix(path, ".rule") {
			return []string{path}, nil
		}
		return nil, fmt.Errorf("file %q does not have .rule extension", path)
	}

	// Directory - find all .rule files
	var files []string
	err = filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(p, ".rule") {
			files = append(files, p)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

// validateFile validates a single rule file
func validateFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	_, err = domain.ParseDSL(string(content))
	return err
}
