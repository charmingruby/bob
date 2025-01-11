package service

import (
	"{{ .SourcePath }}/{{ .Module }}/core/model"
	"{{ .SourcePath }}/shared/custom_err/core_err"
)

type {{ .ServiceName }}Params struct {
	Name string
}

type {{ .ServiceName }}Result struct {
	ID string
}

func (s *Service) {{ .ServiceName }}(params {{ .ServiceName }}Params) ({{ .ServiceName }}Result, error) {
	{{ .LowerCaseModel }} := model.New{{ .CapitalizedModel }}(model.New{{ .CapitalizedModel }}Input{
		Name: params.Name,
	})
	
	if err := s.{{ .LowerCaseModel }}Repository.Store({{ .LowerCaseModel }}); err != nil {
		return {{ .ServiceName }}Result{}, err
	}

	{{ .LowerCaseModel }}.SoftDelete()
	if err := s.{{ .LowerCaseModel }}Repository.Delete({{ .LowerCaseModel }}); err != nil {
		return {{ .ServiceName }}Result{}, err
	}

	{{ .LowerCaseModel }}Found, err := s.{{ .LowerCaseModel }}Repository.FindByID({{ .LowerCaseModel }}.ID)
	if err != nil {
		return {{ .ServiceName }}Result{}, err
	}
	
	if {{ .LowerCaseModel }}Found == nil {
		return {{ .ServiceName }}Result{}, core_err.NewResourceNotFoundErr("{{ .LowerCaseModel }}")
	}

	return {{ .ServiceName }}Result{
		ID: {{ .LowerCaseModel }}Found.ID,
	}, nil
}
