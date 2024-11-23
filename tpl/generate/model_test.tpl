package {{ .Package }}

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New{{ .Name }}(t *testing.T) {
	t.Run("it should be able to create a new {{ .Name }}", func(t *testing.T) {
		in := New{{ .Name }}Input{
			Name: "{{ .Name }}",
		}
		{{ .PackageRegistryIdentifier }} := New{{ .Name }}(in)

		assert.Equal(t, {{ .PackageRegistryIdentifier }}.ID, "id")
		assert.Equal(t, {{ .PackageRegistryIdentifier }}.Name, in.Name)
	})
}

func Test_{{ .Name }}Marshal(t *testing.T) {
	t.Run("it should be able to marshal {{ .Name }}", func(t *testing.T) {
		in := New{{ .Name }}Input{
			Name: "{{ .Name }}",
		}
		{{ .PackageRegistryIdentifier }} := New{{ .Name }}(in)

		data, err := {{ .PackageRegistryIdentifier }}.MarshalJSON()

		assert.Nil(t, err)
		assert.NotNil(t, data)
	})
}
