package decision

import (
	domain "adg/internal/domain/decision"
	"fmt"
)

type AddDecisionsPresenter struct{}

func NewAddPresenter() *AddDecisionsPresenter {
	return &AddDecisionsPresenter{}
}

func (p *AddDecisionsPresenter) Added(successes []*domain.Decision, failures map[string]error) {
	for _, decision := range successes {
		fmt.Printf("Decision %s (%s) added successfully.\n", decision.Title, decision.ID)
	}
	for title, err := range failures {
		fmt.Printf("Failed to add decision %q: %v\n", title, err)
	}
}
