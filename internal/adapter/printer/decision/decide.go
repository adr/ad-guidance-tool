package decision

import "fmt"

type DecidePresenter struct{}

func NewDecidePresenter() *DecidePresenter {
	return &DecidePresenter{}
}

func (p *DecidePresenter) Decided(decisionID string) {
	fmt.Printf("Decision %s has been marked as decided.\n", decisionID)
}
