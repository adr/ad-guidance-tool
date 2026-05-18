package decision

import (
	"adg/internal/application/outputport"
	"adg/internal/domain/decision/madr"

	"fmt"
	"sort"
)

// PrintDecisionsPresenter renders ADR bodies for the `adg view` command.
// The section names below are MADR canonical (no more configurable headers);
// in PR 1b this printer accepts pre-loaded body text and uses madr.ParseBody to
// pluck out individual sections when `--section` filtering is requested.
type PrintDecisionsPresenter struct{}

func NewPrintPresenter() *PrintDecisionsPresenter {
	return &PrintDecisionsPresenter{}
}

// Printed implements outputport.DecisionPrint. `sections` filters by canonical
// MADR section key ("context", "drivers", "options", "outcome", "more", "comments").
// An empty map prints the whole body.
func (p *PrintDecisionsPresenter) Printed(bodies []outputport.DecisionBody, sections map[string]bool) {
	sort.Slice(bodies, func(i, j int) bool {
		return bodies[i].ID < bodies[j].ID
	})

	for _, b := range bodies {
		fmt.Printf("===== Decision %s =====\n\n", b.ID)

		if len(sections) == 0 {
			fmt.Println(b.Body)
			continue
		}

		parsed, err := madr.ParseBody(b.Body)
		if err != nil {
			fmt.Printf("(failed to parse body: %v)\n", err)
			continue
		}
		for _, key := range []string{"context", "drivers", "options", "outcome", "pros-cons", "more", "comments"} {
			if !sections[key] {
				continue
			}
			if sec, ok := parsed.Sections[key]; ok {
				fmt.Println(sec)
			}
		}
	}
}
