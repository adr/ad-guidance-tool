package decision

import (
	"fmt"
	"strings"
)

type TagDecisionPresenter struct{}

func NewTagPresenter() *TagDecisionPresenter {
	return &TagDecisionPresenter{}
}

func (p *TagDecisionPresenter) Tagged(decisionID string, tags []string) {
	fmt.Printf("Tags [%s] added to decision %s\n", strings.Join(tags, ", "), decisionID)
}
