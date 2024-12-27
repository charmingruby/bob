package {{ .Database }}

import (
	"errors"

	"{{ .SourcePath }}/{{ .Module }}/core/model"
)

type {{ .RepositoryName }}Repository struct {
}

func New{{ .RepositoryName }}Repository() *{{ .RepositoryName }}Repository {
	return &{{ .RepositoryName }}Repository{}
}

func (r *{{ .RepositoryName }}Repository) Store(model *model.{{ .ModelName }}) error {
	return errors.New("unimplemented method")
}

func (r *{{ .RepositoryName }}Repository) FindByID(id string) (*model.{{ .ModelName }}, error) {
	return nil, errors.New("unimplemented method")
}

func (r *{{ .RepositoryName }}Repository) Save(model *model.{{ .ModelName }}) error {
	return errors.New("unimplemented method")
}

func (r *{{ .RepositoryName }}Repository) Delete(model *model.{{ .ModelName }}) error {
	return errors.New("unimplemented method")
}
