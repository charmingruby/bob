package repository

import "{{ .SourcePath }}/{{ .Module }}/core/model"

type {{ .Name }}Repository interface {
	Store(model model.{{ .Name }}) error
	FindByID(id string) (*model.{{ .Name }}, error)
	Save(model model.{{ .Name }}) error
	Delete(model model.{{ .Name }}) error
}
