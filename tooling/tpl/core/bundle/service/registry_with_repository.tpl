package service

import "{{ .SourcePath }}/{{ .Module }}/core/repository"

type Service struct {
    {{ .LowerCaseRepositoryName }}Repository repository.{{ .CapitalizedRepositoryName }}Repository
}

type Input struct {
    {{ .CapitalizedRepositoryName }}Repository repository.{{ .CapitalizedRepositoryName }}Repository
}

func New(in Input) *Service {
	return &Service{
        {{ .LowerCaseRepositoryName }}Repository: in.{{ .CapitalizedRepositoryName }}Repository,
    }
}
