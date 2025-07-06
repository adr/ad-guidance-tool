package config

type ConfigService interface {
	IsLoaded() bool
	// todo: generic method with anchor as parameter
	GetQuestionHeader() string
	GetCriteriaHeader() string
	GetOptionsHeader() string
	GetCommentsHeader() string
	GetOutcomeHeader() string
	GetAuthor() string
	GetDefaultModelPath() string
	Save(question, criteria, options, comments, outcome, author, modelPath string) (string, error)
	SetConfigPath(customPath string) error
	ResetAll() error
	ResetTemplateHeaders() error
}
