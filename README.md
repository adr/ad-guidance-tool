# ADG
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

> For convenience, feel free to rename the downloaded file to adg (or adg.exe on Windows) so you can run it directly from the terminal.

### Building from source

To build ADG yourself, ensure that [Go](https://go.dev/dl/) is installed on your system. Then run:

```bash
git clone https://github.com/adr/ad-guidance-tool.git
cd ad-guidance-tool
go build -o adg ./main.go
```

> On Windows, be sure to name the output binary with a `.exe` extension (e.g., `adg.exe`) so the terminal can execute it properly.

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
  help         Help about any command
  import       Imports a decision model into an existing model
  init         Initializes a new model
  link         Link two decisions using optional custom tags or default precedes/succeeds logic
  list         Lists decisions in the model, optionally filtering by tag, status, title, or ID
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

To enhance your workflow, ADG supports shell auto-completion. Generate a script with:

```bash
adg completion [shell]
```

For example, to enable auto-completion in PowerShell:

```bash
adg completion powershell
```

Copy the output into your [PowerShell profile](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_profiles?view=powershell-7.5) to enable completions. Follow a similar process for other shells (available: `bash`, `fish`, `powershell`, `zsh`).

### Examples

In the `models/clean` folder, you can see an example of a model created with ADG. This model contains a set of recurring architectural decisions for using Clean Architecture.

## Contributing

If you have a feature request or found a bug, you can [open an issue](https://github.com/adr/ad-guidance-tool/issues) to share your feedback.

Contributions are also welcome. Please submit a [pull request](https://github.com/adr/ad-guidance-tool/pulls) with your changes.

We follow [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/The-Clean-Architecture.html) to organize our codebase. If you're adding a feature, we recommend to:

1. Start with the use case (interactor) of your feature
2. Add any necessary core logic in the domain layer
3. Implement the [Cobra CLI command](https://github.com/spf13/cobra) for input and a *presenter/printer* for output
4. Write unit tests (refer to existing tests for guidance). To simplify mocking, we use [mockery](https://github.com/vektra/mockery), though hand-written mocks are also possible.

## License

ADG is released under the [Apache License, Version 2.0.](https://www.apache.org/licenses/LICENSE-2.0)