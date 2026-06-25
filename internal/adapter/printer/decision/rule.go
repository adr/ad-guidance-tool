package decision

import (
	"fmt"
)

type RulePresenter struct{}

func NewRulePresenter() *RulePresenter {
	return &RulePresenter{}
}

func (p *RulePresenter) RuleGenerated(decisionID, ruleFilePath string) {
	fmt.Printf("Rule file for decision %s created at: %s\n", decisionID, ruleFilePath)
}
