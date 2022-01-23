package template

import (
	"strings"
)

var GinDemo = `
config
docs
middleware
models
routers
service
runtime
pkg
tests
types
request
response
go.mod
go.sum
main.go
CHANGELOG.md
README.md
.gitignore
.pre-commit-config.yaml
`

// ParseTemplate parse template directory and files.
func ParseTemplate(contents string) []string {
	return strings.Split(contents, "\n")
}
