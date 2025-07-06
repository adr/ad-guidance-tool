package interactor

import (
	"adg/internal/domain/decision"
)

func ResolveDecisionByIdOrTitle(modelPath, id, title string, service decision.DecisionService) (*decision.Decision, error) {
	if id != "" {
		return service.GetDecisionByID(modelPath, id)
	}
	return service.GetDecisionByTitle(modelPath, title)
}
