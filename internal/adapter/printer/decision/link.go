package decision

import "fmt"

type LinkPresenter struct{}

func NewLinkPresenter() *LinkPresenter {
	return &LinkPresenter{}
}

func (p *LinkPresenter) Linked(sourceID, targetID, tag, reverseTag string) {
	fmt.Printf("Link added: %s →[%s]→ %s\n", sourceID, tag, targetID)
	if reverseTag != "" {
		fmt.Printf("Reverse link added: %s →[%s]→ %s\n", targetID, reverseTag, sourceID)
	}
}
