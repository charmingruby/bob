package model

import (
	"encoding/json"
	"time"
)

type {{ .Name }} struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
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
		DeletedAt: nil,
	}
}

func (m *{{ .Name }}) MarshalJSON() ([]byte, error) {
	copy := *m

	return json.Marshal(copy)
}
