package domain

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/adr/ad-guidance-tool/internal/ade/parser"
)

// customBlockRE matches the header of a custom block:
//
//	custom "name" {
//
// It captures the rule name in group 1.
var customBlockRE = regexp.MustCompile(
	`(?m)custom\s+"((?:[^"\\]|\\.)*)"` + `\s*\{`,
)

// extractCustomBlocks scans src for custom blocks, extracts each one into a
// RuleIR with IsCustomRule=true, and replaces the block in src with an equal
// number of newlines so that ANTLR line numbers remain correct.
func extractCustomBlocks(src string) (string, []*RuleIR, error) {
	var rules []*RuleIR
	cleaned := src

	for {
		loc := customBlockRE.FindStringIndex(cleaned)
		if loc == nil {
			break
		}

		// Extract header fields
		match := customBlockRE.FindStringSubmatch(cleaned[loc[0]:])
		ruleName := match[1]

		// Compute the line number of the custom keyword for error reporting
		line := strings.Count(cleaned[:loc[0]], "\n") + 1

		// loc[1] points just past the opening '{'. Find the matching '}'.
		braceStart := loc[1] - 1 // position of '{'
		closePos := findMatchingBrace(cleaned, braceStart)
		if closePos < 0 {
			return "", nil, fmt.Errorf("line %d: unterminated custom block %q", line, ruleName)
		}

		// Extract raw body (text between { and })
		rawBody := cleaned[braceStart+1 : closePos]
		rawBody = strings.TrimSpace(rawBody)

		rules = append(rules, &RuleIR{
			Name:         ruleName,
			IsCustomRule: true,
			RawBody:      rawBody,
		})

		// Replace the entire block with newlines to preserve line numbers
		blockText := cleaned[loc[0] : closePos+1]
		newlineCount := strings.Count(blockText, "\n")
		replacement := strings.Repeat("\n", newlineCount)
		cleaned = cleaned[:loc[0]] + replacement + cleaned[closePos+1:]
	}

	return cleaned, rules, nil
}

// findMatchingBrace returns the index of the closing '}' that matches the
// opening '{' at position openPos, respecting nested braces. Returns -1 if
// no matching brace is found.
func findMatchingBrace(src string, openPos int) int {
	depth := 1
	for i := openPos + 1; i < len(src); i++ {
		switch src[i] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	return -1
}

// ParseDSL parses a DSL source string using the ANTLR grammar and returns
// the corresponding protobuf SpecIR. It performs semantic validation after
// building the IR.
func ParseDSL(src string) (*SpecIR, error) {
	// Extract custom blocks before ANTLR parsing. Custom block bodies contain
	// arbitrary text that the lexer cannot tokenize, so we pull them out first
	// and replace each block with newlines to preserve line numbers.
	cleaned, customRules, err := extractCustomBlocks(src)
	if err != nil {
		return nil, err
	}

	// Lex & parse
	input := antlr.NewInputStream(cleaned)
	lexer := parser.NewADELexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewADEParser(stream)

	// Collect syntax errors
	errListener := &dslErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)
	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)

	tree := p.File()

	if len(errListener.errors) > 0 {
		return nil, fmt.Errorf("syntax errors:\n%s", strings.Join(errListener.errors, "\n"))
	}

	// Visit the parse tree and build SpecIR
	visitor := &irVisitor{}
	result := tree.Accept(visitor)

	if visitor.err != nil {
		return nil, visitor.err
	}

	ir, ok := result.(*SpecIR)
	if !ok {
		return nil, fmt.Errorf("visitor did not return SpecIR")
	}

	// Append custom rules extracted during pre-processing
	ir.Rules = append(ir.Rules, customRules...)

	// Semantic validation
	if err := validateIR(ir); err != nil {
		return nil, err
	}

	return ir, nil
}

// ---------- ANTLR error listener ----------

type dslErrorListener struct {
	antlr.DefaultErrorListener
	errors []string
}

