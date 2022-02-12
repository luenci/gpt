package template

import "strings"

// ParseTemplate parse template directory and files.
func ParseTemplate(contents string) []string {
	return strings.Split(contents, "\n")
}
