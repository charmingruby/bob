package data

import "github.com/charmingruby/bob/internal/component/base"

type ExchangeData struct {
	ExchangePackage string
	Exchange        string
	Name            string
}

func NewExchangeData(exchange, name string) ExchangeData {
	return ExchangeData{
		ExchangePackage: base.ModuleFormat(exchange),
		Exchange:        base.PublicNameFormat(exchange),
		Name:            base.PublicNameFormat(name),
	}
}
