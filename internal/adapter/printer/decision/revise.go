package decision

import "fmt"

type ReviseDecisionPresenter struct{}

func NewRevisePresenter() *ReviseDecisionPresenter {
	return &ReviseDecisionPresenter{}
}

func (p *ReviseDecisionPresenter) Revised(originalID, revisedID string) {
	fmt.Printf("Successfully revised decision %s → new decision %s\n", originalID, revisedID)
}