func (l *dslErrorListener) SyntaxError(
	_ antlr.Recognizer, _ interface{},
	line, column int,
	msg string,
	_ antlr.RecognitionException,
) {
	l.errors = append(l.errors, fmt.Sprintf("line %d:%d: %s", line, column, msg))
}

// ---------- IR visitor ----------

type irVisitor struct {
	parser.BaseADEVisitor
	err error
}

// unquote strips the surrounding double quotes from an ANTLR STRING token
func unquote(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	return strings.ReplaceAll(s, `\"`, `"`)
}

// VisitFile builds the complete SpecIR from the file context
func (v *irVisitor) VisitFile(ctx *parser.FileContext) interface{} {
	ir := &SpecIR{}
	selectors := make(map[string]bool)
	ruleNames := make(map[string]bool)

	// Visit ADR declaration
	if adrCtx := ctx.AdrDecl(); adrCtx != nil {
		if adr, ok := adrCtx.(*parser.AdrDeclContext); ok {
			adrIR := v.VisitAdrDecl(adr)
			if v.err != nil {
				return nil
			}
			ir.Adr = adrIR.(*AdrIR)
		}
	}

	// Visit all selector declarations
	for _, selCtx := range ctx.AllSelectorDecl() {
		if sel, ok := selCtx.(*parser.SelectorDeclContext); ok {
			selIR := v.VisitSelectorDecl(sel)
			if v.err != nil {
				return nil
			}
			selector := selIR.(*SelectorIR)

			// Check for duplicates
			if selectors[selector.Name] {
				v.err = fmt.Errorf("line %d: duplicate selector name %q", sel.GetStart().GetLine(), selector.Name)
				return nil
			}
			selectors[selector.Name] = true
			ir.Selectors = append(ir.Selectors, selector)
		}
	}

	// Visit all rule declarations
	for _, ruleCtx := range ctx.AllRuleDecl() {
		if ruleC, ok := ruleCtx.(*parser.RuleDeclContext); ok {
			ruleIR := v.VisitRuleDecl(ruleC)
			if v.err != nil {
				return nil
			}
			r := ruleIR.(*RuleIR)

			// Check for duplicates
			if ruleNames[r.Name] {
				v.err = fmt.Errorf("line %d: duplicate rule name %q", ruleC.GetStart().GetLine(), r.Name)
				return nil
			}
			ruleNames[r.Name] = true
			ir.Rules = append(ir.Rules, r)
		}
	}

	return ir
}

// VisitAdrDecl extracts ADR id and title
func (v *irVisitor) VisitAdrDecl(ctx *parser.AdrDeclContext) interface{} {
	strs := ctx.AllSTRING()
	if len(strs) < 2 {
		v.err = fmt.Errorf("line %d: adr requires id and title", ctx.GetStart().GetLine())
		return nil
	}
	return &AdrIR{
		Id:    unquote(strs[0].GetText()),
		Title: unquote(strs[1].GetText()),
	}
}

// VisitSelectorDecl extracts selector name, kind, and pattern
func (v *irVisitor) VisitSelectorDecl(ctx *parser.SelectorDeclContext) interface{} {
	strs := ctx.AllSTRING()
	if len(strs) < 2 {
		v.err = fmt.Errorf("line %d: selector requires name and pattern", ctx.GetStart().GetLine())
		return nil
	}

	var kind SelectorKind
	switch {
	case ctx.COMPONENT() != nil:
		kind = SelectorKind_SELECTOR_COMPONENT
	case ctx.CLASS() != nil:
		kind = SelectorKind_SELECTOR_CLASS
	case ctx.INTERFACE() != nil:
		kind = SelectorKind_SELECTOR_INTERFACE
	case ctx.PATH() != nil:
		kind = SelectorKind_SELECTOR_PATH
	}

	return &SelectorIR{
		Name:    unquote(strs[0].GetText()),
		Kind:    kind,
		Pattern: unquote(strs[1].GetText()),
	}
}

