package decision

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	util "github.com/adr/ad-guidance-tool/internal/domain"
	domain "github.com/adr/ad-guidance-tool/internal/domain/decision"
	svc_mocks "github.com/adr/ad-guidance-tool/mocks/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestConfig() *svc_mocks.ConfigService {
	cfg := new(svc_mocks.ConfigService)
	cfg.On("GetQuestionHeader").Return("Question")
	cfg.On("GetOptionsHeader").Return("Options")
	cfg.On("GetCriteriaHeader").Return("Criteria")
	cfg.On("GetOutcomeHeader").Return("Outcome")
	cfg.On("GetCommentsHeader").Return("Comments")
	return cfg
}

func TestAppendCommentSection_KeepsCommentAddedByDecide(t *testing.T) {
	modelPath := t.TempDir()
	service := domain.NewDecisionService(NewFileDecisionRepository(newTestConfig()))

	decision, err := service.AddNew(modelPath, "Sample decision")
	require.NoError(t, err)
	require.NoError(t, service.Edit(modelPath, decision, nil, &[]string{"A", "B"}, nil))

	// `adg decide` records its own comment right after deciding
	require.NoError(t, service.Decide(modelPath, decision, "1", "", false))
	require.NoError(t, service.Comment(modelPath, decision, "", "marked decision as decided"))

	// `adg comment` runs as a separate process, so it starts from the persisted decision
	decision, err = service.GetDecisionByID(modelPath, decision.ID)
	require.NoError(t, err)
	require.NoError(t, service.Comment(modelPath, decision, "x", "second comment"))

	decision, err = service.GetDecisionByID(modelPath, decision.ID)
	require.NoError(t, err)
	require.Len(t, decision.Comments, 2)

	content, err := service.GetDecisionContent(modelPath, decision.ID)
	require.NoError(t, err)
	assert.Equal(t, strings.Join([]string{
		util.AnchorForComment(1, "", decision.Comments[0].Date, "marked decision as decided"),
		util.AnchorForComment(2, "x", decision.Comments[1].Date, "second comment"),
	}, "\n"), content.Comments)
}

func TestAppendCommentSection_KeepsSectionsAfterComments(t *testing.T) {
	modelPath := t.TempDir()
	filePath := filepath.Join(modelPath, "AD0001-sample-decision.md")
	frontmatter := "---\nadr_id: \"0001\"\nstatus: decided\ntitle: Sample decision\n---\n"

	require.NoError(t, os.WriteFile(filePath, []byte(frontmatter+`
## <a name="outcome"></a> Outcome
We decided for [Option 1](#option-1).

## <a name="comments"></a> Comments
<a name="comment-1"></a>1. (2026-01-01 00:00:00) : marked decision as decided

## Consequences
Written by hand after the comments.
`), 0644))

	repo := NewFileDecisionRepository(newTestConfig())
	require.NoError(t, repo.AppendCommentSection(modelPath, "0001", "second comment", 2, "x", "2026-01-02 00:00:00"))

	updated, err := os.ReadFile(filePath)
	require.NoError(t, err)
	assert.Equal(t, frontmatter+`
## <a name="outcome"></a> Outcome
We decided for [Option 1](#option-1).

## <a name="comments"></a> Comments
<a name="comment-1"></a>1. (2026-01-01 00:00:00) : marked decision as decided
<a name="comment-2"></a>2. (2026-01-02 00:00:00) x: second comment

## Consequences
Written by hand after the comments.
`, string(updated))
}
