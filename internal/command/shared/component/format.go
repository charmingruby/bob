package component

import "github.com/charmingruby/bob/pkg/formatter"

func PublicNameFormat(name string) string {
	return formatter.ToCamelCase(name)
}

func ModuleFormat(module string) string {
	return formatter.ToSnakeCase(module)
}

func PrivateNameFormat(name string) string {
	return formatter.ToCamelLowerCase(name)
}
