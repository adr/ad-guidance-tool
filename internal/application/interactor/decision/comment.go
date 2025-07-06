package decision

import (
	"adg/internal/application/inputport"
	util "adg/internal/application/interactor"
	"adg/internal/application/outputport"
	domain "adg/internal/domain/decision"
	"fmt"
)

type CommentDecisionInteractor struct {
	service domain.DecisionService
	output  outputport.DecisionComment
}

func NewCommentDecisionInteractor(service domain.DecisionService, output outputport.DecisionComment) inputport.DecisionComment {
	return &CommentDecisionInteractor{
		service: service,
		output:  output,
	}
}

func (i *CommentDecisionInteractor) Comment(modelPath, id, title, author, comment string) error {
	var (
		decision *domain.Decision
		err      error
	)

	decision, err = util.ResolveDecisionByIdOrTitle(modelPath, id, title, i.service)
	if err != nil {
		return err
	}

	if err := i.service.Comment(modelPath, decision, author, comment); err != nil {
		return fmt.Errorf("failed to add comment: %w", err)
	}

	i.output.Commented(decision.ID, author, comment)
	return nil
}
