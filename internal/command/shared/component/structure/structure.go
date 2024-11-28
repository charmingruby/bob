package structure

type PureData struct {
	Name string
}

func NewPureData(name string) PureData {
	return PureData{
		Name: PublicNameFormat(name),
	}
}

type DependentPackageData struct {
	SourcePath string
	Module     string
	Name       string
}

func NewDependentPackageData(sourcePath, module, name string) DependentPackageData {
	return DependentPackageData{
		SourcePath: sourcePath,
		Module:     ModuleFormat(module),
		Name:       PublicNameFormat(name),
	}
}
