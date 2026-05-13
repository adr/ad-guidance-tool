# Adopt structured logging

## Context and Problem Statement

Free-form logs are hard to grep.

## Considered Options

* Keep free-form
* Adopt structured logging (zap)

## Decision Outcome

Chosen option: "Adopt structured logging (zap)", because grep-friendly logs cost less in incidents.

### Consequences

* Good, because incident MTTR drops.
* Bad, because every site needs to be migrated.

### Confirmation

Lint rule rejects `fmt.Println` in production code.
