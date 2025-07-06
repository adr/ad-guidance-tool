---
adr_id: "0010"
title: define-error-handling-strategy
status: open
tags:
    - error-handling
    - exceptions
    - contracts
    - clean-architecture
links:
    precedes:
        - "0011"
    succeeds:
        - "0009"
comments: []
---

## <a name="question"></a> Question

What strategy should be used for handling and propagating errors across architectural layers in a way that preserves separation of concerns and keeps business logic independent of frameworks?

## <a name="options"></a> Options

1. <a name="option-1"></a> Use return types (e.g., error objects or `Result` types) at every boundary, avoiding exceptions or panic flows.
2. <a name="option-2"></a> Allow exceptions or panics in outer layers, but catch and translate them into controlled forms before reaching core layers.
3. <a name="option-3"></a> Propagate exceptions across all layers with a global handler that maps them to user-facing responses.

## <a name="criteria"></a> Criteria

- Transparency and consistency of error behavior
- Coupling introduced between layers through exception types
- Ease of testing and mocking error paths
- Ability to localize and translate error messages for clients
- Support for logging, auditing, and observability
