package domain

import (
	"strings"
	"testing"
)

func TestParseDSL_CleanArchitecture(t *testing.T) {
	input := `
adr "10" "Use Clean Architecture for writes"

component "Domain"         = "CompanyName.MyMeetings.Modules.*.."
component "Application"    = "CompanyName.MyMeetings.Modules.*.Application.."
component "Infrastructure" = "CompanyName.MyMeetings.Modules.*.Infrastructure.."
component "API"            = "CompanyName.MyMeetings.API.."

code "domain_is_inner" {
  Domain must only depend on Domain
  severity error
}

code "application_only_inward" {
  Application must only depend on Application, Domain
  severity error
}
`
	ir, err := ParseDSL(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// ADR
	if ir.Adr == nil {
		t.Fatal("expected ADR to be set")
	}
	if ir.Adr.Id != "10" || ir.Adr.Title != "Use Clean Architecture for writes" {
		t.Errorf("ADR = %q / %q", ir.Adr.Id, ir.Adr.Title)
	}

	// Selectors
	if len(ir.Selectors) != 4 {
		t.Fatalf("expected 4 selectors, got %d", len(ir.Selectors))
	}
	for _, s := range ir.Selectors {
		if s.Kind != SelectorKind_SELECTOR_COMPONENT {
			t.Errorf("selector %q: expected COMPONENT kind, got %v", s.Name, s.Kind)
		}
	}

	// Rules
	if len(ir.Rules) != 2 {
		t.Fatalf("expected 2 rules, got %d", len(ir.Rules))
	}
	r0 := ir.Rules[0]
	if r0.Name != "domain_is_inner" {
		t.Errorf("rule[0].Name = %q", r0.Name)
	}
	if r0.Kind != RuleKind_RULE_DEPEND_ONLY {
		t.Errorf("rule[0].Kind = %v", r0.Kind)
	}
	// Check From (TargetRefIR)
	if r0.From == nil {
		t.Fatal("rule[0].From is nil")
	}
	if r0.From.Value != "Domain" || r0.From.IsInline {
		t.Errorf("rule[0].From = %+v, want selector ref to 'Domain'", r0.From)
	}
	// Check Targets
	if len(r0.Targets) != 1 {
		t.Fatalf("rule[0]: expected 1 target, got %d", len(r0.Targets))
	}
	if r0.Targets[0] == nil || r0.Targets[0].Value != "Domain" || r0.Targets[0].IsInline {
		t.Errorf("rule[0].Targets[0] = %+v, want selector ref to 'Domain'", r0.Targets[0])
	}
	if r0.Severity != Severity_SEVERITY_ERROR {
		t.Errorf("rule[0].Severity = %v", r0.Severity)
	}
}

func TestParseDSL_ModuleIsolation(t *testing.T) {
	input := `
adr "14" "Event-driven communication between modules"

component "Meetings"       = "CompanyName.MyMeetings.Modules.Meetings"
component "Payments"       = "CompanyName.MyMeetings.Modules.Payments"
component "UserAccess"     = "CompanyName.MyMeetings.Modules.UserAccess"
component "Administration" = "CompanyName.MyMeetings.Modules.Administration"

code "meetings_isolated_from_payments" {
  Meetings must not depend on Payments
  exclude class implementing interface "MediatR.INotificationHandler<>"
  severity error
}
`
	ir, err := ParseDSL(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(ir.Selectors) != 4 {
		t.Fatalf("expected 4 selectors, got %d", len(ir.Selectors))
	}
	for _, s := range ir.Selectors {
		if s.Kind != SelectorKind_SELECTOR_COMPONENT {
			t.Errorf("selector %q: expected COMPONENT kind, got %v", s.Name, s.Kind)
		}
	}

	if len(ir.Rules) != 1 {
		t.Fatalf("expected 1 rule, got %d", len(ir.Rules))
	}
	r := ir.Rules[0]
	if r.Kind != RuleKind_RULE_NOT_DEPEND {
		t.Errorf("rule kind = %v, want RULE_NOT_DEPEND", r.Kind)
	}
	// Check From
	if r.From == nil || r.From.Value != "Meetings" || r.From.IsInline {
		t.Errorf("rule.From = %+v, want selector ref to 'Meetings'", r.From)
	}
	// Check Targets
	if len(r.Targets) != 1 {
		t.Fatalf("expected 1 target, got %d", len(r.Targets))
	}
	if r.Targets[0] == nil || r.Targets[0].Value != "Payments" || r.Targets[0].IsInline {
		t.Errorf("rule.Targets[0] = %+v, want selector ref to 'Payments'", r.Targets[0])
	}
	// Check Exclusions
	if len(r.Excludes) != 1 {
		t.Fatalf("expected 1 exclusion, got %d", len(r.Excludes))
	}
	if r.Excludes[0].Kind != ExcludeKind_EXCLUDE_IMPLEMENT_INTERFACE {
		t.Errorf("exclude kind = %v", r.Excludes[0].Kind)
	}
	if r.Excludes[0].Value != "MediatR.INotificationHandler<>" {
		t.Errorf("exclude value = %q", r.Excludes[0].Value)
	}
}

func TestParseDSL_PathChecks(t *testing.T) {
	input := `
adr "17" "Implement Architecture Tests"

file "archtests_project_exists" {
  path "src/Tests/ArchTests/*.csproj" must exist
  severity error
}

file "netarchtest_is_referenced" {
  path "src/Directory.Packages.props" must contain "NetArchTest\\.Rules"
  severity warning
}
`
	ir, err := ParseDSL(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(ir.Rules) != 2 {
		t.Fatalf("expected 2 rules, got %d", len(ir.Rules))
	}

	r0 := ir.Rules[0]
	if len(r0.Checks) != 1 {
		t.Fatalf("rule[0]: expected 1 check, got %d", len(r0.Checks))
	}
	if r0.Checks[0].Kind != CheckKind_CHECK_FS_MUST_EXIST {
		t.Errorf("check kind = %v", r0.Checks[0].Kind)
	}
	if r0.Checks[0].Path != "src/Tests/ArchTests/*.csproj" {
		t.Errorf("check path = %q", r0.Checks[0].Path)
	}

	r1 := ir.Rules[1]
	if len(r1.Checks) != 1 {
		t.Fatalf("rule[1]: expected 1 check, got %d", len(r1.Checks))
	}
	if r1.Checks[0].Kind != CheckKind_CHECK_FS_MUST_CONTAIN {
		t.Errorf("check kind = %v", r1.Checks[0].Kind)
	}
	if r1.Checks[0].Pattern != `NetArchTest\\.Rules` {
		t.Errorf("check pattern = %q", r1.Checks[0].Pattern)
	}
	if r1.Severity != Severity_SEVERITY_WARNING {
		t.Errorf("severity = %v", r1.Severity)
	}
}

func TestParseDSL_SyntaxError(t *testing.T) {
	input := `adr "10"`
	_, err := ParseDSL(input)
	if err == nil {
		t.Fatal("expected error for incomplete adr declaration")
	}
}

func TestParseDSL_DuplicateSelector(t *testing.T) {
	input := `
adr "1" "Test"
component "A" = "foo"
component "A" = "bar"
`
	_, err := ParseDSL(input)
	if err == nil {
		t.Fatal("expected error for duplicate selector")
	}
}

func TestParseDSL_AnnotateVerb(t *testing.T) {
	input := `
adr "20" "Classes must be annotated"

component "Domain" = "MyApp.Domain"

code "domain_classes_annotated" {
  class "Handler" must be annotated with "DomainAttribute"
  severity error
}
`
	ir, err := ParseDSL(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(ir.Rules) != 1 {
		t.Fatalf("expected 1 rule, got %d", len(ir.Rules))
	}
	r := ir.Rules[0]
	if r.Kind != RuleKind_RULE_ANNOTATE {
		t.Errorf("rule kind = %v, want RULE_ANNOTATE", r.Kind)
	}
	if r.From == nil {
		t.Fatal("rule.From is nil")
	}
	if r.From.Value != "Handler" || !r.From.IsInline {
		t.Errorf("rule.From = %+v, want inline class 'Handler'", r.From)
	}
	if r.From.Type != SelectorKind_SELECTOR_CLASS {
		t.Errorf("rule.From.Type = %v, want SELECTOR_CLASS", r.From.Type)
	}
	if len(r.Targets) != 1 {
		t.Fatalf("expected 1 target, got %d", len(r.Targets))
	}
	if r.Targets[0].Value != "DomainAttribute" {
		t.Errorf("target value = %q, want 'DomainAttribute'", r.Targets[0].Value)
	}
}

func TestParseDSL_SubsetInRelation(t *testing.T) {
	input := `
adr "21" "Scoped class rules"

component "Domain" = "MyApp.Domain"

code "classes_in_domain_annotated" {
  class in Domain must be annotated with "DomainAttribute"
  severity error
}
`
	ir, err := ParseDSL(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(ir.Rules) != 1 {
		t.Fatalf("expected 1 rule, got %d", len(ir.Rules))
	}
	r := ir.Rules[0]
	if r.Kind != RuleKind_RULE_ANNOTATE {
		t.Errorf("rule kind = %v, want RULE_ANNOTATE", r.Kind)
	}
	if r.From == nil {
		t.Fatal("rule.From is nil")
	}
	// Subject should be class type with empty value (all classes)
	if r.From.Type != SelectorKind_SELECTOR_CLASS {
		t.Errorf("rule.From.Type = %v, want SELECTOR_CLASS", r.From.Type)
	}
	if r.From.Value != "" {
		t.Errorf("rule.From.Value = %q, want empty (all classes)", r.From.Value)
	}
	// Scope should be a selector ref to "Domain"
	if r.From.Scope == nil {
		t.Fatal("rule.From.Scope is nil, expected 'Domain' scope")
	}
	if r.From.Scope.Value != "Domain" || r.From.Scope.IsInline {
		t.Errorf("rule.From.Scope = %+v, want selector ref to 'Domain'", r.From.Scope)
	}
}

func TestParseDSL_SubsetLiteralInComponent(t *testing.T) {
	input := `
adr "22" "Subset literal"

component "Domain" = "MyApp.Domain"

code "handler_in_domain" {
  class "EventHandler" in Domain must be annotated with "Serializable"
  severity error
}
`
	ir, err := ParseDSL(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := ir.Rules[0]
	if r.From == nil {
		t.Fatal("rule.From is nil")
	}
	if r.From.Value != "EventHandler" {
		t.Errorf("rule.From.Value = %q, want 'EventHandler'", r.From.Value)
	}
	if r.From.Type != SelectorKind_SELECTOR_CLASS {
		t.Errorf("rule.From.Type = %v, want SELECTOR_CLASS", r.From.Type)
	}
	if r.From.Scope == nil {
		t.Fatal("rule.From.Scope is nil")
	}
	if r.From.Scope.Value != "Domain" || r.From.Scope.IsInline {
		t.Errorf("rule.From.Scope = %+v, want selector ref to 'Domain'", r.From.Scope)
	}
}

func TestParseDSL_SubsetMatchInComponent(t *testing.T) {
	input := `
adr "23" "Subset match"

code "handlers_in_domain" {
  class match "regex:.*Handler" in component "MyApp.Domain" must be annotated with "Serializable"
  severity error
}
`
	ir, err := ParseDSL(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := ir.Rules[0]
	if r.From == nil {
		t.Fatal("rule.From is nil")
	}
	if r.From.Value != "regex:.*Handler" {
		t.Errorf("rule.From.Value = %q, want 'regex:.*Handler'", r.From.Value)
	}
	if !r.From.IsMatch {
		t.Error("rule.From.IsMatch should be true")
	}
	if r.From.Scope == nil {
		t.Fatal("rule.From.Scope is nil")
	}
	if r.From.Scope.Value != "MyApp.Domain" || !r.From.Scope.IsInline {
		t.Errorf("rule.From.Scope = %+v, want inline component 'MyApp.Domain'", r.From.Scope)
	}
	if r.From.Scope.Type != SelectorKind_SELECTOR_COMPONENT {
		t.Errorf("scope type = %v, want SELECTOR_COMPONENT", r.From.Scope.Type)
	}
}

func TestParseDSL_SubsetInWithDependOn(t *testing.T) {
	input := `
adr "24" "Subset with depend on"

component "Domain" = "MyApp.Domain"
component "Infra"  = "MyApp.Infrastructure"

code "classes_in_domain_no_infra" {
  class in Domain must not depend on Infra
  severity error
}
`
	ir, err := ParseDSL(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := ir.Rules[0]
	if r.Kind != RuleKind_RULE_NOT_DEPEND {
		t.Errorf("rule kind = %v, want RULE_NOT_DEPEND", r.Kind)
	}
	if r.From == nil || r.From.Scope == nil {
		t.Fatal("rule.From or rule.From.Scope is nil")
	}
	if r.From.Scope.Value != "Domain" {
		t.Errorf("scope value = %q, want 'Domain'", r.From.Scope.Value)
	}
	if len(r.Targets) != 1 || r.Targets[0].Value != "Infra" {
		t.Errorf("targets = %+v, want single target 'Infra'", r.Targets)
	}
}

func TestParseDSL_CustomBlock_Basic(t *testing.T) {
	input := `
adr "99" "Custom rule"

custom "my_check" {
  any text the plugin understands
  can go here with whatever syntax
  the plugin author defines
}
`
	ir, err := ParseDSL(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(ir.Rules) != 1 {
		t.Fatalf("expected 1 rule, got %d", len(ir.Rules))
	}
	r := ir.Rules[0]

	if r.Name != "my_check" {
		t.Errorf("rule.Name = %q, want %q", r.Name, "my_check")
	}
	if !r.IsCustomRule {
		t.Error("rule.IsCustomRule should be true")
	}
	if r.RawBody == "" {
		t.Error("rule.RawBody should not be empty")
	}
	if !containsLine(r.RawBody, "any text the plugin understands") {
		t.Errorf("RawBody missing expected content, got: %q", r.RawBody)
	}
}

func TestParseDSL_CustomBlock_MixedWithOtherRules(t *testing.T) {
	input := `
adr "100" "Mixed rules"

component "Domain" = "MyApp.Domain"

code "domain_is_inner" {
  Domain must only depend on Domain
  severity error
}

custom "plugin_check" {
  check: all-modules
  threshold: 80
}

file "csproj_exists" {
  path "src/*.csproj" must exist
}
`
	ir, err := ParseDSL(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(ir.Rules) != 3 {
		t.Fatalf("expected 3 rules, got %d", len(ir.Rules))
	}

	// First rule is the code rule
	codeRule := ir.Rules[0]
	if codeRule.Name != "domain_is_inner" || codeRule.IsCustomRule {
		t.Errorf("rules[0] should be the code rule, got Name=%q IsCustomRule=%v", codeRule.Name, codeRule.IsCustomRule)
	}

	// Second rule is the file rule
	fileRule := ir.Rules[1]
	if fileRule.Name != "csproj_exists" || fileRule.IsCustomRule {
		t.Errorf("rules[1] should be the file rule, got Name=%q IsCustomRule=%v", fileRule.Name, fileRule.IsCustomRule)
	}

	// Custom rules are appended last
	customRule := ir.Rules[2]
	if customRule.Name != "plugin_check" || !customRule.IsCustomRule {
		t.Errorf("rules[2] should be the custom rule, got Name=%q IsCustomRule=%v", customRule.Name, customRule.IsCustomRule)
	}
	if !containsLine(customRule.RawBody, "check: all-modules") {
		t.Errorf("RawBody missing expected content, got: %q", customRule.RawBody)
	}
}

func TestParseDSL_CustomBlock_NestedBraces(t *testing.T) {
	input := `
adr "101" "Custom block with nested braces"

custom "json_check" {
  {
    "key": "value",
    "nested": { "inner": true }
  }
}
`
	ir, err := ParseDSL(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(ir.Rules) != 1 {
		t.Fatalf("expected 1 rule, got %d", len(ir.Rules))
	}
	r := ir.Rules[0]
	if !r.IsCustomRule {
		t.Error("rule.IsCustomRule should be true")
	}
	if !containsLine(r.RawBody, `"key": "value"`) {
		t.Errorf("RawBody missing expected content, got: %q", r.RawBody)
	}
}

func TestParseDSL_CustomBlock_UnterminatedBlock(t *testing.T) {
	input := `
adr "102" "Unterminated custom block"

custom "bad_check" {
  no closing brace here
`
	_, err := ParseDSL(input)
	if err == nil {
		t.Fatal("expected error for unterminated custom block")
	}
}

func TestParseDSL_CustomBlock_DuplicateName(t *testing.T) {
	input := `
adr "103" "Duplicate custom rule name"

custom "my_check" {
  body one
}

custom "my_check" {
  body two
}
`
	_, err := ParseDSL(input)
	if err == nil {
		t.Fatal("expected error for duplicate rule name")
	}
}

func TestParseDSL_CustomBlock_DuplicateNameWithRegularRule(t *testing.T) {
	input := `
adr "104" "Duplicate name across types"

component "Domain" = "MyApp.Domain"

code "my_rule" {
  Domain must only depend on Domain
}

custom "my_rule" {
  body
}
`
	_, err := ParseDSL(input)
	if err == nil {
		t.Fatal("expected error for duplicate rule name (custom conflicts with code rule)")
	}
}

// containsLine reports whether s contains the given substring.
func containsLine(s, substr string) bool {
	return strings.Contains(s, substr)
}
