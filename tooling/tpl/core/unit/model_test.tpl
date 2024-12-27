package model

import (
	"testing"

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
	})
}

func Test_{{ .ModelName }}Marshal(t *testing.T) {
	t.Run("it should be able to marshal {{ .ModelName }}", func(t *testing.T) {
		in := New{{ .ModelName }}Input{
			Name: "{{ .ModelName }}",
		}
		m := New{{ .ModelName }}(in)

		data, err := m.MarshalJSON()

		assert.Nil(t, err)
		assert.NotNil(t, data)
	})
}
