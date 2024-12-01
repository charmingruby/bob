package structure

import "github.com/charmingruby/bob/internal/command/shared/component/structure"

type DefaultData struct {
	Name string
}

func NewDefaultData(name string) DefaultData {
	return DefaultData{
		Name: structure.PublicNameFormat(name),
	}
}
