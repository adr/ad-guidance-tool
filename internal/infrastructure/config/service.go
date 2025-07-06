package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type ConfigServiceViper struct {
	v      *viper.Viper
	loaded bool
}

func NewConfigService() (*ConfigServiceViper, error) {
	configPath, err := resolveActiveConfigPath()
	if err != nil {
		return &ConfigServiceViper{loaded: false}, err
	}

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return &ConfigServiceViper{v: v, loaded: false}, nil
	}

	return &ConfigServiceViper{v: v, loaded: true}, nil
}

func (c *ConfigServiceViper) IsLoaded() bool {
	return c.loaded
}

func (c *ConfigServiceViper) GetQuestionHeader() string {
	question := c.v.GetString("question_header")
	if question == "" {
		question = "Question"
	}
	return question
}

func (c *ConfigServiceViper) GetCriteriaHeader() string {
	criteria := c.v.GetString("criteria_header")
	if criteria == "" {
		criteria = "Criteria"
	}
	return criteria
}

func (c *ConfigServiceViper) GetOptionsHeader() string {
	options := c.v.GetString("options_header")
	if options == "" {
		options = "Options"
	}
	return options
}

func (c *ConfigServiceViper) GetCommentsHeader() string {
	comments := c.v.GetString("comments_header")
	if comments == "" {
		comments = "Comments"
	}
	return comments
}

func (c *ConfigServiceViper) GetOutcomeHeader() string {
	outcome := c.v.GetString("outcome_header")
	if outcome == "" {
		outcome = "Outcome"
	}
	return outcome
}

func (c *ConfigServiceViper) GetAuthor() string {
	return c.v.GetString("author")
}

func (c *ConfigServiceViper) GetDefaultModelPath() string {
	return c.v.GetString("default_model")
}

func (c *ConfigServiceViper) Save(question, criteria, options, comments, outcome, author, modelPath string) (string, error) {
	actualPath, err := resolveActiveConfigPath()
	if err != nil {
		return "", err
	}

	v := viper.New()
	v.SetConfigFile(actualPath)
	v.SetConfigType("yaml")
	_ = v.ReadInConfig()

	if question != "" {
		v.Set("question_header", question)
	}
	if criteria != "" {
		v.Set("criteria_header", criteria)
	}
	if options != "" {
		v.Set("options_header", options)
	}
	if comments != "" {
		v.Set("comments_header", comments)
	}
	if outcome != "" {
		v.Set("outcome_header", outcome)
	}
	if author != "" {
		v.Set("author", author)
	}
	if modelPath != "" {
		v.Set("default_model", modelPath)
	}

	if err := v.WriteConfigAs(actualPath); err != nil {
		return "", fmt.Errorf("failed to write config to %s: %w", actualPath, err)
	}

	return actualPath, nil
}

func (c *ConfigServiceViper) SetConfigPath(customPath string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not determine home directory: %w", err)
	}

	defaultPath := filepath.Join(homeDir, ".adgconfig.yaml")

	v := viper.New()
	v.SetConfigFile(defaultPath)
	v.SetConfigType("yaml")
	_ = v.ReadInConfig()

	v.Set("custom_config_path", customPath)

	if err := v.WriteConfigAs(defaultPath); err != nil {
		return fmt.Errorf("failed to update config path redirect: %w", err)
	}

	return nil
}

func (c *ConfigServiceViper) ResetAll() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not determine home directory: %w", err)
	}

	defaultConfigPath := filepath.Join(homeDir, ".adgconfig.yaml")

	// just delete the default config (removes redirect), not the custom file
	if _, err := os.Stat(defaultConfigPath); err == nil {
		if err := os.Remove(defaultConfigPath); err != nil {
			return fmt.Errorf("failed to delete default config at %s: %w", defaultConfigPath, err)
		}
	}

	return nil
}

func (c *ConfigServiceViper) ResetTemplateHeaders() error {
	targetPath, err := resolveActiveConfigPath()
	if err != nil {
		return err
	}

	v := viper.New()
	v.SetConfigFile(targetPath)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config at %s: %w", targetPath, err)
	}

	var raw map[string]interface{}
	if err := v.Unmarshal(&raw); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	for _, key := range []string{
		"question_header",
		"criteria_header",
		"options_header",
		"comments_header",
		"outcome_header",
	} {
		delete(raw, key)
	}

	content, err := yaml.Marshal(raw)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(targetPath, content, 0644); err != nil {
		return fmt.Errorf("failed to write updated config to %s: %w", targetPath, err)
	}

	return nil
}

func resolveActiveConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine home directory: %w", err)
	}

	defaultPath := filepath.Join(homeDir, ".adgconfig.yaml")
	activePath := defaultPath

	v := viper.New()
	v.SetConfigFile(defaultPath)
	v.SetConfigType("yaml")
	_ = v.ReadInConfig()

	if custom := v.GetString("custom_config_path"); custom != "" {
		activePath = custom
	}

	return activePath, nil
}
