package base

import "github.com/charmingruby/bob/pkg/formatter"

func CapitalizedFormat(name string) string {
	return formatter.ToCamelCase(name)
}

func SnakeCaseFormat(module string) string {
	return formatter.ToSnakeCase(module)
}

func LowerCaseFormat(name string) string {
	return formatter.ToCamelLowerCase(name)
}

func CamelCaseFormat(path string) string {
	return formatter.ToCamelLowerCase(path)
}
