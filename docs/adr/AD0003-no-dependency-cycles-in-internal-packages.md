---
adr_id: "0003"
comments:
    - author: ""
      comment: "1"
      date: "2026-05-15 18:40:51"
status: decided
title: No dependency cycles in internal packages
---

## <a name="question"></a> Question

How should the internal packages relate to each other to avoid circular dependencies?

## <a name="options"></a> Options

1. <a name="option-1"></a> Enforce acyclic constraint: no import cycle may exist among packages under internal/
2. <a name="option-2"></a> Allow cycles: packages under internal/ may form import cycles

## <a name="criteria"></a> Criteria

Import cycles in Go cause compile errors and indicate confused responsibilities. Maintaining an acyclic dependency graph keeps the package structure navigable and prevents accidental coupling.

## <a name="outcome"></a> Outcome
We decided for [Option 1](#option-1) because: An acyclic dependency graph is a prerequisite for maintainable layered design. Go enforces this at compile time, but making it an explicit architectural decision documents the intent and allows automated verification that no cycles are introduced.

## <a name="comments"></a> Comments
<a name="comment-1"></a>1. (2026-05-15 18:40:51) : marked decision as decided
