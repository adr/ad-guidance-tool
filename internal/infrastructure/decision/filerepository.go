package decision

import (
	domain "adg/internal/domain/decision"
	"adg/internal/domain/decision/madr"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// FileDecisionRepository persists Decisions as MADR 4.0–format files on disk.
// Files are named NNNN-slug.md; the model directory is scanned on every read
// (no index.yaml). Reads error if any file uses the legacy ADG format so
// callers can steer the user to `adg migrate` in PR 4.
type FileDecisionRepository struct{}

func NewFileDecisionRepository() *FileDecisionRepository {
	return &FileDecisionRepository{}
}

func (r *FileDecisionRepository) Create(modelPath, subFolderPath string, d *domain.Decision) (*domain.Decision, error) {
	id, err := r.generateNextID(modelPath)
	if err != nil {
		return nil, err
	}
	d.ID = id
	d.Slug = slugify(d.Title)

	body := madr.RenderNewBody(d.Title)
	out, err := madr.RenderFile(*d, body)
	if err != nil {
		return nil, err
	}

	filename := fmt.Sprintf("%s-%s.md", d.ID, d.Slug)
	fullPath := filepath.Join(modelPath, subFolderPath, filename)
	if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
		return nil, fmt.Errorf("failed to create dir: %w", err)
	}
	if err := os.WriteFile(fullPath, []byte(out), 0o644); err != nil {
		return nil, fmt.Errorf("failed to write file: %w", err)
	}
	return d, nil
}

func (r *FileDecisionRepository) Save(modelPath string, d *domain.Decision, body string) error {
	path, err := r.FindDecisionFile(modelPath, d.ID)
	if err != nil {
		return err
	}
	out, err := madr.RenderFile(*d, body)
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(out), 0o644)
}

func (r *FileDecisionRepository) LoadById(modelPath, id string) (*domain.Decision, error) {
	path, err := r.FindDecisionFile(modelPath, id)
	if err != nil {
		return nil, err
	}
	return r.loadFile(path)
}

func (r *FileDecisionRepository) LoadByTitle(modelPath, title string) (*domain.Decision, error) {
	all, err := r.LoadAll(modelPath)
	if err != nil {
		return nil, err
	}
	slug := slugify(title)
	var exact *domain.Decision
	var partial []*domain.Decision
	for i := range all {
		d := &all[i]
		switch {
		case d.Slug == slug:
			if exact != nil {
				return nil, fmt.Errorf("multiple decisions match title exactly; use id")
			}
			exact = d
		case strings.Contains(d.Slug, slug):
			partial = append(partial, d)
		}
	}
	switch {
	case exact != nil:
		return exact, nil
	case len(partial) == 1:
		return partial[0], nil
	case len(partial) > 1:
		return nil, fmt.Errorf("multiple titles matched %q; be more specific or use id", title)
	default:
		return nil, fmt.Errorf("no decision title matched %q", title)
	}
}

func (r *FileDecisionRepository) LoadAll(modelPath string) ([]domain.Decision, error) {
	var decisions []domain.Decision
	err := filepath.WalkDir(modelPath, func(path string, e fs.DirEntry, err error) error {
		if err != nil || e.IsDir() {
			return err
		}
		if filepath.Ext(e.Name()) != ".md" {
			return nil
		}
		content, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		if madr.IsLegacyADG(path, content) {
			return fmt.Errorf("file %s appears to use legacy ADG format; run 'adg migrate --model %s' to convert", path, modelPath)
		}
		// Skip unrelated markdown (e.g. README.md) whose filename doesn't match NNNN-slug.md
		if _, _, perr := madr.ParseFilename(path); perr != nil {
			return nil
		}
		d, err := r.loadFileContent(path, content)
		if err != nil {
			return err
		}
		decisions = append(decisions, *d)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return decisions, nil
}

func (r *FileDecisionRepository) LoadBody(modelPath, id string) (string, error) {
	path, err := r.FindDecisionFile(modelPath, id)
	if err != nil {
		return "", err
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	_, body, err := madr.SplitFile(raw)
	return body, err
}

func (r *FileDecisionRepository) FindDecisionFile(modelPath, id string) (string, error) {
	var found string
	prefix := id + "-"
	err := filepath.WalkDir(modelPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if strings.HasPrefix(d.Name(), prefix) && strings.HasSuffix(d.Name(), ".md") {
			found = path
			return io.EOF
		}
		return nil
	})
	if err != nil && err != io.EOF {
		return "", err
	}
	if found == "" {
		return "", fmt.Errorf("decision %s not found", id)
	}
	return found, nil
}

func (r *FileDecisionRepository) Copy(srcPath, dstPath, decisionID string) error {
	srcFile, err := r.FindDecisionFile(srcPath, decisionID)
	if err != nil {
		return err
	}
	rel, err := filepath.Rel(srcPath, srcFile)
	if err != nil {
		return fmt.Errorf("failed to compute relative path: %w", err)
	}
	dstFile := filepath.Join(dstPath, rel)
	if err := os.MkdirAll(filepath.Dir(dstFile), 0o755); err != nil {
		return err
	}
	content, err := os.ReadFile(srcFile)
	if err != nil {
		return err
	}
	return os.WriteFile(dstFile, content, 0o644)
}

// Helpers

func (r *FileDecisionRepository) loadFile(path string) (*domain.Decision, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return r.loadFileContent(path, raw)
}

func (r *FileDecisionRepository) loadFileContent(path string, raw []byte) (*domain.Decision, error) {
	if madr.IsLegacyADG(path, raw) {
		return nil, fmt.Errorf("file %s appears to use legacy ADG format; run 'adg migrate' to convert", path)
	}
	fmText, body, err := madr.SplitFile(raw)
	if err != nil {
		return nil, err
	}
	fm, err := madr.ParseFrontmatter(fmText)
	if err != nil {
		return nil, err
	}
	parsed, err := madr.ParseBody(body)
	if err != nil {
		return nil, err
	}
	id, slug, err := madr.ParseFilename(path)
	if err != nil {
		return nil, err
	}
	d := madr.DecisionFromFrontmatter(fm)
	d.ID = id
	d.Slug = slug
	d.Title = parsed.Title
	return &d, nil
}

func (r *FileDecisionRepository) generateNextID(modelPath string) (string, error) {
	maxID := 0
	err := filepath.WalkDir(modelPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		name := d.Name()
		if len(name) < 5 || !strings.HasSuffix(name, ".md") {
			return nil
		}
		if id, err := strconv.Atoi(name[:4]); err == nil && name[4] == '-' {
			if id > maxID {
				maxID = id
			}
		}
		return nil
	})
	if err != nil && !os.IsNotExist(err) {
		return "", fmt.Errorf("scan failed: %w", err)
	}
	return fmt.Sprintf("%04d", maxID+1), nil
}

func slugify(title string) string {
	s := strings.ToLower(title)
	s = strings.ReplaceAll(s, " ", "-")
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			b.WriteRune(r)
		}
	}
	return b.String()
}
