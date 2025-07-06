package decision

import (
	"adg/internal/domain/config"
	domain "adg/internal/domain/decision"

	"fmt"
	"sort"
)

type PrintDecisionsPresenter struct {
	config config.ConfigService
}

func NewPrintPresenter(config config.ConfigService) *PrintDecisionsPresenter {
	return &PrintDecisionsPresenter{config: config}
}

func (p *PrintDecisionsPresenter) Printed(contents []domain.DecisionContent, sections map[string]bool) {
	sort.Slice(contents, func(i, j int) bool {
		return contents[i].ID < contents[j].ID
	})

	for _, d := range contents {
		fmt.Printf("===== Decision %s =====\n\n", d.ID)

		if sections["question"] && d.Question != "" {
			fmt.Println(p.config.GetQuestionHeader())
			fmt.Println(d.Question + "\n")
		}
		if sections["options"] && d.Options != "" {
			fmt.Println(p.config.GetOptionsHeader())
			fmt.Println(d.Options + "\n")
		}
		if sections["criteria"] && d.Criteria != "" {
			fmt.Println(p.config.GetCriteriaHeader())
			fmt.Println(d.Criteria + "\n")
		}
		if sections["outcome"] && d.Outcome != "" {
			fmt.Println(p.config.GetOutcomeHeader())
			fmt.Println(d.Outcome + "\n")
		}
		if sections["comments"] && d.Comments != "" {
			fmt.Println(p.config.GetCommentsHeader())
			fmt.Println(d.Comments + "\n")
		}
	}
}
