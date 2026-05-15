grammar ADE;

// =====================
// Parser rules
// =====================

file
  : adrDecl (selectorDecl | ruleDecl)* EOF
  ;

adrDecl
  : ADR STRING STRING
  ;

selectorDecl
  : (COMPONENT | CLASS | INTERFACE | PATH) STRING EQ STRING
  ;

ruleDecl
  : ruleType STRING LBRACE ruleStmt* RBRACE
  ;

ruleType
  : FILE
  | CODE
  ;

ruleStmt
  : assertionStmt
  | excludeStmt
  | severityStmt
  ;

// ---- Assertions ----
// Main rule assertions using natural language

assertionStmt
  : subjectExpr mustExpr verbPhrase
  ;

// Subject: can be a selector name or inline pattern
// Optionally scoped via "in <targetExpr>" for subset relation
subjectExpr
  : IDENTIFIER                                # SelectorRef
  | selectorType STRING IN targetExpr         # SubsetLiteral
  | selectorType MATCH STRING IN targetExpr   # SubsetMatch
  | selectorType IN targetExpr                # SubsetAll
  | selectorType STRING                       # InlineLiteral
  | selectorType MATCH STRING                 # InlineMatch
  | selectorType                              # InlineType
  ;

// Target: same as subject but simpler (no scope)
targetExpr
  : IDENTIFIER                                # TargetSelectorRef
  | selectorType STRING                       # TargetInlineLiteral
  | selectorType MATCH STRING                 # TargetInlineMatch
  | STRING                                    # TargetStringLiteral
  ;

selectorType
  : COMPONENT | CLASS | INTERFACE | PATH
  ;

// Modality: must, must not, must only
mustExpr
  : MUST NOT
  | MUST ONLY
  | MUST
  ;

// Verb phrases: different verbs take different arguments
// Optional filler words: on, by, with for natural language
verbPhrase
  : DEPEND ON? targetExpr (COMMA targetExpr)*                # DependOnPhrase
  | EXIST                                                    # ExistPhrase
  | CONTAIN STRING                                           # ContainPhrase
  | IMPLEMENT INTERFACE? targetExpr                          # ImplementPhrase
  | EXTEND CLASS? targetExpr                                 # ExtendPhrase
  | BE? ANNOTATED WITH? STRING                               # AnnotatedPhrase
  | BE? ACCESSED BY targetExpr (COMMA targetExpr)*           # AccessedByPhrase
  | BE? ACYCLIC                                              # AcyclicPhrase
  | BE? IN targetExpr                                        # InPhrase
  | MATCH STRING                                             # MatchPhrase
  | BE? visibility                                           # VisibilityPhrase
  | BE? typeConstraint                                       # TypeConstraintPhrase
  ;

visibility
  : PUBLIC
  | INTERNAL
  | PRIVATE
  ;

typeConstraint
  : ABSTRACT
  | SEALED
  | STATIC
  ;

// ---- Exclusions ----
// Standalone exclude statements for dependency rules

excludeStmt
  : EXCLUDE CLASS STRING                            # ExcludeClass
  | EXCLUDE CLASS IMPLEMENTING INTERFACE STRING     # ExcludeClassImplementing
  | EXCLUDE COMPONENT STRING                        # ExcludeComponent
  | EXCLUDE STRING                                  # ExcludePattern
  ;

// ---- Severity ----

severityStmt
  : SEVERITY severityValue
  ;

severityValue
  : ERROR
  | WARNING
  ;

// =====================
// Lexer rules
// =====================

// Keywords
ADR               : 'adr';
FILE              : 'file';
CODE              : 'code';

COMPONENT         : 'component';
CLASS             : 'class';
INTERFACE         : 'interface';
PATH              : 'path';

MUST              : 'must';
NOT               : 'not';
ONLY              : 'only';

// Verbs
DEPEND            : 'depend';
EXIST             : 'exist';
CONTAIN           : 'contain';
IMPLEMENT         : 'implement';
EXTEND            : 'extend';
ANNOTATED         : 'annotated';
ACCESSED          : 'accessed';
ACYCLIC           : 'acyclic';

// Filler words (optional for natural language)
ON                : 'on';
BE                : 'be';
BY                : 'by';
WITH              : 'with';

// Other keywords
MATCH             : 'match';
EXCLUDE           : 'exclude';
IMPLEMENTING      : 'implementing';
IN                : 'in';

// Visibility modifiers
PUBLIC            : 'public';
INTERNAL          : 'internal';
PRIVATE           : 'private';

// Type constraints
ABSTRACT          : 'abstract';
SEALED            : 'sealed';
STATIC            : 'static';

SEVERITY          : 'severity';
ERROR             : 'error';
WARNING           : 'warning';

// Punctuation
EQ                : '=';
COMMA             : ',';
LBRACE            : '{';
RBRACE            : '}';

// Identifier (for selector references)
// Must not collide with keywords
IDENTIFIER
  : [A-Z][a-zA-Z0-9_]*
  ;

// String literal with \" support
STRING
  : '"' ( '\\"' | ~["\r\n] )* '"'
  ;

// Comments
COMMENT
  : '#' ~[\r\n]* -> skip
  ;

// Whitespace
WS
  : [ \t\r\n]+ -> skip
  ;
