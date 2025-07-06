---
adr_id: "0014"
title: structure-usecase-modules
status: open
tags: 
    - "modularity"
    - "use-cases"
    - "feature-modules"
    - "organization"
links:
    precedes:
        - "0015"
    succeeds:
        - "0013"
comments: []
---

## <a name="question"></a> Question

How should feature modules be structured within the use case layer to maintain clarity, modularity, and scalability of the application?

## <a name="options"></a> Options

1. <a name="option-1"></a> Organize use cases by feature verticals (e.g., `user/register`, `invoice/generate`) with dedicated subdirectories per feature.
2. <a name="option-2"></a> Group use cases by technical function (e.g., all use cases, all interfaces) across the application.
3. <a name="option-3"></a> Follow domain-driven design aggregates or bounded contexts to group related business capabilities.

## <a name="criteria"></a> Criteria

- Discoverability and navigation of use case logic
- Encapsulation and independence of features or domains
- Reusability and testability of logic across modules
- Scalability as the application grows
- Alignment with business concepts and responsibilities