// VisitRuleDecl builds a complete rule with all its statements
func (v *irVisitor) VisitRuleDecl(ctx *parser.RuleDeclContext) interface{} {
	rule := &RuleIR{
		Name: unquote(ctx.STRING().GetText()),
	}

	// Determine if this is a file rule or code rule
	if ruleTypeCtx := ctx.RuleType(); ruleTypeCtx != nil {
		if rt, ok := ruleTypeCtx.(*parser.RuleTypeContext); ok {
			rule.IsFileRule = rt.FILE() != nil
		}
	}

	// Process all statements in the rule
	for _, stmtCtx := range ctx.AllRuleStmt() {
		// Assertion statement
		if assertCtx := stmtCtx.AssertionStmt(); assertCtx != nil {
			if assert, ok := assertCtx.(*parser.AssertionStmtContext); ok {
				v.visitAssertionStmt(assert, rule)
				if v.err != nil {
					return nil
				}
			}
		}

		// Exclude statement
		if exclCtx := stmtCtx.ExcludeStmt(); exclCtx != nil {
			excl := v.visitExcludeStmt(exclCtx)
			if v.err != nil {
				return nil
			}
			if excl != nil {
				rule.Excludes = append(rule.Excludes, excl)
			}
		}

		// Severity statement
		if sevCtx := stmtCtx.SeverityStmt(); sevCtx != nil {
			if sev, ok := sevCtx.(*parser.SeverityStmtContext); ok {
				rule.Severity = v.visitSeverityStmt(sev)
			}
		}
	}

	return rule
}

// visitAssertionStmt processes an assertion and updates the rule
func (v *irVisitor) visitAssertionStmt(ctx *parser.AssertionStmtContext, rule *RuleIR) {
	// Extract subject
	rule.From = v.visitSubjectExpr(ctx.SubjectExpr())
	if v.err != nil {
		return
	}

	// Extract modality
	mod := v.visitMustExpr(ctx.MustExpr())

	// Extract verb phrase and set rule kind + targets
	v.visitVerbPhrase(ctx.VerbPhrase(), mod, rule)
}

// visitSubjectExpr converts a subjectExpr to TargetRefIR
func (v *irVisitor) visitSubjectExpr(ctx parser.ISubjectExprContext) *TargetRefIR {
	if ctx == nil {
		return nil
	}

	switch c := ctx.(type) {
	case *parser.SelectorRefContext:
		return &TargetRefIR{
			Value:    c.IDENTIFIER().GetText(),
			IsInline: false,
		}
	case *parser.InlineLiteralContext:
		return &TargetRefIR{
			Value:    unquote(c.STRING().GetText()),
			IsInline: true,
			Type:     v.getSelectorKind(c.SelectorType()),
			IsMatch:  false,
		}
	case *parser.InlineMatchContext:
		return &TargetRefIR{
			Value:    unquote(c.STRING().GetText()),
			IsInline: true,
			Type:     v.getSelectorKind(c.SelectorType()),
			IsMatch:  true,
		}
	case *parser.InlineTypeContext:
		return &TargetRefIR{
			Value:    "",
			IsInline: true,
			Type:     v.getSelectorKind(c.SelectorType()),
			IsMatch:  false,
		}
	case *parser.SubsetAllContext:
		scope := v.visitTargetExpr(c.TargetExpr())
		return &TargetRefIR{
			Value:    "",
			IsInline: true,
			Type:     v.getSelectorKind(c.SelectorType()),
			IsMatch:  false,
			Scope:    scope,
		}
	case *parser.SubsetLiteralContext:
		scope := v.visitTargetExpr(c.TargetExpr())
		return &TargetRefIR{
			Value:    unquote(c.STRING().GetText()),
			IsInline: true,
			Type:     v.getSelectorKind(c.SelectorType()),
			IsMatch:  false,
			Scope:    scope,
		}
	case *parser.SubsetMatchContext:
		scope := v.visitTargetExpr(c.TargetExpr())
		return &TargetRefIR{
			Value:    unquote(c.STRING().GetText()),
			IsInline: true,
			Type:     v.getSelectorKind(c.SelectorType()),
			IsMatch:  true,
			Scope:    scope,
		}
	}
	return nil
}

