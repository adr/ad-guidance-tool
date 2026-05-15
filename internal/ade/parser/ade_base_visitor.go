// Code generated from ADE.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ADE

import "github.com/antlr4-go/antlr/v4"

type BaseADEVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseADEVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitAdrDecl(ctx *AdrDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitSelectorDecl(ctx *SelectorDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitRuleDecl(ctx *RuleDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitRuleType(ctx *RuleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitRuleStmt(ctx *RuleStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitAssertionStmt(ctx *AssertionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitSelectorRef(ctx *SelectorRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitSubsetLiteral(ctx *SubsetLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitSubsetMatch(ctx *SubsetMatchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitSubsetAll(ctx *SubsetAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitInlineLiteral(ctx *InlineLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitInlineMatch(ctx *InlineMatchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitInlineType(ctx *InlineTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitTargetSelectorRef(ctx *TargetSelectorRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitTargetInlineLiteral(ctx *TargetInlineLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitTargetInlineMatch(ctx *TargetInlineMatchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitTargetStringLiteral(ctx *TargetStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitSelectorType(ctx *SelectorTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitMustExpr(ctx *MustExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitDependOnPhrase(ctx *DependOnPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitExistPhrase(ctx *ExistPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitContainPhrase(ctx *ContainPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitImplementPhrase(ctx *ImplementPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitExtendPhrase(ctx *ExtendPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitAnnotatedPhrase(ctx *AnnotatedPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitAccessedByPhrase(ctx *AccessedByPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitAcyclicPhrase(ctx *AcyclicPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitInPhrase(ctx *InPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitMatchPhrase(ctx *MatchPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitVisibilityPhrase(ctx *VisibilityPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitTypeConstraintPhrase(ctx *TypeConstraintPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitVisibility(ctx *VisibilityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitTypeConstraint(ctx *TypeConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitExcludeClass(ctx *ExcludeClassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitExcludeClassImplementing(ctx *ExcludeClassImplementingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitExcludeComponent(ctx *ExcludeComponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitExcludePattern(ctx *ExcludePatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitSeverityStmt(ctx *SeverityStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseADEVisitor) VisitSeverityValue(ctx *SeverityValueContext) interface{} {
	return v.VisitChildren(ctx)
}
