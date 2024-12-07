package data

import "github.com/charmingruby/bob/pkg/formatter"

type HandlerData struct {
	SourcePath   string
	CommonModule string
	Module       string
	Name         string
}

func NewHandlerData(sourcePath, commonModule, module, name string) HandlerData {
	return HandlerData{
		SourcePath:   sourcePath,
		CommonModule: commonModule,
		Module:       module,
		Name:         formatter.ToCamelCase(name),
	}
}
