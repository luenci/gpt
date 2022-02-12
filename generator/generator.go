package generator

import "github.com/luenci/gpt/generator/content"

// TemplateMap is the map of templates register.
var (
	TemplateMap = map[string]string{
		".pre-commit-config.yaml": content.PreCommit,
		".gitignore":              content.GitIgnore,
		"Dockerfile":              content.Dockerfile,
		"main.go":                 content.MainCLT,
		"README.md":               content.GinReadMe,
	}

	ArgsTemplateMap = map[string]string{
		"Makefile": content.Makefile,
	}
)
