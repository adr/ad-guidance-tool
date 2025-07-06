---
adr_id: "0002"
title: decide-entity-boundaries
status: open
tags:
    - entities
    - domain-model
    - business-rules
    - clean-architecture
links:
    precedes:
        - "0003"
    succeeds:
        - "0001"
comments: []
---

## <a name="question"></a> Question

How should the boundaries of core entities be defined to encapsulate business rules in a stable and reusable way?

## <a name="options"></a> Options

1. <a name="option-1"></a> Define dedicated domain entities representing key business concepts, independent of frameworks or infrastructure.
2. <a name="option-2"></a> Use simple data structures or records (e.g., DTOs or structs) and keep business logic external.
3. <a name="option-3"></a> Postpone entity modeling until implementation details are clearer (emergent modeling).

## <a name="criteria"></a> Criteria

- Stability and reusability across different applications or bounded contexts
- Degree of encapsulation and abstraction from external systems
- Clarity and alignment with core business rules
- Development effort and complexity of maintaining separation
