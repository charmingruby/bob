package filesystem

import (
	"github.com/charmingruby/bob/pkg/formatter"
)

func formatTplFile(name string) string {
	return name + ".tpl"
}

func formatGoFile(name string) string {
	return formatter.ToSnakeCase(name) + ".go"
}
