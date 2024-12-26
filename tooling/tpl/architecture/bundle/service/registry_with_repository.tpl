package service

import "{{ .SourcePath }}/{{ .Module }}/core/repository"

type Service struct {
    {{ .PrivateRepositoryName }}Repository repository.{{ .RepositoryName }}Repository
}

type Input struct {
    {{ .RepositoryName }}Repository repository.{{ .RepositoryName }}Repository
}

func New(in Input) *Service {
	return &Service{
        {{ .PrivateRepositoryName }}Repository: in.{{ .RepositoryName }}Repository,
    }
}
