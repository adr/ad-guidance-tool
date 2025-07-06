package decision

import "fmt"

type EditDecisionPresenter struct{}

func NewEditPresenter() *EditDecisionPresenter {
	return &EditDecisionPresenter{}
}

func (p *EditDecisionPresenter) Edited(decisionID string) {
	fmt.Printf("Decision %s updated successfully.\n", decisionID)
}
