package inputport

type ModelCopy interface {
	Copy(modelPath, targetPath string, filters map[string][]string) error
}

type ModelImport interface {
	Import(sourcePath, targetPath string, filters map[string][]string) error
}

type ModelInit interface {
	Init(modelPath string) error
}

type ModelMerge interface {
	Merge(modelAPath, modelBPath, targetPath string, filters map[string][]string) error
}

type ModelRebuildIndex interface {
	RebuildIndex(modelPath string) error
}

type ModelValidate interface {
	Validate(modelPath string) error
}
