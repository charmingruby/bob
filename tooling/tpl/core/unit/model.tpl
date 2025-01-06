package model

import (
	"encoding/json"
	"time"

	"github.com/oklog/ulid/v2"
)

type {{ .ModelName }} struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type New{{ .ModelName }}Input struct {
	Name string
}

func New{{ .ModelName }}(in New{{ .ModelName }}Input) *{{ .ModelName }} {
	return &{{ .ModelName }}{
		ID:        ulid.Make().String(),
		Name:      in.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
}

func (m *{{ .ModelName }}) MarshalJSON() ([]byte, error) {
	copy := *m

	return json.Marshal(copy)
}
