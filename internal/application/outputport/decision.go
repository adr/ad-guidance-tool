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

// DecisionPrint receives the rendered body text for each decision being viewed.
// The legacy DecisionContent struct (one parsed field per section) is gone;
// printers now receive the raw body and either pass it through or run the
// madr.ParseBody helper themselves if they need section-aware output.
type DecisionPrint interface {
	Printed(bodies []DecisionBody, sections map[string]bool)
}

// DecisionBody is what the DecisionPrint port receives — the decision's ID and
// the raw markdown body. Sections filtering happens in the printer.
type DecisionBody struct {
	ID   string
	Body string
}

type DecisionRevise interface {
	Revised(originalID, revisedID string)
}

type DecisionTag interface {
	Tagged(decisionID string, tags []string)
}
