package decision

import (
	util "adg/internal/domain"
	config "adg/internal/domain/config"
	domain "adg/internal/domain/decision"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type FileDecisionRepository struct {
	config config.ConfigService
}

func NewFileDecisionRepository(config config.ConfigService) *FileDecisionRepository {
	return &FileDecisionRepository{config: config}
}

func (r *FileDecisionRepository) Create(modelPath, subFolderPath string, decision *domain.Decision, content *domain.DecisionContent) (*domain.Decision, error) {
	newID, err := r.generateNextID(modelPath)
	if err != nil {
		return nil, err
	}
	decision.ID = newID

	slug := strings.ReplaceAll(strings.ToLower(decision.Title), " ", "-")
	filename := fmt.Sprintf("AD%s-%s.md", decision.ID, slug)
	fullPath := filepath.Join(modelPath, filepath.Join(subFolderPath, filename))

	markdown, err := r.composeDecisionFileContent(decision, content)
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory for %s: %w", fullPath, err)
	}
	if err := os.WriteFile(fullPath, markdown, 0644); err != nil {
		return nil, fmt.Errorf("failed to write decision file: %w", err)
	}
	if err := r.updateIndex(modelPath, decision); err != nil {
		return nil, fmt.Errorf("failed to update index: %w", err)
	}
	return decision, nil
}

func (r *FileDecisionRepository) Save(modelPath string, decision *domain.Decision) error {
	filePath, err := r.FindDecisionFile(modelPath, decision.ID)
	if err != nil {
		return err
	}

	existingMeta, body, err := getFileParts(filePath)
	if err != nil {
		return err
	}

	mergedMeta, err := r.mergeMetadata(existingMeta, decision)
	if err != nil {
		return err
	}

	finalContent := constructMarkdownWithMetaAndBody(mergedMeta, body)

	if err := os.WriteFile(filePath, finalContent, 0644); err != nil {
		return fmt.Errorf("failed to write updated decision: %w", err)
	}

	return r.updateIndex(modelPath, decision)
}

func (r *FileDecisionRepository) Copy(srcPath, dstPath, decisionID string) error {
	// ensure destination path exists
	if err := os.MkdirAll(dstPath, 0755); err != nil {
		return fmt.Errorf("failed to create destination model directory: %w", err)
	}

	// find the source file
	srcFilePath, err := r.FindDecisionFile(srcPath, decisionID)
	if err != nil {
		return err
	}

	// preserve subfolder structure
	relPath, err := filepath.Rel(srcPath, srcFilePath)
	if err != nil {
		return fmt.Errorf("failed to compute relative path: %w", err)
	}
	dstFilePath := filepath.Join(dstPath, relPath)

	// ensure target directories exist
	if err := os.MkdirAll(filepath.Dir(dstFilePath), 0755); err != nil {
		return fmt.Errorf("failed to create destination directories: %w", err)
	}

	// perform file copy
	return copyFileContents(srcFilePath, dstFilePath)
}

func (r *FileDecisionRepository) LoadById(modelPath, id string) (*domain.Decision, error) {
	decisions, err := r.loadDecisionsWithFallback(modelPath)
	if err != nil {
		return nil, err
	}

	for _, d := range decisions {
		if d.ID == id {
			return &d, nil
		}
	}
	return nil, fmt.Errorf("no decision with id %q found", id)
}

func (r *FileDecisionRepository) LoadByTitle(modelPath, title string) (*domain.Decision, error) {
	decisions, err := r.loadDecisionsWithFallback(modelPath)
	if err != nil {
		return nil, err
	}

	titleSlug := slugify(title)
	var exactMatch *domain.Decision
	var partialMatches []*domain.Decision

	for _, d := range decisions {
		slug := slugify(d.Title)
		switch {
		case slug == titleSlug:
			if exactMatch != nil {
				return nil, fmt.Errorf("multiple decisions with exact same name, use id to be more specific")
			}
			exactMatch = &d
		case strings.Contains(slug, titleSlug):
			partialMatches = append(partialMatches, &d)
		}
	}

	switch {
	case exactMatch != nil:
		return exactMatch, nil
	case len(partialMatches) == 1:
		return partialMatches[0], nil
	case len(partialMatches) > 1:
		return nil, fmt.Errorf("multiple decision titles matched %q, be more specific or use id instead", title)
	default:
		return nil, fmt.Errorf("no decision title matched %q", title)
	}
}

