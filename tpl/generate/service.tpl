package {{ .PackageName }}

type {{ .ServiceName }}Params struct{}

type {{ .ServiceName }}Result struct{}

func (s *Service) {{ .ServiceName }}{{ .PackageName }}(params {{ .ServiceName }}Params) ({{ .ServiceName }}Result, error) {
	return {{ .ServiceName }}Result{}, nil
}
