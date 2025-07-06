package decision

import "fmt"

type CommentDecisionPresenter struct{}

func NewCommentPresenter() *CommentDecisionPresenter {
	return &CommentDecisionPresenter{}
}

func (p *CommentDecisionPresenter) Commented(decisionID, author, comment string) {
	fmt.Printf("Comment added by %s to decision %s: \"%s\"\n", author, decisionID, comment)
}
