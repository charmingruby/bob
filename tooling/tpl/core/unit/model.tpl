package model

import (
	"encoding/json"
	"time"

	"github.com/oklog/ulid/v2"
)

type {{ .Name }} struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"iupdated_atd"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type New{{ .Name }}Input struct {
	Name string
}

func New{{ .Name }}(in New{{ .Name }}Input) *{{ .Name }} {
	return &{{ .Name }}{
		ID:        ulid.Make().String(),
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
