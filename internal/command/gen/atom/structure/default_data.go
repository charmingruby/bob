package structure

import (
	"github.com/charmingruby/bob/internal/command/shared/component"
)

type DefaultData struct {
	Name string
}

func NewDefaultData(name string) DefaultData {
	return DefaultData{
		Name: component.PublicNameFormat(name),
	}
}
