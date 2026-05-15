// Code generated from ADE.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ADE

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ADEParser.
type ADEVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ADEParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by ADEParser#adrDecl.
	VisitAdrDecl(ctx *AdrDeclContext) interface{}

	// Visit a parse tree produced by ADEParser#selectorDecl.
	VisitSelectorDecl(ctx *SelectorDeclContext) interface{}

	// Visit a parse tree produced by ADEParser#ruleDecl.
	VisitRuleDecl(ctx *RuleDeclContext) interface{}

	// Visit a parse tree produced by ADEParser#ruleType.
	VisitRuleType(ctx *RuleTypeContext) interface{}

	// Visit a parse tree produced by ADEParser#ruleStmt.
	VisitRuleStmt(ctx *RuleStmtContext) interface{}

	// Visit a parse tree produced by ADEParser#assertionStmt.
	VisitAssertionStmt(ctx *AssertionStmtContext) interface{}

	// Visit a parse tree produced by ADEParser#SelectorRef.
	VisitSelectorRef(ctx *SelectorRefContext) interface{}

	// Visit a parse tree produced by ADEParser#SubsetLiteral.
	VisitSubsetLiteral(ctx *SubsetLiteralContext) interface{}

	// Visit a parse tree produced by ADEParser#SubsetMatch.
	VisitSubsetMatch(ctx *SubsetMatchContext) interface{}

	// Visit a parse tree produced by ADEParser#SubsetAll.
	VisitSubsetAll(ctx *SubsetAllContext) interface{}

	// Visit a parse tree produced by ADEParser#InlineLiteral.
	VisitInlineLiteral(ctx *InlineLiteralContext) interface{}

	// Visit a parse tree produced by ADEParser#InlineMatch.
	VisitInlineMatch(ctx *InlineMatchContext) interface{}

	// Visit a parse tree produced by ADEParser#InlineType.
	VisitInlineType(ctx *InlineTypeContext) interface{}

	// Visit a parse tree produced by ADEParser#TargetSelectorRef.
	VisitTargetSelectorRef(ctx *TargetSelectorRefContext) interface{}

	// Visit a parse tree produced by ADEParser#TargetInlineLiteral.
	VisitTargetInlineLiteral(ctx *TargetInlineLiteralContext) interface{}

	// Visit a parse tree produced by ADEParser#TargetInlineMatch.
	VisitTargetInlineMatch(ctx *TargetInlineMatchContext) interface{}

	// Visit a parse tree produced by ADEParser#TargetStringLiteral.
	VisitTargetStringLiteral(ctx *TargetStringLiteralContext) interface{}

	// Visit a parse tree produced by ADEParser#selectorType.
	VisitSelectorType(ctx *SelectorTypeContext) interface{}

	// Visit a parse tree produced by ADEParser#mustExpr.
	VisitMustExpr(ctx *MustExprContext) interface{}

	// Visit a parse tree produced by ADEParser#DependOnPhrase.
	VisitDependOnPhrase(ctx *DependOnPhraseContext) interface{}

	// Visit a parse tree produced by ADEParser#ExistPhrase.
	VisitExistPhrase(ctx *ExistPhraseContext) interface{}

	// Visit a parse tree produced by ADEParser#ContainPhrase.
	VisitContainPhrase(ctx *ContainPhraseContext) interface{}

	// Visit a parse tree produced by ADEParser#ImplementPhrase.
	VisitImplementPhrase(ctx *ImplementPhraseContext) interface{}

	// Visit a parse tree produced by ADEParser#ExtendPhrase.
	VisitExtendPhrase(ctx *ExtendPhraseContext) interface{}

	// Visit a parse tree produced by ADEParser#AnnotatedPhrase.
	VisitAnnotatedPhrase(ctx *AnnotatedPhraseContext) interface{}

	// Visit a parse tree produced by ADEParser#AccessedByPhrase.
	VisitAccessedByPhrase(ctx *AccessedByPhraseContext) interface{}

	// Visit a parse tree produced by ADEParser#AcyclicPhrase.
	VisitAcyclicPhrase(ctx *AcyclicPhraseContext) interface{}

	// Visit a parse tree produced by ADEParser#InPhrase.
	VisitInPhrase(ctx *InPhraseContext) interface{}

	// Visit a parse tree produced by ADEParser#MatchPhrase.
	VisitMatchPhrase(ctx *MatchPhraseContext) interface{}

	// Visit a parse tree produced by ADEParser#VisibilityPhrase.
	VisitVisibilityPhrase(ctx *VisibilityPhraseContext) interface{}

	// Visit a parse tree produced by ADEParser#TypeConstraintPhrase.
	VisitTypeConstraintPhrase(ctx *TypeConstraintPhraseContext) interface{}

	// Visit a parse tree produced by ADEParser#visibility.
	VisitVisibility(ctx *VisibilityContext) interface{}

	// Visit a parse tree produced by ADEParser#typeConstraint.
	VisitTypeConstraint(ctx *TypeConstraintContext) interface{}

	// Visit a parse tree produced by ADEParser#ExcludeClass.
	VisitExcludeClass(ctx *ExcludeClassContext) interface{}

	// Visit a parse tree produced by ADEParser#ExcludeClassImplementing.
	VisitExcludeClassImplementing(ctx *ExcludeClassImplementingContext) interface{}

	// Visit a parse tree produced by ADEParser#ExcludeComponent.
	VisitExcludeComponent(ctx *ExcludeComponentContext) interface{}

	// Visit a parse tree produced by ADEParser#ExcludePattern.
	VisitExcludePattern(ctx *ExcludePatternContext) interface{}

	// Visit a parse tree produced by ADEParser#severityStmt.
	VisitSeverityStmt(ctx *SeverityStmtContext) interface{}

	// Visit a parse tree produced by ADEParser#severityValue.
	VisitSeverityValue(ctx *SeverityValueContext) interface{}
}
