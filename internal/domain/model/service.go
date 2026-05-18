package model

import (
	decisiondomain "adg/internal/domain/decision"
	"adg/internal/domain/decision/madr"
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type ModelService interface {
	CreateModel(modelPath string) error
	Exists(modelPath string) bool
	Validate(modelPath string) ([]ValidationIssue, error)
}

// ValidationIssue is one finding reported by Validate. Issues are grouped by
// decision ID for per-decision output formatting.
type ValidationIssue struct {
	ID      string
	Message string
}

type ModelServiceImplementation struct {
	modelRepo    ModelRepository
	decisionRepo decisiondomain.DecisionRepository
}

func NewModelService(modelRepo ModelRepository, decisionRepo decisiondomain.DecisionRepository) ModelService {
	return &ModelServiceImplementation{
		modelRepo:    modelRepo,
		decisionRepo: decisionRepo,
	}
}

func (s *ModelServiceImplementation) CreateModel(modelPath string) error {
	return s.modelRepo.CreateModel(modelPath)
}

func (s *ModelServiceImplementation) Exists(modelPath string) bool {
	return s.modelRepo.Exists(modelPath)
}

var (
	statusRe     = regexp.MustCompile(`^(proposed|rejected|accepted|deprecated|superseded by ADR-[0-9]{4})$`)
	supersededRe = regexp.MustCompile(`^superseded by ADR-([0-9]{4})$`)
	idRe         = regexp.MustCompile(`^[0-9]{4}$`)
)

// Validate runs MADR-shape and ADG-extension integrity checks across every ADR
// in the model directory. Each check produces zero or more issues; a non-empty
// result indicates failures (the caller decides exit code).
//
// Rules (per the design spec, §3 — File Format & Data Model):
//
//   1. Filename matches NNNN-slug.md (enforced by repo.LoadAll already, but
//      defended here).
//   2. H1 title present and non-empty.
//   3. Required H2 sections present: Context and Problem Statement,
//      Considered Options, Decision Outcome.
//   4. Considered Options has at least one bullet.
//   5. When status is "accepted" and legacy-outcome is false: Decision Outcome
//      contains `Chosen option: "X"` and X matches a bullet in Considered Options.
//   6. Status is exactly one of the MADR vocabulary entries.
//   7. Supersession forward integrity: if status is "superseded by ADR-X",
//      then ADR-X exists AND has <self> in its supersedes list.
//   8. Supersession reverse integrity: every supersedes entry points to an
//      existing ADR whose status references <self>.
//   9. Comment text non-empty and not solely numeric (defends against §A.1
//      regression).
func (s *ModelServiceImplementation) Validate(modelPath string) ([]ValidationIssue, error) {
	decisions, err := s.decisionRepo.LoadAll(modelPath)
	if err != nil {
		return nil, err
	}

	sort.Slice(decisions, func(i, j int) bool {
		return decisions[i].ID < decisions[j].ID
	})

	byID := map[string]decisiondomain.Decision{}
	for _, d := range decisions {
		byID[d.ID] = d
	}

	var issues []ValidationIssue
	for _, d := range decisions {
		issues = append(issues, s.validateDecision(modelPath, d, byID)...)
	}
	return issues, nil
}

func (s *ModelServiceImplementation) validateDecision(modelPath string, d decisiondomain.Decision, byID map[string]decisiondomain.Decision) []ValidationIssue {
	var issues []ValidationIssue
	add := func(msg string) { issues = append(issues, ValidationIssue{ID: d.ID, Message: msg}) }

	if !idRe.MatchString(d.ID) {
		add("filename does not match NNNN-slug.md")
	}

	if strings.TrimSpace(d.Title) == "" {
		add("H1 title is missing or empty")
	}

	body, err := s.decisionRepo.LoadBody(modelPath, d.ID)
	if err != nil {
		add(fmt.Sprintf("failed to load body: %v", err))
		return issues
	}
	parsed, err := madr.ParseBody(body)
	if err != nil {
		add(fmt.Sprintf("failed to parse body: %v", err))
		return issues
	}

	if _, ok := parsed.Sections["context"]; !ok {
		add("missing required section: Context and Problem Statement")
	}
	if _, ok := parsed.Sections["options"]; !ok {
		add("missing required section: Considered Options")
	} else if len(parsed.Options) == 0 {
		add("Considered Options section has no bullets")
	}
	if _, ok := parsed.Sections["outcome"]; !ok {
		add("missing required section: Decision Outcome")
	}

	if d.Status == "accepted" && !d.LegacyOutcome {
		if parsed.ChosenOption == "" {
			add(`Decision Outcome must contain Chosen option: "..." when status is accepted`)
		} else if !slices.Contains(parsed.Options, parsed.ChosenOption) {
			add(fmt.Sprintf("chosen option %q is not in Considered Options", parsed.ChosenOption))
		}
	}

	if d.Status != "" && !statusRe.MatchString(d.Status) {
		add(fmt.Sprintf("status %q is not valid MADR vocabulary", d.Status))
	}

	if m := supersededRe.FindStringSubmatch(d.Status); m != nil {
		successorID := m[1]
		successor, ok := byID[successorID]
		if !ok {
			add(fmt.Sprintf("status references ADR-%s but no such ADR exists", successorID))
		} else if !slices.Contains(successor.Supersedes, d.ID) {
			add(fmt.Sprintf("ADR-%s status says it supersedes us, but its supersedes list does not include %s", successorID, d.ID))
		}
	}

	for _, predID := range d.Supersedes {
		pred, ok := byID[predID]
		if !ok {
			add(fmt.Sprintf("supersedes %s but no such ADR exists", predID))
			continue
		}
		expected := fmt.Sprintf("superseded by ADR-%s", d.ID)
		if pred.Status != expected {
			add(fmt.Sprintf("supersedes %s but ADR-%s status is %q, not %q", predID, predID, pred.Status, expected))
		}
	}

	for i, c := range d.Comments {
		if strings.TrimSpace(c.Text) == "" {
			add(fmt.Sprintf("comment %d has empty text", i+1))
			continue
		}
		if _, err := strconv.Atoi(strings.TrimSpace(c.Text)); err == nil {
			add(fmt.Sprintf("comment %d has placeholder text %q; run adg migrate to recover", i+1, c.Text))
		}
	}

	return issues
}
