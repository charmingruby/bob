package service

import "{{ .SourcePath }}/{{ .Module }}/core/repository"

type Service struct {
    {{ .PrivateRepositoryName }}Repo repository.{{ .RepositoryName }}Repository
}

func New({{ .PrivateRepositoryName }}Repo repository.{{ .RepositoryName }}Repository) *Service {
	return &Service{
        {{ .PrivateRepositoryName }}Repo: {{ .PrivateRepositoryName }}Repo,
    }
}