func (r *FileDecisionRepository) LoadAllByIndex(modelPath string) ([]domain.Decision, error) {
	indexPath := filepath.Join(modelPath, "index.yaml")
	content, err := os.ReadFile(indexPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read index: %w", err)
	}

	var index struct {
		Decisions map[string]domain.Decision `yaml:"decisions"`
	}
	if err := yaml.Unmarshal(content, &index); err != nil {
		return nil, fmt.Errorf("invalid index format: %w", err)
	}

	decisions := make([]domain.Decision, 0, len(index.Decisions))
	for id, decision := range index.Decisions {
		decision.ID = id
		decisions = append(decisions, decision)
	}
	return decisions, nil
}

func (r *FileDecisionRepository) LoadAllByData(modelPath string) ([]domain.Decision, error) {
	var decisions []domain.Decision

	err := filepath.WalkDir(modelPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing %s: %w", path, err)
		}
		if d.IsDir() || filepath.Ext(d.Name()) != ".md" || !isValidDecisionFilename(d.Name()) {
			return nil
		}

		meta, err := extractMetadataFromFile(path)
		if err != nil {
			return fmt.Errorf("failed to parse %s: %w", path, err)
		}

		decisions = append(decisions, *meta)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(decisions, func(i, j int) bool {
		return decisions[i].ID < decisions[j].ID
	})
	return decisions, nil
}

func (r *FileDecisionRepository) LoadDecisionContentRaw(modelPath, decisionID string) (string, error) {
	filePath, err := r.FindDecisionFile(modelPath, decisionID)
	if err != nil {
		return "", err
	}
	_, body, err := getFileParts(filePath)
	return body, err
}

// todo: reuse in other methods to load the sections, change it for example, and then add a new method for updating with new DecisionContent
func (r *FileDecisionRepository) LoadDecisionContent(modelPath, decisionID string) (*domain.DecisionContent, error) {
	body, err := r.LoadDecisionContentRaw(modelPath, decisionID)
	if err != nil {
		return nil, err
	}

	sections := extractSections(body)

	return &domain.DecisionContent{
		ID:       decisionID,
		Question: stripHeader(sections["question"]),
		Options:  stripHeader(sections["options"]),
		Criteria: stripHeader(sections["criteria"]),
		Outcome:  stripHeader(sections["outcome"]),
		Comments: stripHeader(sections["comments"]),
	}, nil
}

func (r *FileDecisionRepository) UpdateSection(modelPath, decisionID, anchorName string, lines []string) error {
	filePath, err := r.FindDecisionFile(modelPath, decisionID)
	if err != nil {
		return err
	}

	metadata, body, err := getFileParts(filePath)
	if err != nil {
		return err
	}

	linesIn := strings.Split(body, "\n")
	var updated []string
	var skipping bool
	var foundAnchor bool

	for i := 0; i < len(linesIn); i++ {
		line := linesIn[i]

		if strings.Contains(line, fmt.Sprintf(`name="%s"`, anchorName)) {
			foundAnchor = true
			skipping = true

			// write new section
			updated = append(updated, line)
			updated = append(updated, lines...)
			updated = append(updated, "")

			// skip old section lines
			for j := i + 1; j < len(linesIn); j++ {
				if strings.HasPrefix(linesIn[j], "## ") {
					i = j - 1
					skipping = false
					break
				}
			}
			continue
		}

		if !skipping {
			updated = append(updated, line)
		}
	}

	if !foundAnchor {
		header := r.resolveHeader(anchorName)
		newSection := []string{
			"## " + util.AnchorForSection(anchorName) + " " + header,
		}
		newSection = append(newSection, lines...)
		newSection = append(newSection, "")

		insertAt := findSectionInsertIndex(updated, anchorName)
		before := updated[:insertAt]
		after := updated[insertAt:]
		updated = append(before, append(newSection, after...)...)
	}

	return writeFinalContent(filePath, []byte(metadata), updated)
}

