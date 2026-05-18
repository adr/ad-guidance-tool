package outputport

type ModelCopy interface {
	Copied(source, target string, copiedDecisions int)
}

type ModelImport interface {
	Imported(sourcePath, targetPath string, importedDecisions int) error
}

type ModelInit interface {
	Initialized(name string)
}

type ModelMerge interface {
	Merged(modelAPath, modelBPath, targetPath string, mergedDecisions int) error
}

type ModelRebuildIndex interface {
	IndexRebuilt(modelName string)
}

// ModelValidate receives the list of validation issues from `adg validate`.
// The legacy interface returned two error pointers (one for index, one for
// data) which were specific to the index-vs-files reconciliation logic that
// has been removed. The new interface passes a flat issues list; the printer
// is responsible for grouping per-decision and exit-code mapping.
type ModelValidate interface {
	ModelValidated(modelName string, issues []ValidationIssue)
}

// ValidationIssue is the outputport mirror of model.ValidationIssue. The
// outputport layer can't import the domain package without creating a cycle,
// so the issue shape is repeated here.
type ValidationIssue struct {
	ID      string
	Message string
}
