---
status: accepted
date: "2026-05-13"
decision-makers:
    - danielle
consulted:
    - rsmith
tags:
    - infrastructure
    - migration
links:
    related-to:
        - "0004"
supersedes:
    - "0017"
comments:
    - author: danielle
      date: "2026-05-13 14:22:01"
      text: Initial decision; revisit after Q3.
    - author: rsmith
      date: "2026-05-14 09:00:00"
      text: Confirmed in prod load test.
---

# Use Kafka for the event bus

## Context and Problem Statement

We need an event bus that scales to 100k msg/sec.

## Considered Options

* Use Kafka
* Use NATS
* Roll our own

## Decision Outcome

Chosen option: "Use Kafka", because the operations team already runs it.

### Consequences

* Good, because zero new operational burden.
* Bad, because Kafka's footprint is heavyweight for our scale.

## Comments

* **2026-05-13 14:22:01 — @danielle:** Initial decision; revisit after Q3.
* **2026-05-14 09:00:00 — @rsmith:** Confirmed in prod load test.