// visitTargetExpr converts a targetExpr to TargetRefIR
func (v *irVisitor) visitTargetExpr(ctx parser.ITargetExprContext) *TargetRefIR {
	if ctx == nil {
		return nil
	}

	switch c := ctx.(type) {
	case *parser.TargetSelectorRefContext:
		return &TargetRefIR{
			Value:    c.IDENTIFIER().GetText(),
			IsInline: false,
		}
	case *parser.TargetInlineLiteralContext:
		return &TargetRefIR{
			Value:    unquote(c.STRING().GetText()),
			IsInline: true,
			Type:     v.getSelectorKind(c.SelectorType()),
			IsMatch:  false,
		}
	case *parser.TargetInlineMatchContext:
		return &TargetRefIR{
			Value:    unquote(c.STRING().GetText()),
			IsInline: true,
			Type:     v.getSelectorKind(c.SelectorType()),
			IsMatch:  true,
		}
	case *parser.TargetStringLiteralContext:
		return &TargetRefIR{
			Value:    unquote(c.STRING().GetText()),
			IsInline: true,
			IsMatch:  false,
		}
	}
	return nil
}

// getSelectorKind converts a grammar selectorType to protobuf SelectorKind
func (v *irVisitor) getSelectorKind(ctx parser.ISelectorTypeContext) SelectorKind {
	if ctx == nil {
		return SelectorKind_SELECTOR_UNSPECIFIED
	}

	switch {
	case ctx.COMPONENT() != nil:
		return SelectorKind_SELECTOR_COMPONENT
	case ctx.CLASS() != nil:
		return SelectorKind_SELECTOR_CLASS
	case ctx.INTERFACE() != nil:
		return SelectorKind_SELECTOR_INTERFACE
	case ctx.PATH() != nil:
		return SelectorKind_SELECTOR_PATH
	}
	return SelectorKind_SELECTOR_UNSPECIFIED
}

// modality represents the must expression type
type modality int

const (
	modalityMust modality = iota
	modalityMustNot
	modalityMustOnly
)

// visitMustExpr determines the modality
func (v *irVisitor) visitMustExpr(ctx parser.IMustExprContext) modality {
	if ctx == nil {
		return modalityMust
	}

	mustCtx, ok := ctx.(*parser.MustExprContext)
	if !ok {
		return modalityMust
	}

	if mustCtx.NOT() != nil {
		return modalityMustNot
	}
	if mustCtx.ONLY() != nil {
		return modalityMustOnly
	}
	return modalityMust
}

