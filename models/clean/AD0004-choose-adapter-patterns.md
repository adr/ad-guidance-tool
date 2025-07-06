---
adr_id: "0004"
title: choose-adapter-patterns
status: open
tags:
    - pattern
    - translation
    - interface-adapter
    - clean-architecture
links:
    precedes:
        - "0005"
    succeeds:
        - "0003"
comments: []
---

## <a name="question"></a> Question

Which pattern should be used to handle data translation between a use case and an external system, in order to avoid tight coupling and maintain separation of concerns?

## <a name="options"></a> Options

1. <a name="option-1"></a> Use the Data Mapper pattern to convert between internal domain models and external representations.
2. <a name="option-2"></a> Apply the Adapter (or Anti-Corruption Layer) pattern to fully isolate external models behind interface boundaries.
3. <a name="option-3"></a> Embed translation logic directly within the use case or controller to reduce indirection.

## <a name="criteria"></a> Criteria

- Degree of decoupling from external systems
- Maintainability and testability of translation logic
- Complexity introduced by the pattern
- Reusability of transformation logic across use cases
