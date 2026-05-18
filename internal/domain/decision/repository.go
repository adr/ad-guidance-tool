package decision

// DecisionRepository is the persistence boundary for ADR files in MADR 4.0 format.
//
// All path arguments are model paths (directories containing ADR files). Reads
// scan the directory directly — no index.yaml. Implementations must use the
// madr package for parsing and rendering so the round-trip property is preserved.
type DecisionRepository interface {
	// Create writes a new ADR using the canonical MADR template. The Decision's
	// ID is assigned by the repository (next NNNN); Slug is derived from Title.
	Create(modelPath, subFolderPath string, decision *Decision) (*Decision, error)

	// Save writes back a Decision's frontmatter and body. The repository
	// regenerates the ## Comments body section from the frontmatter list on
	// every save.
	Save(modelPath string, decision *Decision, body string) error

	// LoadById returns the Decision with the given 4-digit ID.
	LoadById(modelPath, id string) (*Decision, error)

	// LoadByTitle returns the Decision whose slug matches the (slugified) title.
	// Exact match wins over partial; ambiguous partial matches error.
	LoadByTitle(modelPath, title string) (*Decision, error)

	// LoadAll returns all Decisions in the model directory. Errors if any file
	// uses the legacy ADG format (steers user to `adg migrate` in PR 4).
	LoadAll(modelPath string) ([]Decision, error)

	// LoadBody returns the raw markdown body (no frontmatter, no ## Comments
	// regeneration) for the given ID.
	LoadBody(modelPath, id string) (string, error)

	// FindDecisionFile returns the absolute file path for the given ID.
	FindDecisionFile(modelPath, id string) (string, error)

	// Copy duplicates a decision file from one model directory to another,
	// preserving subdirectory structure.
	Copy(srcPath, dstPath, decisionID string) error
}
