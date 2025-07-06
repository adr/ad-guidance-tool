package domain

import "fmt"

const (
	AnchorSectionQuestion = "question"
	AnchorSectionOptions  = "options"
	AnchorSectionCriteria = "criteria"
	AnchorSectionOutcome  = "outcome"
	AnchorSectionComments = "comments"
)

func AnchorForSection(section string) string {
	return fmt.Sprintf(`<a name="%s"></a>`, section)
}

func AnchorLinkToSection(section string, label string) string {
	return fmt.Sprintf("[%s](#%s)", label, section)
}

func AnchorForOption(number int) string {
	return fmt.Sprintf(`<a name="option-%d"></a>`, number)
}

func AnchorLinkToOption(number int) string {
	return fmt.Sprintf("[Option %d](#option-%d)", number, number)
}

func AnchorForComment(number int, author, date, text string) string {
	return fmt.Sprintf(`<a name="comment-%d"></a>%d. (%s) %s: %s`, number, number, date, author, text)
}

func AnchorLinkToComment(number int) string {
	return fmt.Sprintf("[Comment %d](#comment-%d)", number, number)
}
