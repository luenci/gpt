package content

import _ "embed"

var (
	//go:embed pre-commit.txt
	PreCommit string

	//go:embed docker.txt
	Dockerfile string

	//go:embed gitignore.txt
	GitIgnore string

	//go:embed makefile.txt
	Makefile string

	//go:embed module.txt
	Module string

	//go:embed main.txt
	MainCLT string

	//go:embed ginReadMe.txt
	GinReadMe string
)
