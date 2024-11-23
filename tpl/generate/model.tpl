package {{ .Package }}

import (
	"encoding/json"
	"time"
)

type {{ .Name }} struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type New{{ .Name }}Input struct {
	Name string
}

func New{{ .Name }}(in New{{ .Name }}Input) *{{ .Name }} {
	return &{{ .Name }}{
		ID:        "id",
		Name:      in.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func ({{ .PackageRegistryIdentifier }} *{{ .Name }}) MarshalJSON() ([]byte, error) {
	copy := *{{ .PackageRegistryIdentifier }}

	return json.Marshal(copy)
}
