package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is the golang-project-template Version.
const Version = "1.0.2"

func init() {
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of golang-web",
	Long:  `This is the version of golang-web.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("💻 golang-web version is %s\n", Version)
	},
}
