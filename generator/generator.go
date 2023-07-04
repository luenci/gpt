package generator

import "github.com/luenci/gpt/generator/content"

// TemplateMap is the map of templates register.
var (
	TemplateMap = map[string]string{
		".pre-commit-config.yaml": content.PreCommit,
		".gitignore":              content.GitIgnore,
		"Dockerfile":              content.Dockerfile,
		".golangci.yml":           content.GolangCI,
		"main.go":                 content.MainCLT,
	}

	ArgsTemplateMap = map[string]string{
		"Makefile": content.Makefile,
		"go.mod":   content.GoMod,
	}
)
