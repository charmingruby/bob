package repository

import "{{ .SourcePath }}/{{ .Module }}/core/model"

type {{ .RepositoryName }}Repository interface {
	Store(model *model.{{ .ModelName }}) error
	FindByID(id string) (*model.{{ .ModelName }}, error)
	Delete(model *model.{{ .ModelName }}) error
}
