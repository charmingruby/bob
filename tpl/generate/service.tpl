package {{ .Package }}

type {{ .Name }}Params struct{}

type {{ .Name }}Result struct{}

func ({{ .PackageRegistryIdentifier }} *{{ .PackageRegistry }}) {{ .Name }}(params {{ .Name }}Params) ({{ .Name }}Result, error) {
	return {{ .Name }}Result{}, nil
}
