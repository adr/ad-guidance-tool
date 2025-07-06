package decision

import (
	"adg/internal/domain"
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

type DecisionService interface {
	AddNew(modelPath, title string) (*Decision, error)
	AddExisting(sourceModelPath, targetModelPath string, decision *Decision, content *DecisionContent, increment int) (*Decision, error)
	GetAllDecisions(modelPath string) ([]Decision, error)
	GetDecisionByID(modelPath, id string) (*Decision, error)
	GetDecisionByTitle(modelPath, title string) (*Decision, error)
	GetDecisionContent(modelPath, decisionID string) (*DecisionContent, error)
	Edit(modelPath string, decision *Decision, question *string, options *[]string, criteria *string) error
	Link(modelPath string, source, target *Decision, forwardTag, reverseTag string) error
	Tag(modelPath string, decision *Decision, tag string) error
	FilterDecisions(decisions []Decision, filters map[string][]string) ([]Decision, error)
	Decide(modelPath string, decision *Decision, option, rationale string, enforceOption bool) error
	Revise(modelPath string, original *Decision) (*Decision, error)
	Copy(sourceModelPath, targetPath, decisionID string) error
	Comment(modelPath string, decision *Decision, author, comment string) error
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

	decision := &Decision{
		Title:    title,
		Status:   "open",
		Tags:     []string{},
		Links:    Links{Precedes: []string{}, Succeeds: []string{}},
		Comments: []Comment{},
	}

	content := &DecisionContent{}

	return s.repo.Create(modelPath, "", decision, content)
}

func (s *DecisionServiceImplementation) AddExisting(sourceModelPath, targetModelPath string, decision *Decision, content *DecisionContent, increment int) (*Decision, error) {
	subFolderPath, err := s.getSubFolderPath(sourceModelPath, decision.ID)
	if err != nil {
		return nil, err
	}

	decision.Links.Precedes = adjustIDsBy(decision.Links.Precedes, increment)
	decision.Links.Succeeds = adjustIDsBy(decision.Links.Succeeds, increment)

	for tag, ids := range decision.Links.Custom {
		decision.Links.Custom[tag] = adjustIDsBy(ids, increment)
	}

	return s.repo.Create(targetModelPath, subFolderPath, decision, content)
}

func (s *DecisionServiceImplementation) GetAllDecisions(modelPath string) ([]Decision, error) {
	decisions, err := s.repo.LoadAllByIndex(modelPath)
	if err == nil {
		return decisions, nil
	}

	decisions, fallbackErr := s.repo.LoadAllByData(modelPath)
	if fallbackErr != nil {
		return nil, fallbackErr
	}

	return decisions, nil
}

func (s *DecisionServiceImplementation) GetDecisionByID(modelPath, id string) (*Decision, error) {
	return s.repo.LoadById(modelPath, id)
}

func (s *DecisionServiceImplementation) GetDecisionByTitle(modelPath, title string) (*Decision, error) {
	return s.repo.LoadByTitle(modelPath, title)
}

func (s *DecisionServiceImplementation) GetDecisionContent(modelPath, decisionID string) (*DecisionContent, error) {
	return s.repo.LoadDecisionContent(modelPath, decisionID)
}

func (s *DecisionServiceImplementation) Edit(modelPath string, decision *Decision, question *string, options *[]string, criteria *string) error {
	appendOps := []func() error{}

	if question != nil {
		q := *question
		appendOps = append(appendOps, func() error {
			return s.appendQuestion(modelPath, decision.ID, q)
		})
	}
	if criteria != nil {
		c := *criteria
		appendOps = append(appendOps, func() error {
			return s.appendCriteria(modelPath, decision.ID, c)
		})
	}
	if options != nil {
		opts := *options
		appendOps = append(appendOps, func() error {
			return s.appendOptions(modelPath, decision.ID, opts)
		})
	}

	for _, op := range appendOps {
		if err := op(); err != nil {
			return err
		}
	}

	return nil
}

func (s *DecisionServiceImplementation) Link(
	modelPath string,
	source *Decision,
	target *Decision,
	tag string,
	reverseTag string,
) error {
	if source.Links.Custom == nil {
		source.Links.Custom = make(map[string][]string)
	}
	if target.Links.Custom == nil {
		target.Links.Custom = make(map[string][]string)
	}

	if tag == "precedes" && reverseTag == "succeeds" {
		if s.wouldCreateCycle(modelPath, source.ID, target.ID) {
			return fmt.Errorf("linking %s -> %s would create a cycle", source.ID, target.ID)
		}
		source.Links.Precedes = append(source.Links.Precedes, target.ID)
		target.Links.Succeeds = append(target.Links.Succeeds, source.ID)
	} else {
		addUnique := func(linkMap map[string][]string, key, value string) {
			if !slices.Contains(linkMap[key], value) {
				linkMap[key] = append(linkMap[key], value)
			}
		}
		addUnique(source.Links.Custom, tag, target.ID)
		if reverseTag != "" {
			addUnique(target.Links.Custom, reverseTag, source.ID)
		}
	}

	if err := s.repo.Save(modelPath, source); err != nil {
		return fmt.Errorf("failed to save source decision: %w", err)
	}
	if err := s.repo.Save(modelPath, target); err != nil {
		return fmt.Errorf("failed to save target decision: %w", err)
	}

	return nil
}

func (s *DecisionServiceImplementation) Tag(modelPath string, decision *Decision, tag string) error {
	if slices.Contains(decision.Tags, tag) {
		return fmt.Errorf("tag %q already exists in this decision", tag)
	}

	decision.Tags = append(decision.Tags, tag)

	if err := s.repo.Save(modelPath, decision); err != nil {
		return fmt.Errorf("failed to save decision with new tag: %w", err)
	}
	return nil
}

func (s *DecisionServiceImplementation) FilterDecisions(decisions []Decision, filters map[string][]string) ([]Decision, error) {
	var results []Decision

	// ID filtering
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

	// Title regex filtering
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

func (s *DecisionServiceImplementation) Decide(modelPath string, decision *Decision, option, rationale string, enforceOption bool) error {
	exists, err := s.repo.OptionExists(modelPath, decision.ID, option)
	if err != nil {
		return err
	}

	if !exists {
		if !enforceOption {
			return fmt.Errorf(
				"option does not exist in the decision (provide either a number or name of an existing option or use -f or --force to automatically create new option for the decision)",
			)
		}
		if isNumeric(option) {
			return fmt.Errorf("cannot auto-create numeric option: %q, use a descriptive name when using --force", option)
		}
		if err := s.appendOptions(modelPath, decision.ID, []string{option}); err != nil {
			return fmt.Errorf("failed to append new option: %w", err)
		}
	}

	optionNum, err := s.repo.ResolveOptionNumber(modelPath, decision.ID, option)
	if err != nil {
		return err
	}

	outcome := formatOutcome(optionNum, rationale)

	// TODO: use generic UpdateSection function
	if err := s.repo.AppendOutcomeSection(modelPath, decision.ID, outcome); err != nil {
		return err
	}

	decision.Status = "decided"
	return s.repo.Save(modelPath, decision)
}

func (s *DecisionServiceImplementation) Revise(modelPath string, original *Decision) (*Decision, error) {
	revised := s.buildRevisedDecision(original)

	content, err := s.repo.LoadDecisionContent(modelPath, original.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to load original content: %w", err)
	}

	s.resetContentForRevision(content)

	subFolderPath, err := s.getSubFolderPath(modelPath, original.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get relative path: %w", err)
	}

	return s.repo.Create(modelPath, subFolderPath, revised, content)
}

func (s *DecisionServiceImplementation) Copy(modelPath, targetPath, decisionId string) error {
	return s.repo.Copy(modelPath, targetPath, decisionId)
}

func (s *DecisionServiceImplementation) Comment(modelPath string, decision *Decision, author, commentText string) error {
	commentCount := len(decision.Comments)
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	newComment := Comment{
		Author:  author,
		Date:    timestamp,
		Comment: fmt.Sprintf("%d", commentCount+1),
	}
	decision.Comments = append(decision.Comments, newComment)

	if err := s.repo.Save(modelPath, decision); err != nil {
		return fmt.Errorf("failed to save updated decision: %w", err)
	}

	if err := s.repo.AppendCommentSection(modelPath, decision.ID, commentText, commentCount+1, author, timestamp); err != nil {
		return fmt.Errorf("failed to append comment section: %w", err)
	}

	return nil
}

// Helpers

func (s *DecisionServiceImplementation) getSubFolderPath(modelPath, decisionID string) (string, error) {
	filePath, err := s.repo.FindDecisionFile(modelPath, decisionID)
	if err != nil {
		return "", err
	}

	relPath, err := filepath.Rel(modelPath, filePath)
	if err != nil {
		return "", fmt.Errorf("failed to get relative path: %w", err)
	}

	return filepath.Dir(relPath), nil
}

func (s *DecisionServiceImplementation) appendToSection(modelPath, decisionID, section string, existingContent, newContent string) error {
	existing := strings.Split(existingContent, "\n")
	newLines := strings.Split(newContent, "\n")
	lines := append(existing, newLines...)
	return s.repo.UpdateSection(modelPath, decisionID, section, lines)
}

func (s *DecisionServiceImplementation) appendQuestion(modelPath, decisionID string, question string) error {
	content, err := s.GetDecisionContent(modelPath, decisionID)
	if err != nil {
		return err
	}
	return s.appendToSection(modelPath, decisionID, domain.AnchorSectionQuestion, content.Question, question)
}

func (s *DecisionServiceImplementation) appendCriteria(modelPath, decisionID string, criteria string) error {
	content, err := s.GetDecisionContent(modelPath, decisionID)
	if err != nil {
		return err
	}
	return s.appendToSection(modelPath, decisionID, domain.AnchorSectionCriteria, content.Criteria, criteria)
}

func (s *DecisionServiceImplementation) appendOptions(modelPath, decisionID string, newOptions []string) error {
	content, err := s.GetDecisionContent(modelPath, decisionID)
	if err != nil {
		return err
	}

	lines := strings.Split(content.Options, "\n")
	optionCount := countExistingOptions(lines)

	for i, opt := range newOptions {
		if err := s.validateOptionDoesNotExist(modelPath, decisionID, opt); err != nil {
			return err
		}

		num := optionCount + i + 1
		lines = append(lines, formatOptionLine(num, opt))
	}

	return s.repo.UpdateSection(modelPath, decisionID, domain.AnchorSectionOptions, lines)
}

func countExistingOptions(lines []string) int {
	count := 0
	for _, line := range lines {
		if strings.Contains(line, `name="option-`) {
			count++
		}
	}
	return count
}

func (s *DecisionServiceImplementation) validateOptionDoesNotExist(modelPath, decisionID, option string) error {
	exists, err := s.repo.OptionExists(modelPath, decisionID, option)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("cannot create option %s because it already exists", option)
	}
	return nil
}

func formatOptionLine(num int, option string) string {
	return fmt.Sprintf("%d. %s %s", num, domain.AnchorForOption(num), option)
}

func (s *DecisionServiceImplementation) wouldCreateCycle(modelPath, fromID, toID string) bool {
	visited := make(map[string]bool)
	return s.dfs(modelPath, toID, fromID, visited)
}

func (s *DecisionServiceImplementation) dfs(modelPath, currentID, targetID string, visited map[string]bool) bool {
	if currentID == targetID {
		return true
	}
	if visited[currentID] {
		return false
	}
	visited[currentID] = true

	node, err := s.repo.LoadById(modelPath, currentID)
	if err != nil {
		return false
	}

	for _, next := range node.Links.Precedes {
		if s.dfs(modelPath, next, targetID, visited) {
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

func containsLetter(s string) bool {
	matched, err := regexp.MatchString(`[a-zA-Z]`, s)
	return err == nil && matched
}

func adjustIDsBy(ids []string, delta int) []string {
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

func matchesID(d Decision, idSet map[string]bool) bool {
	return len(idSet) > 0 && idSet[d.ID]
}

func matchesTitle(d Decision, titleRegex *regexp.Regexp) bool {
	return titleRegex != nil && titleRegex.MatchString(d.Title)
}

func matchesTag(d Decision, tags []string) bool {
	for _, filterTag := range tags {
		for _, decisionTag := range d.Tags {
			if filterTag == decisionTag {
				return true
			}
		}
	}
	return false
}

func matchesStatus(d Decision, statuses []string) bool {
	for _, s := range statuses {
		if d.Status == s {
			return true
		}
	}
	return false
}

func isNumeric(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func formatOutcome(optionNum int, rationale string) string {
	optionAnchor := domain.AnchorLinkToOption(optionNum) // todo: use name for the displayed option text
	if rationale != "" {
		return fmt.Sprintf("We decided for %s because: %s", optionAnchor, rationale)
	}
	return fmt.Sprintf("We decided for %s.", optionAnchor)
}

func (s *DecisionServiceImplementation) buildRevisedDecision(original *Decision) *Decision {
	return &Decision{
		Title:  original.Title + " (Revised)",
		Status: "open",
		Tags:   original.Tags,
	}
}

func (s *DecisionServiceImplementation) resetContentForRevision(content *DecisionContent) {
	content.Outcome = ""
	content.Comments = ""
}
