---
adr_id: "0011"
title: manage-cross-cutting-concerns
status: open
tags: 
    - "cross-cutting"
    - "logging"
    - "monitoring"
    - "auth"
    - "clean-architecture"
links:
    precedes:
        - "0012"
    succeeds:
        - "0010"
comments: []
---

## <a name="question"></a> Question

How should cross-cutting concerns such as logging, monitoring, or authentication be handled across layers while preserving the architectural integrity of the system?

## <a name="options"></a> Options

1. <a name="option-1"></a> Use middleware or decorator patterns in the outer layers (e.g., interface adapters) to inject cross-cutting logic.
2. <a name="option-2"></a> Apply aspect-oriented techniques (e.g., interceptors, annotations) to weave in concerns dynamically.
3. <a name="option-3"></a> Embed cross-cutting logic directly into use cases and infrastructure code where needed.

## <a name="criteria"></a> Criteria

- Degree of decoupling and modularity
- Ease of applying concerns uniformly across layers
- Testability of use cases without concern-side effects
- Visibility and traceability of system behavior (e.g., logging context)
- Performance and overhead introduced by concern injection
