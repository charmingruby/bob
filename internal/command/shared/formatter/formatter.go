package formatter

import (
	"strings"

	"github.com/ettle/strcase"
)

func ToSnakeCase(s string) string {
	return strcase.ToSnake(s)
}

func ToCamelCase(s string) string {
	return strcase.ToGoCase(s, strcase.TitleCase, 0)
}

func ToLowerCase(s string) string {
	return strings.ToLower(s)
}
