---
adr_id: "0004"
comments:
    - author: ""
      comment: "1"
      date: "2026-05-15 19:03:03"
status: decided
title: Required project files must exist at the repository root
---

## <a name="question"></a> Question

Which standard project files must always be present at the repository root to ensure the project is discoverable, licensable, and buildable?

## <a name="options"></a> Options

1. <a name="option-1"></a> Enforce required files: README.md, LICENSE, go.mod, and .gitignore must exist at the repository root
2. <a name="option-2"></a> No enforcement: rely on contributor discipline to maintain required files

## <a name="criteria"></a> Criteria

The repository must be immediately usable by a new contributor. Missing a LICENSE file creates legal ambiguity; missing README.md makes the project undiscoverable; missing go.mod prevents the module from being built; missing .gitignore leads to unwanted files being committed.

## <a name="outcome"></a> Outcome
We decided for [Option 1](#option-1).

## <a name="comments"></a> Comments
<a name="comment-1"></a>1. (2026-05-15 19:03:03) : marked decision as decided
