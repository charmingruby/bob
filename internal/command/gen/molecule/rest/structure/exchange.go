package structure

import "github.com/charmingruby/bob/internal/command/shared/component/structure"

type ExchangeData struct {
	ExchangePackage string
	Exchange        string
	Name            string
}

func NewExchangeData(exchange, name string) ExchangeData {
	return ExchangeData{
		ExchangePackage: structure.ModuleFormat(exchange),
		Exchange:        structure.PublicNameFormat(exchange),
		Name:            structure.PublicNameFormat(name),
	}
}
