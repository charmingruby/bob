package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New{{ .ModelName }}(t *testing.T) {
	t.Run("it should be able to create a new {{ .ModelName }}", func(t *testing.T) {
		in := New{{ .ModelName }}Input{
			Name: "{{ .ModelName }} model",
		}
		u := New{{ .ModelName }}(in)

		assert.Equal(t, u.ID, "id")
		assert.Equal(t, u.Name, in.Name)
	})
}

func Test_{{ .ModelName }}Marshal(t *testing.T) {
	t.Run("it should be able to marshal {{ .ModelName }}", func(t *testing.T) {
		in := New{{ .ModelName }}Input{
			Name: "{{ .ModelName }} model",
		}
		u := New{{ .ModelName }}(in)

		data, err := u.MarshalJSON()

		assert.Nil(t, err)
		assert.NotNil(t, data)
	})
}
