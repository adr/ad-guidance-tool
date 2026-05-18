package decision

import (
	"adg/internal/domain/decision/madr"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

// DecisionService is the application-facing API for ADR mutations.
//
// The PR 1b implementation focuses on data-layer correctness:
//   - Comment stores text in Comment.Text (architectural fix for §A.1; the
//     legacy code stored a placeholder count in Comment.Comment and lost the
//     real text). Comments are append-only; the repository regenerates the
//     ## Comments body section from the Decision.Comments list on every Save.
//   - Edit/Decide/Link/Tag/Revise/Copy/Comment all operate on MADR-shaped
//     state. UX changes (renamed flags, status gating, replace-mode editing)
//     are deferred to PR 2 / PR 3.
type DecisionService interface {
	AddNew(modelPath, title string) (*Decision, error)
	AddExisting(sourceModelPath, targetModelPath string, decision *Decision, body string, increment int) (*Decision, error)
	GetAllDecisions(modelPath string) ([]Decision, error)
	GetDecisionByID(modelPath, id string) (*Decision, error)
	GetDecisionByTitle(modelPath, title string) (*Decision, error)
	GetBody(modelPath, id string) (string, error)
	Edit(modelPath string, decision *Decision, context *string, options *[]string, drivers *string) error
	Link(modelPath string, source, target *Decision, forwardTag, reverseTag string) error
	Tag(modelPath string, decision *Decision, tag string) error
	FilterDecisions(decisions []Decision, filters map[string][]string) ([]Decision, error)
	Decide(modelPath string, decision *Decision, option, rationale string, enforceOption bool) error
	Revise(modelPath string, original *Decision) (*Decision, error)
	Copy(sourceModelPath, targetPath, decisionID string) error
	Comment(modelPath string, decision *Decision, author, text string) error
}

type DecisionServiceImplementation struct {
	repo DecisionRepository
}

func NewDecisionService(repo DecisionRepository) DecisionService {
	return &DecisionServiceImplementation{repo: repo}
}

func (s *DecisionServiceImplementation) AddNew(modelPath, title string) (*Decision, error) {
	if !containsLetter(title) {
		return nil, errors.New("title must contain at least one letter")
	}
	today := time.Now().Format("2006-01-02")
	d := &Decision{
		Title:  title,
		Status: "proposed",
		Date:   today,
	}
	return s.repo.Create(modelPath, "", d)
}

// AddExisting writes a Decision and body into a target model directory, shifting
// any ID references in Links/Supersedes by `increment` so the imported decision
// fits the target's numbering. The repository assigns a fresh ID on Create; the
// caller passes that fresh ID back via decision.ID before Save.
func (s *DecisionServiceImplementation) AddExisting(sourceModelPath, targetModelPath string, d *Decision, body string, increment int) (*Decision, error) {
	d.Supersedes = adjustIDsBy(d.Supersedes, increment)
	for tag, ids := range d.Links {
		d.Links[tag] = adjustIDsBy(ids, increment)
	}
	created, err := s.repo.Create(targetModelPath, "", d)
	if err != nil {
		return nil, err
	}
	if body != "" {
		if err := s.repo.Save(targetModelPath, created, body); err != nil {
			return nil, err
		}
	}
	return created, nil
}

func (s *DecisionServiceImplementation) GetAllDecisions(modelPath string) ([]Decision, error) {
	return s.repo.LoadAll(modelPath)
}

func (s *DecisionServiceImplementation) GetDecisionByID(modelPath, id string) (*Decision, error) {
	return s.repo.LoadById(modelPath, id)
}

func (s *DecisionServiceImplementation) GetDecisionByTitle(modelPath, title string) (*Decision, error) {
	return s.repo.LoadByTitle(modelPath, title)
}

func (s *DecisionServiceImplementation) GetBody(modelPath, id string) (string, error) {
	return s.repo.LoadBody(modelPath, id)
}

// Edit appends to the named sections. context/options/drivers are pointers so a
// nil value means "leave untouched". options is a list of new bullet items
// appended to the Considered Options section. Replace-mode editing is PR 3.
func (s *DecisionServiceImplementation) Edit(modelPath string, d *Decision, context *string, options *[]string, drivers *string) error {
	body, err := s.repo.LoadBody(modelPath, d.ID)
	if err != nil {
		return err
	}

	if context != nil {
		body = appendToSection(body, "Context and Problem Statement", *context)
	}
	if drivers != nil {
		body = appendBulletsToSection(body, "Decision Drivers", []string{*drivers})
	}
	if options != nil {
		body = appendBulletsToSection(body, "Considered Options", *options)
	}

	d.Date = time.Now().Format("2006-01-02")
	return s.repo.Save(modelPath, d, body)
}

func (s *DecisionServiceImplementation) Link(modelPath string, source, target *Decision, tag, reverseTag string) error {
	if tag == "supersedes" || tag == "superseded-by" {
		return fmt.Errorf("use 'adg supersede' to record supersession, not 'adg link'")
	}
	if source.Links == nil {
		source.Links = map[string][]string{}
	}
	if !slices.Contains(source.Links[tag], target.ID) {
		source.Links[tag] = append(source.Links[tag], target.ID)
	}
	source.Date = time.Now().Format("2006-01-02")

	bodySrc, err := s.repo.LoadBody(modelPath, source.ID)
	if err != nil {
		return err
	}
	if err := s.repo.Save(modelPath, source, bodySrc); err != nil {
		return fmt.Errorf("failed to save source decision: %w", err)
	}

	if reverseTag != "" {
		if target.Links == nil {
			target.Links = map[string][]string{}
		}
		if !slices.Contains(target.Links[reverseTag], source.ID) {
			target.Links[reverseTag] = append(target.Links[reverseTag], source.ID)
		}
		target.Date = time.Now().Format("2006-01-02")
		bodyTgt, err := s.repo.LoadBody(modelPath, target.ID)
		if err != nil {
			return err
		}
		if err := s.repo.Save(modelPath, target, bodyTgt); err != nil {
			return fmt.Errorf("failed to save target decision: %w", err)
		}
	}
	return nil
}

func (s *DecisionServiceImplementation) Tag(modelPath string, d *Decision, tag string) error {
	if slices.Contains(d.Tags, tag) {
		return fmt.Errorf("tag %q already exists in this decision", tag)
	}
	d.Tags = append(d.Tags, tag)
	d.Date = time.Now().Format("2006-01-02")
	body, err := s.repo.LoadBody(modelPath, d.ID)
	if err != nil {
		return err
	}
	return s.repo.Save(modelPath, d, body)
}

func (s *DecisionServiceImplementation) FilterDecisions(decisions []Decision, filters map[string][]string) ([]Decision, error) {
	var results []Decision

	idSet := make(map[string]bool)
	if idFilters, ok := filters["id"]; ok {
		expandedIDs, err := expandIDFilters(idFilters)
		if err != nil {
			return nil, err
		}
		for _, id := range expandedIDs {
			idSet[id] = true
		}
	}

	var titleRegex *regexp.Regexp
	if titles, ok := filters["title"]; ok && len(titles) > 0 {
		var err error
		titleRegex, err = regexp.Compile(titles[0])
		if err != nil {
			return nil, fmt.Errorf("invalid title regex: %w", err)
		}
	}

	for _, d := range decisions {
		if matchesID(d, idSet) || matchesTitle(d, titleRegex) || matchesTag(d, filters["tag"]) || matchesStatus(d, filters["status"]) {
			results = append(results, d)
		}
	}
	return results, nil
}

// Decide writes the MADR `Chosen option: "X", because Y.` line into Decision
// Outcome and flips status to accepted. enforceOption=true preserves the legacy
// behavior of auto-creating an option that wasn't in Considered Options; the
// MADR-purist rejection of this behavior lands in PR 2.
func (s *DecisionServiceImplementation) Decide(modelPath string, d *Decision, option, rationale string, enforceOption bool) error {
	body, err := s.repo.LoadBody(modelPath, d.ID)
	if err != nil {
		return err
	}
	parsed, err := madr.ParseBody(body)
	if err != nil {
		return err
	}

	chosen, err := resolveOption(parsed.Options, option)
	if err != nil {
		if !enforceOption {
			return fmt.Errorf("%w (use --force to add the option automatically)", err)
		}
		if isNumeric(option) {
			return fmt.Errorf("cannot auto-create numeric option: %q; use a descriptive name when using --force", option)
		}
		body = appendBulletsToSection(body, "Considered Options", []string{option})
		chosen = option
	}

	outcome := fmt.Sprintf("## Decision Outcome\n\nChosen option: %q", chosen)
	if rationale != "" {
		outcome += fmt.Sprintf(", because %s", rationale)
	}
	outcome += ".\n"
	body = replaceSection(body, "Decision Outcome", outcome)

	d.Status = "accepted"
	d.Date = time.Now().Format("2006-01-02")
	d.LegacyOutcome = false
	return s.repo.Save(modelPath, d, body)
}

func (s *DecisionServiceImplementation) Revise(modelPath string, original *Decision) (*Decision, error) {
	body, err := s.repo.LoadBody(modelPath, original.ID)
	if err != nil {
		return nil, err
	}
	// Clear Decision Outcome.
	body = replaceSection(body, "Decision Outcome", "## Decision Outcome\n\n{...}\n")

	today := time.Now().Format("2006-01-02")
	revised := &Decision{
		Title:          original.Title + " (Revised)",
		Status:         "proposed",
		Date:           today,
		DecisionMakers: original.DecisionMakers,
		Consulted:      original.Consulted,
		Informed:       original.Informed,
		Tags:           original.Tags,
		// Links/Supersedes/Comments deliberately not copied; the revised
		// decision starts fresh.
	}
	created, err := s.repo.Create(modelPath, "", revised)
	if err != nil {
		return nil, err
	}
	body = replaceH1Title(body, created.Title)
	if err := s.repo.Save(modelPath, created, body); err != nil {
		return nil, err
	}
	return created, nil
}

func (s *DecisionServiceImplementation) Copy(modelPath, targetPath, decisionID string) error {
	return s.repo.Copy(modelPath, targetPath, decisionID)
}

// Comment appends an entry to Decision.Comments (with the actual text in
// Comment.Text — the §A.1 architectural fix) and re-saves so the repository's
// renderer regenerates the ## Comments body section.
func (s *DecisionServiceImplementation) Comment(modelPath string, d *Decision, author, text string) error {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	d.Comments = append(d.Comments, Comment{Author: author, Date: timestamp, Text: text})
	d.Date = time.Now().Format("2006-01-02")

	body, err := s.repo.LoadBody(modelPath, d.ID)
	if err != nil {
		return err
	}
	return s.repo.Save(modelPath, d, body)
}

// Helpers

func containsLetter(s string) bool {
	matched, err := regexp.MatchString(`[a-zA-Z]`, s)
	return err == nil && matched
}

func isNumeric(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func adjustIDsBy(ids []string, delta int) []string {
	if delta == 0 || len(ids) == 0 {
		return ids
	}
	var updated []string
	for _, id := range ids {
		if num, err := strconv.Atoi(id); err == nil {
			updated = append(updated, fmt.Sprintf("%04d", num+delta))
		} else {
			updated = append(updated, id)
		}
	}
	return updated
}

func resolveOption(options []string, option string) (string, error) {
	if n, err := strconv.Atoi(option); err == nil {
		if n < 1 || n > len(options) {
			return "", fmt.Errorf("option %d out of range (1..%d)", n, len(options))
		}
		return options[n-1], nil
	}
	target := strings.ToLower(strings.TrimSpace(option))
	for _, o := range options {
		if strings.ToLower(o) == target {
			return o, nil
		}
	}
	return "", fmt.Errorf("option %q is not in Considered Options", option)
}

// appendToSection appends content as a new paragraph inside an existing H2
// section. If the section doesn't exist, it's created at the end of the body.
func appendToSection(body, headerText, content string) string {
	header := "## " + headerText
	lines := strings.Split(body, "\n")
	startIdx := -1
	for i, line := range lines {
		if strings.EqualFold(strings.TrimSpace(line), header) {
			startIdx = i
			break
		}
	}
	if startIdx == -1 {
		// Section absent — append at end.
		return strings.TrimRight(body, "\n") + "\n\n" + header + "\n\n" + content + "\n"
	}
	// Find next H2 (or EOF) to insert before.
	insertAt := len(lines)
	for i := startIdx + 1; i < len(lines); i++ {
		if strings.HasPrefix(strings.TrimSpace(lines[i]), "## ") {
			insertAt = i
			break
		}
	}
	newLines := append([]string{}, lines[:insertAt]...)
	newLines = append(newLines, content)
	newLines = append(newLines, lines[insertAt:]...)
	return strings.Join(newLines, "\n")
}

func appendBulletsToSection(body, headerText string, items []string) string {
	var b strings.Builder
	for _, item := range items {
		b.WriteString("* ")
		b.WriteString(item)
		b.WriteString("\n")
	}
	return appendToSection(body, headerText, b.String())
}

// replaceSection replaces an H2 section's contents (header line included) with
// newSection. If no matching section exists, newSection is appended at the end.
func replaceSection(body, headerText, newSection string) string {
	header := "## " + headerText
	lines := strings.Split(body, "\n")
	startIdx := -1
	for i, line := range lines {
		if strings.EqualFold(strings.TrimSpace(line), header) {
			startIdx = i
			break
		}
	}
	if startIdx == -1 {
		return strings.TrimRight(body, "\n") + "\n\n" + strings.TrimRight(newSection, "\n") + "\n"
	}
	endIdx := len(lines)
	for i := startIdx + 1; i < len(lines); i++ {
		if strings.HasPrefix(strings.TrimSpace(lines[i]), "## ") {
			endIdx = i
			break
		}
	}
	var out []string
	out = append(out, lines[:startIdx]...)
	out = append(out, strings.TrimRight(newSection, "\n"))
	out = append(out, "")
	out = append(out, lines[endIdx:]...)
	return strings.Join(out, "\n")
}

func replaceH1Title(body, newTitle string) string {
	return regexp.MustCompile(`(?m)^# +.+$`).ReplaceAllString(body, "# "+newTitle)
}

func matchesID(d Decision, idSet map[string]bool) bool {
	return len(idSet) > 0 && idSet[d.ID]
}

func matchesTitle(d Decision, titleRegex *regexp.Regexp) bool {
	return titleRegex != nil && titleRegex.MatchString(d.Title)
}

func matchesTag(d Decision, tags []string) bool {
	for _, ft := range tags {
		if slices.Contains(d.Tags, ft) {
			return true
		}
	}
	return false
}

func matchesStatus(d Decision, statuses []string) bool {
	for _, status := range statuses {
		if d.Status == status {
			return true
		}
	}
	return false
}

func expandIDFilters(ids []string) ([]string, error) {
	var result []string
	for _, raw := range ids {
		for _, id := range strings.Split(raw, ",") {
			id = strings.TrimSpace(id)
			if strings.Contains(id, "-") {
				expanded, err := expandRange(id)
				if err != nil {
					return nil, err
				}
				result = append(result, expanded...)
			} else {
				result = append(result, id)
			}
		}
	}
	return result, nil
}

func expandRange(rng string) ([]string, error) {
	parts := strings.Split(rng, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid ID range: %s", rng)
	}
	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil || start > end {
		return nil, fmt.Errorf("invalid ID range: %s", rng)
	}
	var out []string
	for i := start; i <= end; i++ {
		out = append(out, fmt.Sprintf("%04d", i))
	}
	return out, nil
}
