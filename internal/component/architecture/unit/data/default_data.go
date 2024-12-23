package data

import (
	"github.com/charmingruby/bob/internal/component/base"
)

type DefaultData struct {
	Name string
}

func NewDefaultData(name string) DefaultData {
	return DefaultData{
		Name: base.PublicNameFormat(name),
	}
}
