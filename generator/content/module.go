package content

// Module is the go.mod template used for new projects.
var Module = `module {{.}}

go 1.16

require (
	github.com/gin-gonic/gin
)
`
