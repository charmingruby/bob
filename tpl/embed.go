package tpl

import "embed"

//go:embed gen/*.tpl
var GenerateTemplateFS embed.FS
