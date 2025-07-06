package outputport

import domain "adg/internal/domain/decision"

type DecisionAdd interface {
	Added(successes []*domain.Decision, failures map[string]error)
}

type DecisionComment interface {
	Commented(decisionID, author, comment string)
}

type DecisionDecide interface {
	Decided(decisionID string)
}

type DecisionEdit interface {
	Edited(decisionID string)
}

type DecisionLink interface {
	Linked(sourceID, targetID, tag, reverseTag string)
}

type DecisionList interface {
	Listed(decisions []domain.Decision, format string)
}

type DecisionPrint interface {
	Printed(content []domain.DecisionContent, sections map[string]bool)
}

type DecisionRevise interface {
	Revised(originalID, revisedID string)
}

type DecisionTag interface {
	Tagged(decisionID string, tags []string)
}
