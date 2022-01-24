package content

// Module is the go.mod template used for new projects.
var Module = `module {{ProjectName}}

go 1.16

require (
	github.com/gin-gonic/gin
)
`
