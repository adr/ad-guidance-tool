---
status: "proposed"
date: 2026-05-13
---

# Migrate from MySQL to PostgreSQL

## Context and Problem Statement

We need a database with stronger JSONB support.

## Considered Options

* Stay on MySQL with workarounds
* Migrate to PostgreSQL

## Decision Outcome

Chosen option: "Migrate to PostgreSQL", because JSONB support is first-class.

## Risks

* Migration tooling immaturity
* 6-month dual-write window

## Open Questions

* Do we colocate read replicas?