// visitVerbPhrase processes the verb phrase and updates the rule
func (v *irVisitor) visitVerbPhrase(ctx parser.IVerbPhraseContext, mod modality, rule *RuleIR) {
	if ctx == nil {
		return
	}

	switch c := ctx.(type) {
	case *parser.DependOnPhraseContext:
		// Collect targets
		for _, tctx := range c.AllTargetExpr() {
			target := v.visitTargetExpr(tctx)
			if target != nil {
				rule.Targets = append(rule.Targets, target)
			}
		}
		// Set rule kind
		switch mod {
		case modalityMustNot:
			rule.Kind = RuleKind_RULE_NOT_DEPEND
		case modalityMustOnly:
			rule.Kind = RuleKind_RULE_DEPEND_ONLY
		case modalityMust:
			v.err = fmt.Errorf("invalid dependency rule: use 'must not' or 'must only', not plain 'must'")
		}

	case *parser.ExistPhraseContext:
		if mod == modalityMustOnly {
			v.err = fmt.Errorf("invalid 'must only exist' - use 'must exist'")
			return
		}
		kind := CheckKind_CHECK_FS_MUST_EXIST
		if mod == modalityMustNot {
			kind = CheckKind_CHECK_FS_MUST_NOT_EXIST
		}
		rule.Checks = append(rule.Checks, &CheckIR{
			Kind: kind,
			Path: rule.From.Value,
		})

	case *parser.ContainPhraseContext:
		if mod == modalityMustOnly {
			v.err = fmt.Errorf("invalid 'must only contain' - use 'must contain'")
			return
		}
		pattern := unquote(c.STRING().GetText())
		kind := CheckKind_CHECK_FS_MUST_CONTAIN
		if mod == modalityMustNot {
			kind = CheckKind_CHECK_FS_MUST_NOT_CONTAIN
		}
		rule.Checks = append(rule.Checks, &CheckIR{
			Kind:    kind,
			Path:    rule.From.Value,
			Pattern: pattern,
		})

	case *parser.ImplementPhraseContext:
		target := v.visitTargetExpr(c.TargetExpr())
		if target != nil {
			rule.Targets = append(rule.Targets, target)
		}
		switch mod {
		case modalityMust, modalityMustOnly:
			rule.Kind = RuleKind_RULE_IMPLEMENT
		case modalityMustNot:
			rule.Kind = RuleKind_RULE_NOT_IMPLEMENT
		}

	case *parser.ExtendPhraseContext:
		target := v.visitTargetExpr(c.TargetExpr())
		if target != nil {
			rule.Targets = append(rule.Targets, target)
		}
		switch mod {
		case modalityMust, modalityMustOnly:
			rule.Kind = RuleKind_RULE_EXTEND
		case modalityMustNot:
			rule.Kind = RuleKind_RULE_NOT_EXTEND
		}

	case *parser.AnnotatedPhraseContext:
		annotation := unquote(c.STRING().GetText())
		rule.Targets = append(rule.Targets, &TargetRefIR{
			Value:    annotation,
			IsInline: true,
		})
		switch mod {
		case modalityMust:
			rule.Kind = RuleKind_RULE_ANNOTATE
		case modalityMustNot:
			rule.Kind = RuleKind_RULE_NOT_ANNOTATE
		default:
			v.err = fmt.Errorf("invalid annotation rule: use 'must be annotated with' or 'must not be annotated with'")
		}

	case *parser.AccessedByPhraseContext:
		for _, tctx := range c.AllTargetExpr() {
			target := v.visitTargetExpr(tctx)
			if target != nil {
				rule.Targets = append(rule.Targets, target)
			}
		}
		if mod != modalityMustOnly {
			v.err = fmt.Errorf("accessed by rule requires 'must only' modality")
			return
		}
		rule.Kind = RuleKind_RULE_ACCESSED_BY

	case *parser.AcyclicPhraseContext:
		if mod != modalityMust {
			v.err = fmt.Errorf("acyclic rule requires 'must' modality (not 'must not' or 'must only')")
			return
		}
		rule.Kind = RuleKind_RULE_ACYCLIC

	case *parser.InPhraseContext:
		target := v.visitTargetExpr(c.TargetExpr())
		if target != nil {
			rule.Targets = append(rule.Targets, target)
		}
		switch mod {
		case modalityMust:
			rule.Kind = RuleKind_RULE_IN
		case modalityMustNot:
			rule.Kind = RuleKind_RULE_NOT_IN
		default:
			v.err = fmt.Errorf("location rule: use 'must be in' or 'must not be in'")
		}

	case *parser.MatchPhraseContext:
		pattern := unquote(c.STRING().GetText())
		rule.Targets = append(rule.Targets, &TargetRefIR{
			Value:    pattern,
			IsInline: true,
		})
		switch mod {
		case modalityMust:
			rule.Kind = RuleKind_RULE_MATCH
		case modalityMustNot:
			rule.Kind = RuleKind_RULE_NOT_MATCH
		default:
			v.err = fmt.Errorf("naming pattern rule: use 'must match' or 'must not match'")
		}

	case *parser.VisibilityPhraseContext:
		if mod != modalityMust {
			v.err = fmt.Errorf("visibility rule requires 'must' modality")
			return
		}
		rule.Kind = RuleKind_RULE_VISIBILITY
		visCtx := c.Visibility()
		if v, ok := visCtx.(*parser.VisibilityContext); ok {
			switch {
			case v.PUBLIC() != nil:
				rule.Visibility = Visibility_VISIBILITY_PUBLIC
			case v.INTERNAL() != nil:
				rule.Visibility = Visibility_VISIBILITY_INTERNAL
			case v.PRIVATE() != nil:
				rule.Visibility = Visibility_VISIBILITY_PRIVATE
			}
		}

	case *parser.TypeConstraintPhraseContext:
		if mod != modalityMust {
			v.err = fmt.Errorf("type constraint rule requires 'must' modality")
			return
		}
		rule.Kind = RuleKind_RULE_TYPE_CONSTRAINT
		tcCtx := c.TypeConstraint()
		if tc, ok := tcCtx.(*parser.TypeConstraintContext); ok {
			switch {
			case tc.ABSTRACT() != nil:
				rule.TypeConstraint = TypeConstraint_TYPE_CONSTRAINT_ABSTRACT
			case tc.SEALED() != nil:
				rule.TypeConstraint = TypeConstraint_TYPE_CONSTRAINT_SEALED
			case tc.STATIC() != nil:
				rule.TypeConstraint = TypeConstraint_TYPE_CONSTRAINT_STATIC
			}
		}
	}
}

