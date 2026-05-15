# Rule File DSL

The rule file DSL encodes architectural rules from ADRs as enforceable checks. Rule files are parsed to an intermediate representation (IR) and then compiled or verified by language-specific plugins.

## Syntax

### File structure

```dsl
adr "<id>" "<title>"                    # Required header

<type> "<Name>" = "<pattern>"           # Optional selectors

file "<name>" {                         # File system rules
  <file_assertion>
  [exclude <pattern>]*
  severity <error|warning>              # Optional, default: error
}

code "<name>" {                         # Code structure rules
  <code_assertion>
  [exclude <pattern>]*
  severity <error|warning>              # Optional, default: error
}

custom "<name>" {                       # Opaque plugin-defined rules
  <any text the plugin understands>
}
```

#### Example

```dsl
adr "0001" "Clean Architecture Layering"

component "Domain" = "MyApp.Domain"
component "Application" = "MyApp.Application"

code "domain_isolated" {
  Domain must not depend on Application
  severity error
}
```

### Selectors

Define reusable references:

| Type | Purpose | Example pattern |
| ---- | ------- | --------------- |
| `component` | Layers/modules | `"MyApp.Domain.."` |
| `class` | Classes | `"EventHandler"` |
| `interface` | Interfaces | `"IRepository"` |
| `path` | Files/folders | `"src/**/*.cs"` |

Selector names must start with an uppercase letter (`A`-`Z`). Subsequent characters may be letters, digits, or underscores.

Selectors cannot be named after any of the reserved keywords: `must`, `not`, `only`, `depend`, `exist`, `contain`, `implement`, `interface`, `match`, `exclude`, `class`, `component`, `path`, `severity`, `error`, `warning`, `extend`, `accessed`, `annotated`, `acyclic`, `public`, `internal`, `private`, `abstract`, `sealed`, `static`, `in`.

### Rule categories

Rules are explicitly categorized by domain:

- `file`: File system checks (existence, content).
- `code`: Code structure checks (dependencies, types, naming).
- `custom`: Opaque blocks delegated entirely to a named plugin.

The parser validates that file assertions (e.g., `path ... must exist`) only appear in `file` blocks, and code assertions (e.g., `must depend on`, `must implement`) only appear in `code` blocks. Custom block bodies are not parsed by the host at all.

### Natural language syntax

The parser treats certain words as optional filler words for readability:

- `"on"` after `depend`: `must depend on` = `must depend`
- `"be"` after `must`: `must be public` = `must public`
- `"with"` after `annotated`: `must be annotated with "..."` = `must annotated "..."`
- `"by"` after `accessed`: `must only be accessed by` = `must only accessed`

## Assertions by rule type

### File system rules (`file`)

Use `file` blocks for file system checks:

| Rule | Syntax |
| ---- | ------ |
| Existence | `path "<pattern>" must exist` |
| Absence | `path "<pattern>" must not exist` |
| Content match | `path "<pattern>" must contain "<text\|regex:pattern>"` |
| Content exclusion | `path "<pattern>" must not contain "<text\|regex:pattern>"` |

```dsl
file "tests_exist" {
  path "tests/ArchTests/" must exist
  severity error
}

file "no_legacy_deps" {
  path "**/*.csproj" must not contain "regex:OldFramework"
  severity error
}
```

### Code structure rules (`code`)

Use `code` blocks for all code-level architectural checks.

#### Dependency rules

| Rule | Syntax |
| ---- | ------ |
| Forbidden | `<subject> must not depend on <targets>` |
| Allowed only | `<subject> must only depend on <targets>` |
| Incoming access | `<subject> must only be accessed by <targets>` |
| Acyclic | `<subject> must be acyclic` |

```dsl
code "domain_isolated" {
  Domain must not depend on Application, Infrastructure
  severity error
}

code "application_inward" {
  Application must only depend on Domain
  severity error
}

code "core_limited_access" {
  Domain must only be accessed by Application
  severity error
}

code "no_cycles" {
  component match "regex:Modules\\..*" must be acyclic
  severity error
}
```

#### Type relationship rules

| Rule | Syntax |
| ---- | ------ |
| Interface (require) | `<class> must implement interface <interface>` |
| Interface (forbid) | `<class> must not implement interface <interface>` |
| Inheritance (require) | `<class> must extend <class>` |
| Inheritance (forbid) | `<class> must not extend <class>` |

```dsl
code "repos_use_interface" {
  class match "regex:.*Repository$" must implement interface "IRepository"
  severity error
}

code "entities_extend_base" {
  class in Domain must extend class "BaseEntity"
  severity error
}
```

#### Naming and location rules

