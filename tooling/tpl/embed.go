package tpl

import "embed"

//go:embed */*/*.tpl */*/*/*.tpl */*/*/*/*.tpl
var GenerateTemplateFS embed.FS
