package structure

type Pure struct {
	Name string
}

func NewPure(name string) Pure {
	return Pure{
		Name: PublicNameFormat(name),
	}
}

type DependentPackage struct {
	SourcePath string
	Module     string
	Name       string
}

func NewDependentPackage(sourcePath, module, name string) DependentPackage {
	return DependentPackage{
		SourcePath: sourcePath,
		Module:     ModuleFormat(module),
		Name:       PublicNameFormat(name),
	}
}
