package content

import _ "embed"

var (
	//go:embed pre-commit.txt
	PreCommit string

	//go:embed docker.txt
	Dockerfile string

	//go:embed gitignore.txt
	GitIgnore string

	//go:embed golangci.txt
	GolangCI string

	//go:embed makefile.txt
	Makefile string

	//go:embed main.txt
	MainCLT string

	//go:embed ginReadMe.txt
	GinReadMe string

	//go:embed DDDReadMe.txt
	DDDReadMe string

	//go:embed gomod.txt
	GoMod string
)
