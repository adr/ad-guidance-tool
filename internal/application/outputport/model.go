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

type ModelValidate interface {
	ModelValidated(modelName string, indexErr, dataErr error)
}
