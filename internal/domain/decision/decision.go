package decision

import "adg/internal/domain/decision/madr"

// Decision and Comment are aliases for the MADR-shaped types in the madr
// subpackage. The decision package keeps these aliases for the duration of the
// fork refactor so existing callers (service, interactors, adapters, cmd) can
// continue to reference `decision.Decision` while the storage and serialization
// logic lives in `madr`. After PR 1d the alias layer may be flattened entirely.
type Decision = madr.Decision

// Comment is the MADR-shaped comment entry: author + timestamp + text. Replaces
// the legacy Comment struct whose `Comment` field stored the comment number as
// a string instead of the actual text (the §A.1 data-loss bug).
type Comment = madr.Comment
