---
adr_id: "0009"
title: define-usecase-interfaces
status: open
tags:
    - interfaces
    - dependency-inversion
    - use-cases
    - clean-architecture
links:
    precedes:
        - "0010"
    succeeds:
        - "0003"
comments: []
---

## <a name="question"></a> Question

How should the interface contracts between the use case layer and the outer layers (e.g., interface adapters, frameworks) be defined to support dependency inversion and layer independence?

## <a name="options"></a> Options

1. <a name="option-1"></a> Let the use case layer define abstract interfaces that are implemented by external layers (e.g., `UserRepository` interface).
2. <a name="option-2"></a> Define interfaces in a shared contract layer decoupled from both use case and infrastructure.
3. <a name="option-3"></a> Allow outer layers to define concrete APIs and let use cases depend directly on them (inversion not enforced).

## <a name="criteria"></a> Criteria

- Clarity of architectural boundaries and ownership
- Degree of adherence to dependency inversion principle
- Ease of mocking or substituting components in tests
- Maintainability and discoverability of contracts across modules
