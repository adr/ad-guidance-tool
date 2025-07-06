---
adr_id: "0015"
title: define-io-boundary-strategy
status: open
tags:
    - boundaries
    - io
    - interface-adapter
    - clean-architecture
links:
    precedes: []
    succeeds:
        - "0004"
        - "0009"
        - "0010"
comments: []
---

## <a name="question"></a> Question

How should the boundaries between the application core and the outside world (e.g., HTTP, CLI, gRPC, messaging systems) be structured to maintain layer separation and support adaptability?

## <a name="options"></a> Options

1. <a name="option-1"></a> Define explicit input and output interfaces in the interface adapter layer that map external requests/responses to internal models.
2. <a name="option-2"></a> Embed parsing, formatting, and transport logic directly into use case handlers (e.g., use case accepts HTTP request types).
3. <a name="option-3"></a> Use a middleware pipeline that transforms all I/O at the application boundary before passing to the use case.

## <a name="criteria"></a> Criteria

- Separation of concerns between layers
- Reusability of core business logic across channels
- Testability of input/output handling
- Flexibility to support multiple protocols or transport formats (e.g., JSON, Protobuf)
- Simplicity and maintainability of adapters
