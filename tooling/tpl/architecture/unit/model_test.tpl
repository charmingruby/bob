package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New{{ .Name }}(t *testing.T) {
	t.Run("it should be able to create a new {{ .Name }}", func(t *testing.T) {
		in := New{{ .Name }}Input{
			Name: "{{ .Name }}",
		}
		m := New{{ .Name }}(in)

		assert.Equal(t, m.ID, "id")
		assert.Equal(t, m.Name, in.Name)
	})
}

func Test_{{ .Name }}Marshal(t *testing.T) {
	t.Run("it should be able to marshal {{ .Name }}", func(t *testing.T) {
		in := New{{ .Name }}Input{
			Name: "{{ .Name }}",
		}
		m := New{{ .Name }}(in)

		data, err := m.MarshalJSON()

		assert.Nil(t, err)
		assert.NotNil(t, data)
	})
}