// visitExcludeStmt processes an exclude statement
func (v *irVisitor) visitExcludeStmt(ctx parser.IExcludeStmtContext) *ExclusionIR {
	if ctx == nil {
		return nil
	}

	switch c := ctx.(type) {
	case *parser.ExcludeClassContext:
		return &ExclusionIR{
			Kind:  ExcludeKind_EXCLUDE_CLASS,
			Value: unquote(c.STRING().GetText()),
		}
	case *parser.ExcludeClassImplementingContext:
		return &ExclusionIR{
			Kind:  ExcludeKind_EXCLUDE_IMPLEMENT_INTERFACE,
			Value: unquote(c.STRING().GetText()),
		}
	case *parser.ExcludeComponentContext:
		return &ExclusionIR{
			Kind:  ExcludeKind_EXCLUDE_COMPONENT,
			Value: unquote(c.STRING().GetText()),
		}
	case *parser.ExcludePatternContext:
		return &ExclusionIR{
			Kind:  ExcludeKind_EXCLUDE_CLASS,
			Value: unquote(c.STRING().GetText()),
		}
	}
	return nil
}

// visitSeverityStmt extracts severity value
func (v *irVisitor) visitSeverityStmt(ctx *parser.SeverityStmtContext) Severity {
	if ctx == nil {
		return Severity_SEVERITY_UNSPECIFIED
	}

	svCtx := ctx.SeverityValue()
	if svCtx == nil {
		return Severity_SEVERITY_UNSPECIFIED
	}

	switch {
	case svCtx.ERROR() != nil:
		return Severity_SEVERITY_ERROR
	case svCtx.WARNING() != nil:
		return Severity_SEVERITY_WARNING
	}
	return Severity_SEVERITY_UNSPECIFIED
}

// ---------- Validation ----------

// selectorRefRE matches strings that look like IDENTIFIER tokens: they were
// produced by the grammar rule IDENTIFIER = [A-Z][a-zA-Z0-9_]* and therefore
// denote a named selector reference rather than a literal filesystem path.
var selectorRefRE = regexp.MustCompile(`^[A-Z][a-zA-Z0-9_]*$`)

