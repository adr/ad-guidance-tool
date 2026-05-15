---
adr_id: "0002"
comments:
    - author: ""
      comment: "1"
      date: "2026-05-15 18:40:43"
status: decided
title: Application layer does not depend on adapters or infrastructure
---

## <a name="question"></a> Question

Which packages may the application use-case layer depend on?

## <a name="options"></a> Options

1. <a name="option-1"></a> Depend only on domain: application packages may import domain but not adapter or infrastructure
2. <a name="option-2"></a> Allow any dependency: application packages may freely import adapter and infrastructure

## <a name="criteria"></a> Criteria

Use-case logic must remain independent of concrete delivery mechanisms (CLI, HTTP) and persistence implementations. This allows use cases to be unit-tested without the real infrastructure and adapted to new delivery channels without changing business logic.

## <a name="outcome"></a> Outcome
We decided for [Option 1](#option-1) because: Following Clean Architecture, use cases sit between the domain and the delivery/persistence layers. They orchestrate domain logic and define output ports without knowing how those ports are implemented.

## <a name="comments"></a> Comments
<a name="comment-1"></a>1. (2026-05-15 18:40:43) : marked decision as decided
