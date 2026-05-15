# ADE Syntax Highlighting

This extension provides syntax highlighting for Architecture Decision Rule (`.rule`) files.

## Features

- Syntax highlighting across four distinct token groups (see [Color Scheme](#color-scheme) below)
- IntelliSense completion: type keywords and press Ctrl+Space for suggestions
- Code snippets for common patterns:
  - `adr`: ADR header
  - `code`, `file`, `custom`: rule block templates
  - `component`, `class`, `interface`, `path`: selectors
  - `must not depend`, `must only depend`: dependency rule templates
  - `path must exist`, `path must contain`: filesystem check templates
  - `must implement`, `must extend`: type relationship templates
  - `component match`, `class match`: pattern matching
  - `exclude class implementing`, `exclude class`: exclusions
- String highlighting with escape sequence support
- Comment support (`#`)
- Auto-closing pairs for quotes and brackets
- Bracket matching

## File Association

This extension automatically associates `.rule` files with the ADE language.

## Color Scheme

The extension uses four distinct token groups, each mapped to a different standard TextMate scope so they render in a distinct color in every VS Code theme.

- Types (`storage.type`): `adr`, the rule category keywords `file`, `code`, and `custom`, and the selector type keywords `component`, `class`, `interface`, and `path`.
- Subjects (`entity.name.type`): named selector references reused inside rule blocks, for example `Domain` or `Application` in `Domain must not depend on Application`. Subjects always start with an uppercase letter.
- Descriptors (`keyword.control`): keywords that define how a rule is evaluated, including `must`, `not`, `only`, `depend`, `exist`, `contain`, `implement`, `extend`, `accessed`, `acyclic`, `annotated`, `implementing`, `match`, `exclude`, `in`, `severity`, and the optional filler words `on`, `be`, `by`, and `with`.
- Properties (`constant.numeric`): values that specify what is checked or the outcome, including `public`, `internal`, `private`, `abstract`, `sealed`, `static`, `error`, and `warning`.

Strings use the theme's string color, and comments use the theme's comment color.

## Examples

```ade-syntax
adr "0010" "Use Clean Architecture"

component "Domain" = "MyApp.Domain.."
component "Application" = "MyApp.Application.."

code "domain_is_independent" {
  Domain must not depend on Application
  severity error
}

file "config_files_exist" {
  path "**/*.config" must exist
  severity warning
}

code "domain_classes_annotated" {
  class in Domain must be annotated with "DomainAttribute"
  severity error
}

code "handlers_annotated" {
  class "EventHandler" must be annotated with "SerializableAttribute"
  severity error
}

code "specific_class_in_component" {
  class "Handler" in Domain must not depend on component "Infrastructure"
  severity error
}

custom "my_check" {
  any text the plugin understands
  can go here with whatever syntax
  the plugin author defines
}
```
