# Plugin Development

How to write a new enforcement plugin, build from source, and regenerate the parser and protobuf code.

## Prerequisites

- [Go](https://go.dev).
- For grammar changes: [ANTLR](https://www.antlr.org/) 4.13.2 or later and a Java runtime.
- For protobuf changes: `protoc` and `protoc-gen-go` (or another language's protobuf compiler if not writing a Go plugin).

## Writing a plugin

Any executable that speaks the ADG plugin protocol qualifies as a plugin. Plugins are decoupled from the main tool and can live in their own repositories.

### Protocol

1. `adg` builds a `SpecIR` protobuf message (defined in [`proto/rule.proto`](../proto/rule.proto)) describing the rule file and invocation context.
2. It spawns the plugin as a child process and writes the serialized `SpecIR` to the plugin's `stdin`.
3. The plugin:
   - For `adg enforce compile`: writes one or more generated files to the directory named in `SpecIR.OutputDir`.
   - For `adg enforce verify`: performs checks immediately and reports results to `stdout` using the standard `LEVEL  [rule] message` format.
4. A non-zero exit code indicates failure.

A reference copy of the generated Go types is available at [`proto/generated/go/rule.pb.go`](../proto/generated/go/rule.pb.go). Copy this file into a `rule/` subdirectory of your plugin.

### Info flag

When invoked with `--info`, the plugin must print a JSON object to `stdout` and exit zero:

```json
{"modes": ["compile"]}
```

or

```json
{"modes": ["verify"]}
```

ADG calls `--info` before each invocation to verify that the plugin supports the requested mode.

### Custom rule blocks

A `custom` rule block lets plugin authors define entirely new assertions without modifying the grammar. The host stores the raw body text in the `raw_body` field of `RuleIR`, which is forwarded to the plugin unchanged:

```dsl
custom "my_check" {
  any text the plugin understands
  can go here with whatever syntax
  the plugin author defines
}
```

The `is_custom_rule` boolean on `RuleIR` marks these entries. Custom rules are forwarded to the plugin for both `adg enforce compile` and `adg enforce verify`.

### Publishing via GitHub release

To make a plugin installable with `adg enforce plugin install <name> --repo github.com/<owner>/<repo>`, the GitHub release must contain assets whose filenames include the target OS and architecture:

```
<repo>-<goos>-<goarch>          # Unix
<repo>-<goos>-<goarch>.exe      # Windows
```

Example assets for a repository named `my-plugin`:

```
my-plugin-linux-amd64
my-plugin-linux-arm64
my-plugin-darwin-amd64
my-plugin-darwin-arm64
my-plugin-windows-amd64.exe
```

`<goos>` and `<goarch>` must match Go's `runtime.GOOS` and `runtime.GOARCH` strings.

## Logging

All output from `adg` and its plugins is routed through a structured logger that produces consistently formatted lines:

```
LEVEL  message
```

The level label is always six characters wide (e.g., `INFO  `, `WARN  `, `ERROR `) so that lines align regardless of level.

### Log levels

| Level   | When it appears      | Typical content                                           |
| ------- | -------------------- | --------------------------------------------------------- |
| `DEBUG` | Only with `--debug`. | Internal progress steps, file paths, plugin lifecycle.    |
| `INFO`  | Default and above.   | Successful results, generated file names, passing checks. |
| `WARN`  | Default and above.   | Skipped rules, non-fatal notices.                         |
| `ERROR` | Always.              | Fatal errors that stop execution.                         |

### Flags

Three persistent flags control the log level and apply to all subcommands:

| Flag            | Effect                                                                                                               |
| --------------- | -------------------------------------------------------------------------------------------------------------------- |
| `--debug`       | Shows `DEBUG`, `INFO`, `WARN`, and `ERROR`. Use for diagnosing unexpected behaviour.                                 |
| `--no-warnings` | Shows `INFO` and `ERROR` only. Suppresses `WARN` lines. Useful in pipelines where skipped-rule notices are expected. |
| `--quiet`       | Shows `ERROR` only. Produces no output on a fully successful run. Useful in scripts where only failures matter.      |

When none of these flags are set, the default level shows `INFO` and above (including `WARN`). If `--debug` and `--quiet` are both set, `--debug` takes precedence.

### How plugins inherit the log level

Plugins run as separate child processes. `adg` propagates the chosen level to each plugin via the `ADE_LOG_LEVEL` environment variable before spawning the process:

| `ADE_LOG_LEVEL` value | Meaning                              |
| --------------------- | ------------------------------------ |
| *(unset)*             | `INFO` level, default behaviour.     |
| `debug`               | `DEBUG` level.                       |
| `no-warnings`         | `INFO` level with `WARN` suppressed. |
| `quiet`               | `ERROR` level.                       |

A plugin that reads `os.Getenv("ADE_LOG_LEVEL")` at startup and configures its logger accordingly will behave consistently whether `adg` invokes it or a test harness calls it directly.

## Regenerating the parser

Run after modifying [`internal/ade/parser/ADE.g4`](../internal/ade/parser/ADE.g4):

```sh
java -jar antlr-4.13.2-complete.jar -Dlanguage=Go -visitor -no-listener -o internal/ade/parser internal/ade/parser/ADE.g4
```

The generated files are committed, so regeneration is only needed when the grammar changes.

## Regenerating protobuf code

Run after modifying [`proto/rule.proto`](../proto/rule.proto):

```sh
protoc --go_out=./internal/ade/rule --go_opt=paths=source_relative proto/rule.proto
```

Then copy the result to [`proto/generated/go/rule.pb.go`](../proto/generated/go/rule.pb.go) for plugin developers to reference.

The generated files are committed, so regeneration is only needed when the schema changes.
