package decision

type DecisionRepository interface {
	Create(modelPath, subFolderPath string, decision *Decision, decisionContent *DecisionContent) (*Decision, error)
	Save(modelPath string, decision *Decision) error
	Copy(srcPath, dstPath, decisionID string) error
	LoadById(modelPath, id string) (*Decision, error)
	LoadByTitle(modelPath, title string) (*Decision, error)
	LoadAllByIndex(modelPath string) ([]Decision, error)
	LoadAllByData(modelPath string) ([]Decision, error)
	LoadDecisionContentRaw(modelPath, decisionID string) (string, error)
	LoadDecisionContent(modelPath, decisionID string) (*DecisionContent, error)
	UpdateSection(modelPath, decisionID, anchorName string, lines []string) error
	AppendCommentSection(modelPath, decisionID, commentText string, commentNumber int, author, date string) error
	AppendOutcomeSection(modelPath, decisionID, outcome string) error
	OptionExists(modelPath, decisionID, option string) (bool, error)
	ResolveOptionNumber(modelPath, decisionID, option string) (int, error)
	FindDecisionFile(modelPath, decisionID string) (string, error)
}
