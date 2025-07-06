---
adr_id: "0003"
title: select-use-cases
status: open
tags:
    - use-case
    - application-logic
    - workflow
    - clean-architecture
links:
    precedes:
        - "0004"
    succeeds:
        - "0002"
comments: []
---

## <a name="question"></a> Question

How should the use case for handling a specific workflow be structured to encapsulate the application-specific business logic while remaining independent of external systems?

## <a name="options"></a> Options

1. <a name="option-1"></a> Define a dedicated use case class or module that orchestrates the workflow and contains all relevant business logic.
2. <a name="option-2"></a> Implement the workflow as part of a service layer shared across use cases.
3. <a name="option-3"></a> Integrate the workflow logic directly within a controller or adapter component to reduce indirection.

## <a name="criteria"></a> Criteria

- Separation of application logic from infrastructure and delivery concerns
- Maintainability and clarity of workflow responsibilities
- Ease of testing and mocking in isolation
- Reusability of the use case logic across interfaces or delivery mechanisms
