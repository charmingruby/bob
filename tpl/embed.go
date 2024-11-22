package tpl

import "embed"

//go:embed generate/*.tpl
var GenerateTemplateFS embed.FS
