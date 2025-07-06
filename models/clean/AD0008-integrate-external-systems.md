---
adr_id: "0008"
title: integrate-external-systems
status: open
tags:
    - integration
    - external-system
    - interfaces
    - clean-architecture
links:
    precedes: []
    succeeds: []
comments: []
---

## <a name="question"></a> Question

How should an external system be integrated in a way that maintains the independence of core business logic and upholds Clean Architecture principles?

## <a name="options"></a> Options

1. <a name="option-1"></a> Use an interface-based adapter pattern to wrap the external system and inject it via dependency inversion.
2. <a name="option-2"></a> Introduce an intermediary service layer that mediates between core logic and the external system.
3. <a name="option-3"></a> Allow direct calls to the external system from the use case layer, with minimal abstraction.

## <a name="criteria"></a> Criteria

- Degree of decoupling between core logic and the external system
- Complexity of integration and future maintainability
- Adherence to dependency inversion and interface segregation
- Testability of components dependent on external system behavior
