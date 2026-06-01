# ADG

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](./LICENSE)

ADG (Architectural Decision Guidance) is a command-line tool written in Go for modeling, managing, and reusing architectural decisions in a lightweight and structured way.

An architectural decision is a justified design choice addressing a functional or non-functional requirement that is architecturally significant. These decisions can be captured using Architectural Decision Records (ADRs). ADG allows you to create and edit ADRs, group them into *models*, and manage those models. A model can be created, copied, imported, or merged, providing guidance for recurring decisions.

## Getting started

To start using ADG, you can either download the [latest release](https://github.com/adr/ad-guidance-tool/releases) or build it from source.

### Downloading a release

Precompiled executables for major operating systems are available:
- Windows: `adg_win.exe`
- Linux: `adg_linux`
- macOS (Intel): `adg_mac_intel`
- macOS (Apple Silicon): `adg_mac_arm`

> For convenience, feel free to remove the suffix (e.g., `_win`) after you have downloaded the file.

### Installing via Go

If you have [Go](https://go.dev/dl/) installed, you can install ADG directly using:

```bash
go install github.com/adr/ad-guidance-tool/adg@latest
```

This will download, build, and install the `adg` binary to your `$GOPATH/bin` directory (typically `~/go/bin` on Linux/macOS or `%USERPROFILE%\go\bin` on Windows). Make sure this directory is in your system's PATH.

### Building from source

To build ADG yourself, ensure that [Go](https://go.dev/dl/) is installed on your system. Then run:

```bash
git clone https://github.com/adr/ad-guidance-tool.git
cd ad-guidance-tool
go build -o adg ./adg
```
This will generate a binary in your current directory called `adg` (or `adg.exe` on Windows).

### Running the tool

Executing the binary displays the CLI help:

```
CLI tool for managing architectural decision records and models

Usage:
  adg [command]

Available Commands:
  add          Adds one or more decision points to a model
  comment      Add a comment to a decision
  copy         Copies a model, optionally a subset based on filters
  decide       Marks a decision as decided by selecting one of its options
  edit         Edit a decision file
  enforce      Enforce architectural decisions using rule files.
  help         Help about any command
  import       Imports a decision model into an existing model
  init         Initializes a new model
  link         Link two decisions using optional custom tags or default precedes/succeeds logic
  list         Lists decisions in the model, optionally filtering by tag, status, title, or ID
  mcp          MCP server setup for AI tool integration
  merge        Merges two decision models into a new target model
  rebuild      Rebuilds the index file for the given model
  reset-config Reset all configuration (or only template headers with --template)
  revise       Creates a copy of a decision and resets its status to 'open' (if not already)
  set-config   Set persistent configuration values
  tag          Categorizes a decision by adding one or more tags to its metadata
  validate     Validate the models decisions by checking if the files match the index file
  view         Show the full or partial content of one or more decision files

Flags:
  -h, --help   help for adg

Use "adg [command] --help" for more information about a command.
```

### Shell auto-completion

ADG supports shell auto-completion for `bash`, `fish`, `powershell`, and `zsh`. Run the completion command with `--help` to see the installation instructions for your shell:

```bash
adg completion powershell --help
```

For example, to enable auto-completion in PowerShell, add the following line to your [PowerShell profile](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_profiles?view=powershell-7.5):

```powershell
adg completion powershell | Out-String | Invoke-Expression
```

Other shells follow the same pattern: run `adg completion <shell> --help` for the shell-specific instructions.

## User Guide

### Creating a new model

To create a new decision model, use the `init` command:

```bash
adg init <model-name>
```

This creates a new directory (in your current working directory, unless an absolute or relative path is provided) containing an `index` file. This index tracks metadata for all decisions in the model and is continuously updated as decisions change.

### Adding and editing a decision

To add a new decision to the model:

```bash
adg add --model <model-name> --title <decision-title>
```

Or simply provide the title as arguments without quotes:

```bash
adg add --model <model-name> <decision-title-words...>
```

For example:

```bash
adg add --model my-model Favor cloud functions over az logic apps
```

This generates a new Markdown file inside the model directory. Each new file includes a metadata block followed by three sections: *Question*, *Options*, *Criteria*.

> **Tip:** When using positional arguments, the title is automatically assembled from all arguments after the command and flags. The `--title` flag is still available for backwards compatibility or when you need to add multiple decisions at once.

You can edit these sections manually in a text editor, or using the `edit` command of the tool:

```bash
adg edit --model <model-name> --id <decision-id | decision-title> [--question "<section-content>"] [--option "<option-name>"] [--criteria "<section-content>"]
```

> The `--option` flag is repeatable for adding multiple options. Each option is automatically given an anchor tag so it can be referenced.

If you're editing manually, ensure the structure matches the following format:
```markdown
---
adr_id: "0001"
title: your-title
status: open
tags: []
links: []
comments: []
---

## <a name="question"></a> Question

<!-- section content -->

## <a name="options"></a> Options

1. <a name="option-1"></a> Option 1
2. <a name="option-2"></a> Option 2
3. <a name="option-3"></a> Option 3
<!-- and so on -->

## <a name="criteria"></a> Criteria

<!-- section content -->
```

You may change the displayed section header and/or include additional sections, but the tool expects at least the three sections and their anchor tags mentioned above to function properly.

For example:

```markdown
---
adr_id: "0001"
title: your-title
status: open
tags: []
links: []
comments: []
---

## <a name="question"></a> Context and Problem Statement
<!-- section content -->

## <a name="options"></a> Considered Options
1. <a name="option-1"></a> This is my first considered option
2. <a name="option-2"></a> This is my second considered option
<!-- and so on -->

## <a name="criteria"></a> Decision Drivers
<!-- section content -->

## Pros and Cons of the Options
<!-- section content -->

```

### Deciding on an option

To mark a decision as decided:

```bash
adg decide --model <model-name> --id <decision-id | decision-title> --option <option-number | option-title> [--rationale "your-rationale"] 
```

This will add a new section **Outcome** pointing out the chosen option and a rationale if provided to the command.

### Generating rule files for ADRs

ADG can generate `.rule` files based on your architectural decisions. These rule files encode architectural rules in a domain-specific language that can be compiled into architecture tests or verified directly using `adg enforce`.

To generate a rule file for an existing decision:

```bash
adg enforce rule --model <model-name> --id <decision-id>
```

Or using the decision title:

```bash
adg enforce rule --model <model-name> --title <decision-title>
```

By default, the rule file is created in the same directory as the ADR with the same filename but a `.rule` extension (e.g., `AD0001-my-decision.rule`).

You can specify a custom output path:

```bash
adg enforce rule --model <model-name> --id <decision-id> --output path/to/custom.rule
```

Or a custom output directory (the filename will be based on the ADR filename):

```bash
adg enforce rule --model <model-name> --id <decision-id> --output path/to/directory/
```

The generated rule file contains a template with the ADR ID and title pre-filled:

```dsl
adr "0001" "my-decision-title"

# component "MyComponent" = "com.example.mypackage"
# path "MyPath" = "src/mypackage"

code "rule_name" {
  # MyComponent must not depend on MyOtherComponent
  severity error
}

file "rule_name" {
  # path "**/*.go" must exist
  severity error
}
```

You can then customize this file to define specific architectural rules based on your decision. See the [Enforcement](#enforcement) section for how to compile or verify rule files.

### Config

You can customize ADG's behavior using:

```bash
adg set-config [flags]
```
> Run with `-h` to see available configuration flags.

By default configuration is stored in a file called `.adgconfig.yaml` located in your home directory. You can specify a custom path using the `--config-path` flag.

To reset all configuration values (and use the default configuration path again):

```bash
adg reset-config
```

## Enforcement

ADG can enforce architectural decisions via the `adg enforce` command, which is implemented using [ADE (Architectural Decision Enforcement)](https://github.com/phi42/ad-enforcement-tool).

Rules are written in a domain-specific language (DSL) and stored in `.rule` files alongside your ADRs. Use `adg enforce rule` to generate a template from an existing decision, then customize it.

The `adg enforce` command provides:
| Subcommand             | Description                                                      |
| ---------------------- | ---------------------------------------------------------------- |
| `validate` | check rule file syntax                                           |
| `compile`  | compile rules into architecture tests (Go, .NET, …) via a plugin |
| `verify`   | verify rules directly against the filesystem via a plugin        |
| `plugin`   | install, update, and manage enforcement plugins                  |
| `config`   | manage configuration defaults                                    |

See the [ADE repository](https://github.com/phi42/ad-enforcement-tool/tree/main/docs) for more documentation.

### AI Integration (MCP)

ADG includes an MCP (Model Context Protocol) server that lets AI assistants read your ADRs, access the DSL reference, and validate generated rule files. This enables workflows where the AI generates a `.rule` file from an ADR and immediately validates it before presenting the result.

Run the following command to get a config snippet for your model:

```bash
adg mcp --model <model-name>
```

This prints the configuration for VS Code Copilot Chat. Other AI tools that support MCP can be configured similarly by adding a server with the same command and arguments.

#### VS Code

Add the snippet to `.vscode/mcp.json` in your workspace, or run `MCP: Add Server` in the command palette and select `Command (stdio)` to configure it interactively.

```json
{
	"servers": {
		"adg": {
			"command": "./repos/ad-guidance-tool/adg.exe",
			"args": ["mcp", "run", "--model", "<path-to-your-model>"]
		}
	}
}
```

Once configured, open Copilot Chat in Agent mode and ask it to generate a rule file for an ADR, for example:

```
@adg Generate a rule file for ADR 0001.
```

### Available tools

The MCP server exposes the following tools to the AI:

- `list_adrs`: list all decisions in the model with their ID, title, and status
- `get_adr`: retrieve the full content of a specific ADR and the path where its rule file should be placed
- `get_dsl_reference`: retrieve the full ADE rule DSL language reference
- `list_rule_files`: list all existing `.rule` files in the model directory as examples
- `validate_rule`: validate the syntax and semantics of rule content using the ADE parser

## Example Model

In the [models/clean](/models/clean/) directory, you’ll find a sample model containing common architectural decisions based on **Clean Architecture**.

You can use this model as a starting point to get familiar with the tool by copying it:

```bash
adg copy --model models/clean --target <new-model-name>
```

For more commands see the help text output:
```bash
adg -h            # for a general overview
adg <command> -h  # for a specific command
```

## Contributing

If you have a feature request or found a bug, you can [open an issue](https://github.com/adr/ad-guidance-tool/issues) to share your feedback.

Contributions are also welcome. Please submit a [pull request](https://github.com/adr/ad-guidance-tool/pulls) with your changes.

We follow [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) to organize our codebase. If you're adding a feature, we recommend to:
1. Start with the use case (interactor) of your feature
2. Add any necessary core logic in the domain layer
3. Implement the [Cobra CLI command](https://github.com/spf13/cobra) for input and a *presenter/printer* for output
4. Write unit tests (refer to existing tests for guidance). To simplify mocking, we use [mockery](https://github.com/vektra/mockery), though hand-written mocks are also possible.

## References

ADG was developed as part of two theses at the [Eastern Switzerland University of Applied Sciences](https://www.ost.ch/en/)
- [Concept Alternatives for the Management of Architectural Decisions in Clean Architectures](https://eprints.ost.ch/id/eprint/1280/1/MSECS-FS24-CleanArchitectureDecisionsConceptsRS.pdf)
- [A Command-Line Tool for Managing Recurring Architectural Decisions: Design, Implementation, and Empirical Evaluation](https://eprints.ost.ch/id/eprint/1287/1/PA2-Raphael-Schellander.pdf)
