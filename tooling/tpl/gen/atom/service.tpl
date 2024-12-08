package service

type {{ .Name }}Params struct{}

type {{ .Name }}Result struct{}

func (s *Service) {{ .Name }}Service(params {{ .Name }}Params) ({{ .Name }}Result, error) {
	return {{ .Name }}Result{}, nil
}