// TODO: call UpdateSection directly and completely remove these two specific section functions
func (r *FileDecisionRepository) AppendOutcomeSection(modelPath, decisionID, outcome string) error {
	lines := strings.Split(outcome, "\n")
	return r.UpdateSection(modelPath, decisionID, util.AnchorSectionOutcome, lines)
}

func (r *FileDecisionRepository) AppendCommentSection(modelPath, decisionID, commentText string, commentNumber int, author, date string) error {
	commentLine := util.AnchorForComment(commentNumber, author, date, commentText)
	return r.UpdateSection(modelPath, decisionID, util.AnchorSectionComments, []string{commentLine})
}

func (r *FileDecisionRepository) OptionExists(modelPath, decisionID, option string) (bool, error) {
	filePath, err := r.FindDecisionFile(modelPath, decisionID)
	if err != nil {
		return false, err
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	return findOptionInContent(string(content), option)
}

func (r *FileDecisionRepository) ResolveOptionNumber(modelPath, decisionID, option string) (int, error) {
	filePath, err := r.FindDecisionFile(modelPath, decisionID)
	if err != nil {
		return 0, fmt.Errorf("failed to locate decision file: %w", err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to read decision file: %w", err)
	}

	return extractOptionNumberFromContent(string(content), option)
}

func (r *FileDecisionRepository) FindDecisionFile(modelPath, decisionID string) (string, error) {
	var foundPath string
	searchPrefix := "AD" + decisionID

	err := filepath.WalkDir(modelPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasPrefix(d.Name(), searchPrefix) && strings.HasSuffix(d.Name(), ".md") {
			foundPath = path
			return io.EOF
		}
		return nil
	})

	if err != nil && err != io.EOF {
		return "", err
	}
	if foundPath == "" {
		return "", fmt.Errorf("decision %s not found", decisionID)
	}
	return foundPath, nil
}

// Helpers
func (r *FileDecisionRepository) updateIndex(modelPath string, decision *domain.Decision) error {
	indexPath := filepath.Join(modelPath, "index.yaml")

	var data struct {
		Decisions map[string]domain.Decision `yaml:"decisions"`
	}

	content, err := os.ReadFile(indexPath)
	if err == nil {
		_ = yaml.Unmarshal(content, &data)
	}

	if data.Decisions == nil {
		data.Decisions = make(map[string]domain.Decision)
	}
	data.Decisions[decision.ID] = *decision

	out, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal index: %w", err)
	}
	if err := os.WriteFile(indexPath, out, 0644); err != nil {
		return fmt.Errorf("failed to write index: %w", err)
	}
	return nil
}

func getFileParts(filePath string) (metadata string, body string, err error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", "", err
	}

	content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))
	parts := bytes.SplitN(content, []byte("---\n"), 3)

	if len(parts) < 2 {
		return "", "", fmt.Errorf("invalid ADR file format: missing frontmatter")
	}

	meta := bytes.TrimSuffix(parts[1], []byte("---"))
	meta = bytes.TrimSuffix(meta, []byte("---\n"))

	bodyStr := ""
	if len(parts) == 3 {
		bodyStr = string(parts[2])
	}

	return string(meta), bodyStr, nil
}

func (r *FileDecisionRepository) generateNextID(modelPath string) (string, error) {
	maxID := 0
	err := filepath.WalkDir(modelPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if strings.HasPrefix(d.Name(), "AD") && strings.HasSuffix(d.Name(), ".md") && len(d.Name()) >= 6 {
			if id, err := strconv.Atoi(d.Name()[2:6]); err == nil && id > maxID {
				maxID = id
			}
		}
		return nil
	})
	if err != nil && !os.IsNotExist(err) {
		return "", fmt.Errorf("error scanning model dir: %w", err)
	}
	return fmt.Sprintf("%04d", maxID+1), nil
}

