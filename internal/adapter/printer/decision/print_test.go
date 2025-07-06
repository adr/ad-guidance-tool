package decision

import (
	"adg/internal/domain/decision"
	svc_mocks "adg/mocks/service"
	"strings"
	"testing"
)

func TestPrinted_AllSections(t *testing.T) {
	mockConfig := new(svc_mocks.ConfigService)
	presenter := NewPrintPresenter(mockConfig)

	mockConfig.On("GetQuestionHeader").Return("## Question")
	mockConfig.On("GetOptionsHeader").Return("## Options")
	mockConfig.On("GetCriteriaHeader").Return("## Criteria")
	mockConfig.On("GetOutcomeHeader").Return("## Outcome")
	mockConfig.On("GetCommentsHeader").Return("## Comments")

	contents := []decision.DecisionContent{
		{
			ID:       "0002",
			Question: "How should we do it?",
			Options:  "Option A\nOption B",
			Criteria: "Must be easy",
			Outcome:  "Chose Option B",
			Comments: "Comment 1",
		},
		{
			ID:       "0001",
			Question: "What should we do?",
			Options:  "Option A\nOption B",
			Criteria: "Must be fast",
			Outcome:  "Chose Option A",
			Comments: "Comment 1",
		},
	}

	sections := map[string]bool{
		"question": true,
		"options":  true,
		"criteria": true,
		"outcome":  true,
		"comments": true,
	}

	output := captureOutput(func() {
		presenter.Printed(contents, sections)
	})

	for _, expected := range []string{
		"===== Decision 0001 =====",
		"## Question", "What should we do?",
		"## Options", "Option A",
		"## Criteria", "Must be fast",
		"## Outcome", "Chose Option A",
		"## Comments", "Comment 1",
		"===== Decision 0002 =====",
		"## Question", "How should we do it?",
		"## Options", "Option A",
		"## Criteria", "Must be easy",
		"## Outcome", "Chose Option B",
		"## Comments", "Comment 1",
	} {
		if !strings.Contains(output, expected) {
			t.Errorf("Expected output to contain: %q", expected)
		}
	}
}

func TestPrinted_PartialSections(t *testing.T) {
	mockConfig := new(svc_mocks.ConfigService)
	presenter := NewPrintPresenter(mockConfig)

	mockConfig.On("GetQuestionHeader").Return("## Question")
	mockConfig.On("GetOptionsHeader").Return("## Options")
	mockConfig.On("GetOutcomeHeader").Return("## Outcome")

	contents := []decision.DecisionContent{
		{
			ID:       "0002",
			Question: "Another question",
			Options:  "",
			Outcome:  "Outcome",
		},
	}

	sections := map[string]bool{
		"question": true,
		"options":  true, // but empty
		"outcome":  true,
	}

	output := captureOutput(func() {
		presenter.Printed(contents, sections)
	})

	if !strings.Contains(output, "===== Decision 0002 =====") {
		t.Error("Expected section header for decision 0002")
	}
	if !strings.Contains(output, "## Question") || !strings.Contains(output, "Another question") {
		t.Error("Expected question section to be printed")
	}
	if strings.Contains(output, "## Options") {
		t.Error("Did not expect options section to be printed (empty)")
	}
	if !strings.Contains(output, "## Outcome") {
		t.Error("Expected outcome section to be printed")
	}
}
