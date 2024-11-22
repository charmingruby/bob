package model

import (
	"encoding/json"
	"time"
)

type {{ .ModelName }} struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type New{{ .ModelName }}Input struct {
	Name string
}

func New{{ .ModelName }}(in New{{ .ModelName }}Input) *{{ .ModelName }} {
	return &{{ .ModelName }}{
		ID:        "id",
		Name:      in.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m *{{ .ModelName }}) MarshalJSON() ([]byte, error) {
	copy := *m

	return json.Marshal(copy)
}
