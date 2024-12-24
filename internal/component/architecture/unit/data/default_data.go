package data

import (
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
)

type DefaultData struct {
	Name string
}

func NewDefaultData(name string) DefaultData {
	return DefaultData{
		Name: base.PublicNameFormat(name),
	}
}
