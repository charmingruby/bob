package model

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
)

func Test_New{{ .ModelName }}(t *testing.T) {
    t.Run("it should be able to create a new {{ .ModelName }}", func(t *testing.T) {
        in := New{{ .ModelName }}Input{
            Name: "{{ .ModelName }}",
        }
        m := New{{ .ModelName }}(in)

        assert.NotEmpty(t, m.ID)
        assert.Equal(t, in.Name, m.Name)
        assert.NotZero(t, m.CreatedAt)
        assert.NotZero(t, m.UpdatedAt)
        assert.Nil(t, m.DeletedAt)
    })
}

func Test_From{{ .ModelName }}(t *testing.T) {
    t.Run("it should be able to create a {{ .ModelName }} from input", func(t *testing.T) {
        now := time.Now()
        in := {{ .ModelName }}{
            ID:        "01F8MECHZX3TBDSZ7XRADM79XE",
            Name:      "{{ .ModelName }}",
            CreatedAt: now,
            UpdatedAt: now,
            DeletedAt: nil,
        }
        m := From{{ .ModelName }}(in)

        assert.Equal(t, in.ID, m.ID)
        assert.Equal(t, in.Name, m.Name)
        assert.Equal(t, in.CreatedAt, m.CreatedAt)
        assert.Equal(t, in.UpdatedAt, m.UpdatedAt)
        assert.Equal(t, in.DeletedAt, m.DeletedAt)
    })
}
