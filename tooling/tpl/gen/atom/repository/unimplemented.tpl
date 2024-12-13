package {{ .Database }}

import (
	"errors"

	"{{ .SourcePath }}/{{ .Module }}/core/model"
)

type {{ .Name }}Repository struct {
}

func New{{ .Name }}Repository() *{{ .Name }}Repository {
	return &{{ .Name }}Repository{}
}

func (r *{{ .Name }}Repository) Store(model model.{{ .Name }}) error {
	return errors.New("unimplemented method")
}

func (r *{{ .Name }}Repository) FindByID(id string) (*model.{{ .Name }}, error) {
	return nil, errors.New("unimplemented method")
}

func (r *{{ .Name }}Repository) Save(model model.{{ .Name }}) error {
	return errors.New("unimplemented method")
}

func (r *{{ .Name }}Repository) Delete(model model.{{ .Name }}) error {
	return errors.New("unimplemented method")
}
