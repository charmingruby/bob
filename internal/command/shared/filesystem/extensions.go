package filesystem

import "github.com/ettle/strcase"

func formatTplFile(name string) string {
	return name + ".tpl"
}

func formatGoFile(name string) string {
	return strcase.ToSnake(name) + ".go"
}