// validateIR performs semantic checks on the built IR.
func validateIR(ir *SpecIR) error {
	if ir.Adr == nil {
		return fmt.Errorf("missing ADR declaration")
	}

	// Build selector map for validation
	selectors := make(map[string]bool)
	for _, sel := range ir.Selectors {
		selectors[sel.Name] = true
	}

	// Track rule names to detect duplicates across all rule types
	ruleNames := make(map[string]bool)

	for _, r := range ir.Rules {
		// Custom rules: validate header only (name and plugin), skip body
		if r.IsCustomRule {
			if r.Name == "" {
				return fmt.Errorf("custom rule: missing name")
			}
			if ruleNames[r.Name] {
				return fmt.Errorf("duplicate rule name %q", r.Name)
			}
			ruleNames[r.Name] = true
			continue
		}

		// Validate file vs code rule assertions
		hasChecks := len(r.Checks) > 0
		hasCodeAssertion := r.Kind != RuleKind_RULE_UNSPECIFIED && !hasChecks

		if r.IsFileRule && hasCodeAssertion {
			return fmt.Errorf("rule %q: code assertions (depend on, implement, etc.) cannot be used in file rules", r.Name)
		}
		if !r.IsFileRule && hasChecks {
			return fmt.Errorf("rule %q: file system assertions (exist, contain) cannot be used in code rules", r.Name)
		}

		// Code rules with subjects require validation
		codeRulesWithSubject := []RuleKind{
			RuleKind_RULE_DEPEND_ONLY,
			RuleKind_RULE_NOT_DEPEND,
			RuleKind_RULE_ANNOTATE,
			RuleKind_RULE_NOT_ANNOTATE,
			RuleKind_RULE_EXTEND,
			RuleKind_RULE_NOT_EXTEND,
			RuleKind_RULE_IMPLEMENT,
			RuleKind_RULE_NOT_IMPLEMENT,
			RuleKind_RULE_ACCESSED_BY,
			RuleKind_RULE_ACYCLIC,
			RuleKind_RULE_IN,
			RuleKind_RULE_NOT_IN,
			RuleKind_RULE_MATCH,
			RuleKind_RULE_NOT_MATCH,
			RuleKind_RULE_VISIBILITY,
			RuleKind_RULE_TYPE_CONSTRAINT,
		}

		isCodeRuleWithSubject := false
		for _, k := range codeRulesWithSubject {
			if r.Kind == k {
				isCodeRuleWithSubject = true
				break
			}
		}

		if isCodeRuleWithSubject {
			if r.From == nil {
				return fmt.Errorf("rule %q: missing subject", r.Name)
			}

			// Validate selector references (not inline patterns)
			if !r.From.IsInline && !selectors[r.From.Value] {
				return fmt.Errorf("rule %q: unknown selector %q", r.Name, r.From.Value)
			}

			// Validate scope selector references
			if r.From.Scope != nil && !r.From.Scope.IsInline && !selectors[r.From.Scope.Value] {
				return fmt.Errorf("rule %q: unknown selector %q in scope", r.Name, r.From.Scope.Value)
			}

			// Validate target selector references
			for _, t := range r.Targets {
				if !t.IsInline && !selectors[t.Value] {
					return fmt.Errorf("rule %q: unknown selector %q", r.Name, t.Value)
				}
			}
		}

		// Validate selector references used as path subjects in file rules.
		// A path that matches the IDENTIFIER pattern ([A-Z][a-zA-Z0-9_]*) was
		// produced by a SelectorRef in the grammar and must resolve to a
		// declared selector.
		if r.IsFileRule {
			for _, c := range r.Checks {
				if selectorRefRE.MatchString(c.Path) && !selectors[c.Path] {
					// Produce a helpful hint when the selector was probably
					// declared with a lowercase-first name, which can never be
					// referenced because IDENTIFIER requires an uppercase first
					// letter.
					lower := strings.ToLower(c.Path[:1]) + c.Path[1:]
					if selectors[lower] {
						return fmt.Errorf(
							"rule %q: undefined selector %q (selector %q is defined but selector names must start with an uppercase letter to be referenced by name)",
							r.Name, c.Path, lower,
						)
					}
					return fmt.Errorf("rule %q: undefined selector %q", r.Name, c.Path)
				}
			}
		}

		// Track name for cross-type duplicate detection (e.g. custom vs code rule)
		ruleNames[r.Name] = true
	}

	return nil
}
