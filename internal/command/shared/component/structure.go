package component

type Component struct {
	Identifier string
	Directory  string
	Module     string
	Name       string
	FullName   string
	Suffix     string
	Package    Package
	Data       any
	ActionType string
	HasTest    bool
}

type Package struct {
	Name               string
	RegistryIdentifier string
	Registry           string
}
