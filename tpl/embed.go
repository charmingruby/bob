package tpl

import "embed"

//go:embed */*/*.tpl */*/*/*.tpl
var GenerateTemplateFS embed.FS
