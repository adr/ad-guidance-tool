package inputport

type DecisionAdd interface {
	Add(modelPath string, titles []string) error
}

type DecisionComment interface {
	Comment(modelPath, id, title, author, comment string) error
}

type DecisionDecide interface {
	Decide(modelPath, id, title, option, reason, author string, enforceOption bool) error
}

type DecisionEdit interface {
	Edit(modelPath string, id string, title string, question *string, options *[]string, criteria *string) error
}

type DecisionLink interface {
	Link(modelPath, sourceID, sourceTitle, targetID, targetTitle, tag, reverseTag string) error
}

type DecisionList interface {
	ListDecisions(modelPath string, filters map[string][]string, format string) error
}

type DecisionPrint interface {
	Print(modelPath string, ids []string, titles []string, sections map[string]bool) error
}

type DecisionRevise interface {
	ReviseDecision(modelPath, id, title string) error
}

type DecisionTag interface {
	Tag(modelPath, id, title string, tags []string) error
}