func (r *FileDecisionRepository) composeDecisionFileContent(d *domain.Decision, content *domain.DecisionContent) ([]byte, error) {
	meta, err := yaml.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal decision: %w", err)
	}
	var b bytes.Buffer
	b.WriteString("---\n")
	b.Write(meta)
	b.WriteString("---\n\n")

	if content != nil {
		sections := []struct {
			Anchor      string
			Header      string
			Body        string
			SkipIfEmpty bool
		}{
			{"question", r.config.GetQuestionHeader(), content.Question, false},
			{"options", r.config.GetOptionsHeader(), content.Options, false},
			{"criteria", r.config.GetCriteriaHeader(), content.Criteria, false},
			{"outcome", r.config.GetOutcomeHeader(), content.Outcome, true},
			{"comments", r.config.GetCommentsHeader(), content.Comments, true},
		}

		for _, sec := range sections {
			if sec.SkipIfEmpty && strings.TrimSpace(sec.Body) == "" {
				continue
			}
			b.WriteString(fmt.Sprintf("## %s %s\n", util.AnchorForSection(sec.Anchor), sec.Header))
			b.WriteString(sec.Body + "\n\n")
		}
	}
	return b.Bytes(), nil
}

func constructMarkdownWithMetaAndBody(meta []byte, body string) []byte {
	var buf bytes.Buffer
	buf.WriteString("---\n")
	buf.Write(meta)
	buf.WriteString("---\n")
	if !strings.HasPrefix(body, "\n") {
		buf.WriteString("\n")
	}
	buf.WriteString(body)
	return buf.Bytes()
}

func (r *FileDecisionRepository) mergeMetadata(existingMeta string, decision *domain.Decision) ([]byte, error) {
	existingMap, err := unmarshalToMap([]byte(existingMeta))
	if err != nil {
		return nil, fmt.Errorf("failed to parse existing metadata: %w", err)
	}

	newMap, err := marshalDecisionToMap(decision)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal decision to map: %w", err)
	}

	merged := mergeMaps(existingMap, newMap)

	return yaml.Marshal(merged)
}

func unmarshalToMap(data []byte) (map[string]interface{}, error) {
	var m map[string]interface{}
	if err := yaml.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func marshalDecisionToMap(d *domain.Decision) (map[string]interface{}, error) {
	b, err := yaml.Marshal(d)
	if err != nil {
		return nil, err
	}
	return unmarshalToMap(b)
}

func mergeMaps(original, updated map[string]interface{}) map[string]interface{} {
	for k, v := range updated {
		original[k] = v
	}
	return original
}

func copyFileContents(src, dst string) error {
	content, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("failed to read source file %s: %w", src, err)
	}
	if err := os.WriteFile(dst, content, 0644); err != nil {
		return fmt.Errorf("failed to write destination file %s: %w", dst, err)
	}
	return nil
}

func (r *FileDecisionRepository) loadDecisionsWithFallback(modelPath string) ([]domain.Decision, error) {
	decisions, err := r.LoadAllByIndex(modelPath)
	if err == nil {
		return decisions, nil
	}
	return r.LoadAllByData(modelPath)
}

func slugify(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}

func isValidDecisionFilename(name string) bool {
	match := regexp.MustCompile(`^AD\d{4}-.*\.md$`)
	return match.MatchString(name)
}

func extractMetadataFromFile(path string) (*domain.Decision, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		return nil, nil // invalid, skip
	}

	var decision domain.Decision
	if err := yaml.Unmarshal([]byte(parts[1]), &decision); err != nil {
		return nil, err
	}
	return &decision, nil
}

func extractSections(body string) map[string]string {
	lines := strings.Split(body, "\n")
	sections := make(map[string]string)

	var currentKey string
	var buffer []string

	for _, line := range lines {
		if strings.HasPrefix(line, "## ") {
			if currentKey != "" {
				sections[currentKey] = strings.Join(buffer, "\n")
			}
			currentKey = parseSectionHeader(line)
			buffer = []string{line}
		} else if currentKey != "" {
			buffer = append(buffer, line)
		}
	}

	if currentKey != "" {
		sections[currentKey] = strings.Join(buffer, "\n")
	}

	return sections
}

