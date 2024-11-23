package component

import (
	"strings"

	"github.com/ettle/strcase"
)

func toSnakeCase(s string) string {
	return strcase.ToSnake(s)
}

func toCamelCase(s string) string {
	return strcase.ToGoCase(s, strcase.TitleCase, 0)
}

func toLowerCase(s string) string {
	return strings.ToLower(s)
}
