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
go.sum
CHANGELOG.md
README.md`

// ParseTemplate parse template directory and files.
func ParseTemplate(contents string) []string {
	return strings.Split(contents, "\n")
}
