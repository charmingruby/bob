package model

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/oklog/ulid/v2"
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
        assert.Nil(t, m.UpdatedAt)
        assert.Nil(t, m.DeletedAt)
    })
}

func Test_From{{ .ModelName }}(t *testing.T) {
    t.Run("it should be able to create a {{ .ModelName }} from input", func(t *testing.T) {
        in := {{ .ModelName }}{
            ID:        ulid.Make().String(),
            Name:      "{{ .ModelName }}",
            CreatedAt: time.Now(),
            UpdatedAt: nil,
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

func Test_SoftDelete{{ .ModelName }}(t *testing.T) {
    t.Run("it should be able to soft delete a {{ .ModelName }}", func(t *testing.T) {
        now := time.Now()

        in := {{ .ModelName }}{
            ID:        ulid.Make().String(),
            Name:      "{{ .ModelName }}",
            CreatedAt: now,
            UpdatedAt: nil,
            DeletedAt: nil,
        }
        
        m := From{{ .ModelName }}(in)

        m.SoftDelete()

        assert.NotNil(t, m.DeletedAt)
        assert.NotNil(t, m.UpdatedAt)
        assert.True(t, m.DeletedAt.After(now) || m.DeletedAt.Equal(now))
        assert.True(t, m.UpdatedAt.After(now) || m.UpdatedAt.Equal(now))
    })
}