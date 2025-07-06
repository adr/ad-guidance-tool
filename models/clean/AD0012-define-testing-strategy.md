---
adr_id: "0012"
title: define-testing-strategy
status: open
tags: 
    - "testing"
    - "unit-test"
    - "integration-test"
    - "architecture-layers"
links:
    precedes:
        - "0013"
    succeeds:
        - "0011"
comments: []
---

## <a name="question"></a> Question

What testing strategy should be used for the different architectural layers to ensure correctness while maintaining separation of concerns?

## <a name="options"></a> Options

1. <a name="option-1"></a> Write isolated unit tests per layer (e.g., use cases, entities) with integration tests only at adapter boundaries.
2. <a name="option-2"></a> Focus on full-stack integration tests covering end-to-end behavior, with minimal unit testing.
3. <a name="option-3"></a> Apply a hybrid strategy combining unit tests for critical logic and integration tests for workflows.

## <a name="criteria"></a> Criteria

- Test execution speed and feedback cycle
- Test coverage of business-critical paths
- Isolation of architectural concerns during testing
- Maintenance effort and fragility of test suites
- Tooling and framework support for mocking and assertions
