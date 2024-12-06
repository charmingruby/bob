package structure

import "github.com/charmingruby/bob/internal/command/shared/component"

type ExchangeData struct {
	ExchangePackage string
	Exchange        string
	Name            string
}

func NewExchangeData(exchange, name string) ExchangeData {
	return ExchangeData{
		ExchangePackage: component.ModuleFormat(exchange),
		Exchange:        component.PublicNameFormat(exchange),
		Name:            component.PublicNameFormat(name),
	}
}