func stripHeader(section string) string {
	lines := strings.Split(section, "\n")
	if len(lines) <= 1 {
		return ""
	}
	return strings.TrimSpace(strings.Join(lines[1:], "\n"))
}

func parseSectionHeader(header string) string {
	header = strings.ToLower(header)
	switch {
	case strings.Contains(header, "question"):
		return "question"
	case strings.Contains(header, "options"):
		return "options"
	case strings.Contains(header, "criteria"):
		return "criteria"
	case strings.Contains(header, "outcome"):
		return "outcome"
	case strings.Contains(header, "comment"):
		return "comments"
	default:
		return ""
	}
}

func (r *FileDecisionRepository) resolveHeader(anchor string) string {
	switch anchor {
	case util.AnchorSectionQuestion:
		return r.config.GetQuestionHeader()
	case util.AnchorSectionOptions:
		return r.config.GetOptionsHeader()
	case util.AnchorSectionCriteria:
		return r.config.GetCriteriaHeader()
	case util.AnchorSectionOutcome:
		return r.config.GetOutcomeHeader()
	case util.AnchorSectionComments:
		return r.config.GetCommentsHeader()
	default:
		return fmt.Sprintf("<a name=\"%s\"></a>", anchor)
	}
}

func writeFinalContent(filePath string, metadata []byte, lines []string) error {
	var final bytes.Buffer
	final.WriteString("---\n")
	final.Write(metadata)
	final.WriteString("---\n")
	final.WriteString(strings.Join(lines, "\n"))
	return os.WriteFile(filePath, final.Bytes(), 0644)
}

func findOptionInContent(content, option string) (bool, error) {
	// try by number
	if number, err := strconv.Atoi(option); err == nil {
		anchor := fmt.Sprintf(`name="option-%d"`, number)
		return strings.Contains(content, anchor), nil
	}

	// try by label
	lines := strings.Split(content, "\n")
	pattern := regexp.MustCompile(`(?i)^.*<a name="option-(\d+)"></a>\s*(.+)$`)
	target := strings.ToLower(strings.TrimSpace(option))

	for _, line := range lines {
		if m := pattern.FindStringSubmatch(line); len(m) == 3 {
			if strings.ToLower(strings.TrimSpace(m[2])) == target {
				return true, nil
			}
		}
	}
	return false, nil
}

func extractOptionNumberFromContent(content, option string) (int, error) {
	// option is a number
	if number, err := strconv.Atoi(option); err == nil {
		if strings.Contains(content, fmt.Sprintf(`name="option-%d"`, number)) {
			return number, nil
		}
		return 0, fmt.Errorf("option number %d not found", number)
	}

	// option is a label
	lines := strings.Split(content, "\n")
	pattern := regexp.MustCompile(`(?i)^.*<a name="option-(\d+)"></a>\s*(.+)$`)
	target := strings.ToLower(strings.TrimSpace(option))

	for _, line := range lines {
		if m := pattern.FindStringSubmatch(line); len(m) == 3 {
			if strings.ToLower(strings.TrimSpace(m[2])) == target {
				num, _ := strconv.Atoi(m[1])
				return num, nil
			}
		}
	}
	return 0, fmt.Errorf("could not resolve option number for label %q", option)
}

func findSectionInsertIndex(lines []string, newAnchor string) int {
	// section order
	sectionOrder := []string{"question", "options", "criteria", "outcome", "comments"}

	// build map of existing anchors and their line numbers
	anchorLines := map[string]int{}
	for i, line := range lines {
		for _, anchor := range sectionOrder {
			if strings.Contains(line, util.AnchorForSection(anchor)) {
				anchorLines[anchor] = i
			}
		}
	}

	// find where in the canonical order the new anchor belongs
	newIndex := -1
	for i, anchor := range sectionOrder {
		if anchor == newAnchor {
			newIndex = i
			break
		}
	}

	// find the next existing section that comes *after* the newAnchor in order
	for i := newIndex + 1; i < len(sectionOrder); i++ {
		if pos, ok := anchorLines[sectionOrder[i]]; ok {
			return pos
		}
	}

	// if no later sections exist, insert at end
	return len(lines)
}
