---
adr_id: "0001"
comments:
    - author: ""
      comment: "1"
      date: "2026-05-15 18:40:35"
status: decided
title: Domain layer is independent from upper layers
---

## <a name="question"></a> Question

How should the core domain model relate to the layers above it?

## <a name="options"></a> Options

1. <a name="option-1"></a> Keep domain independent: domain packages may not import application, adapter, or infrastructure packages
2. <a name="option-2"></a> Allow upward imports: domain packages may freely import any layer

## <a name="criteria"></a> Criteria

Domain logic must be reusable independently of any delivery mechanism or persistence technology. Coupling domain to upper layers makes it impossible to test or reuse in isolation.

## <a name="outcome"></a> Outcome
We decided for [Option 1](#option-1) because: The domain layer encodes business rules that must remain stable regardless of how the tool is delivered or how data is stored. Keeping domain independent ensures that use-case logic and infrastructure can change without affecting the core model.

## <a name="comments"></a> Comments
<a name="comment-1"></a>1. (2026-05-15 18:40:35) : marked decision as decided
