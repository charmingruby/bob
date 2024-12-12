package data

import "github.com/charmingruby/bob/pkg/formatter"

type HandlerData struct {
	SourcePath string
	Module     string
	Name       string
}

func NewHandlerData(sourcePath, module, name string) HandlerData {
	return HandlerData{
		SourcePath: sourcePath,
		Module:     module,
		Name:       formatter.ToCamelCase(name),
	}
}