| Rule | Syntax |
| ---- | ------ |
| Naming pattern | `<subject> must match "regex:<pattern>"` |
| Anti-pattern | `<subject> must not match "regex:<pattern>"` |
| Location | `<class> must be in <component>` |

```dsl
code "interface_naming" {
  interface must match "regex:^I[A-Z].*"
  severity error
}

code "repos_in_infra" {
  class match "regex:.*Repository$" must be in Infrastructure
  severity error
}
```

#### Annotation rules

| Rule | Syntax |
| ---- | ------ |
| Required | `<subject> must be annotated with "<name>"` |
| Forbidden | `<subject> must not be annotated with "<name>"` |

```dsl
code "domain_marked" {
  class in Domain must be annotated with "DomainAttribute"
  exclude class implementing interface "IValueObject"
  severity error
}
```

#### Visibility rules

| Rule | Syntax |
| ---- | ------ |
| Public | `<class> must be public` |
| Internal | `<class> must be internal` |
| Private | `<class> must be private` |

```dsl
code "entities_public" {
  class in component "Domain.Entities" must be public
  severity error
}
```

#### Type constraint rules

| Rule | Syntax |
| ---- | ------ |
| Abstract | `<class> must be abstract` |
| Sealed | `<class> must be sealed` |
| Static | `<class> must be static` |

```dsl
code "base_classes_abstract" {
  class match "regex:^Base.*" must be abstract
  severity error
}
```

## Subject expressions

Six ways to specify rule subjects:

| Type | Syntax | Example |
| ---- | ------ | ------- |
| Named selector | `Name` | `Domain` |
| Inline literal | `<type> "<pattern>"` | `component "MyApp.Domain"` |
| Inline regex | `<type> match "regex:..."` | `class match "regex:.*Handler"` |
| Scoped all | `<type> in <target>` | `class in Domain` |
| Scoped literal | `<type> "<name>" in <target>` | `class "Handler" in Domain` |
| Scoped regex | `<type> match "..." in <target>` | `class match "regex:.*Handler" in Domain` |

The `in` keyword creates subset relations, filtering subjects to elements within the specified scope.

## Exclusions

Filter false positives within any rule:

```dsl
exclude class "<name>"
exclude class implementing interface "<name>"
exclude component "<pattern>"
exclude "<pattern>"     # for path rules
```

## Patterns

- Glob: path matching (e.g., `src/**/*.cs`).
- Regex: prefix with `regex:` (e.g., `regex:.*\.Domain\..*`).
- Namespace: double-dot suffix for sub-packages (e.g., `MyApp.Domain..`).

## Complete example

```dsl
adr "0001" "Implement Clean Architecture"

# Define layers
component "Domain"         = "MyApp.Domain"
component "Application"    = "MyApp.Application"
component "Infrastructure" = "MyApp.Infrastructure"

# Dependency rules
code "domain_isolated" {
  Domain must not depend on Application, Infrastructure
  severity error
}

code "application_inward" {
  Application must only depend on Domain
  severity error
}

code "no_cycles" {
  component match "regex:MyApp\\..*" must be acyclic
  severity error
}

# Type rules
code "repos_implement_interface" {
  class match "regex:.*Repository$" must implement interface "IRepository"
  exclude class "LegacyRepository"
  severity error
}

code "interfaces_prefixed" {
  interface must match "regex:^I[A-Z].*"
  severity error
}

# File rules
file "tests_exist" {
  path "tests/ArchTests/" must exist
  severity error
}

file "no_legacy_frameworks" {
  path "**/*.csproj" must not contain "regex:OldFramework|LegacyLib"
  severity error
}
```

## Custom blocks

A `custom` block lets plugin authors define entirely new assertion vocabularies without modifying the grammar.

```dsl
custom "my_check" {
  any text the plugin understands
  can go here with whatever syntax
  the plugin author defines
}
```

The host parses and validates the block header (`custom` and name) but treats the body as raw, opaque text. The body is stored in the `raw_body` field of the `RuleIR` protobuf message and forwarded to the plugin unchanged. The plugin receives `is_custom_rule = true` on the rule and reads `raw_body` to interpret the body.

The body may contain any characters, including nested `{}` braces (brace depth is tracked). The `custom` keyword and the block's closing `}` are the only reserved tokens the host checks.

## Validation

The parser rejects files with:

- Missing or duplicate ADR header.
- Duplicate selector or rule names (including across `file`, `code`, and `custom` blocks).
- Keyword collisions in names.
- Undefined references.
- Missing required arguments.
- Invalid severity values.
- File assertions in `code` blocks (e.g., `path ... must exist` in code).
- Code assertions in `file` blocks (e.g., `must depend on` in file).
- `custom` block with an empty name.
- Unterminated `custom` block (no matching closing `}`).
